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

const (
	// EmbeddedDriver provides an in-memory storage implementation based on sqlite3's in-memory
	// database. This should not be used in production as data is not persisted.
	EmbeddedDriver Driver = "embedded"
)

type embeddedOptions struct{}

func (d *Driver) configureSQLiteMemory(ctx context.Context) (*ent.Client, error) {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		return nil, fmt.Errorf("open embedded database: %w", err)
	}

	return client, nil
}
