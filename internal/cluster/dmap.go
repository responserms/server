package cluster

// DMapName represents the name of a distributed key/value map. This allows fetching a specific distributed map
// without understanding the underlying string-based variation.
type DMapName string

type dmap struct {}

func (dm *dmap) Get(key Key) ([]byte, error) {
	return nil, nil
}
