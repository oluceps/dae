# Changelogs

Also seen in [GitHub Releases](https://github.com/daeuniverse/dae/releases)

HTML version available at https://dae.v2raya.org/docs/current/changelogs

## Query history releases

```bash
curl --silent "https://api.github.com/repos/daeuniverse/dae/releases" | jq -r '.[] | {tag_name,created_at,prerelease}'
```

## Releases

- [0.1.10rc1 (Pre-release)](#0110rc1-pre-release)
- [0.1.10rc (Pre-release)](#0110rc-pre-release)
- [0.1.9-patch.1 (Current)](#019-patch1-current)
- [0.1.9](#019)
- [0.1.8](#018)
- [0.1.7](#017)
- [0.1.6](#016)
- [0.1.5](#015)
- [0.1.4](#014)
- [0.1.3](#013)
- [0.1.2](#012)
- [0.1.1](#011)
- [0.1.0](#010)

### 0.1.10rc1 (Pre-release)

> Release date: 2023/05/15

#### 功能变更

- patch(geodata): 修复由 #84 导致的错误的 geodata 搜索路径 `/etc/dae/dae` by @mzz2017 in https://github.com/daeuniverse/dae/pull/90

#### Changes

- patch(geodata): fix incorrect geodata search path `/etc/dae/dae` caused by #84 by @mzz2017 in https://github.com/daeuniverse/dae/pull/90

**Full Changelog**: https://github.com/daeuniverse/dae/compare/v0.1.10rc...v0.1.10rc1

### 0.1.10rc (Pre-release)

> Release date: 2023/05/14

#### 功能变更

- feat: 支持 `tcp_check_http_method` by @mzz2017 in https://github.com/daeuniverse/dae/pull/77
- patch: 现在会优先在配置文件同目录搜索 geodata by @mzz2017 in https://github.com/daeuniverse/dae/pull/84
- fix(dns): 修复 0.1.8 版本中 PR #63 导致的 DNS 缓存不会过期的问题 by @mzz2017 in https://github.com/daeuniverse/dae/pull/87

#### 其他变更

- chore(Makefile): 添加 export GOOS=linux 以修复在 macos 上的构建 by @mzz2017 in https://github.com/daeuniverse/dae/pull/78
- chore: 添加 editorconfig 文件以美化 github 上对 go 文件的展示 by @yqlbu in https://github.com/daeuniverse/dae/pull/85
- chore: 添加 PR 模板 by @yqlbu in https://github.com/daeuniverse/dae/pull/86

#### Changes

- feat: support `tcp_check_http_method` by @mzz2017 in https://github.com/daeuniverse/dae/pull/77
- patch: search geodata at same dir with config first by @mzz2017 in https://github.com/daeuniverse/dae/pull/84
- fix(dns): cache would never expire caused by #63 by accident by @mzz2017 in https://github.com/daeuniverse/dae/pull/87

#### Other Changes

- chore(Makefile): add export GOOS=linux to build on macos by @mzz2017 in https://github.com/daeuniverse/dae/pull/78
- chore: add editorconfig by @yqlbu in https://github.com/daeuniverse/dae/pull/85
- chore: add pull_request_template by @yqlbu in https://github.com/daeuniverse/dae/pull/86

**Full Changelog**: https://github.com/daeuniverse/dae/compare/v0.1.9...v0.1.10rc

### 0.1.9-patch.1 (Current)

> Release date: 2023/05/14

#### 功能变更

- 修复(dns): 修复 0.1.8 版本中 PR #63 导致的 DNS 缓存不会过期的问题 by @mzz2017 in https://github.com/daeuniverse/dae/pull/87

#### Changes

- fix(dns): cache would never expire caused by #63 by accident by @mzz2017 in https://github.com/daeuniverse/dae/pull/87

**Full Changelog**: https://github.com/daeuniverse/dae/compare/v0.1.9...v0.1.9patch1

### 0.1.9

> Release date: 2023/05/09

#### 功能变更

- 修复 trojan UDP 不通的问题 by @mzz2017 in https://github.com/daeuniverse/dae/pull/71
- 修复 `curl http://[ipv6]:port` 不通的问题 by @mzz2017 in https://github.com/daeuniverse/dae/pull/70

#### 其他变更

- 修复 docker 镜像构建的 CI 会在特定名称的分支提交时意外地运行的问题 by @mzz2017 in https://github.com/daeuniverse/dae/pull/72

#### Changes

- fix(trojan): udp problem by @mzz2017 in https://github.com/daeuniverse/dae/pull/71
- fix(sniffing): fail to `curl http://[ipv6]:port` by @mzz2017 in https://github.com/daeuniverse/dae/pull/70

#### Other Changes

- fix(ci): PR runs docker action in some cases by @mzz2017 in https://github.com/daeuniverse/dae/pull/72

**Full Changelog**: https://github.com/daeuniverse/dae/compare/v0.1.8...v0.1.9

### 0.1.8 (Current)

> Release date: 2023/04/30

#### 功能变更

- optimize: DNS 缓存空解析和非 A/AAAA 查询，以及 reject 使用 0.0.0.0 和 :: by @mzz2017 in https://github.com/daeuniverse/dae/pull/63
- feat: 支持为 `tcp_check_url` 和 `udp_check_dns` 设定固定 IP 以防止 DNS 污染对 ipv4/ipv6 的支持带来影响 by @mzz2017 in https://github.com/daeuniverse/dae/commit/9493b9a0aa82573fed934bf62cc836f0fe148607

#### 其他变更

- chore: 增加 changelogs by @yqlbu in https://github.com/daeuniverse/dae/pull/55
- chore: 增加 pre-commit 钩子来格式化代码 by @yqlbu in https://github.com/daeuniverse/dae/pull/59
- style: 格式化 golang 代码风格 by @czybjtu in https://github.com/daeuniverse/dae/pull/58
- chore: 增加 issue 模板 by @yqlbu in https://github.com/daeuniverse/dae/pull/62
- chore(codeowner): 更新 ownership by @yqlbu in https://github.com/daeuniverse/dae/pull/64

#### New Contributors

- @czybjtu made their first contribution in https://github.com/daeuniverse/dae/pull/58

**Full Changelog**: https://github.com/daeuniverse/dae/compare/v0.1.7...v0.1.8

### 0.1.7

> Release date: 2023/04/16

#### 特性

支持 `global.sniffing_timeout` 来设定嗅探的超时时间，调大这个值对于时延较高的局域网来说较为有用。

#### 修复

1. 修复无法解析小火箭 shadowrocket 的 vmess+ws+tls 分享链接的问题。
2. 修复域名嗅探失败的问题。

#### PR

- chore: fix doamin regex example by @troubadour-hell in https://github.com/daeuniverse/dae/pull/53
- doc: add badges and contribution guide by @yqlbu in https://github.com/daeuniverse/dae/pull/54

#### New Contributors

- @troubadour-hell made their first contribution in https://github.com/daeuniverse/dae/pull/53

**Full Changelog**: https://github.com/daeuniverse/dae/compare/v0.1.6...v0.1.7

### 0.1.6

> Release date: 2023/04/09

#### 特性

- 支持在 dns 的 request 路由中使用 reject 出站。
- 支持在 routing 中使用 `must_组名` 的出站，该规则将强制作用于 DNS 请求，直接通过特定组发出，而绕过 dns 模块，提供给有特殊用途的用户使用。
- 支持在 routing 中使用 `must_rules` 的出站，命中该出站的 DNS 请求将绕过 dns 模块，直接进行路由并发出，提供给有特殊用途的用户使用。
- 支持 v2rayN 格式的 vmess 分享格式中的不标准 bool 值解析。
- 支持在 dns 中使用 `ipversion_prefer`，设定当域名是双栈时，只返回 ipv4 还是只返回 ipv6。

#### 修复

- 修复在 dns 的 response 路由中对无序 ip 序列的支持问题。
- 修复 trojan 可能的 panic 问题。
- dns 缓存丢失且 dial_mode 为 domain 时将尝试重路由，以缓解 dns 缓存丢失时无法使用 domain 进行路由的问题。
- 修复部分游戏无法进入的问题，该问题是由于 tcp 建立连接时，dae 总是等待客户端发包，但一些游戏场景中，首包是由服务端 push 的，因此陷入无限等待。

**Full Changelog**: https://github.com/daeuniverse/dae/compare/v0.1.5...v0.1.6

### 0.1.5

> Release date: 2023/03/29

#### 更新内容

- 修复 wan_interface 填入 auto 时可能出现的无法启动的问题。
- 修复 https 协议（naiveproxy）的支持问题，新增对 h2 的长连接和多路复用。
- 移除 DNS 抢答检测器，因为它不总是在所有地区都有效，而且在失效时会减慢查询速度。
- 文档（example.dae）：增加通过节点标签精确筛选节点的示例 @yqlbu in https://github.com/daeuniverse/dae/pull/44
- 文档（example.dae）：新增一个 tcp 健康检测 url by @yqlbu in https://github.com/daeuniverse/dae/pull/46

**Full Changelog**: https://github.com/daeuniverse/dae/compare/v0.1.4...v0.1.5

### 0.1.4

> Release date: 2023/03/25

#### 更新内容

- domain routing 给出不标准的域名时将忽略而不是报错。
- 将 config 所在目录加入到 geodata 的搜索路径。
- 优化 udp 的内存占用。
- 忽略 sighup 而使用 sigusr2 作为 suspend 的信号。
- 支持自动配置 sysctl 参数。
- 文档: 更新 debian-kernel-upgrade by @yqlbu in https://github.com/daeuniverse/dae/pull/39

**Full Changelog**: https://github.com/daeuniverse/dae/compare/v0.1.3...v0.1.4

### 0.1.3

> Release date: 2023/03/24

#### 用户相关

- 新增 amd64_v2_sse 和 amd64_v3_avx 的可执行文件构建，使用更高的版本理论上可提高一定性能（这次 Release 的 CI 失败了，等下次吧） by @MarksonHon in https://github.com/daeuniverse/dae/pull/38
- 支持自动侦测 WAN 接口，在 wan_interface 填入 auto 即可。
- 修复热重载失败时的不正确的回滚行为，以及在一定条件下更改 group 配置时可能无法连接新组的问题。
- 修复在有 MAC 地址路由的情况下 bind to WAN 将导致无网络的问题。
- 修改启动时网络联通性检查使用的链接 https://github.com/daeuniverse/dae/commit/c2e02482d0588823d2a3d9cae6998b9a7a5a1fae 。
- 修复在一定条件下可能的针对 DNS upstream 的域名分流失败的问题。

#### 开发者相关

- 打包了包括 go vendor 和 git submodules 在内的源码并随 releases 发布。
- 增加了 export 命令的描述。

**Full Changelog**: https://github.com/daeuniverse/dae/compare/v0.1.2...v0.1.3

### 0.1.2

> Release date: 2023/03/22

1. 优化热重载时的 DNS 缓存行为，解决热重载时 outbound out of range 的问题。
2. 增加高通的 generate_204 作为网络联通性检查的链接，以解决部分用户无法访问`www.msftconnecttest.com`的问题。
3. 支持龙芯 loong64 架构。
4. 修复大并发下可能的崩溃问题。

**Full Changelog**: https://github.com/daeuniverse/dae/compare/v0.1.1...v0.1.2

### 0.1.1

> Release date: 2023/03/16

#### What's Changed

- feat: shorten docker command arguments by leveraging CMD by @kunish in https://github.com/daeuniverse/dae/pull/35

#### New Contributors

- @kunish made their first contribution in https://github.com/daeuniverse/dae/pull/35

**Full Changelog**: https://github.com/daeuniverse/dae/compare/v0.1.0...v0.1.1

### 0.1.0

> Release date: 2023/03/14

Goose out of shell.
