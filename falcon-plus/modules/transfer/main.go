package main

import (
	"flag"
	"fmt"
	"falcon-plus/modules/transfer/g"
	"falcon-plus/modules/transfer/http"
	"falcon-plus/modules/transfer/proc"
	"falcon-plus/modules/transfer/receiver"
	"falcon-plus/modules/transfer/sender"
	"os"
)

func main() {
	cfg := flag.String("c", "cfg.json", "configuration file")
	version := flag.Bool("v", false, "show version")
	versionGit := flag.Bool("vg", false, "show version")
	flag.Parse()

	if *version {
		fmt.Println(g.VERSION)
		os.Exit(0)
	}
	if *versionGit {
		fmt.Println(g.VERSION, g.COMMIT)
		os.Exit(0)
	}

	// global config
	g.ParseConfig(*cfg)
	// proc
	proc.Start()		// 这里只在log中记录了proc.Start()

	sender.Start()		// 初始化数据发送服务
	receiver.Start()	// 接收服务，包括RPC和Socket

	// http
	http.Start()

	select {}
}
