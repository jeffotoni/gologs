// Back-End in Go server
// @jeffotoni
// 2019-01-04

package rabbitqm

import (
	"log"
	"strconv"

	"github.com/streadway/amqp"
)

var (
	err  error
	conn *amqp.Connection
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		return
	}
}

func connect() {
	conn, err = amqp.Dial("amqp://admin:1234#@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
}

func init() {
	connect()
}

func Publish(key_int int, value string) bool {
	key := strconv.Itoa(key_int)
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	//To send, we must declare a queue for us to send to
	//then we can publish a message to the queue
	q, err := ch.QueueDeclare(
		key,   // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	body := value
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")

	return true
}
