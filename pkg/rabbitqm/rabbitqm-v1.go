// Back-End in Go server
// @jeffotoni
// 2019-01-04

package rabbitqm

import (
	"log"
	"strconv"

	"github.com/jeffotoni/gologs/config"
	"github.com/streadway/amqp"
)

var (
	err  error
	conn *amqp.Connection
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Println(msg, err)
		return
	}
}

func connect() {

	conn, err = amqp.Dial(RABBI_DIAL)
	if err != nil {
		log.Println(err)
		return
	}
	//failOnError(err, "Failed to connect to RabbitMQ")
	// defer conn.Close()
}

func init() {
	if config.SERVICE == config.RABBITQM {
		connect()
	}

	// // Get queue connection
	// queue := drivers.NewQueue()
	// queue.Connect()

	// // Declare new queue for submissions 'stuff'
	// _, err = queue.Declare("submissions_queue")
	// if err != nil {
	// 	log.Panic(err)
	// }

	// // We need to declare the responses queue even though we're not actually
	// // going to use it on the bot side as an emitter.
	// _, err = queue.Declare("submissions_responses")

	// // Log out any queue errors
	// if err != nil {
	// 	log.Panic(err)
	// }
}

// func (queue *Queue) Connect(args ...interface{}) {

// 	uri := "amqp://guest:guest@localhost:5672/"

// 	// If connection is successful, return new instance
// 	conn, err = amqp.Dial(uri)

// 	if err == nil {
// 		log.Println("Successfully connected to queue!")
// 		channel, _ := conn.Channel()
// 		queue.Connection = conn
// 		queue.Channel = channel
// 		return
// 	}
// }

// // Declare a new queue
// func (queue *Queue) Declare(queueName string) (amqp.Queue, error) {
// 	return queue.Channel.QueueDeclare(
// 		queueName,
// 		true,
// 		false,
// 		false,
// 		false,
// 		nil,
// 	)
// }

// // Publish a message
// func (queue *Queue) Publish(queueName string, payload []byte) error {
// 	return queue.Channel.Publish(
// 		"",
// 		queueName,
// 		false,
// 		false,
// 		amqp.Publishing{
// 			DeliveryMode: amqp.Persistent,
// 			ContentType:  "application/json",
// 			Body:         payload,
// 		},
// 	)
// }

func Send(key_int int, value string) bool {

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
	//failOnError(err, "Failed to declare a queue")
	if err != nil {
		log.Println(err)
		return false
	}

	body := value
	err = ch.Publish(
		"gologs", // exchange
		q.Name,   // routing key
		false,    // mandatory
		false,    // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(body),
		})
	//failOnError(err, "Failed to publish a message")

	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
