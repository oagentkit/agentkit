// 本文件是已迁移包的别名转发表，手工维护。
//
// 上游新增导出符号时必须在此补上对应别名，否则下游按门面 import 会编译失败。
//
// Package bashguard 是已迁移包的兼容门面。
//
// 新代码请直接 import github.com/oagentkit/tools/bashguard。
// 本仓**已废弃**：v1.x 退役条件已满足，只保留兼容转发，不接受任何新功能。
// 新代码请直接 import 上游模块（见 MIGRATION.md 的映射表）。

package bashguard

import upstream "github.com/oagentkit/tools/bashguard"

type (
	Decision     = upstream.Decision
	FlagKind     = upstream.FlagKind
	InnerCmdLoc  = upstream.InnerCmdLoc
	InnerCmdRule = upstream.InnerCmdRule
	Invocation   = upstream.Invocation
	Policy       = upstream.Policy
	RiskTier     = upstream.RiskTier
	Table        = upstream.Table
)

const (
	FlagBool          = upstream.FlagBool
	FlagNumber        = upstream.FlagNumber
	FlagOptional      = upstream.FlagOptional
	FlagString        = upstream.FlagString
	InnerAfterDashes  = upstream.InnerAfterDashes
	InnerAfterOperand = upstream.InnerAfterOperand
	InnerFirstOperand = upstream.InnerFirstOperand
	InnerFlagValue    = upstream.InnerFlagValue
	MaxCommandLen     = upstream.MaxCommandLen
	RiskDestructive   = upstream.RiskDestructive
	RiskReadOnly      = upstream.RiskReadOnly
	RiskUnevaluated   = upstream.RiskUnevaluated
	RiskWrite         = upstream.RiskWrite
)

var (
	Analyze      = upstream.Analyze
	DefaultTable = upstream.DefaultTable
)

type (
	RedirInfo = upstream.RedirInfo
)
