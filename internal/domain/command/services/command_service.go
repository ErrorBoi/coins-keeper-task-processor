package service

import (
	"taskProcessor/internal/domain/command/contracts"
	"taskProcessor/internal/domain/command/entity"
)

type CommandService struct {
	repository contracts.CommandRepository
}

func NewInstance(repository contracts.CommandRepository) CommandService {
	return CommandService{
		repository: repository,
	}
}

func (this CommandService) FindByName(name string, userId string) *entity.Command {
	return this.repository.FindByName(name, userId)
}
