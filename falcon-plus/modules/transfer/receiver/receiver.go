package receiver

import (
	"falcon-plus/modules/transfer/receiver/rpc"
	"falcon-plus/modules/transfer/receiver/socket"
)

func Start() {
	go rpc.StartRpc()
	go socket.StartSocket()
}
