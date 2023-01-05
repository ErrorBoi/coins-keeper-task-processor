package message_processor

type Transport interface {
	AcceptRequest(
		callback func(message string) string,
	)
}

func Start(transport Transport) {
	transport.AcceptRequest(func(message string) string {
		return processMessage(message)
	})
}

func processMessage(message string) string {
	return "Сообщение обработано"
}
