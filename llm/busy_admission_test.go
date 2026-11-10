package llm

import (
	"errors"
	"testing"
	"time"

	upstream "github.com/oagentkit/llm"
)

// TestSetBusyAdmissionDelaysReachesUpstream 钉 agentkit#156/#166/#169 的契约：
// 经门面 setter 调优的退避表，上游访问器（原子快照）读得到；门面 var 别名
// （Deprecated）读写都不反映生效表。
func TestSetBusyAdmissionDelaysReachesUpstream(t *testing.T) {
	t.Cleanup(func() { _ = upstream.SetBusyAdmissionDelays([]time.Duration{time.Second, 2 * time.Second}) })

	if err := SetBusyAdmissionDelays([]time.Duration{
		111 * time.Millisecond, 7 * time.Second, 9 * time.Second,
	}); err != nil {
		t.Fatalf("SetBusyAdmissionDelays: %v", err)
	}
	if got := upstream.BusyAdmissionDelay(0); got != 111*time.Millisecond {
		t.Fatalf("上游访问器应读到新表第 0 档，实得 %v——调优经门面被无视（#156）", got)
	}
	if got := upstream.BusyAdmissionDelay(5); got != 9*time.Second {
		t.Fatalf("越界应取最后一档（issue 复现点：第三档曾被无视），实得 %v", got)
	}

	// 防御性拷贝：setter 之后的入参改动不得影响生效快照（重试节奏一经设定即稳定）
	tbl := []time.Duration{time.Second, 2 * time.Second}
	if err := SetBusyAdmissionDelays(tbl); err != nil {
		t.Fatalf("SetBusyAdmissionDelays: %v", err)
	}
	tbl[0] = 42 * time.Millisecond
	if got := upstream.BusyAdmissionDelay(0); got != time.Second {
		t.Fatalf("setter 应防御性拷贝入参——共享底层数组会把「部分写法能传播」的坑带回来（#156），实得 %v", got)
	}

	// 空表拒绝（上游对空表取下标 -1 panic）
	if err := SetBusyAdmissionDelays(nil); !errors.Is(err, ErrEmptyBusyAdmissionDelays) {
		t.Fatalf("空表应报 ErrEmptyBusyAdmissionDelays，实得 %v", err)
	}
	if err := SetBusyAdmissionDelays([]time.Duration{}); !errors.Is(err, ErrEmptyBusyAdmissionDelays) {
		t.Fatalf("空表应报 ErrEmptyBusyAdmissionDelays，实得 %v", err)
	}

	// 门面 var 别名（Deprecated）读数停在初始化缺省，不反映 setter 后的表（#169）
	if err := SetBusyAdmissionDelays([]time.Duration{3 * time.Second, 7 * time.Second, 11 * time.Second}); err != nil {
		t.Fatalf("SetBusyAdmissionDelays: %v", err)
	}
	if len(BusyAdmissionDelays) == 3 {
		t.Fatalf("门面 var 别名不应反映 setter 后的表（别名是 init 拷贝；此断言失败说明别名与快照又耦合了）——len=%d", len(BusyAdmissionDelays))
	}
	if got := upstream.BusyAdmissionDelay(0); got != 3*time.Second {
		t.Fatalf("快照应读到 3 档表第 0 档，实得 %v", got)
	}
}
