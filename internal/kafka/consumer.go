package kafka

import (
	"fmt"
	"log"

	"taskProcessor/internal/config"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func InitConsumer(topics []string) *kafka.Consumer {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": config.Get("KAFKA_ADDRESS"),
		"group.id":          "test",
		"auto.offset.reset": "smallest"})

	if err != nil {
		log.Println("Consume failed: ", err)
	}

	err = consumer.SubscribeTopics(topics, nil)

	if err != nil {
		log.Println("Consume failed: ", err)
	}

	return consumer
}

func StartConsumer(consumer *kafka.Consumer, callback func(message *kafka.Message)) {
	for true {
		event := consumer.Poll(100)
		switch eventType := event.(type) {
		case *kafka.Message:
			callback(eventType)
		case kafka.Error:
			log.Fatalf("%% Error: %v\n", eventType)

			break
		}
	}
}

func ProduceMessage(topic *string, value string, key []byte) {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": config.Get("KAFKA_ADDRESS"),
		"acks":              "all"})

	if err != nil {
		log.Println("Produce failed: ", err)
	}

	deliveryChan := make(chan kafka.Event, 10000)

	err = producer.Produce(
		&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: topic, Partition: kafka.PartitionAny},
			Value:          []byte(value),
			Key:            key,
		},
		deliveryChan,
	)

	event := <-deliveryChan

	message := event.(*kafka.Message)

	if message.TopicPartition.Error != nil {
		log.Fatalf("Delivery failed: %v\n", message.TopicPartition.Error)
	} 

	fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
		*message.TopicPartition.Topic, message.TopicPartition.Partition, message.TopicPartition.Offset)

	close(deliveryChan)
}
