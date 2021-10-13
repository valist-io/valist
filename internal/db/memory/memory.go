package memory

import (
	"fmt"
)

type Database struct {
	values map[string][]byte
}

func NewDatabase() *Database {
	return &Database{make(map[string][]byte)}
}

func (db *Database) Set(key string, val []byte) error {
	db.values[key] = val
	return nil
}

func (db *Database) Get(key string) ([]byte, error) {
	val, ok := db.values[key]
	if !ok {
		return nil, fmt.Errorf("key not found")
	}

	return val, nil
}

func (db *Database) Delete(key string) error {
	delete(db.values, key)
	return nil
}

func (db *Database) Close() error {
	return nil
}
