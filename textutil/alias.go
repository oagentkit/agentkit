// 本文件是已迁移包的别名转发表，手工维护。
//
// 上游新增导出符号时必须在此补上对应别名，否则下游按门面 import 会编译失败。
//
// Package textutil 是已迁移包的兼容门面。
//
// 新代码请直接 import github.com/oagentkit/textx/textutil。
// 本仓**已废弃**：v1.x 退役条件已满足，只保留兼容转发，不接受任何新功能。
// 新代码请直接 import 上游模块（见 MIGRATION.md 的映射表）。

package textutil

import upstream "github.com/oagentkit/textx/textutil"

type (
	CJKRatio      = upstream.CJKRatio
	TokenRounding = upstream.TokenRounding
)

const (
	FitInputNote     = upstream.FitInputNote
	RoundingSegment  = upstream.RoundingSegment
	RoundingTotal    = upstream.RoundingTotal
	TokenCharDivisor = upstream.TokenCharDivisor
)

var (
	BigramSet                  = upstream.BigramSet
	ChunkLines                 = upstream.ChunkLines
	ClampRuneBoundary          = upstream.ClampRuneBoundary
	DefaultCJKRatio            = upstream.DefaultCJKRatio
	EstTokens                  = upstream.EstTokens
	EstTokensCJK               = upstream.EstTokensCJK
	EstTokensCJKRunes          = upstream.EstTokensCJKRunes
	EstTokensCJKRunesWithRatio = upstream.EstTokensCJKRunesWithRatio
	EstTokensOf                = upstream.EstTokensOf
	EstTokensWithRatio         = upstream.EstTokensWithRatio
	EstTokensWithRounding      = upstream.EstTokensWithRounding
	FTS5Query                  = upstream.FTS5Query
	FirstNonEmpty              = upstream.FirstNonEmpty
	FitCharsLimit              = upstream.FitCharsLimit
	GlobMatch                  = upstream.GlobMatch
	Jaccard                    = upstream.Jaccard
	NearDuplicate              = upstream.NearDuplicate
	NumberLines                = upstream.NumberLines
	RuneUnitsCJK               = upstream.RuneUnitsCJK
	SanitizeFileStem           = upstream.SanitizeFileStem
	Similarity                 = upstream.Similarity
	SplitRunes                 = upstream.SplitRunes
	StripCodeFence             = upstream.StripCodeFence
	Tokens                     = upstream.Tokens
	TokensJoined               = upstream.TokensJoined
	TruncBytes                 = upstream.TruncBytes
	TruncEllipsis              = upstream.TruncEllipsis
	TruncNote                  = upstream.TruncNote
	TruncRunes                 = upstream.TruncRunes
	UnitsToTokensCJK           = upstream.UnitsToTokensCJK
)
