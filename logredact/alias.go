// 本文件是已迁移包的别名转发表，手工维护。
//
// 上游新增导出符号时必须在此补上对应别名，否则下游按门面 import 会编译失败。
//
// Package logredact 是已迁移包的兼容门面。
//
// 新代码请直接 import github.com/oagentkit/guardx/logredact。
// 本仓**已废弃**：v1.x 退役条件已满足，只保留兼容转发，不接受任何新功能。
// 新代码请直接 import 上游模块（见 MIGRATION.md 的映射表）。

package logredact

import upstream "github.com/oagentkit/guardx/logredact"

type (
	MaskStyle = upstream.MaskStyle
	Masker    = upstream.Masker
	Rule      = upstream.Rule
)

const (
	DefaultHintKeep  = upstream.DefaultHintKeep
	HintMinLenRatio  = upstream.HintMinLenRatio
	MaskStyleHint    = upstream.MaskStyleHint
	MaskStyleLiteral = upstream.MaskStyleLiteral
)

var (
	MaskHint           = upstream.MaskHint
	MaskHintWith       = upstream.MaskHintWith
	NewMasker          = upstream.NewMasker
	NewPayloadRedactor = upstream.NewPayloadRedactor
	NewRedactor        = upstream.NewRedactor
	Redact             = upstream.Redact
	RedactPayloads     = upstream.RedactPayloads
	RedactPayloadsWith = upstream.RedactPayloadsWith
	RedactPayloadValue = upstream.RedactPayloadValue
	RedactSecrets      = upstream.RedactSecrets
	RedactValue        = upstream.RedactValue
	RedactWith         = upstream.RedactWith
	Rules              = upstream.Rules
)
