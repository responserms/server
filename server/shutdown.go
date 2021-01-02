package server

import (
	"context"

	"github.com/responserms/server/internal/log"
)

type shutdownFunc func(ctx context.Context) error

func (s *Server) Shutdown(ctx context.Context) error {
	var shutdownSequence = map[string]shutdownFunc{
		"events": s.shutdownEventsService,
	}

	s.log.Info("shutdown requested")

	for name, shutdown := range shutdownSequence {
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
