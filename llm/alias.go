// 本文件是已迁移包的别名转发表，手工维护。
//
// 上游新增导出符号时必须在此补上对应别名，否则下游按门面 import 会编译失败。
//
// 上游把符号标记 Deprecated 时，门面**仍照实转发**（个别行带 //nolint:staticcheck）：
// 门面是兼容层，移除符号等于对下游做破坏性变更；弃用信息透传给下游，
// 要不要迁移由下游自己决定。
//
// Package llm 是已迁移包的兼容门面。
//
// 新代码请直接 import github.com/oagentkit/llm。
// 本仓将在 v1.x 后退役，见 MIGRATION.md。

package llm

import upstream "github.com/oagentkit/llm"

type (
	Attempt              = upstream.Attempt
	AttemptError         = upstream.AttemptError
	Budget               = upstream.Budget
	BudgetConfig         = upstream.BudgetConfig
	BudgetHolder         = upstream.BudgetHolder
	BudgetInjector       = upstream.BudgetInjector
	ChainLink            = upstream.ChainLink
	Client               = upstream.Client
	CostRecord           = upstream.CostRecord
	CostTracker          = upstream.CostTracker
	FailoverModel        = upstream.FailoverModel
	FallbackChain        = upstream.FallbackChain
	Generator            = upstream.Generator
	MatchKind            = upstream.MatchKind
	ModelSummary         = upstream.ModelSummary
	OpenAIProvider       = upstream.OpenAIProvider
	OpenAIProviderConfig = upstream.OpenAIProviderConfig
	Price                = upstream.Price
	PriceLookup          = upstream.PriceLookup
	PriceStatus          = upstream.PriceStatus
	PriceTable           = upstream.PriceTable
	PricedProvider       = upstream.PricedProvider
	Provider             = upstream.Provider
	Resilient            = upstream.Resilient
	ResilientConfig      = upstream.ResilientConfig
	RetryHint            = upstream.RetryHint
	Sink                 = upstream.Sink
	StageRouter          = upstream.StageRouter
	TokenAccountant      = upstream.TokenAccountant
	UsageCallCounter     = upstream.UsageCallCounter
	UsageRecord          = upstream.UsageRecord
)

const (
	MatchExact               MatchKind   = upstream.MatchExact
	MatchNone                MatchKind   = upstream.MatchNone
	MatchNormalized          MatchKind   = upstream.MatchNormalized
	MaxServerRetryAfter                  = upstream.MaxServerRetryAfter
	MaxStreamRecoveryRetries             = upstream.MaxStreamRecoveryRetries
	PricePriced              PriceStatus = upstream.PricePriced
	PriceUnpriced            PriceStatus = upstream.PriceUnpriced
)

var (
	BusyAdmissionDelay      = upstream.BusyAdmissionDelay
	BusyAdmissionDelays     = upstream.BusyAdmissionDelays
	ClassifyLLMError        = upstream.ClassifyLLMError
	DefaultPriceTable       = upstream.DefaultPriceTable
	ErrPriceUnknown         = upstream.ErrPriceUnknown
	IsContextExceededError  = upstream.IsContextExceededError
	IsContextExceededFinish = upstream.IsContextExceededFinish
	IsProviderBusyCode      = upstream.IsProviderBusyCode
	IsRateLimitError        = upstream.IsRateLimitError
	IsTransientStreamError  = upstream.IsTransientStreamError
	IsTruncatedError        = upstream.IsTruncatedError
	LookupPrice             = upstream.LookupPrice
	LookupPriceStrict       = upstream.LookupPriceStrict
	NewBudget               = upstream.NewBudget
	NewClient               = upstream.NewClient
	NewCostTracker          = upstream.NewCostTracker
	NewCostTrackerWithTable = upstream.NewCostTrackerWithTable
	NewFailoverModel        = upstream.NewFailoverModel
	NewFallbackChain        = upstream.NewFallbackChain
	NewOpenAIProvider       = upstream.NewOpenAIProvider
	NewPriceTable           = upstream.NewPriceTable
	NewResilient            = upstream.NewResilient
	NewStageRouter          = upstream.NewStageRouter
	NewUsageHandler         = upstream.NewUsageHandler
	PriceOf                 = upstream.PriceOf        //nolint:staticcheck // 上游已弃用，门面不移除
	PriceOfMatched          = upstream.PriceOfMatched //nolint:staticcheck // 同上
	RetryAfterFrom          = upstream.RetryAfterFrom
	RetryableLLMError       = upstream.RetryableLLMError
	SelectRetryDelay        = upstream.SelectRetryDelay
	SuspiciousEmptyResult   = upstream.SuspiciousEmptyResult
	WithCallCounter         = upstream.WithCallCounter
	WithRetryAfterSink      = upstream.WithRetryAfterSink
	WithUsageLabels         = upstream.WithUsageLabels
)
