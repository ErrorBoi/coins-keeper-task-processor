package services

import (
	"taskProcessor/internal/domain/operation/contracts"
	"taskProcessor/internal/domain/operation/dto"
	"taskProcessor/internal/domain/operation/entity"
)

type OperationServise struct {
	repository contracts.OperationRepository
}

func NewInstance(repository contracts.OperationRepository) OperationServise {
	return OperationServise{repository: repository}
}

func (this OperationServise) Add(dto dto.AddOperationDto) {

	operation := entity.NewInstance(
		dto.Amount,
		dto.Description,
		//@todo убрать костыль 
		"1",
		dto.UserId,
		dto.IsSpent,
	)

	this.repository.Save(operation)
}
