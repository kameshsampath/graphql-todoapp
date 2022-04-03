package dao

import (
	"context"
	"github.com/kameshsampath/blogapp/graph/model"
	log "github.com/sirupsen/logrus"
	"github.com/uptrace/bun"
)

//FindTodosByStatus finds the todos by its status
func FindTodosByStatus(ctx context.Context, db *bun.DB, todos *[]*model.Todo, status *bool) error {
	if err := db.NewSelect().Model(todos).Where("done = ?", status).Scan(ctx); err != nil {
		log.Errorf("Error querying for todos by status %v", status)
		return err
	}

	return nil
}
