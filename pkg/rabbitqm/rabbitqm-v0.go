// Back-End in Go server
// @jeffotoni
// 2019-01-04

package rabbitqm

import (
	"flag"
	"fmt"
	"log"
	"strconv"

	"github.com/streadway/amqp"
)

var (
	uri          = flag.String("uri", "amqp://guest:guest@localhost:5672/", "AMQP URI")
	exchangeName = flag.String("exchange", "test-exchange", "Durable AMQP exchange name")
	exchangeType = flag.String("exchange-type", "direct", "Exchange type - direct|fanout|topic|x-custom")
	//routingKey   = flag.String("key", "test-key", "AMQP routing key")
	//body     = flag.String("body", "foobar", "Body of message")
	reliable = flag.Bool("reliable", true, "Wait for the publisher confirmation before exiting")
)

type MessagingClient struct {
	conn *amqp.Connection
}

var ci = &MessagingClient{}

func init() {
	flag.Parse()
	ci.ConnectToBroker()
}

func (m *MessagingClient) ConnectToBroker() {

	connectionString := *uri

	if connectionString == "" {
		log.Println("Cannot initialize connection to broker, connectionString not set. Have you initialized?")
		return
	}

	var err error
	m.conn, err = amqp.Dial(fmt.Sprintf("%s/", connectionString))
	if err != nil {
		log.Println("Failed to connect to AMQP compatible broker at: " + connectionString)
		return
	}
}

func (m *MessagingClient) PublishOnQueue(body []byte) error {

	var queueName string
	queueName = "QueueGologs"

	if m.conn == nil {
		log.Println("Tried to send message before connection was initialized. Don't do that.")
		return nil
	}
	ch, err := m.conn.Channel() // Get a channel from the connection
	defer ch.Close()

	// Declare a queue that will be created if not exists with some args
	queue, err := ch.QueueDeclare(
		queueName, // our queue name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)

	// Publishes a message onto the queue.
	err = ch.Publish(
		"gologs_exchange1", // use the default exchange
		queue.Name,         // routing key, e.g. our queue name
		false,              // mandatory
		false,              // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body, // Our JSON body as []byte
		})
	//fmt.Printf("A message was sent to queue %v: %v", queueName, body)
	return err
}

func SendV2(key_int int, body string) bool {

	if ci.PublishOnQueue([]byte(body)) != nil {
		return false
	}

	return true
}

func SendV0(key_int int, body string) bool {

	key := strconv.Itoa(key_int)
	if err := publish(*uri, *exchangeName, *exchangeType, key, body, *reliable); err != nil {
		log.Fatalf("%s", err)
	}
	//log.Printf("published %dB OK", len(body))

	return true
}

func publish(amqpURI, exchange, exchangeType, routingKey, body string, reliable bool) error {

	// This function dials, connects, declares, publishes, and tears down,
	// all in one go. In a real service, you probably want to maintain a
	// long-lived connection as state, and publish against that.
	//log.Printf("dialing %q", amqpURI)
	connection, err := amqp.Dial(amqpURI)
	if err != nil {
		return fmt.Errorf("Dial: %s", err)
	}
	defer connection.Close()

	//log.Printf("got Connection, getting Channel")
	channel, err := connection.Channel()
	if err != nil {
		return fmt.Errorf("Channel: %s", err)
	}

	//log.Printf("got Channel, declaring %q Exchange (%q)", exchangeType, exchange)
	if err := channel.ExchangeDeclare(
		exchange,     // name
		exchangeType, // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // noWait
		nil,          // arguments
	); err != nil {
		return fmt.Errorf("Exchange Declare: %s", err)
	}

	// Reliable publisher confirms require confirm.select support from the
	// connection.
	if reliable {
		//log.Printf("enabling publishing confirms.")
		if err := channel.Confirm(false); err != nil {
			return fmt.Errorf("Channel could not be put into confirm mode: %s", err)
		}
		confirms := channel.NotifyPublish(make(chan amqp.Confirmation, 1))
		defer confirmOne(confirms)
	}

	//log.Printf("declared Exchange, publishing %dB body (%q)", len(body), body)
	if err = channel.Publish(
		exchange,   // publish to an exchange
		routingKey, // routing to 0 or more queues
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			Headers:         amqp.Table{},
			ContentType:     "application/json",
			ContentEncoding: "",
			Body:            []byte(body),
			DeliveryMode:    amqp.Transient, // 1=non-persistent, 2=persistent
			Priority:        0,              // 0-9
			// a bunch of application/implementation-specific fields
		},
	); err != nil {
		return fmt.Errorf("Exchange Publish: %s", err)
	}

	return nil
}

// One would typically keep a channel of publishings, a sequence number, and a
// set of unacknowledged sequence numbers and loop until the publishing channel
// is closed.
func confirmOne(confirms <-chan amqp.Confirmation) {
	log.Printf("waiting for confirmation of one publishing")
	if confirmed := <-confirms; confirmed.Ack {
		log.Printf("confirmed delivery with delivery tag: %d", confirmed.DeliveryTag)
	} else {
		log.Printf("failed delivery of delivery tag: %d", confirmed.DeliveryTag)
	}
}
