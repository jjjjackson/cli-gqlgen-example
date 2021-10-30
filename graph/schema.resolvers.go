package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/jjjjackson/cli-gqlgen-example/graph/generated"
	"github.com/jjjjackson/cli-gqlgen-example/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	// user := model.User{ID: input.UserID, Name: "fooName"}
	return &model.Todo{ID: "created_foo_id", Text: input.Text, Done: false, User: nil}, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return []*model.Todo{{ID: "query_foo_id", Text: "query foo text", Done: false, User: nil}}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
