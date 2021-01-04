// Copyright (c) 2020 Contaim, LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package server

import (
	"context"
	"fmt"

	"github.com/responserms/server/internal/log"
)

type startupFunc func(ctx context.Context, errChan chan error) error

// Start performs the startup sequence in the correct order for all internal services
// used by Response Server. This should be called only once.
//
// When the application is ready to shutdown call Shutdown(ctx) with the same context
// passed into the startup.
func (s *Server) Start(ctx context.Context, errChan chan error) {
	s.start(ctx, errChan)
}

// start performs the actual startup sequence.
func (s *Server) start(ctx context.Context, errChan chan error) {
	var startupSequence = map[string]startupFunc{
		"events":   s.startEventsService,
		"cluster":  s.startClusterService,
		"database": s.startDatabaseService,
		"http":     s.startHTTPService,
	}

	var startupOrder = []string{
		"database",
		"events",
		"cluster",
		"http",
	}

	// startup should happen only once
	s.once.Do(func() {
		s.log.Info("starting")

		for _, name := range startupOrder {
			s.log.Info("starting service", log.Attributes{
				"service": name,
			})

			if err := startupSequence[name](ctx, errChan); err != nil {
				errChan <- fmt.Errorf("startup failed for %s: %w", name, err)
				break
			}
		}
	})
}
