// Copyright (c) 2020 Contaim, LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package database

import (
	"context"
	"fmt"

	"github.com/responserms/server/ent"
)

// Driver is a specifid database driver used to store persisted data in Response.
type Driver string

// Configure configures the provided Database as the implemented driver.
func (d *Driver) Configure(ctx context.Context) (*ent.Client, error) {
	switch *d {
	case EmbeddedDriver:
		return d.configureSQLiteMemory(ctx)
	case SQLiteDriver:
		return d.configureSQLiteFile(ctx)
	case PostgresDriver:
		return d.configurePostgres(ctx)
	default:
		return nil, fmt.Errorf("%s is not a supported database driver", *d)
	}
}
