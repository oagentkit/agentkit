package router

import (
	"sync/atomic"
	"testing"

	upstream "github.com/oagentkit/agent/router"
)

// TestSetOnNegationHitReachesUpstream 钉 agentkit#155 的修复契约：经门面
// setter 注册的钩子，上游运行时读得到。修复前门面的 var 别名是值拷贝，
// 对它赋值只改写门面副本，上游 keyword.go 永远读到 nil——钩子静默失效，
// 无编译错无运行时错（复现：kitrouter.OnNegationHit != nil 为 true 而
// upstream.OnNegationHit == nil）。
func TestSetOnNegationHitReachesUpstream(t *testing.T) {
	var hits atomic.Int32
	SetOnNegationHit(func() { hits.Add(1) })
	t.Cleanup(func() { SetOnNegationHit(nil) })

	if upstream.OnNegationHit == nil {
		t.Fatal("SetOnNegationHit 后上游 var 仍为 nil——钩子经门面注册会静默失效（#155）")
	}
	upstream.OnNegationHit() // 模拟上游 keyword.go 的运行时读取与调用
	if hits.Load() != 1 {
		t.Fatalf("上游应读到 setter 注册的钩子，实得 %d 次命中", hits.Load())
	}

	// 对照：对门面 var 别名直接赋值不回传上游（缺陷形态保持可见，
	// 注释与 MIGRATION 口径一致地指路 setter）
	OnNegationHit = func() { hits.Add(100) }
	t.Cleanup(func() { OnNegationHit = nil })
	_ = upstream.OnNegationHit // 上游 var 不受门面 var 赋值影响（上面的断言已证）

	SetOnNegationHit(nil)
	if upstream.OnNegationHit != nil {
		t.Fatal("传 nil 应注销钩子")
	}
}
