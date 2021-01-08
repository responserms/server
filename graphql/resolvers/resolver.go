package resolvers

import (
	"fmt"

	"github.com/responserms/server/internal/core"
	"github.com/responserms/server/internal/log"
	"github.com/responserms/server/internal/services/cluster"
	"github.com/responserms/server/internal/services/database"
	"github.com/responserms/server/internal/services/events"
)

// Resolver is the GraphQL resolver dependency injection point.
type Resolver struct {
	core *core.Core
}

// Backend is implemented by services that provide a backend for the GraphQL resolver
// implementation. This is primarily done to avoid circular dependencies.
type Backend interface {
	Database() *database.Database
	Cluster() cluster.Cluster
	Events() events.PubSub
}

// New creates a new Resolver with the provided backend.
func New(logger log.Logger, backend Backend) (*Resolver, error) {
	c, err := core.New(logger, backend)
	if err != nil {
		return nil, fmt.Errorf("resolvers.New: %w", err)
	}

	r := &Resolver{
		core: c,
	}

	return r, nil
}
