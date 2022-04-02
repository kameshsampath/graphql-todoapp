package model

import (
	"context"
	"fmt"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/schema"
	"time"
)

//User is the ORM model to save user in the database and return the data back to consumers
type User struct {
	bun.BaseModel `bun:"table: users,alias:u"`
	ID            int       `bun:",pk,autoincrement" json:"id"`
	Name          string    `bun:"name,notnull" json:"name"`
	Gender        string    `bun:"gender,notnull" json:"gender"`
	CreatedAt     time.Time `bun:"created_at,notnull" json:"-"`
	ModifiedAt    time.Time `bun:"modified_at,notnull" json:"-"`
}

var _ bun.BeforeAppendModelHook = (*User)(nil)

func (u *User) BeforeAppendModel(ctx context.Context, query schema.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		u.CreatedAt = time.Now()
	case *bun.UpdateQuery:
		u.ModifiedAt = time.Now()
	}
	return nil
}

func (u *User) String() string {
	return fmt.Sprintf("User: %s", u.Name)
}
