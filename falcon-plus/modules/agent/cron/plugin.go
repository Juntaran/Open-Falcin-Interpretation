package cron

import (
	"falcon-plus/common/model"
	"falcon-plus/modules/agent/g"
	"falcon-plus/modules/agent/plugins"
	"log"
	"strings"
	"time"
)

func SyncMinePlugins() {
	// 判断插件机制是否开启
	if !g.Config().Plugin.Enabled {
		return
	}
	// 判断心跳机制是否开启
	if !g.Config().Heartbeat.Enabled {
		return
	}
	// 判断心跳地址是否填写
	if g.Config().Heartbeat.Addr == "" {
		return
	}

	// 如果以上都没问题，就开始同步插件
	go syncMinePlugins()
}

func syncMinePlugins() {

	var (
		timestamp  int64 = -1
		pluginDirs []string
	)

	duration := time.Duration(g.Config().Heartbeat.Interval) * time.Second

	// 同样，也是每隔duration执行一次，duration为心跳设置的间隔
	for {
		time.Sleep(duration)

		hostname, err := g.Hostname()
		if err != nil {
			continue
		}

		req := model.AgentHeartbeatRequest{
			Hostname: hostname,
		}

		var resp model.AgentPluginsResponse
		// 把hostname放在req里作为Call()的参数
		// Call()告诉这个hostname应该执行哪些plugin
		// 返回来的信息放在resp里
		err = g.HbsClient.Call("Agent.MinePlugins", req, &resp)
		if err != nil {
			log.Println("ERROR:", err)
			continue
		}

		// 时间戳应该在上次的时间戳之后
		if resp.Timestamp <= timestamp {
			continue
		}

		pluginDirs = resp.Plugins
		timestamp = resp.Timestamp

		if g.Config().Debug {
			log.Println(&resp)
		}

		if len(pluginDirs) == 0 {
			plugins.ClearAllPlugins()
		}

		desiredAll := make(map[string]*plugins.Plugin)

		for _, p := range pluginDirs {
			underOneDir := plugins.ListPlugins(strings.Trim(p, "/"))
			for k, v := range underOneDir {
				desiredAll[k] = v
			}
		}

		// 删掉没用的plugin
		plugins.DelNoUsePlugins(desiredAll)
		// 添加执行新的plugin
		plugins.AddNewPlugins(desiredAll)

	}
}
