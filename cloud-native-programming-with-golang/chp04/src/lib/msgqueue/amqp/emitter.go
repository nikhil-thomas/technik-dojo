package amqp

import (
	"encoding/json"
	"os"
	"time"

	"github.com/nikhil-thomas/technik-dojo/cloud-native-programming-with-golang/chp04/src/lib/helper/amqphelper"
	"github.com/nikhil-thomas/technik-dojo/cloud-native-programming-with-golang/chp04/src/lib/msgqueue"
	"github.com/streadway/amqp"
)

type amqpEventEmitter struct {
	connection *amqp.Connection
	exchange   string
	events     chan *emittedEvent
}

type emittedEvent struct {
	event     msgqueue.Event
	errorChan chan error
}

func NewAMQPEventEmitterFromEnvironment() (msgqueue.EventEmitter, error) {
	var url string
	var exchange string

	if url = os.Getenv("AMQP_URL"); url == "" {
		url = "amqp://localhost:5672"
	}

	if exchange = os.Getenv("AMQP_EXCHANGE"); exchange == "" {
		exchange = "example"
	}

	conn := <-amqphelper.RetryConnect(url, 5*time.Second)
	return NewAMQPEventEmitter(conn, exchange)
}

func NewAMQPEventEmitter(conn *amqp.Connection, exchange string) (msgqueue.EventEmitter, error) {
	emitter := amqpEventEmitter{
		connection: conn,
		exchange:   exchange,
	}

	err := emitter.setup()
	if err != nil {
		return nil, err
	}

	return &emitter, nil
}

func (a *amqpEventEmitter) setup() error {
	channel, err := a.connection.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	err = channel.ExchangeDeclare(a.exchange, "topic", true, false, false, false, nil)
	return err
}

func (a *amqpEventEmitter) Emit(event msgqueue.Event) error {
	channel, err := a.connection.Channel()
	if err != nil {
		return err
	}

	defer channel.Close()
	jsonBody, err := json.Marshal(event)
	if err != nil {
		return err
	}

	msg := amqp.Publishing{
		Headers:     amqp.Table{"x-event-name": event.EventName()},
		ContentType: "application/josn",
		Body:        jsonBody,
	}
	err = channel.Publish(a.exchange, event.EventName(), false, false, msg)
	return err
}
