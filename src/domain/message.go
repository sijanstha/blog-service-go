package domain

type CHANGE_TYPE string

const (
	CREATE CHANGE_TYPE = "CREATE"
	UPDATE             = "UPDATE"
)

type Message struct {
	ChangeType CHANGE_TYPE `json:"changeType"`
	Body       interface{} `json:"body"`
}
