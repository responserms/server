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

func (r *metadataResolver) Data(ctx context.Context, obj *ent.Metadata, filter *string) (interface{}, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ApplyMetadata(ctx context.Context, entityID int, data interface{}) (*types.MetadataValidationStatus, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Metadata(ctx context.Context, entityID int) (*ent.Metadata, error) {
	panic(fmt.Errorf("not implemented"))
}

// Metadata returns server.MetadataResolver implementation.
func (r *Resolver) Metadata() server.MetadataResolver { return &metadataResolver{r} }

type metadataResolver struct{ *Resolver }
