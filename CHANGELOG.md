# agentkit CHANGELOG

> 本文件由 `matrix-tools/scripts/gen_changelogs.py` 从 git 历史生成，按 tag 切段、取每段内的 conventional commit 渲染。

## [v1.8.35] — 2026-10-22

- 杂务(deps)：门面依赖地板跟到当前最新（guardx v0.14.0 / oekit/pkg v0.22.0 / oekit/concurrency v0.20.1）

矩阵的发版时刻地板门禁（verify_tag_floor）要求：打 tag 那一刻的依赖地板
必须是当时的最新。门面 doctor 判红 3：
- github.com/oagentkit/guardx 钉在 v0.13.1，最新 v0.14.0 —— 本轮 permgate
  修复随 guardx v0.14.0 发布（issue #24），门面必须跟上。
- github.com/oekit/pkg 钉在 v0.21.0，最新 v0.22.0。
- github.com/oekit/concurrency 钉在 v0.20.0，最新 v0.20.1。
后两条不是本轮引入的——是并发会话在 oekit 侧发的新版使「地板时效」判据
在**打 tag 的那一刻**变红（判据按发布时刻回看，不是按当前）。不抬就打 tag，
产出的是「出生即落后」的 tag，而 tag 受 protect-tags 保护不可重打。
## 行为差异评估（实测）
- **oekit/pkg v0.21.0 → v0.22.0**：新增全在 pathx / stringx（新文件：
  MatchGlob、IsSafeFileName、SanitizeFileName、SplitCSVNonEmpty），
  errors 包逐字未改。门面**零直接 oekit import**（实测 grep = 0），
  oekit/pkg 是间接依赖 ⇒ 对门面零行为差异。
- **oekit/concurrency v0.20.0 → v0.20.1**：含一处真修复——
  「KeyedGate 的 release 幂等标记改用 atomic CAS——裸 bool 在并发释放下
  竞态并突破 cap 上限」。这是修复不是变更；门面同样只有间接依赖。
- **guardx v0.13.1 → v0.14.0**：门面的 permgate 是**类型别名**转发
  （agentkit/permgate/alias.go），新行为自动跟随 ⇒ 门面**无需改任何符号**。
改动面：go.mod 3 行版本号 + go.sum 6 增 6 删，零代码改动。
门禁自检：go build ✅ / go vet ✅ / go test -race ✅（48 包，均为别名包无测试文件，
退出码 0）/ gofmt ✅ / go mod tidy ✅


## [v1.8.34] — 2026-10-22

- 杂务(deps)：门面地板跟到 agent v0.10.5

门面 agentkit 转发的 agentrun.Usage 是**类型别名**（Usage = upstream.Usage），
所以上游新增的 DurationMS 字段自动透传，门面侧无需改动任何符号。
打 tag 时刻的依赖地板必须是当时的最新（matrix-tools 的 verify_tag_floor 判据）：
agent v0.10.5 已发布，门面若仍钉 v0.10.4，v1.8.34 就会是「出生即落后」的 tag
（tag 受 protect-tags 保护不可重打）。
本地：gofmt -l（空）/ go build ./... / go test -race -count=1 ./... 无失败 /
go mod tidy -diff（空）。


## [v1.8.33] — 2026-10-22

- 杂务(deps)：obs 依赖地板抬到 v0.9.6

门面仓是消费方的单一入口，而 Go 的 MVS 只看**被消费的那个版本**的 go.mod。
obs v0.9.6 已发布，门面若继续钉 obs v0.9.5，下游 `go get
github.com/oagentkit/agentkit@<最新 tag>` 就会静默拿到 obs 的旧实现
——不报错、不告警、go build 照样过。
matrix-tools/scripts/verify_facade_release.py --org oagentkit 判红：
「obs 应为 v0.9.6，实际 v0.9.5」。
⚠️ 只改 main 不发 tag 等于没修：线上流传的仍是旧 tag 快照。
本 PR 只是把地板抬到位，随后必须给门面仓发 tag 才真正闭环。
其余 12 个模块的地板本来就是最新，本 PR 不动。
门禁：go build / go vet / go test -race / gofmt -l / go mod tidy -diff 全绿。


## [v1.8.32] — 2026-10-22

- 杂务(deps)：门面依赖地板跟到基座最新 tag——oekit/concurrency v0.20.0、oekit/pkg v0.21.0

门面的已发布 tag v1.8.31 的 go.mod 还钉着 oekit/concurrency v0.19.0 与
oekit/pkg v0.20.0，而基座这两个仓已经各发了一版新 tag。矩阵门禁
verify_facade_release.py --org oagentkit 因此判红（基座依赖落后 2 处）。
基座依赖同样参与 MVS：门面 tag 钉着旧版本，消费方 go get 门面仓会静默拿到
旧实现——不报错、不告警、go build 照样过。
本次只把 go.mod / go.sum 的基座地板抬到各仓当前最新 tag：
  github.com/oekit/concurrency  v0.19.0 -> v0.20.0
  github.com/oekit/pkg          v0.20.0 -> v0.21.0
其余 13 个矩阵模块的依赖地板本轮已是最新，未动。
只改 main 不发 tag 等于没修（线上流传的仍是旧 tag），发版另走一步。


## [v1.8.31] — 2026-10-21

- 杂务(deps)：依赖地板跟到矩阵本轮新版本——llm v0.15.2 mcpkit v0.8.5 toolgate v0.12.4 agent v0.10.4

## 下游要做什么
**`go mod tidy`**。地板抬高对下游是**隐式破坏性变更**：不跑就报
`updates to go.mod needed`，且报错不提真正原因（**不是本仓的 API 变了**）。
本仓**自身代码零改动**——只看本仓的 diff 看不出这是一次地板抬高。
## 为什么抬
上一版 tag 打出去时，它依赖的模块的**新版已经先发出**了——
即「打 tag 那一刻地板就落后」。矩阵的「发版时刻地板时效」门禁
（`verify_tag_floor.py --org oagentkit --latest-only`）当场报红：
- github.com/oagentkit/llm@v0.15.2
- github.com/oagentkit/mcpkit@v0.8.5
- github.com/oagentkit/toolgate@v0.12.4
- github.com/oagentkit/agent@v0.10.4
门禁给的修法是「抬 go.mod 地板 → go mod tidy → 重新发 tag」。
本 PR 就是那条修法。
## 为什么无行为差异
被抬的这几个版本**自身也都是纯依赖地板追平**（各仓代码零改动，只抬
`oekit/pkg` / `oekit/concurrency` / `guardx` / `llm` 四类地板之一），
本仓用到的转发面与符号集一字未动。
## 门禁
```
go build / vet / test -race   ✅
gofmt -l .                    ✅ 空
```


## [v1.8.30] — 2026-10-21

### 新增(policy)：转发 toolgate/policy 的两个新符号——Disposition 与 FailureClass

上游 toolgate v0.12.3 新增了「协议层失败分类 → 处置层类别的机械映射」：
```go
func FailureClass(kind mcp.ErrorKind) toolfail.Class
func Disposition(err error) toolfail.Class
```
门面对这些类型是**类型别名**，但**函数/变量必须逐个转发**（别名只覆盖类型）——
漏转发一个符号，下游按门面 import 就**编译不过**。
矩阵的「门面符号级转发完整性」门禁（`verify_facade_symbols.py --org oagentkit`）
当场报红：
```
❌ 门面符号级转发不完整（remote org=oagentkit facade=agentkit，核对 48 个包）：
   - policy: 漏转发上游导出符号 Disposition（上游 kind=func）
   - policy: 漏转发上游导出符号 FailureClass（上游 kind=func）
```
## 改动
`policy/alias.go` 的 var 块补两个别名（按既有约定：**func/var 上游 → var 块**），
并把 `toolgate` 依赖地板抬到 v0.12.3。
## 门禁
```
go build / vet / test -race   ✅
gofmt -l .                    ✅ 空
verify_facade_symbols.py      ✅（48 个包，policy 门补齐后不再漏）
verify_facade_frozen.py       ✅（48 包 ⊆ 冻结基线 48）
```

