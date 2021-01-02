// Copyright (c) 2021 Contaim, LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package cluster

// DMapName represents the name of a distributed key/value map. This allows fetching a specific distributed map
// without understanding the underlying string-based variation.
type DMapName string

type dmap struct{}

func (dm *dmap) Get(key Key) ([]byte, error) {
	return nil, nil
}
