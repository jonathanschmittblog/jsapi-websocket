package routers

import (
	"jsapi-websocket/utils"

	"github.com/jonathanschmittblog/jsapi-websocket/servers"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(wsServer *servers.Server) error {
	wsServer.Router.GET("/messages", messages)
	// Configura o arquivo html como página inicial
	err := utils.DownloadFile("./public/index.html", "https://github.com/jonathanschmittbr/jsapi-websocket/public/index.html")
	if err != nil {
		return err
	}
	wsServer.Router.Use(static.Serve("/", static.LocalFile("./public", true)))
	wsServer.Router.NoRoute(func(c *gin.Context) {
		c.File("./public/index.html")
	})
	return nil
}