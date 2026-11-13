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
	Stream = upstream.Stream
	// StreamOption 上游 runtime#112 把裸 Option 改名 StreamOption（AGENTS.md
	// 明禁裸 Option）——门面随上游同步导出；Option 保留为它的别名，
	// 旧引用（pacer.Option）编译不破（门面兼容面：存量别名只做转发）。
	StreamOption = upstream.StreamOption
	Option       = upstream.StreamOption
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
