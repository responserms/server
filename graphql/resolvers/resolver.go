package resolvers

import (
	"github.com/responserms/server/internal/cluster"
	"github.com/responserms/server/internal/database"
)

// Resolver is the GraphQL resolver dependency injection point.
type Resolver struct {
	backend Backend
}

// Backend is implemented by services that provide a backend for the GraphQL resolver
// implementation. This is primarily done to avoid circular dependencies.
type Backend interface {
	Database() *database.Database
	Cluster() cluster.Cluster
}

// New creates a new Resolver with the provided backend.
func New(backend Backend) (*Resolver, error) {
	r := &Resolver{
		backend: backend,
	}

	return r, nil
}
