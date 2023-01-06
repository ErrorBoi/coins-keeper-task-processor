package dto

import (
	"errors"
	"strings"
)

type AddOperationDto struct {
	Amount       string
	Description  string
	CurrencyName string
	UserId       string
	IsSpent      bool
}

//@todo не нравится, много параметров 
func FromString(data string, isSpent bool, userId string) (AddOperationDto, error) {
	words := strings.Fields(data)

	var instance AddOperationDto

	if len(words) < 4 {
		return instance, errors.New("Not all fields provided")
	}

	amount := words[1]

	currency := words[2]

	description := words[3]

	instance = AddOperationDto{
		Amount:       amount,
		Description:  description,
		CurrencyName: currency,
		IsSpent:      isSpent,
		UserId:       userId,
	}

	return instance, nil
}
