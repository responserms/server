package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/responserms/server/ent"
	"github.com/responserms/server/graphql/server"
)

func (r *sessionTokenResolver) Expired(ctx context.Context, obj *ent.SessionToken) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *sessionTokenResolver) Blocked(ctx context.Context, obj *ent.SessionToken) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

// SessionToken returns server.SessionTokenResolver implementation.
func (r *Resolver) SessionToken() server.SessionTokenResolver { return &sessionTokenResolver{r} }

type sessionTokenResolver struct{ *Resolver }
