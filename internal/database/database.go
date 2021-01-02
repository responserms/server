// Copyright (c) 2020 Contaim, LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package database

import (
	"context"
	"fmt"
	"sync"

	"github.com/responserms/server/ent"
	"github.com/responserms/server/ent/migrate"
	"github.com/responserms/server/internal/log"

	// sqlite3 is used for the sqlite and embedded drivers
	_ "github.com/mattn/go-sqlite3"

	// pq is used for the postgres driver
	_ "github.com/lib/pq"
)

type Options struct {
	Driver   Driver
	Path     string
	URL      string
	Host     string
	Port     int
	Name     string
	Username string
	Password string
	Options  map[string]string
	Logger   log.Logger
}

// Database holds the database configuration for the internal ORM. This allows a driver-agnostic way
// to configure and startup Response leaving all driver-specific configuration to each individual driver's
// implementation.
type Database struct {
	*ent.Client
	once    sync.Once
	options *Options

	log log.Logger
}

// Configure configures and opens a new connection to the database and returns a Database instance that can be used
// for manipulating the database graph with the generated ent client.
func Configure(ctx context.Context, options *Options) (*Database, error) {
	db := &Database{
		options: options,
		log:     options.Logger,
	}

	db.log.Info("configuring database driver", log.Attributes{
		"driver": options.Driver,
	})

	client, err := options.Driver.Configure(ctx)
	if err != nil {
		return nil, fmt.Errorf("open database: %w", err)
	}

	db.Client = client

	if err := db.init(ctx); err != nil {
		return nil, fmt.Errorf("migrate database: %w", err)
	}

	return db, nil
}

func (d *Database) init(ctx context.Context) error {
	var err error = nil

	// only allow this to happen once
	d.once.Do(func() {
		err = d.Schema.Create(
			ctx,
			migrate.WithGlobalUniqueID(true),
		)
	})

	return err
}
