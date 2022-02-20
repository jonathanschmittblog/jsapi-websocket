package routers

import (
	"jsapi-websocket/rabbitmq"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)
var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// webSocket returns text format
func messages(c *gin.Context) {
	//Upgrade get request to webSocket protocol
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
	  log.Println("Erro ao conectar websocket")
	  log.Fatal(err)
	}
	defer ws.Close()
	//Read data in ws
	mt, message, err := ws.ReadMessage()
	if err != nil {
	  log.Println("erro ao ler mensagem da dashboard")
	  log.Fatal(err)
	}
	log.Println(message)
	//Envia as mensagens do RabbitMQ para o dashboard
	queue := rabbitmq.New("pessoas")
	queue.ConsumeMessages(ws, mt)
	defer queue.Connection.Close()
	defer queue.Channel.Close()
}