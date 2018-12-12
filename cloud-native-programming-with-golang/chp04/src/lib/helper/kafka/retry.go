package helper

import (
	"log"
	"time"

	"github.com/Shopify/sarama"
)

// RetryConnect implements a retry mechanism for establishing the Kafka client.
func RetryConnect(brokers []string, retryInterval time.Duration) chan sarama.Client {
	result := make(chan sarama.Client)
	defer close(result)

	go func() {
		for {
			config := sarama.NewConfig()
			conn, err := sarama.NewClient(brokers, config)
			if err == nil {
				log.Println("Kafka connection successful")
				result <- conn
				return
			}
			log.Printf("Kafka connection failed with error (retrying in %s): %s\n", retryInterval.String(), err)
			time.Sleep(retryInterval)
		}
	}()
	return result
}
