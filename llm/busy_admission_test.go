package llm

import (
	"errors"
	"testing"
	"time"

	upstream "github.com/oagentkit/llm"
)

// TestSetBusyAdmissionDelaysReachesUpstream 钉 agentkit#156 的修复契约：
// 经门面 setter 调优的退避表，上游 BusyAdmissionDelay 访问器读得到。
// 修复前门面的 var 别名是切片头拷贝，reassign/append 只改写门面副本，
// 上游重试路径读自己的 var——调优被静默无视（issue 复现：len(kit 表)=3 而
// len(upstream 表)=2，kit.BusyAdmissionDelay(5)=2s，7s 第三档被无视）。
func TestSetBusyAdmissionDelaysReachesUpstream(t *testing.T) {
	orig := upstream.BusyAdmissionDelays
	t.Cleanup(func() { upstream.BusyAdmissionDelays = orig })

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

	// 防御性拷贝：setter 之后的入参改动不得影响上游（重试节奏一经设定即稳定）
	tbl := []time.Duration{time.Second, 2 * time.Second}
	if err := SetBusyAdmissionDelays(tbl); err != nil {
		t.Fatalf("SetBusyAdmissionDelays: %v", err)
	}
	tbl[0] = 42 * time.Millisecond
	if upstream.BusyAdmissionDelays[0] != time.Second {
		t.Fatal("setter 应拷贝入参——共享底层数组会把「部分写法能传播」的坑带回来（#156）")
	}

	// 空表拒绝（上游对空表取下标 -1 panic）
	if err := SetBusyAdmissionDelays(nil); !errors.Is(err, ErrEmptyBusyAdmissionDelays) {
		t.Fatalf("空表应报 ErrEmptyBusyAdmissionDelays，实得 %v", err)
	}
	if err := SetBusyAdmissionDelays([]time.Duration{}); !errors.Is(err, ErrEmptyBusyAdmissionDelays) {
		t.Fatalf("空表应报 ErrEmptyBusyAdmissionDelays，实得 %v", err)
	}
}
