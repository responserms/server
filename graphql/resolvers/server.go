package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/responserms/server/ent"
	"github.com/responserms/server/graphql/types"
)

func (r *mutationResolver) CreateServer(ctx context.Context, input types.CreateServerInput) (*ent.Server, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateServer(ctx context.Context, id int, input types.UpdateServerInput) (*ent.Server, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteServer(ctx context.Context, id int) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Server(ctx context.Context, id int) (*ent.Server, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Servers(ctx context.Context, paginate types.PaginationInput, orderBy *ent.ServerOrder) (*ent.ServerConnection, error) {
	panic(fmt.Errorf("not implemented"))
}
