// Copyright (c) 2021 Contaim, LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package cluster

import "time"

// LockKey represents a key to track lock and unlocks. The LockKey is implemented separately from Key so that locks
// are never tied to a data-value key.
type LockKey string

// String returns the key as a string instead of the actual LockKey type.
func (l *LockKey) String() string {
	if l == nil {
		return ""
	}

	return string(*l)
}

// Lock locks the given key. When Lock returns the lock has been acquired, otherwise Lock will block until the lock
// can be acquired.
func (l *LockKey) Lock(locker Locker, deadline time.Duration) (*LockContext, error) {
	return locker.Lock(*l, deadline)
}
