falcon-agent
===

This is a linux monitor agent. Just like zabbix-agent and tcollector.


## Installation

It is a golang classic project

```bash
# set $GOPATH and $GOROOT
mkdir -p $GOPATH/src/github.com/open-falcon
cd $GOPATH/src/github.com/open-falcon
git clone https://github.com/open-falcon/falcon-plus.git
cd falcon-plus/modules/agent
go get
./control build
./control start

# goto http://localhost:1988
```

I use [linux-dash](https://github.com/afaqurk/linux-dash) as the page theme.

## Configuration

- heartbeat: heartbeat server rpc address
- transfer: transfer rpc address
- ignore: the metrics should ignore

# Auto deployment

Just look at https://github.com/open-falcon/ops-updater

# Agent结构

| 目录  | 作用  |  
|---|---|  
| cron  | 每隔一段时间要执行的代码都在这里  |
| func  | 采集信息函数  |
| g  | global目录，全局都会用到的东西  |
| http  | 一个http的server，获取单机指标的一些数值  |
| plugins  | 插件  |
| public  | 静态资源文件  |
