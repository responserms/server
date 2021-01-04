// Copyright (c) 2021 Contaim, LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package server

import "context"

func (s *Server) shutdownClusterService(ctx context.Context) error {
	return s.cluster.Shutdown(ctx)
}
