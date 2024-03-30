package kafkatest

import (
	"EXAM3/user_service/kafka-test/kafka"
	"fmt"
	"testing"

	"github.com/IBM/sarama"
	"github.com/IBM/sarama/mocks"
)

func TestConsumerKafka(t *testing.T) {
	mockConsumer := mocks.NewConsumer(t, nil)
	defer func() {
		if err := mockConsumer.Close(); err != nil {
			t.Error(err)
		}
	}()

	consumer := kafka.NewConsumerMock(mockConsumer)

	expectedMessage := "Kafka Mock Test"

	mockMsg := &sarama.ConsumerMessage{
		Topic:     "kafka-producer-mock",
		Partition: 0,
		Offset:    0,
		Value:     []byte(expectedMessage),
	}

	mockConsumer.ExpectConsumePartition("kafka-producer-mock", 0, 0).YieldMessage(mockMsg)

	consumer.ConsumeMessageFromKafka()

	// Print the offset received
	offset := mockMsg.Offset
	fmt.Printf("Offset received: %d\n", offset)

}
