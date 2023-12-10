package store

import (
	"testing"
)

func TestDataStore_SetAndGet(t *testing.T) {
	ds := NewDataStore[string]()
	key, value := "testKey", "testValue"

	ds.Set(key, value)
	if got, exists := ds.Get(key); !exists || got != value {
		t.Errorf("Set(%q, %q), got %q, exists %v; want %q, exists true", key, value, got, exists, value)
	}
}

func TestDataStore_Delete(t *testing.T) {
	ds := NewDataStore[string]()
	key, value := "testKey", "testValue"

	ds.Set(key, value)
	ds.Delete(key)
	if _, exists := ds.Get(key); exists {
		t.Errorf("Delete(%q) failed, key %q still exists", key, key)
	}
}
