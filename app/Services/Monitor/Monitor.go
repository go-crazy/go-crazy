// 监控记录
package Monitor

import (
	"go-crazy/util/logger"
	"os"
)

type Monitor struct {
	// 运行次数
	RunningTimes int
	// 编译失败次数
	CompileFailTimes int
	// 运行失败次数
	RunFailTimes int
	// 最大失败次数
	MaxFailTimes int
	//版本
	Version string
}

var (
	monitor = Monitor{
		RunningTimes:     0,
		CompileFailTimes: 0,
		RunFailTimes:     0,
		MaxFailTimes:     100,
		Version:          "2.0",
	}
)

func (this *Monitor) AddRunning() {
	this.RunningTimes += 1
}
func (this *Monitor) AddCompileFail() {
	this.CompileFailTimes += 1
}
func (this *Monitor) AddRunFail() {
	this.RunFailTimes += 1
	this.TryRestart()
}

func (this *Monitor) TryRestart() {
	// 到达最大错误个数触发系统退出 -> 重启
	if this.RunFailTimes > this.MaxFailTimes {
		pid := os.Getpid()
		ps, _ := os.FindProcess(pid)
		ps.Signal(os.Interrupt)
		logger.Info("到达最大错误个数触发系统退出 -- 执行重启")
	}
}

func AddRunning() {
	monitor.AddRunning()
}

func AddCompileFail() {
	monitor.AddCompileFail()
}

func AddRunFail() {
	monitor.AddRunFail()
}

func GetMonitorInfo() *Monitor {
	return &monitor
}