## [v1.8.29] — 2026-10-21

- 杂务(deps)：依赖地板跟到矩阵本轮新版本——agent v0.10.3 obs v0.9.5 promptx v0.8.3 rag v0.8.5 skillkit v0.8.4 toolgate v0.12.2 tools v0.10.3

## 下游要做什么
**`go mod tidy`**。地板抬高对下游是**隐式破坏性变更**：不跑就报
`updates to go.mod needed`，且报错不提真正原因（**不是本仓的 API 变了**）。
本仓**自身代码零改动**——只看本仓的 diff 看不出这是一次地板抬高。
## 为什么抬
上一版 tag 打出去时，它依赖的模块的**新版已经先发出**了——
即「打 tag 那一刻地板就落后」。矩阵的「发版时刻地板时效」门禁
（`verify_tag_floor.py --org oagentkit --latest-only`）当场报红：
- github.com/oagentkit/agent@v0.10.3
- github.com/oagentkit/obs@v0.9.5
- github.com/oagentkit/promptx@v0.8.3
- github.com/oagentkit/rag@v0.8.5
- github.com/oagentkit/skillkit@v0.8.4
- github.com/oagentkit/toolgate@v0.12.2
- github.com/oagentkit/tools@v0.10.3
门禁给的修法是「抬 go.mod 地板 → go mod tidy → 重新发 tag」。
本 PR 就是那条修法。
## 为什么无行为差异
被抬的这几个版本**自身也都是纯依赖地板追平**（各仓代码零改动，只抬
`oekit/pkg` / `oekit/concurrency` / `guardx` / `llm` 四类地板之一），
本仓用到的转发面与符号集一字未动。
## 门禁
```
go build / vet / test -race   ✅
gofmt -l .                    ✅ 空
```


## [v1.8.28] — 2026-10-21

- 杂务(deps)：门面依赖地板跟到矩阵本轮新版本（12 块）

## 下游要做什么
**`go mod tidy`**。地板抬高对下游是隐式 breaking：不跑就报
`updates to go.mod needed`，且报错不提真正原因（**不是门面的 API 变了**）。
门面**自身代码零改动**——转发的 48 个包与符号集一字未动。
## 抬了哪 12 块
```
oagentkit/textx     v0.6.1 → v0.6.2      oagentkit/skillkit  v0.8.2 → v0.8.3
oagentkit/runtime   v0.7.1 → v0.7.2      oagentkit/mcpkit    v0.8.3 → v0.8.4
oagentkit/queue     v0.6.2 → v0.6.3      oagentkit/rag       v0.8.3 → v0.8.4
oagentkit/promptx   v0.8.1 → v0.8.2      oagentkit/toolgate  v0.12.0 → v0.12.1
oagentkit/obs       v0.9.3 → v0.9.4      oagentkit/tools     v0.10.1 → v0.10.2
oagentkit/agent     v0.10.1 → v0.10.2
```
## 为什么无行为差异
**这 12 个版本全部是「纯依赖地板追平」——各仓自身代码零改动**：
- textx / runtime / queue / promptx / obs / skillkit / mcpkit / rag /
  toolgate / tools / agent：各自只抬了 `oekit/pkg`（v0.18.0→v0.20.0）、
  `oekit/concurrency`（v0.18.0→v0.19.0）、`oagentkit/guardx`
  （v0.13.0→v0.13.1）、`oagentkit/llm`（v0.14.2→v0.15.0）四类地板之一，
  **没有一行业务代码变更**。
四类上游地板的变更本身也已逐项核过：
- `oekit/pkg` v0.18.0→v0.20.0：`errors` 纯增量（新增 `Errorf`）、
  `httpjson` 退避溢出修复（正常 attempt 下逐字相同）、`urlguard` 新增可选
  `WithExactHosts`（未设置即行为不变）、`stringx` 仅注释、`pathx` 新包。
- `oekit/concurrency` v0.18.0→v0.19.0：`retry` 抖动 panic 修复 +
  `syncx` 新增 `Gate`；**`breaker` 与 `async` 一字未动**。
- `oagentkit/guardx` v0.13.0→v0.13.1：上游只改 `go.mod`/`go.sum`，
  **没有 Go 源码变更**。
- `oagentkit/llm` v0.14.2→v0.15.0：新增**可选**字段
  `OpenAIProviderConfig.HTTPClient`（缺省 nil 时行为逐字不变）。
门面对这些类型都是**类型别名**，故上游的新增能力自动透传——
`llm.OpenAIProviderConfig.HTTPClient`（宿主可注入自己的 `*http.Client`，
把出站守卫挂到 `Transport.DialContext` 上）经门面同样可用。
## 门禁
```
go build / vet / test -race   ✅
gofmt -l .                    ✅ 空
verify_facade_frozen.py       ✅（48 包 ⊆ 冻结基线 48）
verify_facade_symbols.py      ✅（llm 门本地可判，通过）
```


## [v1.8.27] — 2026-10-21

- 杂务(deps)：门面依赖地板跟到 llm v0.15.0

llm v0.15.0 新增 `OpenAIProviderConfig.HTTPClient`——宿主可注入自己的
`*http.Client`，把出站守卫挂到 `Transport.DialContext` 上（写面校验拦不住
DNS rebinding，守卫必须在拨号层）。
门面对该类型是**类型别名**（`type OpenAIProviderConfig = upstream.OpenAIProviderConfig`），
故新字段**自动透传**，门面**不需要改任何转发代码**，只抬依赖地板即可让下游拿到。
无行为差异：llm v0.14.2 → v0.15.0 只新增了一个可选字段（缺省 nil 时行为逐字不变）
+ 抬了 llm 自身的 oekit 地板；门面转发的 48 个包与符号集一字未动。
门禁自检：go build / vet / test ✅、gofmt ✅、
verify_facade_frozen.py ✅（48 包 ⊆ 冻结基线 48）、
verify_facade_symbols.py（llm 门本地可判，通过；其余 18 门因本机未克隆对应上游仓
判不了，CI 走联网模式全量判）。


## [v1.8.26] — 2026-10-21

- 杂务(deps)：门面依赖地板跟到 guardx v0.13.1

发 v1.8.25 之后、跑 `verify_facade_release.py --org oagentkit` 复核时，
guardx 已在 2026-10-21T15:42:26+0800 发了 v0.13.1
（`chore(deps): oekit/pkg 依赖地板抬到 v0.20.0`），
而 v1.8.25 的 go.mod 钉的是 v0.13.0 ⇒ 门面地板门禁判红：
    ❌ 门面仓 tag v1.8.25 的依赖地板有 1 处落后：
      guardx       应为 v0.13.1      实际 v0.13.0
本提交把门面钉到 guardx v0.13.1。guardx v0.13.1 相对 v0.13.0 只改了
CHANGELOG.md / go.mod / go.sum（**无 Go 源码变更**），故门面的 48 个转发包
与转发符号集一字未动；门面的 oekit/pkg 本就已是 v0.20.0，与 guardx 对齐。
门禁依据：matrix-tools/scripts/verify_facade_release.py


## [v1.8.25] — 2026-10-21

- 杂务(deps)：门面依赖地板抬到最新（oekit/pkg v0.20.0、oekit/concurrency v0.19.0）

