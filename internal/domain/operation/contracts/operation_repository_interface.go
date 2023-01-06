package contracts

import "taskProcessor/internal/domain/operation/entity"

type OperationRepository interface {
	Save(operation entity.Operation)
}
