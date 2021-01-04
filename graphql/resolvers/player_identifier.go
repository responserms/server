package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/responserms/server/ent"
	"github.com/responserms/server/graphql/types"
)

func (r *queryResolver) PlayerByIdentifiers(ctx context.Context, input []*types.PlayerIdentifierConstraintInput) (*ent.Player, error) {
	panic(fmt.Errorf("not implemented"))
}
