package dns

import (
	"taskProcessor/internal/config"
	"taskProcessor/internal/mysql/connection"
)

func GetDns() connection.Dns {
	var payload connection.Dns

	payload.Host = config.Get("DBHOST")
	payload.Password = config.Get("DBPASSWORD")
	payload.Port = config.Get("DBPORT")
	payload.User = config.Get("DBUSER")

	return payload
}
