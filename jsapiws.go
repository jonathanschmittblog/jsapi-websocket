package main

import (
	"jsapi-websocket/routers"
	"jsapi-websocket/wsservers"
)

func main() {
	wsServer, err := wsservers.New()
	if err != nil {
		println("Não foi possível iniciar o servidor." + err.Error())
		return
	}
	routers.ApplyAccountRoutes(wsServer)
	wsServer.Start()
}