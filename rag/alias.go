// 本文件是已迁移包的别名转发表，手工维护。
//
// 上游新增导出符号时必须在此补上对应别名，否则下游按门面 import 会编译失败。
//
// Package rag 是已迁移包的兼容门面。
//
// 新代码请直接 import github.com/oagentkit/rag。
// 本仓将在 v1.x 后退役，见 MIGRATION.md。

package rag

import upstream "github.com/oagentkit/rag"

type (
	Chunk            = upstream.Chunk
	Embedder         = upstream.Embedder
	EmbedderOption   = upstream.EmbedderOption
	Filter           = upstream.Filter
	Hybrid           = upstream.Hybrid
	KnowledgeService = upstream.KnowledgeService
	Local            = upstream.Local
	MilvusConfig     = upstream.MilvusConfig
	MilvusStore      = upstream.MilvusStore
	Noop             = upstream.Noop
	OpenAIEmbedder   = upstream.OpenAIEmbedder
)

var (
	ChunkPreview      = upstream.ChunkPreview
	FuseRRF           = upstream.FuseRRF
	NewHybrid         = upstream.NewHybrid
	NewLocal          = upstream.NewLocal
	NewMilvusStore    = upstream.NewMilvusStore
	NewNoop           = upstream.NewNoop
	NewOpenAIEmbedder = upstream.NewOpenAIEmbedder
	NewSearchTool     = upstream.NewSearchTool
	WithHTTPClient    = upstream.WithHTTPClient
	WithRetry         = upstream.WithRetry
	WithTimeout       = upstream.WithTimeout
)
