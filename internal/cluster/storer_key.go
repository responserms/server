// Copyright (c) 2021 Contaim, LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package cluster

import "time"

// Key is a type-safe key used by the Cluster Storer and Locker to ensure key-consistency and avoid typos by
// promoting a reusable key definition.
type Key string

// String gets the string-based implementation fo the Key.
func (k Key) String() string {
	return string(k)
}

// Put puts the value at this Key.
func (k Key) Put(store Storer, val interface{}) error {
	return store.Put(k, val)
}

// PutWithTTL puts the value for this Key, however, the Key will bre destroyed after the ttl has been
// reached, meaning the key essentially self-destructs. This is primarily for tiem-based caching where data
// changes and you want to re-fetch stale data after X time.
func (k Key) PutWithTTL(store Storer, val interface{}, ttl time.Duration) error {
	return store.PutWithTTL(k, val, ttl)
}

// Get ruturns the value for this Key.
func (k Key) Get(store Storer) (interface{}, error) {
	return store.Get(k)
}

// GetEntry returns the StoreEntry (metadata) at the given Key.
func (k Key) GetEntry(store Storer) (*StoreEntry, error) {
	return store.GetEntry(k)
}

// UpdateTTL updates the TTL for the Key that either has a TTL (thus using a new TTL voiding the original), or settings
// a TTL if the Key did not contain a TTL prior to it being set.
func (k Key) UpdateTTL(store Storer, newTTL time.Duration) error {
	return store.UpdateTTL(k, newTTL)
}

// Delete deletes this Key. This will result in this Key no longer being distributed until Put() is called again.
func (k Key) Delete(store Storer) error {
	return store.Delete(k)
}

// Incr increments the key atomically, returning the value post-increment.
func (k Key) Incr(store Storer, delta int) (int, error) {
	return store.Incr(k, delta)
}

// Decr decrements the key atomically, returning the value post-decrement.
func (k Key) Decr(store Storer, delta int) (int, error) {
	return store.Decr(k, delta)
}
