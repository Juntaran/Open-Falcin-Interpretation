package main

import (
	"flag"
	"fmt"
	"os"

	"falcon-plus/modules/nodata/collector"
	"falcon-plus/modules/nodata/config"
	"falcon-plus/modules/nodata/g"
	"falcon-plus/modules/nodata/http"
	"falcon-plus/modules/nodata/judge"
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
	g.StartProc()

	// config
	config.Start()
	// collector
	collector.Start()
	// judge
	judge.Start()

	// http
	http.Start()

	select {}
}
