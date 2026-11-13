# AGENTS.md — oagentkit/agentkit（已退役门面）

> 🪦 **本模块已于 v1.9.0（2026-11-13）退役**：48 个兼容转发子包全部删除，
> 仓里**不再有任何 `.go` 代码**。go.mod 只保留 module 声明、零依赖。
> module：`github.com/oagentkit/agentkit` · org：[oagentkit = Open Agent Kit](https://github.com/oagentkit)

## 本仓现在是什么

退役墓碑 + 迁移指引：

- `README.md` —— 退役公告 + 48 包 → 上游真身对照表（机器可读源：
  [matrix-tools/modules.json](https://github.com/oagentkit/matrix-tools/blob/main/modules.json)）。
- `CHANGELOG.md` —— 历史版本记录（v1.9.0 为终版）。
- `go.mod` —— 只有 `module github.com/oagentkit/agentkit`，无任何 require。

## 硬规则（退役后依然有效）

1. **禁止引入任何 `.go` 文件**——退役的语义就是「不再提供 API」。复活某个包
   的正确做法是去上游模块仓提 issue，不是在这里加回 alias。
2. **禁止往 go.mod 加依赖**——零依赖是墓碑的形状；有依赖说明又长出了代码。
3. 后续提交仅限文档勘误（README / CHANGELOG / 本文件）。
4. 归档（archive）待旧单仓 `git.enjoye.top/enjoydream/agentkit` 下线后执行；
   归档前已有 tag 的 `go get` 照常解析（v1.8.x 及以前仍带兼容转发）。

## 上游归属（旧包 → 真身仓）

| 能力 | 归属 |
|---|---|
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
| 熔断（旧 `breaker` 包） | `github.com/oekit/concurrency/breaker` |
