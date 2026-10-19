// 本文件是已迁移包的别名转发表，手工维护。
//
// 上游新增导出符号时必须在此补上对应别名，否则下游按门面 import 会编译失败。
//
// Package toolprior 是已迁移包的兼容门面。
//
// 新代码请直接 import github.com/oagentkit/toolgate/toolprior。
// 本仓将在 v1.x 后退役，见 MIGRATION.md。

package toolprior

import upstream "github.com/oagentkit/toolgate/toolprior"

type (
	Entry = upstream.Entry
	Table = upstream.Table
)

const (
	PriorityCore     = upstream.PriorityCore
	PriorityExternal = upstream.PriorityExternal
	PrioritySupport  = upstream.PrioritySupport
)

var (
	LimitCalls = upstream.LimitCalls
	NewTable   = upstream.NewTable
)
