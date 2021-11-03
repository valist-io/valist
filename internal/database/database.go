package database

import (
	"errors"
)

var ErrKeyNotFound = errors.New("Key not found")

// Database is a key value database.
type Database interface {
	// Set puts the value under the given key.
	Set(string, []byte) error
	// Get returns the value of the entry with the given key.
	Get(string) ([]byte, error)
	// Delete removes the entry with the given key.
	Delete(string) error
	// Close releases database resources.
	Close() error
}
