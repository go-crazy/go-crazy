// 流量控制 利用通道技术快速实现
// 其中运行消耗1，编译消耗4
package FlowControl

import (
	"go-crazy/util/logger"
	"strconv"
	"sync"
	"sync/atomic"

	"go-crazy/Config"
)

type FlowControl struct {
	// 最大并发
	MaxConcurrency int32     // 默认 50
	queue          chan bool // 队列
	count          int32     // 当前队列长度
}

// 定义变量
var (
	flow = FlowControl{
		MaxConcurrency: 200,
		queue:          make(chan bool, 200),
		count:          0,
	}
	mutex sync.Mutex

	CompileConsume = 10
)

// 初始化队列
func Init() {
	if Config.MaxConcurrency > 0 {
		flow.MaxConcurrency = int32(Config.MaxConcurrency)
		flow.queue = make(chan bool, Config.MaxConcurrency)
		logger.Info("最大并发数 = " + strconv.Itoa(Config.MaxConcurrency))
	}
}

// 增加n个
func Add(sum int) {
	logger.Info("进入执行 001" + strconv.Itoa(int(flow.count)))
	mutex.Lock()
	if Config.OpenFlowControl {
		if (flow.count + 4) >= flow.MaxConcurrency {
			logger.Info("进入队列并等待！")
		}
		for index := 0; index < sum; index++ {
			flow.queue <- true
			atomic.AddInt32(&flow.count, 1)
		}
	}
	mutex.Unlock()
	logger.Info("进入执行")
}

// 释放n个
func Del(sum int) {
	if Config.OpenFlowControl {
		for index := 0; index < sum; index++ {
			<-flow.queue
			atomic.AddInt32(&flow.count, -1)
		}
	}
}
