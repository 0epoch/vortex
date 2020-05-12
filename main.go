package main

import (
	"vortex/im"
	"vortex/model"
	"vortex/pkg/config"
	"vortex/router"
)

func main() {
	config.Load()
	model.Load()
	r := router.Load()

	go im.NewWsConn().Run() //websocket服务
	r.Run("0.0.0.0:8888") // 监听并在 0.0.0.0:8080 上启动服务
}

