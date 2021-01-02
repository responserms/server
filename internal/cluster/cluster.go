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

func New()