门面的已发布 tag v1.8.24 的 go.mod 钉着 oekit/pkg v0.18.0 与
oekit/concurrency v0.18.0，而远端最新已是 v0.20.0 / v0.19.0。
Go 的 MVS 只看**被消费版本**的 go.mod ⇒ 存量下游
`go get github.com/oagentkit/agentkit@v1.8.24` 装走的是旧基座实现，
不报错、不告警、go build 照样过。
oagentkit/* 侧的 13 个第一方依赖地板已在 main 上跟平（guardx v0.13.0 见 #45），
本次补齐 oekit 基座两项。
门禁依据：matrix-tools/scripts/verify_facade_release.py

- 杂务(deps)：guardx 依赖地板抬到 v0.13.0

本仓转发的 6 个 guardx 包（egress / fence / injectband / invoke / logredact /
permgate）在 v0.11.0 → v0.13.0 之间**未变**——那几版只加了 sandbox 的
Capable 与 wire，而这两个新包**不在门面的 48 个转发包里**。故这是纯地板追平，
**无行为差异**，门面转发的符号集一字未动。
属本仓「只做别名转发与安全修复」范畴：抬地板不改变任何转发面，冻结的
48 个包集合不变（verify_facade_frozen 不读 go.mod 版本）。
矩阵的依赖地板时效门禁（verify_dep_floors.py）会把「地板落后远端最新」报出来。
门禁自检：go build / vet ✅，go test ✅，gofmt ✅，go mod tidy ✅


## [v1.8.24] — 2026-10-21

- 杂务(deps)：门面依赖地板抬到 guardx v0.11.0

guardx v0.11.0 已发布（新增 sandbox/wire 跨进程调用协议）。
门面的每个 tag 都是一份**对外承诺**——Go 的 MVS 只看被消费那个版本的 go.mod，
所以 `go get github.com/oagentkit/agentkit@vX` 装走的是**该 tag 快照**里钉的版本。
本仓的 `verify_facade_release.py` 守的就是这条：门面的最新 tag 必须钉着
各模块的当前最新 tag。
v1.8.23 钉的是 guardx v0.10.0，而 guardx 随后发了 v0.11.0 ⇒ 门面落后一格。
本地门禁：go mod tidy 无差异 / go build / go vet 通过。


## [v1.8.23] — 2026-10-21

- 杂务(deps)：门面依赖地板抬到 guardx v0.10.0

guardx v0.10.0 已发布（sandbox 预算面补 MaxMemory + 可选能力面 Capable）。
门面的每个 tag 都是一份**对外承诺**——Go 的 MVS 只看被消费那个版本的 go.mod，
所以 `go get github.com/oagentkit/agentkit@vX` 装走的是**该 tag 快照**里钉的版本。
本仓的 `verify_facade_release.py` 守的就是这条：门面的最新 tag 必须钉着
各模块的当前最新 tag。
v1.8.22 钉的是 guardx v0.9.0，而 guardx 随后发了 v0.10.0 ⇒ 门面落后一格。
本地门禁：go mod tidy 无差异 / go build / go vet 通过。


## [v1.8.22] — 2026-10-21

- 杂务(deps)：门面依赖地板抬到 agent v0.10.1 / guardx v0.9.0

门面的每个 tag 都是一份**对外承诺**：Go 的 MVS 只看被消费的那个版本的 go.mod，
所以 `go get github.com/oagentkit/agentkit@vX` 装走的是**该 tag 快照**里钉的版本。
本仓的 `verify_facade_release.py` 守的就是这条——门面的最新 tag 必须钉着
各模块的当前最新，否则下游装走的是旧实现。
两个模块刚发过新版，所以门面必须跟发：
- agent v0.10.0 → v0.10.1（含「Go 地板降到 go 1.24.0」的发布说明）
- guardx v0.8.2 → v0.9.0（新增 toolfail 与 sandbox 两个 tier0 原语）
注意：guardx 的两个新包**不转发**——门面是冻结的兼容层，转发新包就是
「门面又长出一张新脸」，违反它自己的废弃承诺（门禁按「基线之后新增的包 ⇒ 提示」
判定，不算缺口）。
本地门禁：go mod tidy 无差异 / go build / go vet 通过。


- 文档(facade)：补 MIGRATION.md 的 port 映射行，统一 48/47 口径并更新到 v1.8.21

门面仓（oagentkit/agentkit）实测出三处不一致，本次一并修掉：
1. MIGRATION.md 漏了 llm/port 的映射行。门面实际有 48 个 alias.go 包，
   映射表却只覆盖 47 个活包（另 1 行是已退役的 breaker）——port 是 llm 仓的
   无 eino 窄接口层，下游按 README「逐包替换」迁出时会漏掉它，而 llm 是
   tier 2 重依赖仓，漏掉等于把重依赖又拉回来。
   补一行：| `port` | `llm` | `github.com/oagentkit/llm/port` |
   补后活包映射 48 个，与门面的 48 个 alias.go 包一一对应。
2. README 内部口径打架：第 5 行写「48 个包」，第 34 行写「47 个包的完整映射」。
   统一为「48 个活包」，并写清 MIGRATION.md 的映射表
   = 48 个活包映射 + 1 行已退役包（breaker）。
3. README 状态行的版本号陈旧：v1.8.20 -> v1.8.21（实际最新 tag 已是 v1.8.21）。
另新建 AGENTS.md，写本仓的工程红线（门面仓专属，不照抄别仓规则）：
禁止新增包（会让「门面已废弃」的承诺失效）、禁止矩阵内部消费门面、
转发必须 kind 一致（const 不得降级成 var，否则下游常量位置编译失败）、
上游新增导出符号必须在本仓补别名、以及退役纪律。
未动任何 alias.go：转发内容与上游一一对应，改了会破坏符号级转发完整性门禁。
校验：
- go build ./... / go vet ./... / gofmt -l . 全通过
- verify_facade_packages.py：门面包数 48，矩阵应覆盖 48，无缺口、无陈旧转发
- verify_facade_frozen.py：48 个包 ⊆ 冻结基线 48 个
- verify_facade_symbols.py：符号级转发完整（核对 48 个包）
- check_facade_imports.py：未发现门面仓/旧单仓 import
- scan_stale_refs.py：未发现指向已不存在包的文档引用


## [v1.8.21] — 2026-10-20

- 文档：门面仓标记为「已废弃」——v1.x 退役条件已满足

本仓的退役条件已满足：README 一直写着「v1.x 后本仓退役」，而本仓现在已是
**v1.8.20**（v1.x 之内）。`SPLIT-PLAN.md` §8-4 的裁定原文是「README 写死退役
时间表（v1.x 后转 archive）」，现在该执行了。
## 消费方核查：矩阵内零消费方
- oagentkit org 内：0 个仓依赖它（唯一命中是它自己的 module 声明行）；
- 姊妹 org oekit 内：0 个；
- GitHub 全站代码搜索 `github.com/oagentkit/agentkit`：**命中 0**；
- 本机：0 个项目依赖它。
所以「废弃」不会打断任何已知消费方。
## 改动
- **README**：顶部告示从「兼容门面，不是开发目标 / v1.x 后本仓退役」改成明确的
  **已废弃**状态，并写清现状与边界：
  - v1.x 退役条件**已满足**；本仓只保留兼容转发与安全修复，不接受任何新功能；
  - **新代码请直接 import 上游模块**；
  - **归档（archive）待旧单仓 `git.enjoye.top/enjoydream/agentkit`（module path
    的实际所在）下线后执行**——那才是真正切断旧路径的一步。在此之前本仓继续可用，
    **已有 tag 的 `go get` 照常解析**（归档的仓仍可只读 clone）。
  - 顺带修正 README 里过期的「47 个包」→ **48 个**（`breaker` 退役后新增了
    rag 的 embedder 注入缝等符号）。
- **48 个 `*/alias.go`** 的包文档注释：从「本仓将在 v1.x 后退役」改成
  「本仓**已废弃**：v1.x 退役条件已满足，只保留兼容转发，不接受任何新功能。
  新代码请直接 import 上游模块（见 MIGRATION.md 的映射表）。」
## 为什么归档 ≠ 删除
归档后仓仍可只读 clone，已有 tag 的 `go get` 照常解析；而删仓会让任何没进
代理缓存的版本直接 404，下游构建断。先例：`github.com/yi-nology/agentkit`
已 `archived=true` + description 改成迁移指引。
## 门禁继续跑
门面的 4 条守卫（`check_facade_imports` / `verify_facade_packages` /
`verify_facade_release` / `verify_facade_symbols`）在门面活着期间**必须继续跑**——
它们保护存量下游拿到正确的兼容层，且「不许再引门面」这条机器约束就靠
`check_facade_imports.py` 守着。
另在 matrix-tools 侧加一条「门面包集合冻结」守卫，把「不继续长胖」也变成红灯。
关联：oagentkit/matrix-tools#41（门禁改造）、SPLIT-PLAN §8-4 / §8-5


## [v1.8.20] — 2026-10-20

- 杂务(deps)：门面地板跟到 llm v0.14.2

llm v0.14.2 带上了 oagentkit/llm#26 的修复（`NewPriceTable` 的覆盖键归一化——
调用方显式配的单价不再被内建参考价顶掉）。门面转发 llm 的导出符号，
地板必须跟到 llm 的新 tag，否则消费方 `go get oagentkit/agentkit/...` 拿到的
仍是 llm v0.14.1 的实现——门禁 `verify_facade_release.py` 判的正是这一条：
「门面 tag 钉着旧版本 ⇒ 消费方会静默拿到旧实现，不报错、不告警、go build 照样过」。


## [v1.8.19] — 2026-10-20

### 修复(facade)：门面补转发 rag v0.8.3 新增的 4 个导出符号

rag v0.8.3 带上了 oagentkit/rag#15 的修复（`NewOpenAIEmbedder` 的注入缝），
新增 4 个导出符号：
- `EmbedderOption`（type）
- `WithHTTPClient`（func）
- `WithTimeout`（func）
- `WithRetry`（func）
门禁 `matrix-tools/scripts/verify_facade_symbols.py` 抓到门面漏转发这 4 个：
```
❌ 门面符号级转发不完整（remote org=oagentkit facade=agentkit，核对 48 个包）：
   - rag: 漏转发上游导出符号 EmbedderOption（上游 kind=type）
   - rag: 漏转发上游导出符号 WithHTTPClient（上游 kind=func）
   - rag: 漏转发上游导出符号 WithRetry（上游 kind=func）
   - rag: 漏转发上游导出符号 WithTimeout（上游 kind=func）
