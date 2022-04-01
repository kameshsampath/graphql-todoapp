package model

import "fmt"

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Sex  string `json:"sex"`
}

func (u *User) String() string {
	return fmt.Sprintf("User: %s", u.Name)
}
