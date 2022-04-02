package dao

import (
	"context"
	"github.com/kameshsampath/blogapp/graph/model"
	log "github.com/sirupsen/logrus"
	"github.com/uptrace/bun"
)

//InsertTodo inserts the todo in the DB
func InsertTodo(ctx context.Context, db *bun.DB, todo *model.Todo) error {
	//Begin Transaction
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Errorf("Error opening transaction,%s", err)
		return err
	}

	//Do insert record
	if _, err := tx.NewInsert().Model(todo).Exec(ctx); err != nil {
		log.Errorf("Error inserting todo %v,%s", todo, err)
		return err
	}

	if err := tx.Commit(); err != nil {
		log.Errorf("Error committing todo %v,%s", todo, err)
		defer tx.Rollback()
		return err
	}

	return nil
}

//UpdateTodo update the todo in the DB
func UpdateTodo(ctx context.Context, db *bun.DB, todo *model.Todo) error {
	//Begin Transaction
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Errorf("Error opening transaction,%s", err)
		return err
	}

	//Do insert record
	if _, err := tx.NewUpdate().Model(todo).WherePK().Exec(ctx); err != nil {
		log.Errorf("Error inserting todo %v,%s", todo, err)
		return err
	}

	if err := tx.Commit(); err != nil {
		log.Errorf("Error committing todo %v,%s", todo, err)
		defer tx.Rollback()
		return err
	}

	return nil
}

//DeleteTodo delete the todo in the DB
func DeleteTodo(ctx context.Context, db *bun.DB, todo *model.Todo) error {
	if _, err := db.NewDelete().Model(todo).WherePK().Exec(ctx); err != nil {
		log.Errorf("Error deleting todo %d,%s", todo.ID, err)
		return err
	}

	return nil
}

//SelectTodos selects all todos from the database
func SelectTodos(ctx context.Context, db *bun.DB, todos *[]*model.Todo, last *int) error {
	var err error
	if last != nil {
		//TODO add ordering
		err = db.NewSelect().Model(todos).Limit(*last).Scan(ctx)
	} else {
		err = db.NewSelect().Model(todos).Scan(ctx)
	}

	if err != nil {
		log.Errorf("Error querying all todos, %s", err)
		return err
	}

	return nil
}

//FindTodoByID find the todo by its primary key
func FindTodoByID(ctx context.Context, db *bun.DB, todo *model.Todo) error {
	if err := db.NewSelect().WherePK().Model(todo).Scan(ctx); err != nil {
		log.Errorf("Error getting user with id %d, %s", todo.ID, err)
		return err
	}

	return nil
}

//FindTodosByStatus finds the todos by its status
func FindTodosByStatus(ctx context.Context, db *bun.DB, todos *[]*model.Todo, status *bool) error {
	if err := db.NewSelect().Model(todos).Where("status = ?", status).Scan(ctx); err != nil {
		log.Errorf("Error querying for todos by status %b", status)
		return err
	}

	return nil
}
