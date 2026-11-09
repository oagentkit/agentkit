// 本文件是已迁移包的别名转发表，手工维护。
//
// 上游新增导出符号时必须在此补上对应别名，否则下游按门面 import 会编译失败。
//
// Package router 是已迁移包的兼容门面。
//
// 新代码请直接 import github.com/oagentkit/agent/router。
// 本仓**已废弃**：v1.x 退役条件已满足，只保留兼容转发，不接受任何新功能。
// 新代码请直接 import 上游模块（见 MIGRATION.md 的映射表）。

package router

import upstream "github.com/oagentkit/agent/router"

type (
	Config           = upstream.Config
	Decision         = upstream.Decision
	DecisionFace     = upstream.DecisionFace
	DecisionObserver = upstream.DecisionObserver
	Route            = upstream.Route
	Router           = upstream.Router
	SlotSpec         = upstream.SlotSpec
)

const (
	SourceLLM      = upstream.SourceLLM
	SourceDecision = upstream.SourceDecision
)

var (
	KeywordHit               = upstream.KeywordHit
	KeywordHitBoundary       = upstream.KeywordHitBoundary
	KeywordHitBoundaryExcept = upstream.KeywordHitBoundaryExcept
	KeywordHitExcept         = upstream.KeywordHitExcept
	KeywordPostNegated       = upstream.KeywordPostNegated
	New                      = upstream.New
	// ⚠️ 经门面给 OnNegationHit 赋值是**死写**（agentkit#155/#166）：本别名是
	// init 时的值拷贝，赋值只改写门面副本；上游 keyword.go 运行时读的是
	// SetOnNegationHit 维护的原子快照——本别名读写都不反映已生效的钩子。
	// Deprecated: use SetOnNegationHit instead（agentkit#166）。
	OnNegationHit    = upstream.OnNegationHit
	SetOnNegationHit = upstream.SetOnNegationHit
)
