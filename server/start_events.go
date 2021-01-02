// Copyright (c) 2020 Contaim, LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package server

import (
	"context"
	"fmt"

	"github.com/responserms/server/internal/events"
)

func (s *Server) startEventsService(ctx context.Context) error {
	e, err := events.New(ctx, &events.Options{
		Logger:  s.log,
		Weight:  10,
		Driver:  events.Driver(s.cfg.Events.Type),
		URL:     s.cfg.Events.URL,
		Subject: s.cfg.Events.Subject,
	})

	if err != nil {
		return fmt.Errorf("start pubsub: %w", err)
	}

	s.events = e

	return nil
}
