// 本文件是已迁移包的别名转发表，手工维护。
//
// 上游新增导出符号时必须在此补上对应别名，否则下游按门面 import 会编译失败。
//
// Package permgate 是已迁移包的兼容门面。
//
// 新代码请直接 import github.com/oagentkit/guardx/permgate。
// 本仓将在 v1.x 后退役，见 MIGRATION.md。

package permgate

import upstream "github.com/oagentkit/guardx/permgate"

type (
	Behavior   = upstream.Behavior
	Capability = upstream.Capability
	Decision   = upstream.Decision
	Mode       = upstream.Mode
	Rule       = upstream.Rule
	Ruleset    = upstream.Ruleset
	Service    = upstream.Service
)

const (
	Allow        = upstream.Allow
	Ask          = upstream.Ask
	Deny         = upstream.Deny
	ModeDefault  = upstream.ModeDefault
	ModeReadOnly = upstream.ModeReadOnly
	ModeYolo     = upstream.ModeYolo
)
