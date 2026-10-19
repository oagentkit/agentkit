// 本文件是已迁移包的别名转发表，手工维护。
//
// 上游新增导出符号时必须在此补上对应别名，否则下游按门面 import 会编译失败。
//
// Package hookx 是已迁移包的兼容门面。
//
// 新代码请直接 import github.com/oagentkit/runtime/hookx。
// 本仓将在 v1.x 后退役，见 MIGRATION.md。

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
