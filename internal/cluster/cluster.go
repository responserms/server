// Copyright (c) 2021 Contaim, LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package cluster

import (
	"time"

	"github.com/buraksezer/olric"
)

type StoreEntry struct {
	Key       string
	Value     []byte
	TTL       int64
	Timestamp int64
}

type Storer interface {
	Put(key string, val []byte) error
	PutWithTTL(key string, val []byte, ttl time.Duration)
	Get(key string) ([]byte, error)
	GetEntry(key string) (*StoreEntry, error)
	UpdateTTL(key string, newTTL time.Duration) error
	Delete(key string)
}

type Locker interface {
	Lock(key string) error
}

type cluster struct {
	impl *olric.Olric
}

func New() (*cluster, error) {
	return nil, nil
}
