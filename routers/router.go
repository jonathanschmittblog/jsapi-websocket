package routers

import (
	"jsapi-websocket/wsservers"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func ApplyAccountRoutes(wsServer *wsservers.Server) {
	wsServer.Router.GET("/messages", messages)
	// Configura o arquivo html como página inicial
	wsServer.Router.Use(static.Serve("/", static.LocalFile("./public", true)))
	wsServer.Router.NoRoute(func(c *gin.Context) {
		c.File("./public/index.html")
	})
}