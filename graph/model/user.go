package model

import (
	"fmt"
	"github.com/uptrace/bun"
	"time"
)

//User is the ORM model to save user in the database and return the data back to consumers
type User struct {
	bun.BaseModel `bun:"table: users,alias:u"`
	ID            int       `bun:",pk,autoincrement" json:"id"`
	Name          string    `bun:"name,notnull" json:"name"`
	Gender        string    `bun:"gender,notnull" json:"gender"`
	CreatedAt     time.Time `bun:"modified_at,notnull,default:'current_timestamp'" json:"-"`
	ModifiedAt    time.Time `bun:"modified_at,notnull,default:'create_timestamp'" json:"-"`
}

func (u *User) String() string {
	return fmt.Sprintf("User: %s", u.Name)
}
