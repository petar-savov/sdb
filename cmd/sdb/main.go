package main

import (
	"fmt"

	store "github.com/petar-savov/sdb/pkg"
)

func main() {
	// Create a new instance of DataStore for storing strings.
	ds := store.NewDataStore[string]()

	// Set some key-value pairs.
	ds.Set("key1", "value1")
	ds.Set("key2", "value2")
	ds.Set("key3", "value3")

	// Get and print a value.
	if val, found := ds.Get("key2"); found {
		fmt.Println("Got:", val)
	} else {
		fmt.Println("Key not found.")
	}

	// Delete a key.
	ds.Delete("key3")

	// Try to get the deleted key.
	if _, found := ds.Get("key3"); !found {
		fmt.Println("Key3 deleted, not found.")
	}
}
