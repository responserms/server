package core

import (
	"github.com/responserms/server/internal/log"
	"github.com/responserms/server/internal/services/cluster"
	"github.com/responserms/server/internal/services/database"
	"github.com/responserms/server/internal/services/events"
)

// Core is a wrapper around database and application logic, queries, and event
// subscribers.
type Core struct {
	log  log.Logger
	svcs Backend

	Auth     AuthService
	Users    UsersService
	Sessions SessionsService
}

// Backend describes the methods requires to properly initialize Core.
type Backend interface {
	Database() *database.Database
	Events() events.PubSub
	Cluster() cluster.Cluster
}

// New initializes a new Core instance using the services implementation.
func New(logger log.ComponentLogger, svcs Backend) (*Core, error) {
	c := &Core{
		log:  logger.Component("core"),
		svcs: svcs,
	}

	c.Auth = &authService{c}
	c.Users = &usersService{c}
	c.Sessions = &sessionsService{c}

	return c, nil
}
