package receiver

import (
	"falcon-plus/modules/gateway/receiver/rpc"
	"falcon-plus/modules/gateway/receiver/socket"
)

func Start() {
	go rpc.StartRpc()
	go socket.StartSocket()
}
