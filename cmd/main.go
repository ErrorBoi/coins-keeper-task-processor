package main

import (
	commandServices "taskProcessor/internal/domain/command/services"
	messageProcessorServices "taskProcessor/internal/domain/message_processor/services"
	operationServices "taskProcessor/internal/domain/operation/services"
	commandRepositories "taskProcessor/internal/implementation/command/repositories"
	kafkaMessageProcessors "taskProcessor/internal/implementation/message_processor/transport/kafka"
	operationRepositories "taskProcessor/internal/implementation/operation/repositories"
)

func main() {
	makeMessageProcessor().Start()
}

func makeMessageProcessor() messageProcessorServices.MessageProcessorService {
	transport := kafkaMessageProcessors.NewInstance()

	commandRepository := commandRepositories.NewInstance()

	commandService := commandServices.NewInstance(commandRepository)

	operationRepository := operationRepositories.NewInstance()

	operationService := operationServices.NewInstance(operationRepository)

	return messageProcessorServices.NewInstance(
		commandService,
		transport,
		operationService,
	)
}
