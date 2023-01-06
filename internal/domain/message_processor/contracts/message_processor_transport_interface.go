package contracts

type Transport interface {
	AcceptRequest(
		callback func(message string, userId string) string,
	)
}
