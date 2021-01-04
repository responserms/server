package server

import (
	"errors"
	"sync"

	"github.com/responserms/server/internal/cluster"
	"github.com/responserms/server/internal/database"
	"github.com/responserms/server/internal/events"
	"github.com/responserms/server/internal/http"
	"github.com/responserms/server/internal/log"
	"github.com/responserms/server/pkg/config"
)

// Server is the Response Server application.
type Server struct {
	once sync.Once

	// runtime
	log  log.Logger
	cfg  *config.Config
	http *http.Server
	db   *database.Database

	// cluster store and member manager
	cluster cluster.Cluster

	// services
	events events.PubSub
}

// New creates a new Server instance. The Config instance provided can be created using
// the config package.
func New(cfg *config.Config) (*Server, error) {
	server := &Server{}

	// Assign the config to the Server.
	server.cfg = cfg

	if server.cfg.EncryptionKey == "" {
		return nil, errors.New("encryption key is not set")
	}

	// create the logger with the configured log level
	logger, err := log.New(log.WithLogLevel(server.cfg.LogLevel))
	if err != nil {
		return nil, err
	}

	server.log = logger.Component("server")

	return server, nil
}
