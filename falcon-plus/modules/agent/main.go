package main

import (
	"flag"
	"fmt"
	"falcon-plus/modules/agent/cron"
	"falcon-plus/modules/agent/funcs"
	"falcon-plus/modules/agent/g"
	"falcon-plus/modules/agent/http"
	"os"
)

func main() {

	// 读参数
	cfg := flag.String("c", "cfg.json", "configuration file")
	version := flag.Bool("v", false, "show version")
	check := flag.Bool("check", false, "check collector")

	flag.Parse()

	if *version {
		fmt.Println(g.VERSION)
		os.Exit(0)
	}

	if *check {
		funcs.CheckCollector()
		os.Exit(0)
	}

	g.ParseConfig(*cfg)         // 解析参数

	// 初始化全局变量
	if g.Config().Debug {			// 设置日志级别
		g.InitLog("debug")
	} else {
		g.InitLog("info")
	}

	g.InitRootDir()
	g.InitLocalIp()
	g.InitRpcClients()

	funcs.BuildMappers()        // funcs里是采集信息函数

	go cron.InitDataHistory()

	// 与心跳相关以下四行
	cron.ReportAgentStatus()    // 汇报agent状态
	cron.SyncMinePlugins()      // 同步我要执行的插件
	cron.SyncBuiltinMetrics()   // 同步metric
	cron.SyncTrustableIps()     // 获取内置ip

	// 收集信息
	cron.Collect()

	// 默认1988端口
	go http.Start()

	select {}

}
