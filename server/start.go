// Copyright (c) 2020 Contaim, LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package server

import (
	"context"
	"os"

	"github.com/responserms/server/internal/log"
)

type startupFunc func(ctx context.Context) error

// Start performs the startup sequence in the correct order for all internal services
// used by Response Server. This should be called only once.
//
// When the application is ready to shutdown call Shutdown(ctx) with the same context
// passed into the startup.
func (s *Server) Start(ctx context.Context) {
	s.start(ctx)
}

// start performs the actual startup sequence.
func (s *Server) start(ctx context.Context) {
	var startupSequence = map[string]startupFunc{
		"events": s.startEventsService,
		"http":   s.startHTTPService,
	}

	// startup should happen only once
	s.once.Do(func() {
		s.log.Info("starting")

		for name, start := range startupSequence {
			s.log.Info("starting service", log.Attributes{
				"service": name,
			})

			if err := start(ctx); err != nil {
				s.log.Error("starting service failed", log.Attributes{
					"service": name,
					"error":   err.Error(),
				})

				s.log.Info("startup sequence exited", log.Attributes{
					"service": name,
				})

				os.Exit(1)
			}
		}
	})
}
