package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/responserms/server/ent"
	"github.com/responserms/server/graphql/types"
)

func (r *mutationResolver) UpsertPlayers(ctx context.Context, server int, input []*types.UpsertPlayerInput) ([]*ent.Player, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Player(ctx context.Context, id int) (*ent.Player, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Players(ctx context.Context, paginate types.PaginationInput, orderBy *ent.PlayerOrder, filter *types.PlayerFilter) (*ent.PlayerConnection, error) {
	panic(fmt.Errorf("not implemented"))
}
