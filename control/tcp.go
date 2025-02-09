/*
 * SPDX-License-Identifier: AGPL-3.0-only
 * Copyright (c) 2022-2023, daeuniverse Organization <dae@v2raya.org>
 */

package control

import (
	"fmt"
	"net"
	"net/netip"
	"strings"
	"time"

	"github.com/daeuniverse/dae/common"
	"github.com/daeuniverse/dae/common/consts"
	"github.com/daeuniverse/dae/component/outbound/dialer"
	"github.com/daeuniverse/dae/component/sniffing"
	"github.com/mzz2017/softwind/netproxy"
	"github.com/mzz2017/softwind/pkg/zeroalloc/io"
	"github.com/sirupsen/logrus"
	"golang.org/x/sys/unix"
)

const (
	TcpSniffBufSize = 4096
)

func (c *ControlPlane) handleConn(lConn net.Conn) (err error) {
	defer lConn.Close()

	// Sniff target domain.
	sniffer := sniffing.NewConnSniffer(lConn, TcpSniffBufSize, c.sniffingTimeout)
	// ConnSniffer should be used later, so we cannot close it now.
	defer sniffer.Close()
	domain, err := sniffer.SniffTcp()
	if err != nil && !sniffing.IsSniffingError(err) {
		return err
	}

	// Get tuples and outbound.
	src := lConn.RemoteAddr().(*net.TCPAddr).AddrPort()
	dst := lConn.LocalAddr().(*net.TCPAddr).AddrPort()
	routingResult, err := c.core.RetrieveRoutingResult(src, dst, unix.IPPROTO_TCP)
	if err != nil {
		// WAN. Old method.
		var value bpfDstRoutingResult
		ip6 := src.Addr().As16()
		if e := c.core.bpf.TcpDstMap.Lookup(bpfIpPort{
			Ip:   struct{ U6Addr8 [16]uint8 }{U6Addr8: ip6},
			Port: common.Htons(src.Port()),
		}, &value); e != nil {
			return fmt.Errorf("failed to retrieve target info %v: %v, %v", src.String(), err, e)
		}
		routingResult = &value.RoutingResult

		dstAddr, ok := netip.AddrFromSlice(common.Ipv6Uint32ArrayToByteSlice(value.Ip))
		if !ok {
			return fmt.Errorf("failed to parse dest ip: %v", value.Ip)
		}
		dst = netip.AddrPortFrom(dstAddr, common.Htons(value.Port))
	}
	src = common.ConvergeAddrPort(src)
	dst = common.ConvergeAddrPort(dst)

	// Get outbound.
	var outboundIndex = consts.OutboundIndex(routingResult.Outbound)
	if c.dialMode == consts.DialMode_DomainCao && domain != "" {
		outboundIndex = consts.OutboundControlPlaneRouting
	}

	dialTarget, shouldReroute := c.ChooseDialTarget(outboundIndex, dst, domain)
	if shouldReroute {
		outboundIndex = consts.OutboundControlPlaneRouting
	}

	switch outboundIndex {
	case consts.OutboundDirect:
	case consts.OutboundControlPlaneRouting:
		if outboundIndex, routingResult.Mark, _, err = c.Route(src, dst, domain, consts.L4ProtoType_TCP, routingResult); err != nil {
			return err
		}
		routingResult.Outbound = uint8(outboundIndex)

		if c.log.IsLevelEnabled(logrus.TraceLevel) {
			c.log.Tracef("outbound: %v => %v",
				consts.OutboundControlPlaneRouting.String(),
				outboundIndex.String(),
			)
		}
		// Reset dialTarget.
		dialTarget, _ = c.ChooseDialTarget(outboundIndex, dst, domain)
	default:
	}
	// TODO: Set-up ip to domain mapping and show domain if possible.
	if outboundIndex < 0 || int(outboundIndex) >= len(c.outbounds) {
		return fmt.Errorf("outbound id from bpf is out of range: %v not in [0, %v]", outboundIndex, len(c.outbounds)-1)
	}
	outbound := c.outbounds[outboundIndex]
	networkType := &dialer.NetworkType{
		L4Proto:   consts.L4ProtoStr_TCP,
		IpVersion: consts.IpVersionFromAddr(dst.Addr()),
		IsDns:     false,
	}
	d, _, err := outbound.Select(networkType)
	if err != nil {
		return fmt.Errorf("failed to select dialer from group %v (%v): %w", outbound.Name, networkType.String(), err)
	}

	if c.log.IsLevelEnabled(logrus.InfoLevel) {
		c.log.WithFields(logrus.Fields{
			"network":  networkType.String(),
			"outbound": outbound.Name,
			"policy":   outbound.GetSelectionPolicy(),
			"dialer":   d.Property().Name,
			"sniffed":  domain,
			"ip":       RefineAddrPortToShow(dst),
			"pid":      routingResult.Pid,
			"pname":    ProcessName2String(routingResult.Pname[:]),
			"mac":      Mac2String(routingResult.Mac[:]),
		}).Infof("%v <-> %v", RefineSourceToShow(src, dst.Addr(), consts.LanWanFlag_NotApplicable), dialTarget)
	}

	// Dial and relay.
	rConn, err := d.Dial(MagicNetwork("tcp", routingResult.Mark), dialTarget)
	if err != nil {
		return fmt.Errorf("failed to dial %v: %w", dst, err)
	}
	defer rConn.Close()

	if err = RelayTCP(sniffer, rConn); err != nil {
		switch {
		case strings.HasSuffix(err.Error(), "write: broken pipe"),
			strings.HasSuffix(err.Error(), "i/o timeout"):
			return nil // ignore
		default:
			return fmt.Errorf("handleTCP relay error: %w", err)
		}
	}
	return nil
}

type WriteCloser interface {
	CloseWrite() error
}

func RelayTCP(lConn, rConn netproxy.Conn) (err error) {
	eCh := make(chan error, 1)
	go func() {
		_, e := io.Copy(rConn, lConn)
		if rConn, ok := rConn.(WriteCloser); ok {
			rConn.CloseWrite()
		}
		rConn.SetReadDeadline(time.Now().Add(10 * time.Second))
		eCh <- e
	}()
	_, e := io.Copy(lConn, rConn)
	if lConn, ok := lConn.(WriteCloser); ok {
		lConn.CloseWrite()
	}
	lConn.SetReadDeadline(time.Now().Add(10 * time.Second))
	if e != nil {
		<-eCh
		return e
	}
	return <-eCh
}
