// 本文件是已迁移包的别名转发表，手工维护。
//
// 上游新增导出符号时必须在此补上对应别名，否则下游按门面 import 会编译失败。
//
// Package mcp 是已迁移包的兼容门面。
//
// 新代码请直接 import github.com/oagentkit/mcpkit/mcp。
// 本仓将在 v1.x 后退役，见 MIGRATION.md。

package mcp

import upstream "github.com/oagentkit/mcpkit/mcp"

type (
	OAuthClientConfig = upstream.OAuthClientConfig
	OutputSchemer     = upstream.OutputSchemer
	Pool              = upstream.Pool
	ServerConfig      = upstream.ServerConfig
	ToolHints         = upstream.ToolHints
	ToolResultError   = upstream.ToolResultError
	ToolSpec          = upstream.ToolSpec
	ToolsResult       = upstream.ToolsResult
)

const (
	DefaultProbeTimeout   = upstream.DefaultProbeTimeout
	DefaultRequestTimeout = upstream.DefaultRequestTimeout
	DefaultTimeout        = upstream.DefaultTimeout
)

var (
	CallMeta               = upstream.CallMeta
	ErrNotConnected        = upstream.ErrNotConnected
	ExpandEnvRef           = upstream.ExpandEnvRef
	ExpandEnvRefs          = upstream.ExpandEnvRefs
	IsToolResultError      = upstream.IsToolResultError
	NewPool                = upstream.NewPool
	OutputSchemaOf         = upstream.OutputSchemaOf
	UnwrapMCPText          = upstream.UnwrapMCPText
	WithCallMeta           = upstream.WithCallMeta
	WrapErrorAsObservation = upstream.WrapErrorAsObservation
)
