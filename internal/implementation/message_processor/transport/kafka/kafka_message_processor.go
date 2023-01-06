package kafka

import (
	"taskProcessor/internal/application/transport/kafka"

	kafkaPackage "github.com/confluentinc/confluent-kafka-go/kafka"
)

type Kafka struct {
	TaskTypeChat      string
	TaskTypeChatReady string
}

func (this Kafka) AcceptRequest(callback func(message string, userId string) string) {
	consumer := kafka.InitConsumer([]string{this.TaskTypeChat})

	kafka.StartConsumer(consumer, func(message *kafkaPackage.Message) {
		response := callback(string(message.Value), string(message.Key))

		kafka.ProduceMessage(&this.TaskTypeChatReady, response, message.Key)
	})

}

func NewInstance() Kafka {
	return Kafka{"chatTask", "chatTaskReady"}
}
