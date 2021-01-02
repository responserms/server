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
