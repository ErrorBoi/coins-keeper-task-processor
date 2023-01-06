package services

import (
	"strings"
	"taskProcessor/internal/domain/command/entity"
	commandService "taskProcessor/internal/domain/command/services"
	"taskProcessor/internal/domain/message_processor/contracts"
	"taskProcessor/internal/domain/operation/dto"
	operationService "taskProcessor/internal/domain/operation/services"
)

type MessageProcessorService struct {
	commandService   commandService.CommandService
	transport        contracts.Transport
	operationService operationService.OperationServise
}

func NewInstance(
	commandService commandService.CommandService,
	transport contracts.Transport,
	operationService operationService.OperationServise,
) MessageProcessorService {
	return MessageProcessorService{
		commandService:   commandService,
		transport:        transport,
		operationService: operationService,
	}
}

func (this MessageProcessorService) Start() {
	this.transport.AcceptRequest(func(message string, userId string) string {
		return this.processMessage(message, userId)
	})
}

func (this MessageProcessorService) processMessage(message string, userId string) string {
	action := strings.Fields(message)[0]

	command := this.commandService.FindByName(action, userId)

	if command == nil {
		return "Я не знаю, что мне делать, но вы можете меня научить."
	}

	dto, err := dto.FromString(message, command.CommandType == entity.TypeSpent, userId)

	if err != nil {
		return "Сообщение должно быть написано по шаблону: {Операция} {сумма} {валюта} {комментарий}"
	}

	this.operationService.Add(dto)

	return command.Answer
}
