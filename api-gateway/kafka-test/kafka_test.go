package kafkatest

import (
	"EXAM3/api-gateway/kafka-test/kafka"
	"testing"

	"github.com/IBM/sarama/mocks"
)

func TestProducer(t *testing.T) {
	mockProducer := mocks.NewSyncProducer(t, nil)
	defer func() {
		if err := mockProducer.Close(); err != nil {
			t.Error(err)
		}
	}()

	producer := kafka.NewMockKafka(mockProducer)

	messageToProduce := "Kafka Mock Test"

	mockProducer.ExpectSendMessageAndSucceed()

	err := producer.SendMessageToKafka(messageToProduce)
	if err != nil {
		t.Errorf("Error sending message to Kafka: %v", err)
	}

}
