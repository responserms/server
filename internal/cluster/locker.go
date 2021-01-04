// Copyright (c) 2021 Contaim, LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package cluster

import (
	"fmt"
	"time"

	"github.com/buraksezer/olric"
)

// LockContext is the context returned by Lock when the underlying lock is acquired. LockContext
// includes an Unlock method.
type LockContext struct {
	ctx *olric.LockContext
}

// Locker is the interface to be implemented by services that offer locking/unlocking of a key.
type Locker interface {
	Lock(key LockKey, deadline time.Duration) (*LockContext, error)
}

type locker struct {
	dmap *olric.DMap
}

// NewLockStore locks the
func (s *cluster) NewLockStore(name string) (Locker, error) {
	dmap, err := s.impl.NewDMap(name)
	if err != nil {
		return nil, fmt.Errorf("new lock store: %w", err)
	}

	return &locker{
		dmap: dmap,
	}, nil
}

// Lock acquires a lock for the given key and deadline and returns a context to be used to release
// the Lock. If the Lock cannot be acquired an error will be returned as the second parameter. If
// an error is returned the lock was not acquired.
func (l *locker) Lock(key LockKey, deadline time.Duration) (*LockContext, error) {
	ctx, err := l.dmap.Lock(key.String(), deadline)
	if err != nil {
		return nil, fmt.Errorf("lock: %w", err)
	}

	return &LockContext{
		ctx: ctx,
	}, nil
}

// Unlock releases the lock attached to the LockContext. Unlock returns an error if the lock
// cannot be released.
func (c *LockContext) Unlock() error {
	return c.ctx.Unlock()
}
