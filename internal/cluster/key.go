// Copyright (c) 2021 Contaim, LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package cluster

// Key is a type-safe key used by the Cluster Storer and Locker to ensure key-consistency and avoid typos by
// promoting a reusable key definition.
type Key string

// String gets the string-based implementation fo the Key.
func (k *Key) String() string {
	return string(*k)
}

func (k *Key) Get(store Storer) ([]byte, error) {
	// return store.Get(k)
	return nil, nil
}
