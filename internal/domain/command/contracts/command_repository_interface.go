package contracts

import "taskProcessor/internal/domain/command/entity"

type CommandRepository interface {
	FindByName(name string, userId string) *entity.Command
}