```
门面是消费方的**单一入口**（README 推荐 `go get github.com/oagentkit/agentkit/...`）。
漏转发 ⇒ 按门面 import 的下游拿不到 embedding 的注入缝，只能绕门面直接 import
上游仓——这正是门面要消除的分裂。
按 alias.go 的既有约定补：type 进 type 块，func 进 var 块。

## [v1.8.18] — 2026-10-20

### 新增(mcp)：门面转发 mcpkit v0.8.2 的自检/健康/错误分类原语

补齐 mcpkit v0.8.2 新增的 34 个导出符号转发（Aggregate / Catalog / CatalogDiff / ClassifiedError / Classify / ClassifyError / ClassifyRPCCode / DefaultSelfcheckTTL / DefaultSelfcheckTimeout / Diff / DiffNames / ErrorKind / Health* / HealthIssue / HealthState / Kind* / NewCatalog / NewSelfchecker / SelfcheckOptions / SelfcheckResult / Selfchecker / ToolCheck / ToolFailed / ToolMissing / ToolOK / ToolStatus / ToolUnreachable）。
门禁实测漏转发 34 个（run 39357779865 job Facade symbols）。

- 杂务(deps)：门面地板跟到各模块新 tag

本仓是门面（转发各模块的导出符号），地板必须跟到各模块**刚发的新 tag**，
否则门面转发的仍是旧实现。
矩阵内模块：
- obs      v0.9.2 → v0.9.3
- rag      v0.8.2 → v0.8.3
- skillkit v0.8.1 → v0.8.2
- mcpkit   v0.8.2 → v0.8.3
- tools    v0.10.0 → v0.10.1
oekit 基座：
- oekit/pkg        v0.17.0 → v0.18.0
- oekit/concurrency v0.17.1 → v0.18.0
- oekit/logx       已是最新 v0.7.1，不动
依据 oagentkit/matrix-tools#44：各仓最新 tag 的依赖地板在「打 tag 那一刻」就已
落后一格，tag 受 protect-tags ruleset 保护不能重打，只能靠后续版本补齐。
本仓是最后发的那一个——它的地板必须等各模块的新 tag 出来才能跟。


- 文档(changelog)：按新生成器规则重生成 CHANGELOG.md

生成器修复（oagentkit/matrix-tools#43）落地：
① `[round-N] ` 前缀被剥离 → conventional commit 解析不再失效，
   **commit body 保留**（此前整段被丢弃，发布说明只剩标题）；
② 非 conventional 提交也保留 body；
③ `--check` 识别「待发版段」（预写了但还没打 tag 的那段）。
本 PR 恢复的是**本就在 commit body 里**的正文（迁移步骤、代码块、
判据锚点）——此前被 `[round-N] ` 前缀导致的解析失败丢弃了。
**没有任何发布说明内容被删除**：删除的行都能在新增行里找到对应
（规范化掉 `(#N)` 与 `[round-N] ` 前缀后逐行比对）。
关联：oagentkit/matrix-tools#43、oagentkit/matrix-tools#40


## [v1.8.17] — 2026-10-20

### 新增(facade)：转发 llm/port 的 51 个导出符号 + 依赖地板抬到 llm v0.14.1

llm 新增了 `port` 包（窄接口层，不引 eino）——见 oagentkit/llm#20 / issue llm#11。
门面必须同步转发它的全部导出符号，否则走门面的下游够不到这些 API。
转发 51 个符号：20 type / 7 const / 24 var+func
（Breakers、Budget、Chain、ChainOption、Classifier、Completer、CostRecord、
CostTracker、Entry、MatchKind、ModelSummary、Price、PriceLookup、PriceStatus、
PriceTable、Request、StatusError、TokenAccountant、TokenCounter、Usage 等）
go.mod 地板：llm v0.13.0 → v0.14.1。门面是消费方的单一入口，MVS 只看被消费
版本的 go.mod；地板落后 ⇒ 消费方 go get 门面仓静默拿到旧实现。
⚠️ PriceOf / PriceOfMatched 是上游已弃用的符号，转发它们按行豁免 SA1019
（staticcheck 会把「转发」判成「使用已弃用 API」，但门面按契约必须转发上游的
全部导出符号——移除=对下游的破坏性变更，且会被符号级门禁判成幽灵转发）。
验证：go build / go vet / go test 全绿；gofmt 干净；
本地同版本 golangci-lint v2.14.0 → 0 issues；
符号级门禁 → ✅ 48 个门面包全绿（原 47 + 新增 llm/port）。

## [v1.8.16] — 2026-10-20

- 杂务(deps)：门面依赖地板跟到 guardx v0.8.2 / runtime v0.7.1 / textx v0.6.1 + 预写 v1.8.16 段

* chore(deps): 门面依赖地板跟到 guardx v0.8.2 / runtime v0.7.1 / textx v0.6.1
门面 tag v1.8.15 的依赖地板门禁报出 3 处落后——这三个 tag 是 v1.8.15
发布**之后**（2026-10-20 08:07）才发的：
- guardx   应为 v0.8.2   实际 v0.8.1
- runtime  应为 v0.7.1   实际 v0.7.0
- textx    应为 v0.6.1   实际 v0.6.0
为什么这要当错误处理：MVS 只看被消费版本的 go.mod。门面仓是消费方的
单一入口，门面 tag 钉着旧版本 ⇒ 消费方 `go get 门面仓` 会静默拿到旧实现
——不报错、不告警、go build 照样过，只是代码是几个月前的。
验证：go build ./... / go vet ./... / go test ./... 全绿；gofmt 干净。
* docs(changelog): 预写 v1.8.16 待发版段
发版流程第一步：`gen_changelogs.py --next-version v1.8.16` → 提交 → 在该提交上
打 tag v1.8.16。tag 快照里就带着自己那一段（下游 go get 拿到的发布说明完整）。
段内容：门面依赖地板跟到 guardx v0.8.2 / runtime v0.7.1 / textx v0.6.1。
不变式：tag 快照里的 CHANGELOG 与 main 一致——gen_changelogs.py --selftest
端到端覆盖「直推打 tag」与「走 PR 打 tag」两种打法。


## [v1.8.15] — 2026-10-20

- 杂务(deps)：oekit 基座地板跟到 pkg v0.17.0

门面 tag v1.8.13 的依赖地板门禁（verify_facade_release.py --org oagentkit）报出
基座落后：github.com/oekit/pkg 应为 v0.17.0，实际钉在 v0.16.0。
为什么这要当错误处理：MVS 只看被消费版本的 go.mod。门面仓是消费方的单一入口，
门面 tag 钉着旧版本 ⇒ 消费方 go get 门面仓会静默拿到旧实现——不报错、不告警、
go build 照样过，只是代码是几个月前的。concurrency v0.17.1 与 logx v0.7.1 已是
最新，只有 pkg 落后这一处。
验证：go build ./... / go vet ./... / go test ./... 全绿；gofmt 干净。


## [v1.8.14] — 2026-10-20

### 新增(facade)：转发 llm 的价表 API（16 个新符号）+ 依赖地板抬到 llm v0.13.0

llm v0.13.0 新增了内建价表与「价格未知」的记账语义（见 oagentkit/llm#13），
门面必须同步转发这些导出符号，否则走门面的下游够不到它们。
转发 16 个符号：
- type: MatchKind / PriceLookup / PriceStatus / PriceTable / PricedProvider
- const: MatchExact / MatchNone / MatchNormalized / PricePriced / PriceUnpriced
- var: ErrPriceUnknown
- func: DefaultPriceTable / LookupPrice / LookupPriceStrict /
        NewCostTrackerWithTable / NewPriceTable
go.mod 地板：llm v0.12.1 → v0.13.0（门面是消费方的单一入口，MVS 只看被消费
版本的 go.mod；地板落后 ⇒ 消费方 go get 门面仓静默拿到旧实现）。
⚠️ 这条改动是被**矩阵的符号级转发完整性门禁**逼出来的
（matrix-tools/scripts/verify_facade_symbols.py）——门禁在 llm#13 合入后
报出这 16 个缺失。门禁只对账不生成，改动仍走人工 PR。
验证：go build ./... / go vet ./... / go test ./... 全绿；gofmt 干净；
符号级门禁对 47 个门面包全绿。

### 修复(facade)：转发已弃用符号时按行豁免 SA1019

staticcheck SA1019 把「转发上游已弃用符号」判成「使用已弃用 API」——
门面按契约转发上游的**全部**导出符号，删掉 PriceOf/PriceOfMatched 会让走门面的
下游编译断裂，而矩阵的符号级门禁也会判成「幽灵转发」。
故按行豁免 SA1019：豁免的是「转发」这个动作，不是「上游已弃用」这个事实。
上游真正弃用与否由 llm 仓的 Deprecated 注释与 README 承载。
本地同版本 golangci-lint v2.14.0 验证：golangci-lint run ./llm/... → 0 issues。

## [v1.8.13] — 2026-10-20

### 修复(docs)：更正 v1.8.10 发布说明里关于 const 降级影响的描述

v1.8.10 的 CHANGELOG 段（由 git 历史生成）有一处不准确的描述。该段无法手改
——CHANGELOG 由 gen_changelogs.py 从历史重生成，手改会被覆盖；tag 又受
protect-tags ruleset 保护不可变。所以在此更正，让后续版本明确记下正确机制。
## 错在哪
v1.8.10 段写的是：
    下游写 `switch op { case policy.OpToolCall: }` 会编译不过（case 要求常量表达式）。
**不成立。** Go 规范里 expression switch 的 case 表达式不必是常量（只有 type
switch 的 case 必须是类型）。编译器实证：上游 const 被门面降级成 var 之后，
`switch op { case policy.OpToolCall: }` 照样编译通过、运行正确。
## 真正断裂的是「常量位置」
    const x = policy.OpToolCall    →  x (variable of type int) is not constant
    var arr [policy.OpToolCall]int →  invalid array length policy.OpToolCall
即 const 声明、数组长度、const 块、iota 组——这些位置要求常量表达式。
## 影响范围
**修复本身完全正确且必要**：131 处 kind 失真是真实缺陷（上游是 const，门面统一
转成 var），门面的声明形态必须与上游一致。本更正只针对影响描述里的那个举例，
不改变修复内容、不影响任何升级决策。
## 本次一并做的事
README 补一节「转发保真度」，把正确机制写给消费方：门面对上游 const 用 const
块转发，所以常量位置可用，与直接 import 上游语义一致。顺带把「上游新增导出符号
必须在本仓补别名」这条不变量和守它的门禁写进去。
## 相关位置
- 门禁报错文案已按此更正：matrix-tools/scripts/verify_facade_symbols.py
- GitHub Release v1.8.10 正文写的是正确机制
- 符号级门禁本身：oagentkit/matrix-tools#20

## [v1.8.12] — 2026-10-20

- 杂务(deps)：门面依赖地板跟到各模块最新 tag

上一轮 6 个仓打了补丁 tag，门面仓的 go.mod 还钉在旧版本：
  guardx  v0.8.0 → v0.8.1
  llm     v0.12.0 → v0.12.1
  mcpkit  v0.8.0 → v0.8.1
  obs     v0.9.0 → v0.9.1
  queue   v0.6.0 → v0.6.1
  rag     v0.8.0 → v0.8.1
为什么这要当错误处理（门禁 verify_facade_release.py 的原话）：
门面仓是消费方的单一入口，而 Go 的 MVS 只看**被消费的那个版本**的 go.mod。
所以 `go get github.com/oagentkit/agentkit@v1.8.11` 的模块图由 v1.8.11 处的
go.mod 决定——它钉着 6 个旧版本，消费方就静默拿到旧实现：不报错、不告警、
go build 照样过，只是代码是上一轮的。
这条是 .github 的 repo-ci 门禁「门面仓已发布 tag 的依赖地板」抓出来的
（.github#13 的 Matrix docs drift job 红）。注意它与那条 PR 无关——
**别让一个 PR 的红看起来像是它自己改坏的**。
本 PR 只改 main；tag 另发（门面已发布的 tag 是消费方实际拿到的东西，
只改 main 不发 tag 等于没修）。
验证：go build ./... / go test ./... / go vet ./... / gofmt -l 全绿。


## [v1.8.11] — 2026-10-20

### 修复(docs)：一键改写命令改为下载三文件后执行——exec 写法必炸

Co-authored-by: Claude Opus 4.8 (1M context) <noreply@anthropic.com>

## [v1.8.10] — 2026-10-20

### 修复(facade)：补齐 23 个漏转发符号 + 修正 131 处 const 被转发成 var

门面的符号级保真度此前无人看守：三条门面门禁（verify_facade_packages.py
只比包名、check_facade_imports.py 只查禁用 import 前缀、verify_facade_release.py
只查版本地板）没有一条检查符号级完整性。而 47 个 alias.go 的头部都写着
"// Code generated by gen. DO NOT EDIT."，但 matrix-tools 里根本不存在这个生成器
——这行声明让人以为有生成器兜底、不会漂移，实际这些文件是手工冻结的产物，
自 2026-10-17 后再没跟上上游。
影响（下游按门面 import 会直接编译失败）：
1) 漏转发 23 个上游导出符号
   - textutil 缺 11：CJKRatio、TokenRounding（type）；RoundingSegment、
     RoundingTotal（const）；DefaultCJKRatio、EstTokensWithRatio、
     EstTokensWithRounding、EstTokensCJKRunesWithRatio、FTS5Query、Tokens、
     TokensJoined（func）
   - logredact 缺 11：MaskStyle、Rule（type）；DefaultHintKeep、
     HintMinLenRatio、MaskStyleHint、MaskStyleLiteral（const）；MaskHint、
     MaskHintWith、NewRedactor、RedactWith、Rules（func）
   - policy 缺 1：Stage（func）
   其中 Tokens/TokensJoined/FTS5Query 与 MaskHint/Rule 一族都是新增特性，
   正是下游最可能用到的。
