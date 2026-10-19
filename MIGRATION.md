# oagentkit 模块迁移映射表

| 旧包路径 | 新模块 | 新 import 路径 |
|---|---|---|
| `codeintel` | `textx` | `github.com/oagentkit/textx/codeintel` |
| `compact` | `textx` | `github.com/oagentkit/textx/compact` |
| `conversation` | `textx` | `github.com/oagentkit/textx/conversation` |
| `diffx` | `textx` | `github.com/oagentkit/textx/diffx` |
| `reportutil` | `textx` | `github.com/oagentkit/textx/reportutil` |
| `textutil` | `textx` | `github.com/oagentkit/textx/textutil` |
| `breaker` | **已退役** | `github.com/oekit/concurrency/breaker` |
| `hookx` | `runtime` | `github.com/oagentkit/runtime/hookx` |
| `hotplug` | `runtime` | `github.com/oagentkit/runtime/hotplug` |
| `pacer` | `runtime` | `github.com/oagentkit/runtime/pacer` |
| `procx` | `runtime` | `github.com/oagentkit/runtime/procx` |
| `progress` | `runtime` | `github.com/oagentkit/runtime/progress` |
| `audit` | `guardx` | `github.com/oagentkit/guardx/audit` |
| `bashguard` | `guardx` | `github.com/oagentkit/guardx/bashguard` |
| `egress` | `guardx` | `github.com/oagentkit/guardx/egress` |
| `fence` | `guardx` | `github.com/oagentkit/guardx/fence` |
| `injectband` | `guardx` | `github.com/oagentkit/guardx/injectband` |
| `logredact` | `guardx` | `github.com/oagentkit/guardx/logredact` |
| `permgate` | `guardx` | `github.com/oagentkit/guardx/permgate` |
| `worker` | `queue` | `github.com/oagentkit/queue/worker` |
| `worker/pglease` | `queue` | `github.com/oagentkit/queue/pglease` |
| `filestate` | `promptx` | `github.com/oagentkit/promptx/filestate` |
| `steering` | `promptx` | `github.com/oagentkit/promptx/steering` |
| `sysprompt` | `promptx` | `github.com/oagentkit/promptx/sysprompt` |
| `langfuse` | `obs` | `github.com/oagentkit/obs/langfuse` |
| `obsx` | `obs` | `github.com/oagentkit/obs/obsx` |
| `lineage` | `skillkit` | `github.com/oagentkit/skillkit/lineage` |
| `pack` | `skillkit` | `github.com/oagentkit/skillkit/pack` |
| `skill` | `skillkit` | `github.com/oagentkit/skillkit/skill` |
| `mcp` | `mcpkit` | `github.com/oagentkit/mcpkit/mcp` |
| `knowledge/rag` | `rag` | `github.com/oagentkit/rag` |
| `llm` | `llm` | `github.com/oagentkit/llm` |
| `llm/llmtest` | `llm` | `github.com/oagentkit/llm/llmtest` |
| `llm/wrap` | `llm` | `github.com/oagentkit/llm/wrap` |
| `invoke` | `toolgate` | `github.com/oagentkit/toolgate/invoke` |
| `policy` | `toolgate` | `github.com/oagentkit/toolgate/policy` |
| `toolprior` | `toolgate` | `github.com/oagentkit/toolgate/toolprior` |
| `toolsched` | `toolgate` | `github.com/oagentkit/toolgate/toolsched` |
| `acpx` | `tools` | `github.com/oagentkit/tools/acpx` |
| `sast` | `tools` | `github.com/oagentkit/tools/sast` |
| `websearch` | `tools` | `github.com/oagentkit/tools/websearch` |
| `workcopy` | `tools` | `github.com/oagentkit/tools/workcopy` |
| `agentrun` | `agent` | `github.com/oagentkit/agent/agentrun` |
| `blackboard` | `agent` | `github.com/oagentkit/agent/blackboard` |
| `clarify` | `agent` | `github.com/oagentkit/agent/clarify` |
| `dispatch` | `agent` | `github.com/oagentkit/agent/dispatch` |
| `reflection` | `agent` | `github.com/oagentkit/agent/reflection` |
| `router` | `agent` | `github.com/oagentkit/agent/router` |

## 泛型函数（门面无法转发，需直接引用上游）

- `hotplug.NewHolder` → `github.com/oagentkit/runtime/hotplug.NewHolder`
- `progress.NewBus` → `github.com/oagentkit/runtime/progress.NewBus`
- `reportutil.Aggregate` → `github.com/oagentkit/textx/reportutil.Aggregate`

## 已退役的包

### `breaker` —— 三态熔断器

runtime v0.3.0 移除了本包，熔断能力改由上游 `oekit/concurrency/breaker` 承担
（llm 仓 v0.4.0 起全量使用）。

```go
// 旧
import "github.com/oagentkit/runtime/breaker"
// 新
import "github.com/oekit/concurrency/breaker"
```

⚠️ **语义不同**：上游是「速率型 or 计数型并联触发」，显式三态 + 代绑定 +
冷却期/探测时限独立配置；不是原包那种纯计数型的三态熔断器。
配置项映射见上游 [README](https://github.com/oekit/concurrency/blob/main/breaker/README.md)。
