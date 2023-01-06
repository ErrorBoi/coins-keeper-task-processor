package connection

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var openedConnection *sql.DB = nil

type Dns struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

func GetConnection(dnsData Dns) *sql.DB {

	if openedConnection != nil {
		return openedConnection
	}

	pattern := "%s:%s@tcp(%s:%s)/%s"

	dns := fmt.Sprintf(
		pattern,
		dnsData.User,
		dnsData.Password,
		dnsData.Host,
		dnsData.Port,
		dnsData.Database,
	)

	connection, err := sql.Open("mysql", dns)

	if err != nil {
		log.Fatal(err)
	}

	openedConnection = connection

	return openedConnection
}

func CloseConnection() {
	
}
