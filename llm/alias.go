// 本文件是已迁移包的别名转发表，手工维护。
//
// 上游新增导出符号时必须在此补上对应别名，否则下游按门面 import 会编译失败。
//
// Package llm 是已迁移包的兼容门面。
//
// 新代码请直接 import github.com/oagentkit/llm。
// 本仓**已废弃**：v1.x 退役条件已满足，只保留兼容转发，不接受任何新功能。
// 新代码请直接 import 上游模块（见 MIGRATION.md 的映射表）。

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
	MatchExact               = upstream.MatchExact
	MatchNone                = upstream.MatchNone
	MatchNormalized          = upstream.MatchNormalized
	MaxServerRetryAfter      = upstream.MaxServerRetryAfter
	MaxStreamRecoveryRetries = upstream.MaxStreamRecoveryRetries
	PricePriced              = upstream.PricePriced
	PriceUnpriced            = upstream.PriceUnpriced
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
	// 门面按契约转发上游的**全部**导出符号；上游标 Deprecated 的符号在此仍必须转发
	// （删掉会让走门面的下游编译断裂，而符号级门禁也会判成「幽灵转发」）。
	// 故此处按行豁免 SA1019——豁免的是「转发」这个动作，不是「上游已弃用」这个事实。
	PriceOf               = upstream.PriceOf        //nolint:staticcheck // 上游已弃用：改用 PriceTable.Lookup
	PriceOfMatched        = upstream.PriceOfMatched //nolint:staticcheck // 上游已弃用：改用 PriceTable.Lookup
	RetryAfterFrom        = upstream.RetryAfterFrom
	RetryableLLMError     = upstream.RetryableLLMError
	SelectRetryDelay      = upstream.SelectRetryDelay
	SuspiciousEmptyResult = upstream.SuspiciousEmptyResult
	WithCallCounter       = upstream.WithCallCounter
	WithRetryAfterSink    = upstream.WithRetryAfterSink
	WithUsageLabels       = upstream.WithUsageLabels
)
