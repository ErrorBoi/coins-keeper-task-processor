package repositories

import (
	"taskProcessor/internal/application/mysql/builder"
	"taskProcessor/internal/domain/operation/entity"
)

type OperationRepository struct {
}

func NewInstance() OperationRepository {
	return OperationRepository{}
}

func (this OperationRepository) Save(operation entity.Operation) {
	builder.NewInstance().Insert("operation", operation).Execute()
}
