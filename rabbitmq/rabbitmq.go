package rabbitmq

import (
	"jsapi-websocket/utils"
	"log"
	"os"

	"github.com/gorilla/websocket"
	"github.com/streadway/amqp"
)

type RabbitMq struct {
	Connection *amqp.Connection
	Channel *amqp.Channel
	Queue *amqp.Queue
}

// Cria um novo objeto de RabbitMq
func New(nome string) *RabbitMq {
	rabbitMq := &RabbitMq{}
	// Conecta no server RabbitMQ
	rabbitMq.connect()
	rabbitMq.createChannel()
	rabbitMq.createQueue(nome)
	return rabbitMq
}

func (r *RabbitMq) connect() {
	conn, err := amqp.Dial(os.Getenv("JSAPI_RABBITMQ_DIAL"))
	utils.FailOnError(err, "Falha ao conectar no servidor RabbitMQ.")
	r.Connection = conn
}

// Cria o canal de comunicação com a API
func (r *RabbitMq) createChannel() {
	ch, err := r.Connection.Channel()
	utils.FailOnError(err, "Falha ao criar canal no servidor RabbitMQ.")
	r.Channel = ch
}

// Cria o canal de comunicação com a API
func (r *RabbitMq) createQueue(nome string) {
	// Envia mensagem para a fila
	queue, err := r.Channel.QueueDeclare(
		nome, 
		false,   
		false,   
		false,   
		false,   
		nil,     
	)
	utils.FailOnError(err, "Falha ao criar a fila.")
	r.Queue = &queue
}

func (r *RabbitMq)ConsumeMessages(ws *websocket.Conn, mt int) {
	msgs, err := r.Channel.Consume(
		r.Queue.Name, 
		"",     
		true,   
		false, 
		false,  
		false,  
		nil,    
	)
	utils.FailOnError(err, "Falha ao receber mensagens da fila.")

	forever := make(chan bool)

	go func() {
	for d := range msgs {
		log.Printf(": %s", d.Body)
		err := ws.WriteMessage(mt, d.Body)
		if err != nil {
			log.Printf("Erro: %s", err.Error())
		}
	}
	}()

	log.Printf(" [*] Aguardando por mensagens. Para sair pressione CTRL+C")
	<-forever
}