package message_processor

import (
	"taskProcessor/internal/kafka"

	externalKafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

const (
	TaskTypeChat      string = "chatTask"
	TaskTypeChatReady string = "chatTaskReady"
)

func Start() {
	consumer := kafka.InitConsumer([]string{TaskTypeChat})

	kafka.StartConsumer(consumer, func(message *externalKafka.Message) {
		topic := TaskTypeChatReady

		response := processMessage(string(message.Value))

		kafka.ProduceMessage(&topic, response, message.Key)
	})
}

func processMessage(message string) string {
	return "Сообщение обработано"
}
