package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strconv"

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
	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		log.Errorf("Error opening transaction,%s", err)
		return nil, err
	}

	//Do insert record
	if _, err := tx.NewInsert().Model(todo).Exec(ctx); err != nil {
		log.Errorf("Error inserting todo %v,%s", todo, err)
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		log.Errorf("Error committing todo %v,%s", todo, err)
		defer tx.Rollback()
		return nil, err
	}

	return todo, nil
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, id string, text string, done bool) (*model.Todo, error) {
	todoID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	todo := &model.Todo{
		ID:   todoID,
		Text: text,
		Done: done,
	}
	//TODO transactions
	if _, err := r.DB.NewDelete().Model(todo).WherePK().Exec(ctx); err != nil {
		log.Errorf("Error updating todo %d,%s", todoID, err)
		return nil, err
	}

	return todo, nil
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, id string) (*model.Todo, error) {
	todoID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	todo := &model.Todo{
		ID: todoID,
	}
	if _, err := r.DB.NewDelete().Model(todo).WherePK().Exec(ctx); err != nil {
		log.Errorf("Error deleting todo %d,%s", todoID, err)
		return nil, err
	}
	return todo, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	log.Printf("Creating new user %v", input)
	user := &model.User{
		Name:   input.Name,
		Gender: input.Gender.String(),
	}

	//Begin Transaction
	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		log.Errorf("Error opening transaction for user %v insert,%s", user, err)
		return nil, err
	}

	//Do insert record
	if _, err := tx.NewInsert().Model(user).Exec(ctx); err != nil {
		log.Errorf("Error inserting user %s", err)
	}

	if err := tx.Commit(); err != nil {
		log.Errorf("Error committing user %v,%s", user, err)
		defer tx.Rollback()
		return nil, err
	}

	return user, nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (*model.User, error) {
	userID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		ID: userID,
	}

	if _, err := r.DB.NewDelete().Model(user).WherePK().Exec(ctx); err != nil {
		log.Errorf("Error deleting todo %d,%s", userID, err)
		return nil, err
	}
	return user, nil
}

func (r *queryResolver) AllTodos(ctx context.Context, last *int) ([]*model.Todo, error) {
	log.Println("Resolve all todos")
	var todos []*model.Todo
	var err error
	if last != nil {
		//TODO add ordering
		err = r.DB.NewSelect().Model(&todos).Limit(*last).Scan(ctx)
	} else {
		err = r.DB.NewSelect().Model(&todos).Scan(ctx)
	}

	if err != nil {
		log.Errorf("Error querying all todos, %s", err)
		return nil, err
	}

	return todos, nil
}

func (r *queryResolver) Todo(ctx context.Context, id int) (*model.Todo, error) {
	log.Printf("Querying for TODO with id %d", id)
	todo := &model.Todo{
		ID: id,
	}

	if _, err := r.DB.NewSelect().Model(todo).WherePK().Exec(ctx); err != nil {
		log.Errorf("Error getting todo %d,%s", id, err)
		return nil, err
	}

	return todo, nil
}

func (r *queryResolver) TodosByStatus(ctx context.Context, status *bool) ([]*model.Todo, error) {
	log.Printf("Resolve todos by status %b\n", status)
	var todos []*model.Todo
	err := r.DB.NewSelect().Model(&todos).Where("status = ?", status).Scan(ctx)
	if err != nil {
		log.Errorf("Error querying for todos by status %b", status)
		return nil, err
	}

	return todos, nil
}

func (r *queryResolver) AllUsers(ctx context.Context, last *int) ([]*model.User, error) {
	log.Print("Resolving all users")
	var users []*model.User
	var err error
	if last != nil {
		//TODO add ordering
		err = r.DB.NewSelect().Model(&users).Limit(*last).Scan(ctx)
	} else {
		err = r.DB.NewSelect().Model(&users).Scan(ctx)
	}

	if err != nil {
		log.Errorf("Error querying all users, %s", err)
		return nil, err
	}

	return users, nil
}

func (r *queryResolver) User(ctx context.Context, id int) (*model.User, error) {
	log.Printf("Querying for User with id %d", id)
	user := &model.User{
		ID: id,
	}

	if _, err := r.DB.NewSelect().Model(user).WherePK().Exec(ctx); err != nil {
		log.Errorf("Error getting user %d,%s", id, err)
		return nil, err
	}

	return user, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
