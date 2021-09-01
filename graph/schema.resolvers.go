package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/dondoko-susumu/gqlgen-todos/graph/generated"
	"github.com/dondoko-susumu/gqlgen-todos/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	rand, _ := rand.Int(rand.Reader, big.NewInt(100))
	todo := &model.Todo{
		Text:    input.Text,
		ID:      fmt.Sprintf("T%d", rand),
		SubText: input.SubText,
		UserID:  input.UserID,
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	// 例) SQL, 3rdパーティ、別サービスからとる
	return &model.User{
		ID:   obj.UserID,
		Name: "高橋",
	}, nil
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
