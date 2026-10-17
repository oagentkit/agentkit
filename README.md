# agentkit（门面仓）

> ⚠️ **本仓是兼容门面，不是开发目标。**
> 原单仓已拆分为 [`github.com/oagentkit/*`](https://github.com/oagentkit) 下 **13 个独立发版的仓**。
> 本仓只把 46 个包的导出符号做成别名转发，让下游**只需改 import 前缀、无需改调用点**。
> **v1.x 后本仓退役。**

## 为什么还有这个仓

原 module path 是 `git.enjoye.top/enjoydream/agentkit`。拆分意味着所有下游的
`go.mod` 都得改。直接 archive 会让下游在 `go get` 时拿到 404，破坏面比留一个门面大得多 ——
留门面至少给一个平滑过渡期。

代价是本仓会成为垃圾场：每个人都往里加东西，一年后又变回单体。所以**设死退役时间表**，
并且本仓**不接受任何新功能**，只做别名转发。

## 迁移

```
# 迁移前
import "git.enjoye.top/enjoydream/agentkit/textutil"

# 迁移后（46 个包的完整映射见 MIGRATION.md）
import "github.com/oagentkit/textx/textutil"
```

包名全部保留，只加了模块前缀，所以调用点一行都不用改。

一次性改写：

```bash
# 逐包替换 import 前缀
grep -rl 'git.enjoye.top/enjoydream/agentkit' . \
  | xargs sed -i '' 's|git.enjoye.top/enjoydream/agentkit/\([a-z]*\)|github.com/oagentkit/...|g'
```

**推荐做法**：按 [MIGRATION.md](MIGRATION.md) 的映射表直接换到最终路径，不要先引门面 ——
门面只是过渡手段。

## 门面覆盖不到的东西

**3 个泛型函数无法用变量别名转发**（Go 的泛型函数不能赋值给变量），
需要直接引用上游：

| 符号 | 上游路径 |
|---|---|
| `hotplug.NewHolder` | `github.com/oagentkit/runtime/hotplug` |
| `progress.NewBus` | `github.com/oagentkit/runtime/progress` |
| `reportutil.Aggregate` | `github.com/oagentkit/textx/reportutil` |

未导出符号（内部实现）不在门面范围内，这是 Go 包机制的固有边界。

## 模块矩阵

13 个仓，按 tier 严格自底向上：

| tier | 模块 |
|---|---|
| 0 | `textx` `runtime` `guardx` `queue` |
| 1 | `promptx` `obs` `skillkit` `mcpkit` `rag` |
| 2 | `llm` `toolgate` `tools` |
| 3 | `agent` |

## License

MIT
