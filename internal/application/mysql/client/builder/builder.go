package builder

import (
	"database/sql"
	"log"
	"taskProcessor/internal/application/mysql/connection"
	"taskProcessor/internal/application/mysql/dns"
)

/** Make query */
func MakeQuery(query string) *sql.Rows {

	dnsData := dns.GetDns()

	result, err := connection.GetConnection(dnsData).Query(query)

	if err != nil {
		log.Fatal(err)
	}

	return result
}
