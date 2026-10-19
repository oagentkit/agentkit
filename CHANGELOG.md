# agentkit CHANGELOG

> 本文件由 `matrix-tools/scripts/gen_changelogs.py` 从 git 历史生成，按 tag 切段、取每段内的 conventional commit 渲染。

## [v1.8.4] — 2026-10-19

- 杂务(deps)：依赖地板对齐远端最新（agent v0.10.0 / guardx v0.5.0 / llm v0.11.0 / mcpkit v0.8.0 / obs v0.7.0 / promptx v0.8.0 / queue v0.6.0 / rag v0.8.0 / runtime v0.7.0 / skillkit v0.8.0 / textx v0.5.0 / toolgate v0.11.0 / tools v0.9.0） (13)

## [v1.8.3] — 2026-10-19

### 修复(facade)：补 invoke / injectband 两个 re-export + 抬依赖地板到各仓最新 tag

两处独立缺陷，都属于「门面没跟上矩阵」。
## 1. 缺两个 re-export 包
矩阵已 47 个包，门面只有 45 个——缺 `toolgate/invoke` 与 `guardx/injectband`。
后果不是编译失败，而是**走门面的下游根本够不到这两个包**：
`go get github.com/oagentkit/agentkit/invoke` 直接 404。
按既有 alias.go 模式补齐（转发面与上游导出符号逐一对齐，go doc 交叉核过）：
- invoke：11 个类型 + 4 个函数 + 4 个哨兵错误
- injectband：2 个函数
上游的 `stagePrecondition` / `stageInternal` / `stageNil` 是未导出常量，不在门面范围。
README 的「45 个」与 MIGRATION.md 的映射表同步更新为 47 并补两行。
## 2. 依赖地板落后
agent v0.8.0 → v0.9.1、toolgate v0.9.0 → v0.10.0，
oekit/concurrency v0.17.0 → v0.17.1、oekit/logx v0.7.0 → v0.7.1、
oekit/pkg v0.14.0 → v0.15.1。
门面是消费方的单一入口，而 MVS 只看被消费版本的 go.mod：门面 tag 钉着旧版本，
消费方 go get 门面仓会静默拿到旧实现——不报错、不告警、go build 照样过。
**只改 main 不发 tag 等于没修**，所以本次必须重新发 tag。
本地门禁：go mod tidy -diff 无输出 / go build / go vet / gofmt / go test -race 全通过。
另用独立消费方模块实跑，确认 invoke 与 injectband 的别名可被下游真正 import 并调用。

## [v1.8.2] — 2026-10-19

- 杂务(deps)：依赖地板跟进 skillkit v0.6.1 / tools v0.7.1

## [v1.8.1] — 2026-10-19

- 杂务(deps)：guardx v0.4.0 → v0.4.1（egress 补 0.0.0.0/8 拦截）

## [v1.8.0] — 2026-10-19

### 重构(agentkit)!：门面地板跟到 llm v0.9.0 / rag v0.6.0 / agent v0.8.0

BREAKING CHANGE: 上游三个仓本轮为破坏性发布（包上移至 module 根），门面的
re-export 必须指向新版本，否则消费方从门面仓装走的是旧实现。
- llm/alias.go、rag/alias.go 的 upstream import 指向新路径
- go.mod 地板：agent v0.7.0→v0.8.0、llm v0.8.0→v0.9.0、rag v0.5.0→v0.6.0
- MIGRATION.md / README.md 的映射表同步
verify_facade_release.py 守的是 tag 不是 main，本次随 tag 一并发布。

## [v1.7.0] — 2026-10-19

### 修复(deps)：模块仓 tag 移走后重算 go.sum

agent / llm / toolgate 三个 tag 已移到修复 go.sum 的提交上（原先指向的提交里，
对 llm@v0.7.0、obs@v0.4.0、mcpkit@v0.5.0 的 h1 是 tag 重写之前算出来的，
消费方 go get 会直接撞 SECURITY ERROR）。
tag 一动，依赖它的仓记的哈希就失效——本仓 go.sum 记的还是那几个 tag 移动**之前**
的内容，于是门面仓自己也开始构建失败。这是「同名 tag 对应新内容」这条规则的下游
成本：tag 不是只对它自己负责的东西。
仍用空 GOMODCACHE + GOPRIVATE 直连 VCS 重算：复用共享模块缓存会把陈旧副本原样
记回去（命令成功、依据是错的）；CI 的 setup-go 缓存也会把重写前的副本回退恢复。
go.mod 地板不变（版本号没动），verify_facade_release 应仍为通过。
验证：go build ./... 通过、go test ./... 0 失败、go mod tidy -diff 无差异。

