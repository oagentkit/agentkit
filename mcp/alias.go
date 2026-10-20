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
	Catalog           = upstream.Catalog
	CatalogDiff       = upstream.CatalogDiff
	ClassifiedError   = upstream.ClassifiedError
	ErrorKind         = upstream.ErrorKind
	Health            = upstream.Health
	HealthIssue       = upstream.HealthIssue
	HealthState       = upstream.HealthState
	OAuthClientConfig = upstream.OAuthClientConfig
	OutputSchemer     = upstream.OutputSchemer
	Pool              = upstream.Pool
	SelfcheckOptions  = upstream.SelfcheckOptions
	SelfcheckResult   = upstream.SelfcheckResult
	Selfchecker       = upstream.Selfchecker
	ServerConfig      = upstream.ServerConfig
	ToolCheck         = upstream.ToolCheck
	ToolHints         = upstream.ToolHints
	ToolResultError   = upstream.ToolResultError
	ToolSpec          = upstream.ToolSpec
	ToolStatus        = upstream.ToolStatus
	ToolsResult       = upstream.ToolsResult
)

const (
	DefaultProbeTimeout     = upstream.DefaultProbeTimeout
	DefaultRequestTimeout   = upstream.DefaultRequestTimeout
	DefaultSelfcheckTTL     = upstream.DefaultSelfcheckTTL
	DefaultSelfcheckTimeout = upstream.DefaultSelfcheckTimeout
	DefaultTimeout          = upstream.DefaultTimeout
	HealthDegraded          = upstream.HealthDegraded
	HealthDown              = upstream.HealthDown
	HealthHealthy           = upstream.HealthHealthy
	KindInvalidArgs         = upstream.KindInvalidArgs
	KindProtocol            = upstream.KindProtocol
	KindToolMissing         = upstream.KindToolMissing
	KindUnknown             = upstream.KindUnknown
	KindUnreachable         = upstream.KindUnreachable
	ToolFailed              = upstream.ToolFailed
	ToolMissing             = upstream.ToolMissing
	ToolOK                  = upstream.ToolOK
	ToolUnreachable         = upstream.ToolUnreachable
)

var (
	Aggregate              = upstream.Aggregate
	CallMeta               = upstream.CallMeta
	Classify               = upstream.Classify
	ClassifyError          = upstream.ClassifyError
	ClassifyRPCCode        = upstream.ClassifyRPCCode
	Diff                   = upstream.Diff
	DiffNames              = upstream.DiffNames
	ErrNotConnected        = upstream.ErrNotConnected
	ExpandEnvRef           = upstream.ExpandEnvRef
	ExpandEnvRefs          = upstream.ExpandEnvRefs
	IsToolResultError      = upstream.IsToolResultError
	KindOf                 = upstream.KindOf
	NewCatalog             = upstream.NewCatalog
	NewPool                = upstream.NewPool
	NewSelfchecker         = upstream.NewSelfchecker
	OutputSchemaOf         = upstream.OutputSchemaOf
	UnwrapMCPText          = upstream.UnwrapMCPText
	WithCallMeta           = upstream.WithCallMeta
	WrapErrorAsObservation = upstream.WrapErrorAsObservation
)
