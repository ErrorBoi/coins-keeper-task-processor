package main

import (
	"taskProcessor/internal/application/transport/kafka"
	"taskProcessor/internal/implementation/message_processor"
)

func main() {
	transport := kafka.NewInstance()

	message_processor.Start(transport)
}
