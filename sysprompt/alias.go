// 本文件是已迁移包的别名转发表，手工维护。
//
// 上游新增导出符号时必须在此补上对应别名，否则下游按门面 import 会编译失败。
//
// Package sysprompt 是已迁移包的兼容门面。
//
// 新代码请直接 import github.com/oagentkit/promptx/sysprompt。
// 本仓将在 v1.x 后退役，见 MIGRATION.md。

package sysprompt

import upstream "github.com/oagentkit/promptx/sysprompt"

type (
	Block    = upstream.Block
	Boundary = upstream.Boundary
	Builder  = upstream.Builder
	EnvInfo  = upstream.EnvInfo
	GitInfo  = upstream.GitInfo
	Result   = upstream.Result
	Section  = upstream.Section
	Target   = upstream.Target
)

const (
	DefaultUserContextDisclaimer = upstream.DefaultUserContextDisclaimer
	DefaultUserContextIntro      = upstream.DefaultUserContextIntro
	Dynamic                      = upstream.Dynamic
	Stable                       = upstream.Stable
	TargetSystem                 = upstream.TargetSystem
	TargetUserContext            = upstream.TargetUserContext
)

var (
	DateSection     = upstream.DateSection
	DetectEnv       = upstream.DetectEnv
	EnvSection      = upstream.EnvSection
	GitSection      = upstream.GitSection
	IdentitySection = upstream.IdentitySection
	NewSection      = upstream.NewSection
)
