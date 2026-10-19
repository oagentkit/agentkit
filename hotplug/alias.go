// 本文件是已迁移包的别名转发表，手工维护。
//
// 上游新增导出符号时必须在此补上对应别名，否则下游按门面 import 会编译失败。
//
// Package hotplug 是已迁移包的兼容门面。
//
// 新代码请直接 import github.com/oagentkit/runtime/hotplug。
// 本仓将在 v1.x 后退役，见 MIGRATION.md。
//
// 注意：泛型函数无法用变量别名转发，这些符号请直接引用上游：NewHolder

package hotplug

import upstream "github.com/oagentkit/runtime/hotplug"

type (
	Holder[T any] = upstream.Holder[T]
	Plugboard     = upstream.Plugboard
)

var (
	NewPlugboard = upstream.NewPlugboard
)
