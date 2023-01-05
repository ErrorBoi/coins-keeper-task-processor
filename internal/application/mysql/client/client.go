package client

import (
	"database/sql"
	"log"
	"taskProcessor/internal/application/mysql/client/builder"
)

func SelectTableData(databaseName string, tableName string) []map[string]interface{} {
	result := builder.MakeQuery("SELECT * FROM " + databaseName + "." + tableName)

	return makeRawResult(result)
}

func makeSingleFieldList(result *sql.Rows) []string {
	var listData []string

	for result.Next() {
		var name string
		err := result.Scan(&name)

		if err != nil {
			log.Fatal(err)
		}

		listData = append(listData, name)
	}

	return listData
}

func makeRawResult(rows *sql.Rows) []map[string]interface{} {
	cols, err := rows.Columns()

	if err != nil {
		panic(err)
	}

	allgeneric := make([]map[string]interface{}, 0)

	colvals := make([]interface{}, len(cols))

	for rows.Next() {
		colassoc := make(map[string]interface{}, len(cols))

		for i := range colvals {
			colvals[i] = new(interface{})
		}

		if err := rows.Scan(colvals...); err != nil {
			panic(err)
		}

		for i, col := range cols {
			colassoc[col] = *colvals[i].(*interface{})
		}

		allgeneric = append(allgeneric, colassoc)
	}

	return allgeneric
}
