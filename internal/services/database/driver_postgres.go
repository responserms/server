// Copyright (c) 2020 Contaim, LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package database

import (
	"context"

	"github.com/responserms/server/ent"
)

const (
	// PostgresDriver implements storage using a supported PostgreSQL version (currently only 10, 11, or 12).
	// This driver is safe for production use and is the preferred production database.
	PostgresDriver Driver = "postgres"
)

type postgresOptions struct{}

func (d *Driver) configurePostgres(ctx context.Context) (*ent.Client, error) {
	return nil, nil
}
