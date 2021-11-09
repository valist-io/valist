package memory

import (
	"github.com/valist-io/valist/internal/database"
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
		return nil, database.ErrKeyNotFound
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