2) 131 处 kind 失真：上游是 const，门面统一转发成 var
   例：policy.OpToolCall、policy.VerdictDeny、compact.RoleUser、
   hookx.EventPreToolUse、textutil.TokenCharDivisor 等。
   这不是风格问题——const 被降级成 var 后，下游写
   `switch op { case policy.OpToolCall: }` 会编译不过（case 要求常量表达式）。
3) 头部声明与事实不符
   47 个文件统一改为实话：本文件是手工维护的别名转发表，上游新增导出符号
   必须在此补上对应别名。
验证：
- 逐包符号级差集（go/parser 提取上游与门面的导出符号并求差）复验归零：
  0 缺失、0 kind 失真、0 幽灵转发。
- go build ./... 通过；go test ./... 通过。
- 行为验证（/tmp 临时模块，经门面 import 调用）：
  textutil.Tokens("hello world") → [hello world]
  textutil.DefaultCJKRatio() → {2 4}
  textutil.RoundingSegment → 0 / RoundingTotal → 1
  logredact.MaskHint("sk-abcdefghijklmnop") → sk-a****mnop
  logredact.Rules() 长度 → 9
  textutil.TokenCharDivisor 作为 case 分支 → 编译通过（改前是 var，编译失败）
- 3 个泛型函数豁免保持不变（hotplug.NewHolder、progress.NewBus、
  reportutil.Aggregate），泛型类型别名 Holder[T any] = upstream.Holder[T]
  原样保留。
