package store

import (
	"sync"
)

// KeyValue represents a single key-value pair.
// T is the type of the value.
type KeyValue[T any] struct {
	Value T
}

// DataStore represents the in-memory data store.
// T is the type of the values stored.
type DataStore[T any] struct {
	store map[string]KeyValue[T]
	mu    sync.RWMutex // Read Write mutex, this will help in concurrent access of the map.
}

// NewDataStore initializes a new DataStore.
func NewDataStore[T any]() *DataStore[T] {
	return &DataStore[T]{
		store: make(map[string]KeyValue[T]),
	}
}

// Set inserts or updates the value for a key.
func (ds *DataStore[T]) Set(key string, value T) {
	ds.mu.Lock()
	defer ds.mu.Unlock()
	ds.store[key] = KeyValue[T]{Value: value}
}

// Get retrieves the value for a key.
func (ds *DataStore[T]) Get(key string) (T, bool) {
	ds.mu.RLock()
	defer ds.mu.RUnlock()
	val, exists := ds.store[key]
	return val.Value, exists
}

// Delete removes a key from the store.
func (ds *DataStore[T]) Delete(key string) {
	ds.mu.Lock()
	defer ds.mu.Unlock()
	delete(ds.store, key)
}
