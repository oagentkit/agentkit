// 本文件是已迁移包的别名转发表，手工维护。
//
// 上游新增导出符号时必须在此补上对应别名，否则下游按门面 import 会编译失败。
//
// Package worker 是已迁移包的兼容门面。
//
// 新代码请直接 import github.com/oagentkit/queue/worker。
// 本仓将在 v1.x 后退役，见 MIGRATION.md。

package worker

import upstream "github.com/oagentkit/queue/worker"

type (
	LeaderElector = upstream.LeaderElector
	LeaderOption  = upstream.LeaderOption
	LeaseStore    = upstream.LeaseStore
	Pool          = upstream.Pool
	TaskQueue     = upstream.TaskQueue
)

const (
	HeartbeatInterval = upstream.HeartbeatInterval
	StaleRunningAfter = upstream.StaleRunningAfter
)

var (
	NewLeaderElector = upstream.NewLeaderElector
	WithLeaderLogger = upstream.WithLeaderLogger
	WithOnGained     = upstream.WithOnGained
	WithOnLost       = upstream.WithOnLost
)
