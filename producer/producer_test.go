package producer_test

import (
	"fmt"
	"go_kafka/producer"
	"testing"

	"github.com/Shopify/sarama/mocks"
)

func TestSendMessage(t *testing.T) {
	t.Run("Send Message OK", func(t *testing.T) {
		mockedProducer := mocks.NewSyncProducer(t, nil)
		mockedProducer.ExpectSendMessageAndSucceed()
		kafka := &producer.KafkaProducer{
			Producer: mockedProducer,
		}

		msg := "Message 1"

		err := kafka.SendMessage("test_topic", msg)
		if err != nil {
			t.Errorf("Send message should not be errorbut have: %v", err)
		}
	})

	t.Run("Send Mesage NOK", func(t *testing.T) {
		mockedProducer := mocks.NewSyncProducer(t, nil)
		mockedProducer.ExpectSendMessageAndFail(fmt.Errorf("Error"))
		kafka := &producer.KafkaProducer{
			Producer: mockedProducer,
		}

		msg := "Message 1"

		err := kafka.SendMessage("test_topic", msg)
		if err != nil {
			t.Errorf("this should be error")
		}
	})
}
