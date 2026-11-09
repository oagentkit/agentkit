package router

import (
	"sync/atomic"
	"testing"

	upstream "github.com/oagentkit/agent/router"
)

// TestSetOnNegationHitReachesUpstream 钉 agentkit#155/#168 的修复契约：经门面
// setter 注册的钩子，上游热路径（KeywordHit 的否定作废点）真的调用得到。
// 上游 agent v0.14.8 起钩子存 atomic.Pointer 快照——门面 SetOnNegationHit
// 是上游 setter 的纯转发；OnNegationHit var 别名是 init 值拷贝（Deprecated），
// 对它赋值不回传上游，测试用「读快照」口径验证生效面。
func TestSetOnNegationHitReachesUpstream(t *testing.T) {
	var hits atomic.Int32
	SetOnNegationHit(func() { hits.Add(1) })
	t.Cleanup(func() { SetOnNegationHit(nil) })

	upstream.KeywordHit("没有磁盘问题", "磁盘") // 否定作废点触发钩子
	if hits.Load() != 1 {
		t.Fatalf("上游否定作废应触发经门面注册的钩子 1 次，实得 %d", hits.Load())
	}

	// 注销后不再触发
	SetOnNegationHit(nil)
	upstream.KeywordHit("没有磁盘问题", "磁盘")
	if hits.Load() != 1 {
		t.Fatalf("注销后不应再触发，实得 %d", hits.Load())
	}

	// 对照（死写形态保持可见）：对门面 var 别名赋值不改变生效快照
	OnNegationHit = func() { hits.Add(100) }
	t.Cleanup(func() { OnNegationHit = nil })
	SetOnNegationHit(nil)
	upstream.KeywordHit("没有磁盘问题", "磁盘")
	if hits.Load() != 1 {
		t.Fatalf("经 var 别名赋值不应影响生效快照（死写形态），实得 %d", hits.Load())
	}
}
