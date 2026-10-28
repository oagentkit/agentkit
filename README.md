# agentkit（门面仓 · 已废弃）

> 🗄️ **本仓已废弃：只保留兼容转发，不是开发目标，新代码不要引它。**
> 原单仓已拆分为 [`github.com/oagentkit/*`](https://github.com/oagentkit) 下 **13 个独立发版的仓**。
> 本仓只把 **48 个活包**的导出符号做成别名转发，让**存量**下游只需改 import 前缀、无需改调用点。
> 另有 1 个包（`breaker`）已退役，不在此列——见 [MIGRATION.md](MIGRATION.md#已退役的包)。
> 所以 [MIGRATION.md](MIGRATION.md) 的映射表 = **48 个活包**的迁移映射 **+ 1 行已退役包**。
>
> **状态**：v1.x 退役条件**已满足**（当前 v1.8.54）。本仓不接受任何新功能，
> 只做别名转发与安全修复。归档（archive）待旧单仓
> `git.enjoye.top/enjoydream/agentkit`（module path 的实际所在）下线后执行——
> 那才是真正切断旧路径的一步；在此之前本仓继续可用，**已有 tag 的 `go get` 照常解析**。
>
> **新代码请直接 import 上游模块**（见下方「迁移」与 [MIGRATION.md](MIGRATION.md)）。

## 为什么还有这个仓

原 module path 是 `git.enjoye.top/enjoydream/agentkit`。拆分意味着所有下游的
`go.mod` 都得改。直接 archive 会让下游在 `go get` 时拿到 404，破坏面比留一个门面大得多 ——
留门面至少给一个平滑过渡期。

代价是本仓会成为垃圾场：每个人都往里加东西，一年后又变回单体。所以**设死退役时间表**，
并且本仓**不接受任何新功能**，只做别名转发。

> 过渡期已经走完：矩阵内**零消费方**（org 内 0、姊妹 org oekit 内 0、
> GitHub 全站代码搜索命中 0），门面符号级门禁与「谁还在引门面」门禁都在守。
> 剩下的唯一阻塞是 GitCode 侧旧单仓尚未归档。

## 迁移

```
# 迁移前
import "git.enjoye.top/enjoydream/agentkit/textutil"

# 迁移后（48 个包的完整映射见 MIGRATION.md）
import "github.com/oagentkit/textx/textutil"
```

包名全部保留，只加了模块前缀，所以调用点一行都不用改。

一次性改写：

```bash
# 建议按 MIGRATION.md 的映射表逐包替换，逐条核对而不是一条通配。
# 通配 sed 的坑：新路径的模块段不是旧包名（worker/pglease -> queue/pglease、
# knowledge/rag -> rag），正则很难覆盖；且改写会打乱 import 组内字母序，
# 必须让 gofmt 重排，否则挂 CI 门禁。
#
# 用 matrix-tools 的脚本（只改 import 块内的引号字符串，改完自动 gofmt）：
#
#   三个文件必须放在同一目录：rewrite_imports.py 依赖同目录的 matrix_paths.py，
#   矩阵定义 modules.json 用 --modules 显式指定。
#   ⚠️ 不能用 `python3 -c "exec(urlopen(...).read())"` 的写法：exec 的代码里
#   没有 __file__，脚本定位同目录依赖与 spawn 子进程都会炸（NameError）。
#   脚本必须作为文件运行。
curl -sSfL https://raw.githubusercontent.com/oagentkit/matrix-tools/main/scripts/rewrite_imports.py -o rewrite_imports.py
curl -sSfL https://raw.githubusercontent.com/oagentkit/matrix-tools/main/scripts/matrix_paths.py -o matrix_paths.py
curl -sSfL https://raw.githubusercontent.com/oagentkit/matrix-tools/main/modules.json -o modules.json

# 看映射表（不动文件）
python3 rewrite_imports.py --org oagentkit --modules ./modules.json --print-rules

# 真改写（--root 指向你的工程根；只改 import 块内的引号字符串，改完自动 gofmt）
python3 rewrite_imports.py --org oagentkit --modules ./modules.json --root . --apply
```

改写后按脚本提示收尾：`go mod tidy && go build ./... && gofmt -l .`。

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

## 转发保真度

门面对每个包做 **kind 一致** 的转发：上游的 `type` 用 type 块、`const` 用 const 块、
`func`/`var` 用 var 块。所以：

- 上游的 `const` 在门面这边**仍是 const**——可以用于常量位置（const 声明、数组长度、
  const 块、iota 组），语义与直接 import 上游完全一致。（不是「降级成 var」：
  那样会丢掉常量位置，编译不过。）
- 上游新增导出符号时，必须在本仓对应包的 `alias.go` 里补上对应别名，否则下游按门面
  import 会编译失败。矩阵侧有符号级门禁
  （`matrix-tools/scripts/verify_facade_symbols.py`）守着这条不变量。

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
