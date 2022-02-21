package main

import (
	"jsapi-websocket/routers"
	"jsapi-websocket/servers"
)

func main() {
	wsServer, err := servers.New()
	if err != nil {
		println("Não foi possível iniciar o servidor." + err.Error())
		return
	}
	routers.ApplyAccountRoutes(wsServer)
	wsServer.Start()
}