package entity

const (
	TypeSpent   string = "spent"
	TypeAcquire string = "acquire"
)

type Command struct {
	Id int
	Name string
	CommandType string
	Answer string
	UserId *int
}