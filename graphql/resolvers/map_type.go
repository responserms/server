package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/responserms/server/ent"
	"github.com/responserms/server/graphql/types"
)

func (r *mutationResolver) CreateMapType(ctx context.Context, input types.CreateMapTypeInput) (*ent.MapType, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateMapType(ctx context.Context, id int, input types.UpdateMapTypeInput) (*ent.MapType, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteMapType(ctx context.Context, id int) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) MapType(ctx context.Context, id int) (*ent.MapType, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) MapTypes(ctx context.Context) ([]*ent.MapType, error) {
	panic(fmt.Errorf("not implemented"))
}
