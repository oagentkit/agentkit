// 本文件是已迁移包的别名转发表，手工维护。
//
// 上游新增导出符号时必须在此补上对应别名，否则下游按门面 import 会编译失败。
//
// Package wrap 是已迁移包的兼容门面。
//
// 新代码请直接 import github.com/oagentkit/llm/wrap。
// 本仓将在 v1.x 后退役，见 MIGRATION.md。

package wrap

import upstream "github.com/oagentkit/llm/wrap"

type (
	ContinueModel = upstream.ContinueModel
	FitAction     = upstream.FitAction
	FitModel      = upstream.FitModel
	FitResult     = upstream.FitResult
	RateModel     = upstream.RateModel
)

const (
	FitMicro      = upstream.FitMicro
	FitMicroTrunc = upstream.FitMicroTrunc
	FitPass       = upstream.FitPass
	FitTruncated  = upstream.FitTruncated
)

var (
	FitMessages      = upstream.FitMessages
	NewContinueModel = upstream.NewContinueModel
	NewFitModel      = upstream.NewFitModel
	NewRateModel     = upstream.NewRateModel
)
