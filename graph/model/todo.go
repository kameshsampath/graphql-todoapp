package model

import "fmt"

//Todo holds the todo object
type Todo struct {
	ID     string `json:"id"`
	Text   string `json:"text"`
	Done   bool   `json:"done"`
	UserID string `json:"user"`
}

func (t *Todo) String() string {
	return fmt.Sprintf("TODO: %s for user %s", t.Text, t.UserID)
}
