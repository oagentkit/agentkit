// 本文件是已迁移包的别名转发表，手工维护。
//
// 上游新增导出符号时必须在此补上对应别名，否则下游按门面 import 会编译失败。
//
// Package hookx 是已迁移包的兼容门面。
//
// 新代码请直接 import github.com/oagentkit/runtime/hookx。
// 本仓**已废弃**：v1.x 退役条件已满足，只保留兼容转发，不接受任何新功能。
// 新代码请直接 import 上游模块（见 MIGRATION.md 的映射表）。

package hookx

import upstream "github.com/oagentkit/runtime/hookx"

type (
	Event     = upstream.Event
	Hook      = upstream.Hook
	Input     = upstream.Input
	Output    = upstream.Output
	RunResult = upstream.RunResult
	Runner    = upstream.Runner
	Specific  = upstream.Specific
)

const (
	DefaultTimeout          = upstream.DefaultTimeout
	EventPermissionRequest  = upstream.EventPermissionRequest
	EventPostToolUse        = upstream.EventPostToolUse
	EventPostToolUseFailure = upstream.EventPostToolUseFailure
	EventPreToolUse         = upstream.EventPreToolUse
	EventSessionStart       = upstream.EventSessionStart
	EventStop               = upstream.EventStop
	EventUserPromptSubmit   = upstream.EventUserPromptSubmit
)
