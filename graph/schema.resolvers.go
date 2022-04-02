package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/kameshsampath/blogapp/graph/dao"
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

	if err := dao.InsertTodo(ctx, r.DB, todo); err != nil {
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

	if err := dao.UpdateTodo(ctx, r.DB, todo); err != nil {
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

	if err := dao.DeleteTodo(ctx, r.DB, todo); err != nil {
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

	if err := dao.InsertUser(ctx, r.DB, user); err != nil {
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

	if err := dao.DeleteUser(ctx, r.DB, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (r *queryResolver) AllTodos(ctx context.Context, last *int) ([]*model.Todo, error) {
	log.Println("Resolve all todos")
	var todos []*model.Todo

	if err := dao.SelectTodos(ctx, r.DB, &todos, last); err != nil {
		return nil, err
	}

	return todos, nil
}

func (r *queryResolver) Todo(ctx context.Context, id int) (*model.Todo, error) {
	log.Printf("Querying for TODO with id %d", id)
	todo := &model.Todo{
		ID: id,
	}

	if err := dao.FindTodoByID(ctx, r.DB, todo); err != nil {
		return nil, err
	}

	return todo, nil
}

func (r *queryResolver) TodosByStatus(ctx context.Context, status *bool) ([]*model.Todo, error) {
	log.Printf("Resolve todos by status %b\n", status)
	var todos []*model.Todo

	if err := dao.FindTodosByStatus(ctx, r.DB, &todos, status); err != nil {
		log.Errorf("Error querying for todos by status %b", status)
		return nil, err
	}

	return todos, nil
}

func (r *queryResolver) AllUsers(ctx context.Context, last *int) ([]*model.User, error) {
	log.Print("Resolving all users")
	var users []*model.User

	if err := dao.SelectUsers(ctx, r.DB, &users, last); err != nil {
		log.Errorf("Error querying all users, %s", err)
		return nil, err
	}

	log.Printf("Users %v", users)

	return users, nil
}

func (r *queryResolver) User(ctx context.Context, id int) (*model.User, error) {
	log.Printf("Querying for User with id %d", id)
	user := &model.User{
		ID: id,
	}

	if err := dao.FindUserByID(ctx, r.DB, user); err != nil {
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
