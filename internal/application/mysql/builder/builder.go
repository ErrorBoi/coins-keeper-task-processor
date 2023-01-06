package builder

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"strings"
	"taskProcessor/internal/application/helpers"
	"taskProcessor/internal/application/mysql/connection"
	"taskProcessor/internal/application/mysql/dns"
)

type QueryBuilder struct {
	Query    string
	Bindings []any
}

func NewInstance() QueryBuilder {
	return QueryBuilder{}
}

func (this QueryBuilder) Select(table string) QueryBuilder {
	this.Query = "SELECT * FROM " + table + " WHERE 1 = 1"

	return this
}

func (this QueryBuilder) Where(field string, operator string, value any) QueryBuilder {
	this.Query += " AND " + field + " " + operator + " ?"

	this.Bindings = append(this.Bindings, value)

	return this
}

func (this QueryBuilder) Insert(table string, data interface{}) QueryBuilder {

	fields, values := mapStructToDatabaseValues(data)

	valuesPlaceholder := strings.Repeat("?,", len(values))

	valuesPlaceholderTrimmed := strings.TrimRight(valuesPlaceholder, ",")

	columnsPlaceholder := strings.Join(fields[:], ",")

	this.Query += "INSERT INTO " + table + " (" + columnsPlaceholder + ")" + "VALUES (" + valuesPlaceholderTrimmed + ")"

	this.Bindings = values

	return this
}

func (this QueryBuilder) setBindings(value []any) QueryBuilder {
	this.Bindings = value

	return this
}

func (builder QueryBuilder) GetResult(callback func(rows *sql.Rows)) {
	rows, err := getConenction().Query(builder.Query, builder.Bindings...)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		callback(rows)
	}
}

func (builder QueryBuilder) Execute() {
	_, err := getConenction().Exec(builder.Query, builder.Bindings...)

	if err != nil {
		log.Fatal(err)
	}
}

func getConenction() *sql.DB {
	dnsData := dns.GetDns()

	return connection.GetConnection(dnsData)
}

func mapStructToDatabaseValues(data interface{}) ([]string, []any) {
	v := reflect.ValueOf(data)
	n := v.NumField()

	st := reflect.TypeOf(data)

	fields := make([]string, 0, n)
	values := make([]any, 0, n)

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)

		fieldName := st.Field(i).Name

		fields = append(fields, helpers.ToSnakeCase(fieldName))

		s := fmt.Sprintf("%v", field.Interface())

		if field.Type() == reflect.TypeOf("") {
			s = `"` + s + `"`
		}

		values = append(values, s)
	}

	return fields, values
}
