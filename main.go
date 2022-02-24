package main

import (
	"github.com/jonathanschmittblog/jsapi-websocket/routers"
	"github.com/jonathanschmittblog/jsapi-websocket/servers"
)

func main() {
	wsServer, err := servers.New()
	if err != nil {
		println("Não foi possível iniciar o servidor." + err.Error())
		return
	}
	routers.ApplyRoutes(wsServer)
	wsServer.Start()
}