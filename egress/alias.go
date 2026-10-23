// 本文件是已迁移包的别名转发表，手工维护。
//
// 上游新增导出符号时必须在此补上对应别名，否则下游按门面 import 会编译失败。
//
// Package egress 是已迁移包的兼容门面。
//
// 新代码请直接 import github.com/oagentkit/guardx/egress。
// 本仓**已废弃**：v1.x 退役条件已满足，只保留兼容转发，不接受任何新功能。
// 新代码请直接 import 上游模块（见 MIGRATION.md 的映射表）。

package egress

import upstream "github.com/oagentkit/guardx/egress"

type (
	BlockedError = upstream.BlockedError
	Options      = upstream.Options
	Policy       = upstream.Policy
)

var (
	Check         = upstream.Check
	CheckHost     = upstream.CheckHost
	CheckHostWith = upstream.CheckHostWith
	CheckWith     = upstream.CheckWith

	// 具名放行面（上游 guardx v0.15.0 起）：同一个值同时派生写面校验与拨号层守卫。
	PrivateNoLoopback = upstream.PrivateNoLoopback
	PublicOnly        = upstream.PublicOnly
	SelfService       = upstream.SelfService
)
