package kafka

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/IBM/sarama"
)

type ConsumerMock struct {
	consumer sarama.Consumer
}

func NewConsumerMock(consumer sarama.Consumer) *ConsumerMock {
	return &ConsumerMock{consumer: consumer}
}

func (c *ConsumerMock) ConsumeMessageFromKafka() {
	if c.consumer == nil {
		log.Println("consumer not initialized")
		return
	}

	partitionConsumer, err := c.consumer.ConsumePartition("kafka-producer-mock", 0, 0)
	if err != nil {
		log.Printf("error while consuming partition: %v", err)
		return
	}
	defer partitionConsumer.Close()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case msg := <-partitionConsumer.Messages():
			fmt.Printf("Received message: %s\n", string(msg.Value))
			return
		case <-signals:
			return
		}
	}
}
