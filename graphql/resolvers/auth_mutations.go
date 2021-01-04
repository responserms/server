package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/responserms/server/ent"
	"github.com/responserms/server/graphql/server"
	"github.com/responserms/server/graphql/types"
)

func (r *mutationResolver) Login(ctx context.Context, input types.LoginInput) (*types.LoginResult, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Register(ctx context.Context, input types.RegisterInput) (*types.RegisterResult, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ActivateUser(ctx context.Context, userID int, comments *string) (*ent.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns server.MutationResolver implementation.
func (r *Resolver) Mutation() server.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
