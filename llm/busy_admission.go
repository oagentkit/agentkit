package llm

import (
	"errors"
	"time"

	upstream "github.com/oagentkit/llm"
)

// ErrEmptyBusyAdmissionDelays 空退避表被 setter 拒绝：上游 BusyAdmissionDelay
// 对空表会算出下标 -1（`retry >= len(t)` ⇒ `retry = len(t)-1`）直接 panic——
// 宁可在 setter 层报配置错，不给上游埋雷。
var ErrEmptyBusyAdmissionDelays = errors.New("agentkit/llm: BusyAdmissionDelays 不能为空（上游按序号取档，空表会越界 panic）")

// SetBusyAdmissionDelays 设置上游 llm 的 busy 准入退避表（agentkit#156）。
//
// ⚠️ 经门面调优退避**必须走本 setter**，不要 reassign / append 下面的
// BusyAdmissionDelays var：门面的 var 别名是 init 时的**切片头拷贝**——
// reassign 或 append（重分配后）只改写门面副本，上游 recovery.go 的
// BusyAdmissionDelay 访问器运行时读的是上游包的 var，调优被静默无视
// （provider busy 的重试节奏悄悄回到默认 1s/2s）；而未 append 前的索引写
// 恰因共享底层数组能传播——部分写法有效、部分写法无效，最容易误判正常。
// 本函数直接写**上游包的 var**，并做防御性拷贝（调用方此后对入参切片的
// 任何改动都不再影响上游，重试节奏一经设定即稳定）。
//
// 空表拒绝（ErrEmptyBusyAdmissionDelays）。
func SetBusyAdmissionDelays(d []time.Duration) error {
	if len(d) == 0 {
		return ErrEmptyBusyAdmissionDelays
	}
	upstream.BusyAdmissionDelays = append([]time.Duration(nil), d...)
	return nil
}
