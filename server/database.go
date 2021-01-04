// Copyright (c) 2021 Contaim, LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package server

import (
	"github.com/responserms/server/internal/database"
)

// Database returns a pointer to the ent.Client used for database access.
func (s *Server) Database() *database.Database {
	return s.db
}
