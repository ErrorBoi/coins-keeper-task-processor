package entity

import (
	"log"
	"strconv"
)

type Operation struct {
	Amount      float64
	Description string
	CurrencyId  int64
	UserId      int64
	IsSpent     bool
}

func NewInstance(
	amount string,
	description string,
	currencyId string,
	userId string,
	isSpent bool,
) Operation {

	amountAsFloat, err := strconv.ParseFloat(amount, 32)

	if err != nil {
		log.Fatal(err.Error())
	}

	currencyIdAsInt, err := strconv.ParseInt(currencyId, 6, 32)

	if err != nil {
		log.Fatal(err.Error())
	}

	userIdAsInt, err := strconv.ParseInt(userId, 6, 32)

	if err != nil {
		log.Fatal(err.Error())
	}

	return Operation{
		Amount:      amountAsFloat,
		Description: description,
		CurrencyId:  currencyIdAsInt,
		UserId:      userIdAsInt,
		IsSpent:     isSpent,
	}
}
