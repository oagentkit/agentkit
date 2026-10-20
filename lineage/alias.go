// 本文件是已迁移包的别名转发表，手工维护。
//
// 上游新增导出符号时必须在此补上对应别名，否则下游按门面 import 会编译失败。
//
// Package lineage 是已迁移包的兼容门面。
//
// 新代码请直接 import github.com/oagentkit/skillkit/lineage。
// 本仓**已废弃**：v1.x 退役条件已满足，只保留兼容转发，不接受任何新功能。
// 新代码请直接 import 上游模块（见 MIGRATION.md 的映射表）。

package lineage

import upstream "github.com/oagentkit/skillkit/lineage"

type (
	Expert       = upstream.Expert
	ExpertEdges  = upstream.ExpertEdges
	Grant        = upstream.Grant
	Hub          = upstream.Hub
	Impact       = upstream.Impact
	Lineage      = upstream.Lineage
	LineageEdge  = upstream.LineageEdge
	LineageFocus = upstream.LineageFocus
	LineageNode  = upstream.LineageNode
	MCPEdges     = upstream.MCPEdges
	Skill        = upstream.Skill
	SkillEdges   = upstream.SkillEdges
)

var (
	Build         = upstream.Build
	Diff          = upstream.Diff
	NewHub        = upstream.NewHub
	SkillFromMeta = upstream.SkillFromMeta
)
