// 本文件是已迁移包的别名转发表，手工维护。
//
// 上游新增导出符号时必须在此补上对应别名，否则下游按门面 import 会编译失败。
//
// Package langfuse 是已迁移包的兼容门面。
//
// 新代码请直接 import github.com/oagentkit/obs/langfuse。
// 本仓将在 v1.x 后退役，见 MIGRATION.md。

package langfuse

import upstream "github.com/oagentkit/obs/langfuse"

type (
	Batch       = upstream.Batch
	Client      = upstream.Client
	Observation = upstream.Observation
	Query       = upstream.Query
	Trace       = upstream.Trace
)

var (
	NewClient = upstream.NewClient
)
