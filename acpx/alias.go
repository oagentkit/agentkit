// 本文件是已迁移包的别名转发表，手工维护。
//
// 上游新增导出符号时必须在此补上对应别名，否则下游按门面 import 会编译失败。
//
// Package acpx 是已迁移包的兼容门面。
//
// 新代码请直接 import github.com/oagentkit/tools/acpx。
// 本仓**已废弃**：v1.x 退役条件已满足，只保留兼容转发，不接受任何新功能。
// 新代码请直接 import 上游模块（见 MIGRATION.md 的映射表）。

package acpx

import upstream "github.com/oagentkit/tools/acpx"

type (
	Agent        = upstream.Agent
	Capability   = upstream.Capability
	ClaudeCode   = upstream.ClaudeCode
	Codex        = upstream.Codex
	Event        = upstream.Event
	Gemini       = upstream.Gemini
	GenericAgent = upstream.GenericAgent
	Kimi         = upstream.Kimi
	Mcode        = upstream.Mcode
	Mimo         = upstream.Mimo
	OpenCode     = upstream.OpenCode
	Pi           = upstream.Pi
	Registry     = upstream.Registry
	RunEvent     = upstream.RunEvent
	RunRequest   = upstream.RunRequest
	RunResult    = upstream.RunResult
	Usage        = upstream.Usage
)

const (
	EventError       = upstream.EventError
	EventResult      = upstream.EventResult
	EventText        = upstream.EventText
	EventThinking    = upstream.EventThinking
	EventToolCall    = upstream.EventToolCall
	EventToolResult  = upstream.EventToolResult
	SandboxFull      = upstream.SandboxFull
	SandboxReadonly  = upstream.SandboxReadonly
	SandboxWorkspace = upstream.SandboxWorkspace
)

var (
	CapsNames         = upstream.CapsNames
	NewClaudeCode     = upstream.NewClaudeCode
	NewCodex          = upstream.NewCodex
	NewDsh            = upstream.NewDsh
	NewGemini         = upstream.NewGemini
	NewGenericAgent   = upstream.NewGenericAgent
	NewKimi           = upstream.NewKimi
	NewMcode          = upstream.NewMcode
	NewMimo           = upstream.NewMimo
	NewOpenCode       = upstream.NewOpenCode
	NewPi             = upstream.NewPi
	NewQwen           = upstream.NewQwen
	NewRegistry       = upstream.NewRegistry
	NewZCode          = upstream.NewZCode
	UnsupportedFields = upstream.UnsupportedFields
)
