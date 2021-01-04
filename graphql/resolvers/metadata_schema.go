package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/responserms/server/ent"
	"github.com/responserms/server/graphql/server"
)

func (r *metadataSchemaResolver) Schema(ctx context.Context, obj *ent.MetadataSchema) (map[string]interface{}, error) {
	panic(fmt.Errorf("not implemented"))
}

// MetadataSchema returns server.MetadataSchemaResolver implementation.
func (r *Resolver) MetadataSchema() server.MetadataSchemaResolver { return &metadataSchemaResolver{r} }

type metadataSchemaResolver struct{ *Resolver }