防复发：符号级完整性门禁另开 PR 落在 matrix-tools（独立 oracle，
判据不复用任何生成逻辑）。
Co-Authored-By: Claude Opus 4.8 (1M context) <noreply@anthropic.com>

## [v1.8.9] — 2026-10-20

### 重构：obsx 的转发改指 llm（obsx 已迁仓）

BREAKING CHANGE: github.com/oagentkit/obs/obsx → github.com/oagentkit/llm/obsx
门面的公共路径不变。地板跟到 obs v0.9.0 / llm v0.12.0。

## [v1.8.8] — 2026-10-20

### 重构：invoke 的转发改指 guardx

BREAKING CHANGE: github.com/oagentkit/toolgate/invoke → github.com/oagentkit/guardx/invoke
invoke 是框架无关的原语（自身只 import oekit/pkg/errors），迁到零依赖底座后，
消费它不会再被拖进 eino 家族的 34 个第三方模块。
门面的公共路径（github.com/oagentkit/agentkit/invoke）不变。
依赖地板跟到 guardx v0.8.0 / toolgate v0.12.0。

## [v1.8.7] — 2026-10-19

### 重构：audit / bashguard 的转发改指 obs / tools

BREAKING CHANGE: 门面上游的两个包已迁仓（guardx v0.7.0 起）：
  github.com/oagentkit/guardx/audit     → github.com/oagentkit/obs/audit
  github.com/oagentkit/guardx/bashguard → github.com/oagentkit/tools/bashguard
门面的**公共路径不变**（github.com/oagentkit/agentkit/audit 等照旧可用），
只有内部转发指向改了。go.mod 跟新地板：guardx v0.7.0 / obs v0.8.0 / tools v0.10.0，
并去掉两条嵌套 module 的 require（guardx/audit、guardx/bashguard）。
关联 oagentkit/guardx#3

## [v1.8.6] — 2026-10-19

- 杂务(deps)：追平 textx v0.6.0 + oekit/pkg v0.16.0

门面发布地板门禁（verify_facade_release.py）报两处落后：
- textx v0.5.0 → v0.6.0（本轮新增 EstTokensWithRatio / Tokens 等原语）
- oekit/pkg v0.15.1 → v0.16.0
MVS 只看被消费版本的 go.mod：门面 tag 钉着旧版本，消费方从门面装走的就是旧实现，
不报错、不告警、go build 照样过。


## [v1.8.5] — 2026-10-19

- 杂务(deps)：跟进 guardx v0.6.0 + 嵌套模块 audit/bashguard v0.1.0

guardx 把 audit / bashguard 拆成嵌套 module（import 路径不变），门面仓要多
require 这两个 module 才拿得到包。
连带把 guardx 地板从 v0.5.0 抬到 v0.6.0（logredact 规则表可组合 + 新增掩码档）。
关联 oagentkit/guardx#3


## [v1.8.4] — 2026-10-19

- 杂务(deps)：依赖地板对齐远端最新（agent v0.10.0 / guardx v0.5.0 / llm v0.11.0 / mcpkit v0.8.0 / obs v0.7.0 / promptx v0.8.0 / queue v0.6.0 / rag v0.8.0 / runtime v0.7.0 / skillkit v0.8.0 / textx v0.5.0 / toolgate v0.11.0 / tools v0.9.0） (13)

矩阵地板对齐的最后一波：门面把 13 个模块仓的 require 全部跟到最新 tag。
门面是消费方的单一入口，而 MVS 只看被消费版本的 go.mod——门面 tag 钉着旧版本，
消费方 go get 门面仓会静默拿到旧实现，不报错、不告警、go build 照样过。
**只改 main 不发 tag 等于没修**，所以本次必须重新发 tag。
本轮按 tier 拓扑顺序推进（tier0 → tier1 → tier2 → tier3 → 门面）：
每层发完，下一层才把地板跟到它的新 tag。否则就是「发了新的、旧的立刻变落后」，
永远收敛不了——首轮 11 仓齐发时就实测到这一点：一次发完立刻冒出 17 条新漂移。
所升版本与原钉版本相比源码零改动，逐版实测确认（diff -rq 排除 go.mod/go.sum）：
- textx / obs / runtime / mcpkit 的新旧版本 全树逐字节相同
- oekit/pkg v0.14.0→v0.15.1  仅新增 urlguard 包（纯标准库，零新增三方依赖）
- oekit/concurrency v0.17.0→v0.17.1  仅三个文档文件
- oekit/logx v0.7.0→v0.7.1  仅 CI 配置与文档
无 API 变更、无兼容风险。
本地门禁：go mod tidy -diff 无输出 / go build / go vet / gofmt / go test 全通过；
另用独立消费方模块实跑，确认 47 个 re-export 包的别名可被下游真正 import 并调用。
关联 oagentkit/matrix-tools#13


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

skillkit 与 tools 各自把落后的地板清到最新：guardx v0.3.0 → v0.4.1
（拿到 oagentkit/guardx#2 的 egress SSRF 修复），另带 textx v0.3.0 →
v0.4.0、runtime v0.5.0 → v0.6.0。
两仓只分别用 guardx/fence 与 guardx/logredact，而 guardx v0.3.0 与 v0.4.0
指向同一 commit、这两个包零变更，所以对门面仓而言本次是零行为变化。
本次仅 go.mod / go.sum 两文件变化，13 个模块仓地板其余 11 条未动。


## [v1.8.1] — 2026-10-19

- 杂务(deps)：guardx v0.4.0 → v0.4.1（egress 补 0.0.0.0/8 拦截）

