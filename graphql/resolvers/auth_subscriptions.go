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

func (r *subscriptionResolver) UserActivated(ctx context.Context, userID int) (<-chan *types.UserActivatedResult, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) UserRegistered(ctx context.Context) (<-chan *ent.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Subscription returns server.SubscriptionResolver implementation.
func (r *Resolver) Subscription() server.SubscriptionResolver { return &subscriptionResolver{r} }

type subscriptionResolver struct{ *Resolver }
