package router

import upstream "github.com/oagentkit/agent/router"

// SetOnNegationHit 注册上游 router 的否定命中钩子（agentkit#155）。
//
// ⚠️ 经门面注册钩子**必须走本 setter**，不要给 OnNegationHit var 赋值：
// 门面的 var 别名是 init 时的**值拷贝**（上游初始值，通常为 nil），对它赋值
// 只改写门面自己的副本，上游 keyword.go 运行时读的是上游包的 var——钩子
// 静默失效（观测面注册如 metrics 计数，丢了没有任何报错）。本函数直接写
// **上游包的 var**（跨包写导出变量合法），上游运行时读得到。
//
// 传 nil 注销钩子（与上游「可选钩子」语义一致）。
func SetOnNegationHit(f func()) { upstream.OnNegationHit = f }
