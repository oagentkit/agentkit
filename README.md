# agentkit（已退役 · 终版 v1.9.0）

> 🪦 **本模块已退役（v1.9.0，2026-11-13）**：48 个兼容转发子包**全部删除**，
> `github.com/oagentkit/agentkit/*` 的所有 import 路径在 v1.9.0 起编译失败。
> 本仓保留 module 声明与零依赖的 go.mod，仅作墓碑与迁移指引，**不再有任何代码**。
>
> **下游要做什么**：把 `github.com/oagentkit/agentkit/<pkg>` 的 import 改到下表
> 上游真身（符号同名，调用点零改动），并把 go.mod 的
> `require github.com/oagentkit/agentkit` 换成对应上游模块后 `go mod tidy`。
> 旧单仓 module path `git.enjoye.top/enjoydream/agentkit` 同步作废。

## 历史

本仓原是单仓 `git.enjoye.top/enjoydream/agentkit`，拆分为
[`github.com/oagentkit/*`](https://github.com/oagentkit) 下 13 个独立发版仓后，
作为**兼容门面**存在：48 个活包各配一个 `alias.go`，让存量下游只需改 import
前缀、无需改调用点。门面自 v1.8.x 起废弃（不接受新功能），矩阵内零消费验证
通过后，于 v1.9.0 整仓退役（兼容逻辑清仓车道，BREAKING）。

## 48 个包 → 上游真身对照表

| 旧门面包 | 上游模块 | 新 import 路径 |
|---|---|---|
| `acpx` | tools | `github.com/oagentkit/tools/acpx` |
| `agentrun` | agent | `github.com/oagentkit/agent/agentrun` |
| `audit` | obs | `github.com/oagentkit/obs/audit` |
| `bashguard` | tools | `github.com/oagentkit/tools/bashguard` |
| `blackboard` | agent | `github.com/oagentkit/agent/blackboard` |
| `clarify` | agent | `github.com/oagentkit/agent/clarify` |
| `codeintel` | textx | `github.com/oagentkit/textx/codeintel` |
| `compact` | textx | `github.com/oagentkit/textx/compact` |
| `conversation` | textx | `github.com/oagentkit/textx/conversation` |
| `diffx` | textx | `github.com/oagentkit/textx/diffx` |
| `dispatch` | agent | `github.com/oagentkit/agent/dispatch` |
| `egress` | guardx | `github.com/oagentkit/guardx/egress` |
| `fence` | guardx | `github.com/oagentkit/guardx/fence` |
| `filestate` | promptx | `github.com/oagentkit/promptx/filestate` |
| `hookx` | runtime | `github.com/oagentkit/runtime/hookx` |
| `hotplug` | runtime | `github.com/oagentkit/runtime/hotplug` |
| `injectband` | guardx | `github.com/oagentkit/guardx/injectband` |
| `invoke` | guardx | `github.com/oagentkit/guardx/invoke` |
| `langfuse` | obs | `github.com/oagentkit/obs/langfuse` |
| `lineage` | skillkit | `github.com/oagentkit/skillkit/lineage` |
| `llm` | llm | `github.com/oagentkit/llm` |
| `llmtest` | llm | `github.com/oagentkit/llm/llmtest` |
| `logredact` | guardx | `github.com/oagentkit/guardx/logredact` |
| `mcp` | mcpkit | `github.com/oagentkit/mcpkit/mcp` |
| `obsx` | llm | `github.com/oagentkit/llm/obsx` |
| `pacer` | runtime | `github.com/oagentkit/runtime/pacer` |
| `pack` | skillkit | `github.com/oagentkit/skillkit/pack` |
| `permgate` | guardx | `github.com/oagentkit/guardx/permgate` |
| `pglease` | queue | `github.com/oagentkit/queue/pglease` |
| `policy` | toolgate | `github.com/oagentkit/toolgate/policy` |
| `port` | llm | `github.com/oagentkit/llm/port` |
| `procx` | runtime | `github.com/oagentkit/runtime/procx` |
| `progress` | runtime | `github.com/oagentkit/runtime/progress` |
| `rag` | rag | `github.com/oagentkit/rag` |
| `reflection` | agent | `github.com/oagentkit/agent/reflection` |
| `reportutil` | textx | `github.com/oagentkit/textx/reportutil` |
| `router` | agent | `github.com/oagentkit/agent/router` |
| `sast` | tools | `github.com/oagentkit/tools/sast` |
| `skill` | skillkit | `github.com/oagentkit/skillkit/skill` |
| `steering` | promptx | `github.com/oagentkit/promptx/steering` |
| `sysprompt` | promptx | `github.com/oagentkit/promptx/sysprompt` |
| `textutil` | textx | `github.com/oagentkit/textx/textutil` |
| `toolprior` | toolgate | `github.com/oagentkit/toolgate/toolprior` |
| `toolsched` | toolgate | `github.com/oagentkit/toolgate/toolsched` |
| `websearch` | tools | `github.com/oagentkit/tools/websearch` |
| `workcopy` | tools | `github.com/oagentkit/tools/workcopy` |
| `worker` | queue | `github.com/oagentkit/queue/worker` |
| `wrap` | llm | `github.com/oagentkit/llm/wrap` |

另有 1 个早已退役的包 `breaker` → `github.com/oekit/concurrency/breaker`。

迁移映射的机器可读源见 [matrix-tools/modules.json](https://github.com/oagentkit/matrix-tools/blob/main/modules.json)
（`retiringPackages` / `droppedPackages` 账本）。批量改写可用
[matrix-tools/scripts/rewrite_imports.py](https://github.com/oagentkit/matrix-tools/blob/main/scripts/rewrite_imports.py)
（`--print-rules` 先看映射表，`--apply` 真改写）。

## 状态

- **当前 v1.9.0**（退役终版，BREAKING：全部 48 个转发包删除）。
- 本仓不再发含代码的版本；后续提交仅限文档勘误。
- 归档（archive）待旧单仓 `git.enjoye.top/enjoydream/agentkit` 下线后执行。
