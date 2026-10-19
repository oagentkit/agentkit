// 本文件是已迁移包的别名转发表，手工维护。
//
// 上游新增导出符号时必须在此补上对应别名，否则下游按门面 import 会编译失败。
//
// Package reportutil 是已迁移包的兼容门面。
//
// 新代码请直接 import github.com/oagentkit/textx/reportutil。
// 本仓将在 v1.x 后退役，见 MIGRATION.md。
//
// 注意：泛型函数无法用变量别名转发，这些符号请直接引用上游：Aggregate

package reportutil

import upstream "github.com/oagentkit/textx/reportutil"

type (
	Group[T any] = upstream.Group[T]
)

const (
	High           = upstream.High
	Low            = upstream.Low
	Medium         = upstream.Medium
	PatternTextMax = upstream.PatternTextMax
)

var (
	BetaConfidence       = upstream.BetaConfidence
	Fingerprint          = upstream.Fingerprint
	Higher               = upstream.Higher
	McNemarExact         = upstream.McNemarExact
	Normalize            = upstream.Normalize
	NormalizeComment     = upstream.NormalizeComment
	NormalizePatternText = upstream.NormalizePatternText
	PatternKey           = upstream.PatternKey
	PatternSample        = upstream.PatternSample
	PatternSamples       = upstream.PatternSamples
	Rank                 = upstream.Rank
	Valid                = upstream.Valid
	WilsonCI             = upstream.WilsonCI
)
