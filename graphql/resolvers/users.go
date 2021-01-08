package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/responserms/server/ent"
	"github.com/responserms/server/graphql/server"
	"github.com/responserms/server/graphql/types"
)

func (r *queryResolver) Me(ctx context.Context) (*ent.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) User(ctx context.Context, id int) (*ent.User, error) {
	return r.core.Users.GetUserByID(ctx, id)
}

func (r *queryResolver) Users(ctx context.Context, paginate *types.PaginationInput, orderBy *ent.UserOrder, filter *types.UserFilter) (*ent.UserConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) Active(ctx context.Context, obj *ent.User) (bool, error) {
	return obj.ActivatedAt != nil && obj.ActivatedAt.Before(time.Now()), nil
}

// User returns server.UserResolver implementation.
func (r *Resolver) User() server.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
