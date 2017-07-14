package cron

import (
	"falcon-plus/common/model"
	"falcon-plus/modules/agent/g"
	"log"
	"strconv"
	"strings"
	"time"
)

func SyncBuiltinMetrics() {
	if g.Config().Heartbeat.Enabled && g.Config().Heartbeat.Addr != "" {
		go syncBuiltinMetrics()
	}
}

func syncBuiltinMetrics() {

	var timestamp int64 = -1
	var checksum string = "nil"

	duration := time.Duration(g.Config().Heartbeat.Interval) * time.Second

	for {
		time.Sleep(duration)

		var ports = []int64{}			// 监听哪些端口
		var paths = []string{}			// 看目录大小变化
		var procs = make(map[string]map[int]string)	// 监听哪些进程
		var urls = make(map[string]string)	// 监听url是否正常

		hostname, err := g.Hostname()
		if err != nil {
			continue
		}

		req := model.AgentHeartbeatRequest{
			Hostname: hostname,
			Checksum: checksum,
		}

		var resp model.BuiltinMetricResponse
		err = g.HbsClient.Call("Agent.BuiltinMetrics", req, &resp)
		if err != nil {
			log.Println("ERROR:", err)
			continue
		}

		if resp.Timestamp <= timestamp {
			continue
		}

		if resp.Checksum == checksum {
			continue
		}

		timestamp = resp.Timestamp
		checksum = resp.Checksum

		for _, metric := range resp.Metrics {

			// 查看URL是否正常
			if metric.Metric == g.URL_CHECK_HEALTH {
				arr := strings.Split(metric.Tags, ",")
				if len(arr) != 2 {
					continue
				}
				url := strings.Split(arr[0], "=")
				if len(url) != 2 {
					continue
				}
				stime := strings.Split(arr[1], "=")
				if len(stime) != 2 {
					continue
				}
				if _, err := strconv.ParseInt(stime[1], 10, 64); err == nil {
					urls[url[1]] = stime[1]
				} else {
					log.Println("metric ParseInt timeout failed:", err)
				}
			}

			// 监听端口
			if metric.Metric == g.NET_PORT_LISTEN {
				arr := strings.Split(metric.Tags, "=")
				if len(arr) != 2 {
					continue
				}

				if port, err := strconv.ParseInt(arr[1], 10, 64); err == nil {
					ports = append(ports, port)
				} else {
					log.Println("metrics ParseInt failed:", err)
				}

				continue
			}

			// 获取某个目录的大小，看log的大小是否异常
			// metric: du.bs tags: path=/home/work/logs
			// 相当于执行 du -bs /home/work/logs
			if metric.Metric == g.DU_BS {
				arr := strings.Split(metric.Tags, "=")
				if len(arr) != 2 {
					continue
				}
				// arr[1]就相当于/home/work/logs
				paths = append(paths, strings.TrimSpace(arr[1]))
				continue
			}

			// metric: proc.num tags: cmdline=falcon-agent-cfg.json
			// 也可能是name=...
			// cmdline 和 name不能同时使用
			// 所有的java进程，进程名都是java，那我们是没法通过name字段做区分的。
			// 怎么办呢？此时就要求助于这个/proc/$pid/cmdline文件的内容了
			if metric.Metric == g.PROC_NUM {
				arr := strings.Split(metric.Tags, ",")

				tmpMap := make(map[int]string)

				for i := 0; i < len(arr); i++ {
					if strings.HasPrefix(arr[i], "name=") {
						tmpMap[1] = strings.TrimSpace(arr[i][5:])
					} else if strings.HasPrefix(arr[i], "cmdline=") {
						tmpMap[2] = strings.TrimSpace(arr[i][8:])
					}
				}

				procs[metric.Tags] = tmpMap
			}
		}

		g.SetReportUrls(urls)
		g.SetReportPorts(ports)
		g.SetReportProcs(procs)
		g.SetDuPaths(paths)

	}
}
