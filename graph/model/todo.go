package model

import (
	"context"
	"fmt"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/schema"
	"time"
)

//Todo is the ORM model to save todos in the database and used when returning data for queries
type Todo struct {
	bun.BaseModel `bun:"table: todos,alias:t"`
	ID            int       `bun:",pk,autoincrement" json:"id"`
	Text          string    `bun:"text,notnull" json:"text"`
	Done          bool      `bun:"done,notnull,default:'false'" json:"done"`
	UserID        int       `bun:"user_id,notnull" json:"user"`
	User          *User     `bun:"rel:belongs-to,join:user_id=id"`
	CreatedAt     time.Time `bun:"created_at,notnull" json:"-"`
	ModifiedAt    time.Time `bun:"modified_at,notnull" json:"-"`
}

var _ bun.BeforeAppendModelHook = (*Todo)(nil)

func (t *Todo) BeforeAppendModel(ctx context.Context, query schema.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		t.CreatedAt = time.Now()
	case *bun.UpdateQuery:
		t.ModifiedAt = time.Now()
	}
	return nil
}

func (t *Todo) String() string {
	return fmt.Sprintf("TODO: %s for user %d", t.Text, t.UserID)
}
