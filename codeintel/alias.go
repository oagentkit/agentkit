// 本文件是已迁移包的别名转发表，手工维护。
//
// 上游新增导出符号时必须在此补上对应别名，否则下游按门面 import 会编译失败。
//
// Package codeintel 是已迁移包的兼容门面。
//
// 新代码请直接 import github.com/oagentkit/textx/codeintel。
// 本仓**已废弃**：v1.x 退役条件已满足，只保留兼容转发，不接受任何新功能。
// 新代码请直接 import 上游模块（见 MIGRATION.md 的映射表）。

package codeintel

import upstream "github.com/oagentkit/textx/codeintel"

type (
	Index      = upstream.Index
	PolySymbol = upstream.PolySymbol
	Ref        = upstream.Ref
	Symbol     = upstream.Symbol
)

var (
	NewIndex = upstream.NewIndex
)
