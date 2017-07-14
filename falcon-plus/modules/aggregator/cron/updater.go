package cron

import (
	"falcon-plus/modules/aggregator/db"
	"falcon-plus/modules/aggregator/g"
	"time"
)

func UpdateItems() {
	for {
		updateItems()
		d := time.Duration(g.Config().Database.Interval) * time.Second
		time.Sleep(d)
	}
}

func updateItems() {
	items, err := db.ReadClusterMonitorItems()
	if err != nil {
		return
	}

	deleteNoUseWorker(items)
	createWorkerIfNeed(items)
}
