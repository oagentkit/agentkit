// 本文件是已迁移包的别名转发表，手工维护。
//
// 上游新增导出符号时必须在此补上对应别名，否则下游按门面 import 会编译失败。
//
// Package policy 是已迁移包的兼容门面。
//
// 新代码请直接 import github.com/oagentkit/toolgate/policy。
// 本仓将在 v1.x 后退役，见 MIGRATION.md。

package policy

import upstream "github.com/oagentkit/toolgate/policy"

type (
	Arbiter  = upstream.Arbiter
	Decider  = upstream.Decider
	Decision = upstream.Decision
	Gate     = upstream.Gate
	Mode     = upstream.Mode
	Op       = upstream.Op
	OpType   = upstream.OpType
	Policy   = upstream.Policy
	Rule     = upstream.Rule
	Verdict  = upstream.Verdict
)

const (
	DeciderHuman     = upstream.DeciderHuman
	DeciderLLM       = upstream.DeciderLLM
	DeciderRule      = upstream.DeciderRule
	DeciderSystem    = upstream.DeciderSystem
	ModeAuto         = upstream.ModeAuto
	ModeConfirm      = upstream.ModeConfirm
	ModeFull         = upstream.ModeFull
	ModePlan         = upstream.ModePlan
	OpExecuteStep    = upstream.OpExecuteStep
	OpPlanSubmit     = upstream.OpPlanSubmit
	OpToolCall       = upstream.OpToolCall
	VerdictAuto      = upstream.VerdictAuto
	VerdictDeny      = upstream.VerdictDeny
	VerdictNeedHuman = upstream.VerdictNeedHuman
	VerdictPlanOnly  = upstream.VerdictPlanOnly
	VerdictUnknown   = upstream.VerdictUnknown
)

var (
	AuditGate     = upstream.AuditGate
	DefaultPolicy = upstream.DefaultPolicy
	LoadOverrides = upstream.LoadOverrides
	NewGate       = upstream.NewGate
	Normalize     = upstream.Normalize
	Stage         = upstream.Stage
	ValidMode     = upstream.ValidMode
)
