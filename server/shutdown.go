package server

import (
	"context"
	"time"

	"github.com/responserms/server/internal/log"
)

type shutdownFunc func(ctx context.Context) error

// Shutdown initiates a clean shutdown of all services.
func (s *Server) Shutdown(ctx context.Context) error {
	var shutdownSequence = map[string]shutdownFunc{
		"events":  s.shutdownEventsService,
		"cluster": s.shutdownClusterService,
	}

	s.log.Info("shutdown requested")

	for name, shutdown := range shutdownSequence {
		ctx, cancel := context.WithDeadline(ctx, time.Now().Add(2*time.Second))
		defer cancel()

		if err := shutdown(ctx); err != nil {
			s.log.Error("service shutdown failed", log.Attributes{
				"service": name,
				"error":   name,
			})
		}

		s.log.Info("service is shutdown", log.Attributes{
			"service": name,
		})
	}

	s.log.Info("shutdown has completed")

	return nil
}
