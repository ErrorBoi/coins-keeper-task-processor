package repositories

import (
	"database/sql"
	"log"
	"strings"
	"taskProcessor/internal/application/mysql/builder"
	"taskProcessor/internal/domain/command/entity"
)

type CommandRepository struct {
}

func NewInstance() CommandRepository {
	return CommandRepository{}
}

func (this CommandRepository) FindByName(name string, userId string) *entity.Command {
	var commandModel entity.Command

	empty := true

	builder.NewInstance().Select("command").Where("name", "=", strings.ToLower(name)).GetResult(func(rows *sql.Rows) {
		mapTableIntoModel(&commandModel, rows)

		empty = false
	})

	if empty {
		return nil
	}

	return &commandModel
}

func mapTableIntoModel(commandModel *entity.Command, rows *sql.Rows) {
	err := rows.Scan(
		&commandModel.Id,
		&commandModel.Name,
		&commandModel.CommandType,
		&commandModel.Answer,
		&commandModel.UserId,
	)

	if err != nil {
		log.Fatal("Error while select comand data: " + err.Error())
	}

}
