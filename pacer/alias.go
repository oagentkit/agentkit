// 本文件是已迁移包的别名转发表，手工维护。
//
// 上游新增导出符号时必须在此补上对应别名，否则下游按门面 import 会编译失败。
//
// Package pacer 是已迁移包的兼容门面。
//
// 新代码请直接 import github.com/oagentkit/runtime/pacer。
// 本仓将在 v1.x 后退役，见 MIGRATION.md。

package pacer

import upstream "github.com/oagentkit/runtime/pacer"

type (
	Delta  = upstream.Delta
	Option = upstream.Option
	Stream = upstream.Stream
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
