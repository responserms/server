// Copyright (c) 2021 Contaim, LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package server

import (
	"context"
	"fmt"

	"github.com/responserms/server/internal/database"
)

func (s *Server) startDatabaseService(ctx context.Context, _ chan error) error {
	options := &database.Options{
		Driver:   database.Driver(s.cfg.Database.Type),
		Path:     s.cfg.Database.Path,
		URL:      s.cfg.Database.URL,
		Host:     s.cfg.Database.Host,
		Port:     s.cfg.Database.Port,
		Name:     s.cfg.Database.Name,
		Username: s.cfg.Database.Username,
		Password: s.cfg.Database.Password,
		Options:  s.cfg.Database.Options,
		Logger:   s.log,
	}

	db, err := database.Configure(ctx, options)
	if err != nil {
		return fmt.Errorf("configure database: %w", err)
	}

	// setup the service
	s.db = db

	return nil
}