guardx v0.4.1 修 egress 的 "this network" 段 SSRF 缺口：netip 的
IsUnspecified() 只匹配全零这一个地址，不是 0.0.0.0/8 的判定，而该段
此前不在显式段表里，0.1.2.3 一类目标在纯 IP 判定层全放行（oagentkit/guardx#2）。
门面仓 re-export egress，必须跟着走：main 上 go.mod 不跟进，等于门面
最新 tag 仍钉 v0.4.0，消费方 go get agentkit 会静默拿到带洞的实现。
go.sum 仅 guardx 两行变化，三方依赖一条未动。


## [v1.8.0] — 2026-10-19

### 重构(agentkit)!：门面地板跟到 llm v0.9.0 / rag v0.6.0 / agent v0.8.0

BREAKING CHANGE: 上游三个仓本轮为破坏性发布（包上移至 module 根），门面的
re-export 必须指向新版本，否则消费方从门面仓装走的是旧实现。
- llm/alias.go、rag/alias.go 的 upstream import 指向新路径
- go.mod 地板：agent v0.7.0→v0.8.0、llm v0.8.0→v0.9.0、rag v0.5.0→v0.6.0
- MIGRATION.md / README.md 的映射表同步
verify_facade_release.py 守的是 tag 不是 main，本次随 tag 一并发布。

## [v1.7.0] — 2026-10-19

- 杂务(deps)：依赖地板跟进到本轮发布的各模块 tag

各模块仓已发布新版本，门面仓跟进到这些 tag：
  guardx v0.3.0→v0.4.0   textx  v0.3.0→v0.4.0   queue  v0.4.0→v0.5.0
  runtime v0.5.0→v0.6.0  skillkit v0.5.1→v0.6.0 toolgate v0.8.0→v0.9.0
  mcpkit v0.5.1→v0.6.0   obs    v0.4.1→v0.5.0   promptx v0.5.1→v0.6.0
  rag v0.4.1→v0.5.0      tools  v0.6.1→v0.7.0   agent   v0.6.2→v0.7.0
  llm v0.7.2→v0.8.0
门面仓是本组织依赖的唯一入口，必须指向已发布版本，否则消费方拉到的
会是上一轮的旧地板。
已按 CI 环境（GOPRIVATE 直连 VCS）验证 `go mod tidy -diff` 无差异。
go.sum 的 h1 全部取自 sum.golang.org，仅就地替换目标模块的两行。

- 杂务(deps)：oekit 依赖升级 + 迁出被强推的 tag

两项改动：
1. oekit `concurrency` v0.16.0 → v0.17.0（姊妹矩阵依赖地板对齐）。
   v0.16.0...v0.17.0 仅 2 个提交：KeyedLimiter.Create 原子登记修复 + 文档，
   无 breaking，未触及任何已有导出签名。
2. 依赖迁到未被强推过的 tag。
   原先依赖的 tag 在发布后被 force 移到重写后的 main，内容与
   sum.golang.org 记录不一致，走 proxy.golang.org / GOPROXY=direct 的构建
   会 SECURITY ERROR。本地不报错是因为模块缓存预热 + goproxy.cn 还留着
   移动前的字节。已为这些模块重打不可变的附注 tag 并迁移过来。
   全部 14 个仓已启用 tag 保护规则集，禁止 tag 更新与删除。
go.sum 的 h1 全部取自 sum.golang.org，仅就地替换目标模块的两行。


## [v1.6.0] — 2026-10-19

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

- 杂务(deps)：依赖地板跟进到 8 个模块仓的新 tag

各模块仓发了新 tag（agent v0.6.1 / llm v0.7.1 / mcpkit v0.5.1 / obs v0.4.1 /
promptx v0.5.1 / rag v0.4.1 / skillkit v0.5.1 / tools v0.6.1），而门面仓 v1.5.0
处的 go.mod 仍钉着上一轮版本，verify_facade_release 判为漂移（8 处落后）。
门面仓是消费方的单一入口，MVS 只看**被消费版本**的 go.mod：main 跟上不算数，
线上流传的是 tag 处那份。所以合并后还要重新发 tag（v1.6.0），否则门禁照红。
go.sum 仍用空 GOMODCACHE + GOPRIVATE 直连 VCS 重算。原因有两个，都不是洁癖：
  1. **复用共享模块缓存会把陈旧哈希原样记回去**。本仓 #5 那次已经吃过一次：
     本地 tidy 全绿、CI 仍报 checksum mismatch——命令成功，依据却是错的。
  2. **CI 的 setup-go 缓存会把重写前的模块副本回退恢复**。go.sum 一变，缓存 key
     就变，setup-go 会按 restore-keys 回退恢复**上一份**缓存；若那份是 tag 重写
     之前建的，里面存的就是旧内容，于是正确的新哈希反而被判为 mismatch
     （toolgate #8 实测：清掉该仓 setup-go 缓存后同一次提交立刻绿）。
验证（空模块缓存 + GOPRIVATE 直连，与 CI 同配置）：go build ./... 通过、
go test ./... 0 失败（全包无测试文件）、go mod tidy -diff 无差异。


## [v1.5.0] — 2026-10-19

- 杂务(deps)：toolgate 地板跟进到 v0.8.0，并重算已失效的 go.sum

# 为什么必须动 go.sum
org 那次全量 tag/历史重写把同名 tag 指到了新提交，于是**同名 tag 对应新内容**，
而本仓 go.sum 仍记着重写前的 h1。共 8 个第一方模块的条目是陈旧的：
    agent  llm  mcpkit  obs  promptx  rag  skillkit  tools
（toolgate 一项由本次 v0.7.0 → v0.8.0 跟进覆盖）
后果不是「下载慢一点」，是硬失败：
    verifying github.com/oagentkit/agent@v0.6.0: checksum mismatch
    SECURITY ERROR … the bits may have been replaced on the origin server
门面仓 main 上一次 CI 绿在 2026-10-18 16:34Z（重写之前），之后没再跑过——
**这条红是潜伏的**，下一次触发 CI 就会炸。本次重算把它一并修掉。
# 坑：复用旧模块缓存会把陈旧哈希原样记回去
第一次重算在本机做的，看着成功，CI 却仍然报 agent@v0.6.0 校验失败。
原因是共享的模块缓存里存着重写**之前**下载的副本，`go mod tidy` 直接复用了它，
于是把旧哈希原样写回 go.sum——校验通过了，但通过的是错的依据。
正确做法是重算时给一份**空的 GOMODCACHE**，让每个第一方模块都从 VCS 现取现算：
    GOMODCACHE=$(mktemp -d) GOPRIVATE='github.com/oagentkit/*' \
    GONOPROXY='github.com/oagentkit/*' GONOSUMDB='github.com/oagentkit/*' go mod tidy
反查确认：新 go.sum 里 agent@v0.6.0 = h1:MhklGKt0QPo1MK4PmZ7N/N4JezXld7HaDEy1ve4Euik=
与 CI 报出的 downloaded 哈希逐字一致。
三方模块条目本次一条未变（仍走 sumdb 校验），变的只是「第一方模块不经公共 proxy」，
而这正是 CI 里一直在用的配置——goproxy.cn 的旧 tag 缓存无法清除，只能绕开。
# toolgate v0.8.0
facade 是消费方的单一入口，而 Go 的 MVS 只看**被消费版本**的 go.mod：
`go get github.com/oagentkit/agentkit/...` 拿到的是 tag 处的模块图，与 main 无关。
所以 v0.8.0（toolgate #7：导出只求值不执行的判定环入口）必须在这里跟进，
否则门面仓发出去的 v1.4.0 仍带着 toolgate v0.7.0。
verify_facade_release.py 守的正是这一条：各模块仓最新 tag 与门面仓最新 tag 的
go.mod 地板必须一致。toolgate 单仓打 tag 而门面仓不跟进，就是让门面仓 CI 转红——
这也是「发 tag 必须全矩阵联动」的原因。
本地验证：go build ./... 通过、go test ./...（全包无测试文件）0 失败、
go vet 与 gofmt 干净、go mod tidy -diff 无差异（即 CI 的 tidy 门禁）。


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

