package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"

	"github.com/responserms/server/ent"
	"github.com/responserms/server/graphql/server"
	"github.com/responserms/server/internal/cluster"
)

const (
	pingStore = "ping"
	pingsKey  = cluster.Key("pings")
)

func (r *queryResolver) Node(ctx context.Context, id int) (ent.Noder, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Nodes(ctx context.Context, ids []int) ([]ent.Noder, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Ping(ctx context.Context) (string, error) {
	s, err := r.backend.Cluster().NewStore(pingStore)
	if err != nil {
		return "", fmt.Errorf("cannot create store: %w", err)
	}

	val, err := pingsKey.Incr(s, 1)
	if err != nil {
		return "", fmt.Errorf("incr: %w", err)
	}

	return strconv.Itoa(val), nil
}

// Query returns server.QueryResolver implementation.
func (r *Resolver) Query() server.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
