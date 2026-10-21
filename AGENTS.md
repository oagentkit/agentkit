# AGENTS.md — oagentkit/agentkit（门面仓）

> 已废弃的**兼容门面**——把 13 个真实模块仓的 48 个包做成**别名转发**，
> 让**存量**下游只需改 import 前缀、无需改调用点。
> module：`github.com/oagentkit/agentkit` · org：[oagentkit = Open Agent Kit](https://github.com/oagentkit)
> **本仓不是开发目标**：v1.x 退役条件已满足，只做别名转发与安全修复。

## 本仓包含

```
每个包目录 = 一个门面包 = **只有一个** `alias.go`（门面零自研实现）
共 48 个活包：`textutil/` `diffx/` `conversation/` … `llmtest/` `wrap/` `port/` …
```

| 面 | 内容 |
|---|---|
| 转发 | 48 个活包的 `alias.go`：type / const / var 逐符号别名到上游 |
| 文档 | `README.md`（门面状态与退役承诺）、`MIGRATION.md`（旧→新映射表） |
| **不做** | 新增包、新功能、自研实现、矩阵内部消费门面 |

## 硬规则

### 1. 禁止新增包（本仓最硬的红线）

**本仓不得再长出新包目录。** 门面已被钉死为「已废弃的兼容层」——每多一个包，
就多一份「门面还在被当成正经依赖」的证据，也让「门面已废弃」的承诺失效：
下游会继续把新代码引到这里，而它随时会被 archive。

- 机器约束：`matrix-tools/scripts/verify_facade_frozen.py`
  （基线快照 `matrix-tools/scripts/facade_package_baseline.json`，**单向判据**：
  门面的实际包集合必须 ⊆ 基线；基线里有、门面上没有 = 正常，不报）。
- **基线只能人工改**，且改它 = 「有人决定让这个已废弃的仓继续变大」，需要 review。
- 需要新能力？**去上游模块仓加**（textx / runtime / guardx / queue / promptx / obs /
  skillkit / mcpkit / rag / llm / toolgate / tools / agent），门面只负责转发。

### 2. 禁止矩阵内部消费门面

矩阵 13 个仓**一律不得** import `github.com/oagentkit/agentkit/...`（或旧单仓
`git.enjoye.top/enjoydream/agentkit`）。门面是给**存量下游**的过渡层，
矩阵自己消费它 = 把重依赖拉回来，并让「零消费方」的退役前提失效。

- 机器约束：`matrix-tools/scripts/check_facade_imports.py`（禁用的 import 前缀）。
- 本仓**自己除外**——门面的 alias.go 本来就 import 上游，不是 import 自己。

### 3. 转发必须 kind 一致（type / const / var 三块分明）

门面对每个包做 **kind 一致** 的转发：上游的 `type` 用 `type` 块、`const` 用 `const` 块、
`func`/`var` 用 `var` 块。

- 上游的 `const` 在门面这边**必须仍是 const**。降级成 `var` 会丢掉常量位置，
  下游按门面 import **编译失败**（`const x = policy.OpToolCall` 报 "is not constant"；
  常量数组长度等常量位置同理）。
- 上游是 `func` 的用 `var` 块转发是**正常**的（`var F = pkg.F` 调用点无差别）——
  只有 type/const 降级成 var 才会破坏下游用法。
- ⚠️ 常见误解：以为 `switch op { case policy.OpToolCall: }` 会挂——**不会**。
  Go 规范里 expression switch 的 case 表达式**不必是常量**（只有 type switch 的 case
  必须是类型）。真正挂的是 const 表达式、常量数组长度这些位置。
- **泛型函数不能转发**（Go 不允许 `var F = pkg.F[T]`）：上游的导出泛型函数必须
  **不转发**，并登记进显式白名单（`matrix-tools/scripts/verify_facade_symbols.py` 的
  `GENERIC_FUNC_WHITELIST`，当前 3 个：`hotplug.NewHolder`、`progress.NewBus`、
  `reportutil.Aggregate`）。泛型**类型**可以用类型别名转发，**必须转发**。

### 4. 上游新增导出符号 → 必须在本仓补别名

上游模块仓新增导出符号时，**必须**在本仓对应包的 `alias.go` 里补上对应别名，
否则下游按门面 import 会编译失败（`undefined: X`）。

- 补的时候按 §3 分块，保持 kind 一致。
- 上游**删除**符号时门面也要跟着删——留着就是**幽灵转发**（指向已不存在的上游 API）。
- 机器约束：`matrix-tools/scripts/verify_facade_symbols.py`
  （不漏 / 不错配 / 不幽灵 / 泛型白名单不漂 / 文档与白名单一致 / 路径与矩阵一致）。

### 5. 文档必须与包集合同步

`MIGRATION.md` 的映射表与门面的 48 个活包**一一对应**：
上游包新增 / 改名 / 退役时，同步改门面的 alias.go **和** MIGRATION.md 那一行。

> 2026-10-21 实测教训：`llm/port` 的门面转发（`port/alias.go` → `github.com/oagentkit/llm/port`）
> 早就存在，MIGRATION.md 却**漏了这一行**。下游照表逐包替换会漏掉 `port`——
> 而 llm 是 tier 2 重依赖仓，漏掉等于把重依赖又拉回来。

- 机器约束：`matrix-tools/scripts/verify_facade_packages.py`（包级覆盖 + 陈旧转发）。
- README 的包数口径也要跟着改：门面转发的是**活包**数（48），
  MIGRATION.md = 活包映射 **+ 1 行已退役包**（`breaker`）。

### 6. 退役纪律

- 门面 README 首行「门面仓 · 已废弃」、状态行「v1.x 退役条件**已满足**」——
  这两行是**门禁的判据来源**，改动前先确认不会把承诺改松。
- **本仓不接受任何新功能**：只做别名转发与安全修复。
- 归档（archive）待旧单仓 `git.enjoye.top/enjoydream/agentkit` 下线后执行；
  在此之前本仓继续可用，**已有 tag 的 `go get` 照常解析**。
- **门面已废弃 ≠ 已转发的兼容面可以烂掉。** 门面活着期间，§3 / §4 的符号级门禁
  必须继续绿——那 48 个 `alias.go` 是存量下游**唯一编译得过的那张脸**。
- 门面**零自研实现**：任何不是「转发上游」的 `.go` 代码都是越界
  （门面自研包比多转发一个上游更严重）。

## 明确不属于本仓的

| 能力 | 归属 |
|---|---|
| 任何真实实现 | 对应的上游模块仓（本仓只转发） |
| 文本/差异/会话/报告 | `github.com/oagentkit/textx` |
| 事件总线/流控/子进程/热插拔/钩子 | `github.com/oagentkit/runtime` |
| 脱敏/出口围栏/权限/沙箱解析/审计 | `github.com/oagentkit/guardx` |
| DB 队列与选主 | `github.com/oagentkit/queue` |
| 提示词/上下文文件/引导 | `github.com/oagentkit/promptx` |
| 观测/审计/Langfuse | `github.com/oagentkit/obs` |
| 技能/打包/血缘 | `github.com/oagentkit/skillkit` |
| MCP 接入 | `github.com/oagentkit/mcpkit` |
| 检索/知识库 | `github.com/oagentkit/rag` |
| 模型客户端/降级链/预算/成本 | `github.com/oagentkit/llm`（窄接口层 `llm/port`） |
| 工具策略/调度/优先级 | `github.com/oagentkit/toolgate` |
| 外部工具/SAST/git 副本 | `github.com/oagentkit/tools` |
| Agent 运行/编排/路由 | `github.com/oagentkit/agent` |