- 杂务(deps)：依赖地板跟进到本轮发布的各模块 tag
- 杂务(deps)：oekit 依赖升级 + 迁出被强推的 tag

## [v1.6.0] — 2026-10-19

- 杂务(deps)：依赖地板跟进到 8 个模块仓的新 tag

## [v1.5.0] — 2026-10-19

- 杂务(deps)：toolgate 地板跟进到 v0.8.0，并重算已失效的 go.sum

## [v1.4.0] — 2026-10-18

### 新增：接 org 共享 CI + 依赖对齐当前 tag

本仓此前没有 CI —— 46 个 re-export 包坏掉不会有人知道。本提交接上
oagentkit/.github 的共享门禁（gofmt / vet / golangci-lint / test -race /
govulncheck / tidy -diff），与 13 个模块仓同一套规则。
覆盖率门关闭：本仓是纯 re-export，没有任何自研实现，覆盖率对它无意义。
同时把 13 个依赖对齐到当前 tag（llm v0.4.0 / runtime v0.3.1 /
其余 v0.2.1）。breaker 门面已随上一提交移除（runtime v0.3.0 删除了
该包，迁移说明见 MIGRATION.md）。

### 新增：agentkit 门面仓 v1.0.0——46 包 re-export 兼容层

原单仓 git.enjoye.top/enjoydream/agentkit 已拆分为 github.com/oagentkit 下 13 个独立发版仓。
本仓仅做别名转发，让下游只需改 import 前缀、无需改调用点；v1.x 后退役。
内容：
- 46 个门面包，595 个别名（类型别名 + 函数/变量别名，泛型类型保留类型参数）
- MIGRATION.md 完整 46 条旧→新包路径映射表
- 3 个泛型函数（hotplug.NewHolder / progress.NewBus / reportutil.Aggregate）
  无法用变量别名转发，已在生成物与 MIGRATION.md 中标注需直接引用上游
生成方式：go/parser 读各模块真实源码生成 alias.go，非手写（tools/gen）。
门禁：go build / go vet 通过，gofmt 干净。

### 修复(docs)：re-export 包数 46 → 45（breaker 退役后门面仓不再转发它） (#1)

`README.md` 写「本仓只把 46 个包的导出符号做成别名转发」，
两处数字在 `runtime/breaker` 退役后已经不成立。
四源交叉核实（本仓 .go 包目录数 / .go 里的 import 目标数 /
modules.json 登记总包数 / MIGRATION.md 映射表行数）：
| 来源 | 数量 |
| --- | --- |
| modules.json 登记总包数 | 45 |
| 门面仓有 .go 的包目录 | 45 |
| 门面仓 .go 里的 import 目标 | 45 |
| MIGRATION.md 映射表行 | 46（含 1 条已退役的 breaker 映射） |
所以真实数字是 **45 个活跃包 + 1 条退役映射 = 46 行**。
原表述「46 个包」把退役那条也算成了活跃转发包。
改为「45 个包」，并单独点出 breaker 已退役、MIGRATION.md 里另有说明。

### 重构：移除已退役的 breaker 门面

runtime v0.3.0 删除了本包（三态熔断器已迁往
oekit/concurrency/breaker，由 llm 仓消费）。
迁移方改用 github.com/oekit/concurrency/breaker ——
注意语义：上游是「速率型 or 计数型并联触发」，
不是本包原来的纯计数型三态。

- 杂务(deps)：依赖地板升到「历史最新 + 0.1.0」 (#4)
- 杂务(deps)：跟进最新 tag (#3)
- 杂务(deps)：oekit 依赖升级 (#2)

- 文档：去掉迁移示例里的字面量 ... 占位与不可用的通配 sed

- 回滚：go.mod/go.sum 回退到 v1.0.0 的自洽版本
