package rabbitConnection

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

const ConnectionRabbit = "amqp://guest:guest@localhost:5672/"
const QueueName = "onboarding-avenue"

func CreateConnection() (conn *amqp.Connection, ch *amqp.Channel) {
	var err error
	conn, err = amqp.Dial(ConnectionRabbit)
	logFatalError("Falha ao conectar %v", err)
	ch, err = conn.Channel()
	logFatalError("Falha ao abrir canal %v", err)

	return
}

func DeclareQueue(conn *amqp.Connection, ch *amqp.Channel) amqp.Queue {
	q, err := ch.QueueDeclare(
		QueueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)

	logFatalError("Erro ao criar fila %v", err)

	return q
}

func PublishMessage(ch *amqp.Channel, obj interface{}) {

	body, err := json.Marshal(obj)
	logFatalError("Erro ao serializar Json %v", err)

	err = ch.Publish(
		"",        // exchange
		QueueName, // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(body),
		})

	logFatalError("Erro ao criar publicar mensagem %v", err)
}

func ConsumeMessages(ch *amqp.Channel, q amqp.Queue) <-chan amqp.Delivery {

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	logError("Erro ao criar consumir mensagens:", err)

	return msgs
}

func logFatalError(description string, err error) {
	if err != nil {
		log.Fatalf(description, err.Error())
	}
}

func logError(description string, err error) {
	if err != nil {
		log.Println(description, err.Error())
	}
}
