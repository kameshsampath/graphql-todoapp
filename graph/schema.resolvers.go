package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/kameshsampath/blogapp/graph/generated"
	"github.com/kameshsampath/blogapp/graph/model"
	log "github.com/sirupsen/logrus"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		Text:   input.Text,
		UserID: input.UserID,
	}

	//Begin Transaction
	tx, err := r.DB.Begin()
	if err != nil {
		log.Errorf("Error opening transaction,%s", err)
		return nil, err
	}
	defer tx.Rollback()

	//Do insert record
	stmt, err := tx.Prepare("INSERT INTO todos(text,userId,done) VALUES($1,$2,$3)")
	if err != nil {
		log.Errorf("Error preparing todo %s statement ,%s", todo, err)
		return nil, err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(input.Text, input.UserID, false); err != nil {
		log.Errorf("Error inserting todo %s,%s", todo, err)
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		log.Errorf("Error committing todo %s,%s", todo, err)
		return nil, err
	}

	return todo, nil
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, id string, text string, done bool) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, id string) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	log.Println("Creating new user")
	user := &model.User{
		Name: input.Name,
		Sex:  input.Sex.String(),
	}

	//Begin Transaction
	tx, err := r.DB.Begin()
	if err != nil {
		log.Errorf("Error opening transaction for user %s ,%s", user, err)
		return nil, err
	}
	defer tx.Rollback()

	//Do insert record
	stmt, err := tx.Prepare("INSERT INTO users(name,sex) VALUES($1,$2)")
	if err != nil {
		log.Errorf("Error preparing statement for user %s ,%s", user, err)
		return nil, err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(input.Name, input.Sex); err != nil {
		log.Errorf("Error inserting user %s", err)
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		log.Errorf("Error committing user %s,%s", user, err)
		return nil, err
	}

	return user, nil
}

func (r *queryResolver) AllTodos(ctx context.Context, last *int) ([]*model.Todo, error) {
	log.Println("Resolve all todos")
	var todos []*model.Todo
	query := "SELECT id,text,done,userid FROM todos ORDER BY modifiedat"
	if *last > 0 {
		query = fmt.Sprintf("SELECT id,text,done,userid FROM todos ORDER BY modifiedat LIMIT %d", last)
	}

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var todo model.Todo
		if err := rows.Scan(&todo.ID, &todo.Text, &todo.Done, &todo.UserID); err == nil {
			todos = append(todos, &todo)
		}
	}

	return todos, nil
}

func (r *queryResolver) TodosByStatus(ctx context.Context, status *bool) ([]*model.Todo, error) {
	log.Printf("Resolve todos by status %b\n", status)
	var todos []*model.Todo
	stmt, err := r.DB.Prepare("SELECT id,text,done,userid from todos where done=$1")
	if err != nil {
		log.Errorf("Error preparing query for todos by status %b", status)
		return nil, err
	}
	rows, err := stmt.Query(status)
	if err != nil {
		log.Errorf("Error querying todos with status %b", status)
		return nil, err
	}
	for rows.Next() {
		var todo model.Todo
		if err := rows.Scan(&todo.ID, &todo.Text, &todo.Done, &todo.UserID); err == nil {
			todos = append(todos, &todo)
		}
	}

	return todos, nil
}

func (r *queryResolver) AllUsers(ctx context.Context, last *int) ([]*model.User, error) {
	log.Println("Resolve users")
	var users []*model.User
	query := "SELECT id,name,sex from users ORDER BY modifiedat"
	if *last > 0 {
		query = fmt.Sprintf("SELECT id,name,sex from users ORDER BY modifiedat LIMIT %d", last)
	}

	rows, err := r.DB.Query(query)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Sex); err == nil {
			users = append(users, &user)
		}
	}

	return users, nil
}

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	log.Printf("Resolve user from todo,%s\n", obj)
	var user model.User
	stmt, err := r.DB.Prepare("SELECT id,name,sex from users where id=$1")
	if err != nil {
		log.Errorf("Error preaparing statement to query user %s", err)
		return nil, err
	}

	rows, err := stmt.Query(obj.UserID)
	if err != nil {
		log.Errorf("Error while querying for users %s", err)
		return nil, err
	}
	if err := rows.Scan(&user.ID, &user.Name, &user.Sex); err == nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		log.Infof("No user found with id %s", obj.UserID)
	}

	return &user, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
