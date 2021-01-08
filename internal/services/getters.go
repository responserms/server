package services

import (
	"github.com/responserms/server/internal/services/cluster"
	"github.com/responserms/server/internal/services/database"
	"github.com/responserms/server/internal/services/events"
)

// Database returns the database instance. This is implemented in favor of exported structs to ensure New()
// is utilized and to prevent cyclical dependencies by allowing the use of interface implementations.
func (s *Services) Database() *database.Database {
	return s.database
}

// Cluster returns the cluster store instance. This is implemented in favor of exported structs to ensure New()
// is utilized and to prevent cyclical dependencies by allowing the use of interface implementations.
func (s *Services) Cluster() cluster.Cluster {
	return s.cluster
}

// Events returns the event bus instance. This is implemented in favor of exported structs to ensure New()
// is utilized and to prevent cyclical dependencies by allowing the use of interface implementations.
func (s *Services) Events() events.PubSub {
	return s.events
}
