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
	// SQLiteDriver uses an SQLite3 implementation that stores a file at ./data/db.sqlite by default.
	// This path can be configured. This driver is not intended for production and should only be used
	// when persisting is necessary for development or testing purposes.
	SQLiteDriver Driver = "sqlite"
)

type sqliteOptions struct{}

func (d *Driver) configureSQLiteFile(ctx context.Context) (*ent.Client, error) {
	return nil, nil
}
