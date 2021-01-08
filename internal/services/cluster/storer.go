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

// StoreEntry contains metadata about a particular key in a Storer.
type StoreEntry struct {
	Key       string
	Value     interface{}
	TTL       int64
	Timestamp int64
}

// ElseFunc is a function used on *Else() methods. ElseFunc should return an interface and optional error. If error is not nil
// the error will be passed through to the calling method and the returned interface will not be cached.
type ElseFunc func() (interface{}, error)

// Storer is implemented by services offering distributed Key-Value store.
type Storer interface {
	// Has determines if the Key is set on the store already.
	Has(key Key) (bool, error)

	// Put puts the value at the given Key.
	Put(key Key, val interface{}) error

	// PutWithTTL puts the value at the given Key, however, the Key will bre destroyed after the ttl has been
	// reached, meaning the key essentially self-destructs. This is primarily for tiem-based caching where data
	// changes and you want to re-fetch stale data after X time.
	PutWithTTL(key Key, val interface{}, ttl time.Duration) error

	// Get returns the value for the given Key.
	Get(key Key) (interface{}, error)

	// GetEntry returns the StoreEntry (metadata) at the given Key.
	GetEntry(key Key) (*StoreEntry, error)

	// UpdateTTL updates the TTL for the Key that either has a TTL (thus using a new TTL voiding the original), or settings
	// a TTL if the Key did not contain a TTL prior to it being set.
	UpdateTTL(key Key, newTTL time.Duration) error

	// Delete removes the Key and it's associated value.
	Delete(key Key) error

	// Incr increments the key atomically, returning the value post-increment.
	Incr(key Key, delta int) (int, error)

	// Decr decrements the key atomically, returning the value post-decrement.
	Decr(key Key, delta int) (int, error)

	// GetOrPut ruturns the value if the Key is already set otherwise it Puts the value and returns the value that
	// was Put. This is essentially a cache-first approach. elseFunc will not be called unless retrieving the key results
	// in a miss.
	GetOrPut(key Key, elseFunc ElseFunc) (interface{}, error)

	// GetOrPutWithTTL acts as GetOrPut but sets a TTL on the key forcing it to periodically be updated when retrieved.
	GetOrPutWithTTL(key Key, ttl time.Duration, elseFunc ElseFunc) (interface{}, error)
}

// Store represents the name of a distributed key/value map. This allows fetching a specific distributed map
// without understanding the underlying string-based variation.
type Store string

// String returns the string variant of the Store.
func (s Store) String() string {
	return string(s)
}

type store struct {
	dmap *olric.DMap
}

// NewStore creates a new store for reading/writing in the distributed cluster. A store should be created to
// logically separate keys by their purpose. This will avoid shadowing of existing keys with new ones that may
// unintentionally be named the same thing.
func (c *cluster) NewStore(name Store) (Storer, error) {
	dmap, err := c.impl.NewDMap(name.String())
	if err != nil {
		return nil, fmt.Errorf("new store: %w", err)
	}

	return &store{
		dmap: dmap,
	}, nil
}

func (s *store) Has(key Key) (bool, error) {
	_, err := s.dmap.Get(key.String())
	switch {
	case err == olric.ErrKeyNotFound:
		return false, nil
	case err != nil:
		return false, fmt.Errorf("has: %w", err)
	}

	return true, nil
}

func (s *store) Put(key Key, val interface{}) error {
	return s.dmap.Put(key.String(), val)
}

func (s *store) PutWithTTL(key Key, val interface{}, ttl time.Duration) error {
	return s.dmap.PutEx(key.String(), val, ttl)
}

func (s *store) Get(key Key) (interface{}, error) {
	return s.dmap.Get(key.String())
}

func (s *store) GetOrPut(key Key, elseFunc ElseFunc) (interface{}, error) {
	if hit, _ := s.Has(key); hit {
		return s.Get(key)
	}

	val, err := elseFunc()
	if err != nil {
		return nil, err
	}

	if err := s.Put(key, val); err != nil {
		return nil, err
	}

	return val, nil
}

func (s *store) GetOrPutWithTTL(key Key, ttl time.Duration, elseFunc ElseFunc) (interface{}, error) {
	if hit, _ := s.Has(key); hit {
		return s.Get(key)
	}

	val, err := elseFunc()
	if err != nil {
		return nil, err
	}

	if err := s.PutWithTTL(key, val, ttl); err != nil {
		return nil, err
	}

	return val, nil
}

func (s *store) GetEntry(key Key) (*StoreEntry, error) {
	entry, err := s.dmap.GetEntry(key.String())
	if err != nil {
		return nil, fmt.Errorf("get entry: %w", err)
	}

	return &StoreEntry{
		Key:       entry.Key,
		Value:     entry.Value,
		TTL:       entry.TTL,
		Timestamp: entry.Timestamp,
	}, nil
}

func (s *store) UpdateTTL(key Key, newTTL time.Duration) error {
	return s.dmap.Expire(key.String(), newTTL)
}

func (s *store) Delete(key Key) error {
	return s.dmap.Delete(key.String())
}

func (s *store) Incr(key Key, delta int) (int, error) {
	return s.dmap.Incr(key.String(), delta)
}

func (s *store) Decr(key Key, delta int) (int, error) {
	return s.dmap.Decr(key.String(), delta)
}
