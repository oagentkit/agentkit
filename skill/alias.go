// 本文件是已迁移包的别名转发表，手工维护。
//
// 上游新增导出符号时必须在此补上对应别名，否则下游按门面 import 会编译失败。
//
// Package skill 是已迁移包的兼容门面。
//
// 新代码请直接 import github.com/oagentkit/skillkit/skill。
// 本仓将在 v1.x 后退役，见 MIGRATION.md。

package skill

import upstream "github.com/oagentkit/skillkit/skill"

type (
	AliasResolver = upstream.AliasResolver
	Decl          = upstream.Decl
	Deprecated    = upstream.Deprecated
	FileProvider  = upstream.FileProvider
	LibEntry      = upstream.LibEntry
	LibMeta       = upstream.LibMeta
	Library       = upstream.Library
	Lister        = upstream.Lister
	LoadOption    = upstream.LoadOption
	MCPDep        = upstream.MCPDep
	Meta          = upstream.Meta
	Provider      = upstream.Provider
	Ref           = upstream.Ref
	Skill         = upstream.Skill
	ToolOption    = upstream.ToolOption
)

const (
	DefaultContentCap    = upstream.DefaultContentCap
	DefaultListBudget    = upstream.DefaultListBudget
	DefaultVersion       = upstream.DefaultVersion
	MaturityDeprecated   = upstream.MaturityDeprecated
	MaturityExperimental = upstream.MaturityExperimental
	MaturityFrozen       = upstream.MaturityFrozen
	MaturityStable       = upstream.MaturityStable
	ModeOnDemand         = upstream.ModeOnDemand
	ModeStatic           = upstream.ModeStatic
)

var (
	AsSkillTool          = upstream.AsSkillTool
	Format               = upstream.Format
	ListPrompt           = upstream.ListPrompt
	LoadFromFS           = upstream.LoadFromFS
	NewFileProvider      = upstream.NewFileProvider
	ParseRichFrontmatter = upstream.ParseRichFrontmatter
	RenderList           = upstream.RenderList
	RewriteBody          = upstream.RewriteBody
	RewriteMode          = upstream.RewriteMode
	VersionInRange       = upstream.VersionInRange
	WithContentCap       = upstream.WithContentCap
	WithDefaultMaturity  = upstream.WithDefaultMaturity
	WithExtraPrefix      = upstream.WithExtraPrefix
)