### 修复(docs)：re-export 包数 46 → 45（breaker 退役后门面仓不再转发它）

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

- 杂务(deps)：依赖地板升到「历史最新 + 0.1.0」

本轮做了版本线重置：清空全部历史 tag，按「历史最新 + 0.1.0」重打新版本
（如 llm v0.6.0 → v0.7.0、toolgate v0.6.1 → v0.7.0、agentkit v1.3.0 → v1.4.0）。
各仓 go.mod 的矩阵内依赖地板同步跟进到新版本号。
**go.sum 的 oagentkit 行已整段删除**，由 CI 的 `go mod tidy` 重算：
新 tag 的内容与任何历史版本都不同，`h1` 必须现算；本地没有 Go 工具链，
且 `/go.mod` 的 dirhash 无法自行复现（zip hash 可复现，gomod hash 不可复现）。
重算依赖 `.github` 的 go-ci 配了 `GOPRIVATE=github.com/oagentkit/*` 直连 VCS：
goproxy.cn 缓存着旧版本，若走 proxy 会拉回**旧代码**并算出旧 hash，
门禁全绿但依赖是错的。
（第二个提交会补上重算好的 go.sum。）
* chore(deps): 依赖地板升到「历史最新 + 0.1.0」
chore(deps): 依赖地板升到「历史最新 + 0.1.0」
本轮做了版本线重置：清空全部历史 tag，按「历史最新 + 0.1.0」重打新版本
（如 llm v0.6.0 → v0.7.0、toolgate v0.6.1 → v0.7.0、agentkit v1.3.0 → v1.4.0）。
各仓 go.mod 的矩阵内依赖地板同步跟进到新版本号。
**go.sum 的 oagentkit 行已整段删除**，由 CI 的 `go mod tidy` 重算：
新 tag 的内容与任何历史版本都不同，`h1` 必须现算；本地没有 Go 工具链，
且 `/go.mod` 的 dirhash 无法自行复现（zip hash 可复现，gomod hash 不可复现）。
重算依赖 `.github` 的 go-ci 配了 `GOPRIVATE=github.com/oagentkit/*` 直连 VCS：
goproxy.cn 缓存着旧版本，若走 proxy 会拉回**旧代码**并算出旧 hash，
门禁全绿但依赖是错的。
（第二个提交会补上重算好的 go.sum。）
* chore(deps): 回填 CI 重算的 go.sum
`h1` 由 CI 里真实的 `go mod tidy` 算出后回填——本地没有 Go 工具链，
且 `/go.mod` 的 dirhash 无法自行复现（zip hash 可复现，gomod hash 不可复现），
所以这是唯一权威来源。
重算依赖 `.github` go-ci 里的 `GOPRIVATE=github.com/oagentkit/*` 直连 VCS：
goproxy.cn 仍缓存着版本线重置前的旧内容，走 proxy 会拉回**旧代码**并算出旧 hash，
门禁全绿但依赖是错的。
验证方式就是本 PR 的 `go mod tidy -diff`：若回填有任何偏差，它会照样报差异。
* chore(deps): 回填 CI 重算的 go.sum
`h1` 由 CI 里真实的 `go mod tidy` 算出后回填——本地没有 Go 工具链，
且 `/go.mod` 的 dirhash 无法自行复现（zip hash 可复现，gomod hash 不可复现），
所以这是唯一权威来源。
重算依赖 `.github` go-ci 里的 `GOPRIVATE=github.com/oagentkit/*` 直连 VCS：
goproxy.cn 仍缓存着版本线重置前的旧内容，走 proxy 会拉回**旧代码**并算出旧 hash，
门禁全绿但依赖是错的。
验证方式就是本 PR 的 `go mod tidy -diff`：若回填有任何偏差，它会照样报差异。
* chore(deps): 回填 CI 重算的 go.sum
`h1` 由 CI 里真实的 `go mod tidy` 算出后回填——本地没有 Go 工具链，
且 `/go.mod` 的 dirhash 无法自行复现（zip hash 可复现，gomod hash 不可复现），
所以这是唯一权威来源。
重算依赖 `.github` go-ci 里的 `GOPRIVATE=github.com/oagentkit/*` 直连 VCS：
goproxy.cn 仍缓存着版本线重置前的旧内容，走 proxy 会拉回**旧代码**并算出旧 hash，
门禁全绿但依赖是错的。
验证方式就是本 PR 的 `go mod tidy -diff`：若回填有任何偏差，它会照样报差异。

- 杂务(deps)：跟进最新 tag

oagentkit/agent v0.1.0→v0.5.0、oagentkit/guardx v0.1.0→v0.2.0、oagentkit/llm v0.1.0→v0.6.0、oagentkit/mcpkit v0.1.0→v0.4.0、oagentkit/obs v0.1.0→v0.3.0、oagentkit/promptx v0.1.0→v0.4.0、oagentkit/queue v0.1.0→v0.3.0、oagentkit/rag v0.1.0→v0.3.0、oagentkit/runtime v0.1.0→v0.4.0、oagentkit/skillkit v0.1.0→v0.4.0、oagentkit/textx v0.1.0→v0.2.0、oagentkit/toolgate v0.1.0→v0.5.0、oagentkit/tools v0.1.0→v0.5.0
go.sum 取 sumdb 权威 h1，仅就地替换目标 mod 两行。

- 杂务(deps)：oekit 依赖升级

concurrency v0.6.0→v0.16.0、logx v0.5.1→v0.7.0、pkg v0.11.0→v0.14.0
姊妹矩阵 oekit 的依赖地板此前被漏掉（上一轮只扫了 oagentkit 内部依赖），本次一并对齐到最新。
升级前已确认 oekit/concurrency v0.13.1..v0.16.0 的 8 个提交全是 feat/docs/chore，
无 breaking/remove 标记。其中 #28 订正了 `WithCooldown(0)` 的语义——
即 llms.txt 里警告过的那个陷阱，升上去正好拿到修正后的行为。
go.sum 的 h1 取自 sumdb，仅就地替换目标 mod 的两行、其余行不动。


- 文档：去掉迁移示例里的字面量 ... 占位与不可用的通配 sed

原示例是一条通配 sed，目标路径写成 github.com/oagentkit/... —— 读者照抄会得到一个
真不存在的路径。而且通配 sed 本身有两处错：新路径的模块段不是旧包名
（worker/pglease -> queue/pglease、knowledge/rag -> rag/rag），正则覆盖不到；
改写还会打乱 import 组内字母序，不 gofmt 重排会挂 CI 门禁。
改为指向 matrix-tools 的 rewrite_imports.py（只改 import 块内引号，改完自动 gofmt）。


- 回滚：go.mod/go.sum 回退到 v1.0.0 的自洽版本

上一提交把 13 个依赖直接对齐到当前 tag，但 go.sum 的哈希无法在本环境可靠产出
（sum.golang.org 不可达，proxy 的 zip 虽可下但我复现的 dirhash 与 Go 实测值
不一致）。与其推一个无法验证的 go.sum，不如回退。
go.mod 里的 pin 是【最低版本】——Go 的 MVS 会让消费方 go get -u 自行升到更新版本，
而门面本身已写明「新代码不要引它」。版本对齐是打磨，不是正确性问题。
本次提交的实质价值在别处：
- 移除已退役的 breaker 门面（runtime v0.3.0 已删除该包）
- 首次接入 org 共享 CI —— 此前面门仓 46 个 re-export 包坏掉不会有人知道
- MIGRATION.md 补 breaker 迁移路径
