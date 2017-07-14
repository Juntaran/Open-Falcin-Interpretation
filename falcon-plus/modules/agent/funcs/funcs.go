package funcs

import (
	"falcon-plus/common/model"
	"falcon-plus/modules/agent/g"
)

// FuncsAndInterval 类型，包含一个func和一个Interval
// Fs 为一个匿名函数切片，该匿名函数的返回值类型为 *model.MetricValue 类型的切片
// Interval 为间隔时间，由cfg直接定义好的
type FuncsAndInterval struct {
	Fs       []func() []*model.MetricValue
	Interval int
}

var Mappers []FuncsAndInterval		// Mappers是一个slice，基本结构是 FuncsAndInterval 类型

func BuildMappers() {
	interval := g.Config().Transfer.Interval		// 在cfg中，Transfer的interval被配置成了60
	Mappers = []FuncsAndInterval{
		{
			// 这里为什么要分开啊。。。应该可以都放在一个Fs里的。。。

			Fs: []func() []*model.MetricValue{
				AgentMetrics,
				CpuMetrics,
				NetMetrics,
				KernelMetrics,
				LoadAvgMetrics,
				MemMetrics,
				DiskIOMetrics,
				IOStatsMetrics,
				NetstatMetrics,
				ProcMetrics,
				UdpMetrics,
			},
			Interval: interval,
		},
		{
			Fs: []func() []*model.MetricValue{
				DeviceMetrics,
			},
			Interval: interval,
		},
		{
			Fs: []func() []*model.MetricValue{
				PortMetrics,
				SocketStatSummaryMetrics,
			},
			Interval: interval,
		},
		{
			Fs: []func() []*model.MetricValue{
				DuMetrics,
			},
			Interval: interval,
		},
		{
			Fs: []func() []*model.MetricValue{
				UrlMetrics,
			},
			Interval: interval,
		},
	}
}
