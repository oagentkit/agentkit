// 本文件是**新包**的别名转发表，手工维护。
//
// 上游新增导出符号时必须在此补上对应别名，否则下游按门面 import 会编译失败；
// 反过来，删掉这里的转发会让下游编译断裂，而矩阵的**符号级转发完整性门禁**
// 也会判成「幽灵转发」——两者是同一个契约的两面。
//
// Package port 是已迁移包的兼容门面。
//
// 新代码请直接 import github.com/oagentkit/llm/port。
// 本仓**已废弃**：v1.x 退役条件已满足，只保留兼容转发，不接受任何新功能。
// 新代码请直接 import 上游模块（见 MIGRATION.md 的映射表）。
//
// ⚠️ 转发**已弃用**的上游符号时按行豁免 SA1019：staticcheck 会把「转发」判成
// 「使用已弃用 API」，但门面按契约必须转发上游的全部导出符号（移除=对下游的
// 破坏性变更）。豁免的是「转发」这个动作，不是「上游已弃用」这个事实。
//
// 关联：oagentkit/llm#20（port 包）、oagentkit/llm#11（窄接口层需求）

package port

import upstream "github.com/oagentkit/llm/port"

type (
	Breakers        = upstream.Breakers
	Budget          = upstream.Budget
	Chain           = upstream.Chain
	ChainOption     = upstream.ChainOption
	Classifier      = upstream.Classifier
	Completer       = upstream.Completer
	CostRecord      = upstream.CostRecord
	CostTracker     = upstream.CostTracker
	Entry           = upstream.Entry
	MatchKind       = upstream.MatchKind
	ModelSummary    = upstream.ModelSummary
	Price           = upstream.Price
	PriceLookup     = upstream.PriceLookup
	PriceStatus     = upstream.PriceStatus
	PriceTable      = upstream.PriceTable
	Request         = upstream.Request
	StatusError     = upstream.StatusError
	TokenAccountant = upstream.TokenAccountant
	TokenCounter    = upstream.TokenCounter
	Usage           = upstream.Usage
)

const (
	DefaultBreakerCooldown = upstream.DefaultBreakerCooldown
	DefaultBreakerTrip     = upstream.DefaultBreakerTrip
	MatchExact             = upstream.MatchExact
	MatchNone              = upstream.MatchNone
	MatchNormalized        = upstream.MatchNormalized
	PricePriced            = upstream.PricePriced
	PriceUnpriced          = upstream.PriceUnpriced
)

var (
	ErrBudgetExhausted      = upstream.ErrBudgetExhausted
	ErrCircuitOpen          = upstream.ErrCircuitOpen
	ErrEmptyChain           = upstream.ErrEmptyChain
	ErrPriceUnknown         = upstream.ErrPriceUnknown
	ClassifyStatus          = upstream.ClassifyStatus
	DefaultPriceTable       = upstream.DefaultPriceTable
	LookupPrice             = upstream.LookupPrice
	LookupPriceStrict       = upstream.LookupPriceStrict
	NewBreakers             = upstream.NewBreakers
	NewBudget               = upstream.NewBudget
	NewChain                = upstream.NewChain
	NewCostTracker          = upstream.NewCostTracker
	NewCostTrackerWithTable = upstream.NewCostTrackerWithTable
	NewPriceTable           = upstream.NewPriceTable
	PriceOf                 = upstream.PriceOf        //nolint:staticcheck // 上游已弃用：改用 PriceTable.Lookup
	PriceOfMatched          = upstream.PriceOfMatched //nolint:staticcheck // 上游已弃用：改用 PriceTable.Lookup
	WithBackoff             = upstream.WithBackoff
	WithBreaker             = upstream.WithBreaker
	WithBudget              = upstream.WithBudget
	WithClassifier          = upstream.WithClassifier
	WithCompleter           = upstream.WithCompleter
	WithRetries             = upstream.WithRetries
	WithTokenCounter        = upstream.WithTokenCounter
	WithTracker             = upstream.WithTracker
)
