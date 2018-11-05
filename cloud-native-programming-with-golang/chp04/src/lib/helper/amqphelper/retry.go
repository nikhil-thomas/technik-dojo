package amqphelper

import (
	"log"
	"time"

	"github.com/streadway/amqp"
)

// RetryConnect implements a retry mechanism for establishing the AMQP connection.
func RetryConnect(amqpURL string, retryInterval time.Duration) chan *amqp.Connection {
	result := make(chan *amqp.Connection)
	defer close(result)

	go func() {
		for {
			conn, err := amqp.Dial(amqpURL)
			if err == nil {
				log.Println("AMQP connection successful")
				result <- conn
				return
			}
			log.Printf("AMQP connection failed with error (retrying in %s): %s\n", retryInterval.String(), err)
			time.Sleep(retryInterval)
		}
	}()
	return result
}
