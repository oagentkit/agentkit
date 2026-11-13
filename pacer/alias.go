// 本文件是已迁移包的别名转发表，手工维护。
//
// 上游新增导出符号时必须在此补上对应别名，否则下游按门面 import 会编译失败。
//
// Package pacer 是已迁移包的兼容门面。
//
// 新代码请直接 import github.com/oagentkit/runtime/pacer。
// 本仓**已废弃**：v1.x 退役条件已满足，只保留兼容转发，不接受任何新功能。
// 新代码请直接 import 上游模块（见 MIGRATION.md 的映射表）。

package pacer

import upstream "github.com/oagentkit/runtime/pacer"

type (
	Delta = upstream.Delta
	// StreamOption：runtime v0.13.0 起裸 Option 改名 StreamOption（runtime#110
	// 禁裸 Option）——门面只转发上游现存符号，旧名 Option 不再保留（幽灵转发
	// 门禁），下游改引 StreamOption 或直接用上游 runtime/pacer。
	StreamOption = upstream.StreamOption
	Stream       = upstream.Stream
)

const (
	DefaultBytes    = upstream.DefaultBytes
	DefaultCapBytes = upstream.DefaultCapBytes
	DefaultInterval = upstream.DefaultInterval
)

var (
	NewStream = upstream.NewStream
	WithClock = upstream.WithClock
)
