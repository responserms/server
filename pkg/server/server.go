package server

import (
	"context"
	"fmt"

	"github.com/getsentry/sentry-go"
	"github.com/responserms/server/internal/log"
	"github.com/responserms/server/internal/services"
	"github.com/responserms/server/pkg/config"
)

const (
	dsn = "https://a166929708e94a929edfbe3deded771a@o439664.ingest.sentry.io/5579346"
)

// Server is the entrypoint to the Response Server component. This instance controls the
// embedded and external service connections and links them to our internal and external
// APIs.
type Server struct {
	log  log.Logger
	svcs *services.Services
}

// New creates a new Server to launch the Response Server component.
func New(ctx context.Context, cfg *config.Config) (*Server, error) {
	logger, err := log.New(log.WithLogLevel(cfg.LogLevel))
	if err != nil {
		return nil, fmt.Errorf("server.New: %w", err)
	}

	svcs, err := services.New(ctx, logger.Component("server.services"), cfg)
	if err != nil {
		return nil, fmt.Errorf("server.New: %w", err)
	}

	// initialize error reporting
	err = sentry.Init(sentry.ClientOptions{
		Dsn:        dsn,
		ServerName: "",
	})
	if err != nil {
		return nil, fmt.Errorf("server.New: %w", err)
	}

	s := &Server{
		log:  logger.Component("server"),
		svcs: svcs,
	}

	return s, nil
}

// Start the Response Server component.
func (s *Server) Start(ctx context.Context, errChan chan error) {
	s.svcs.Start(ctx, errChan)
}

// Shutdown gracefully stops the Response Server component.
func (s *Server) Shutdown(ctx context.Context) error {
	return s.svcs.Shutdown(ctx)
}
