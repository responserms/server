package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/responserms/server/ent"
	"github.com/responserms/server/graphql/types"
)

func (r *mutationResolver) CreateMapLayer(ctx context.Context, input types.CreateMapLayerInput) (*ent.MapLayer, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateMapLayer(ctx context.Context, id int, input types.UpdateMapLayerInput) (*ent.MapLayer, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteMapLayer(ctx context.Context, id int) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) MapLayer(ctx context.Context, id int) (*ent.MapLayer, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) MapLayers(ctx context.Context) ([]*ent.MapLayer, error) {
	panic(fmt.Errorf("not implemented"))
}
