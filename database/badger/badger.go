package badger

import (
	"time"

	"github.com/dgraph-io/badger/v3"

	"github.com/valist-io/valist/database"
)

type Database struct {
	badger *badger.DB
	tick   *time.Ticker
	done   chan bool
}

// NewDatabase returns a badger backed database using the given storage path.
func NewDatabase(storagePath string) (*Database, error) {
	opts := badger.DefaultOptions(storagePath)
	opts = opts.WithLoggingLevel(badger.ERROR)

	bdb, err := badger.Open(opts)
	if err != nil {
		return nil, err
	}

	tick := time.NewTicker(5 * time.Minute)
	done := make(chan bool)

	db := &Database{bdb, tick, done}
	go db.gcLoop()

	return db, nil
}

func (db *Database) Set(key string, val []byte) error {
	txn := db.badger.NewTransaction(true)
	defer txn.Discard()

	if err := txn.Set([]byte(key), val); err != nil {
		return err
	}

	return txn.Commit()
}

func (db *Database) Get(key string) ([]byte, error) {
	txn := db.badger.NewTransaction(false)
	defer txn.Discard()

	item, err := txn.Get([]byte(key))
	if err == badger.ErrKeyNotFound {
		return nil, database.ErrKeyNotFound
	}

	if err != nil {
		return nil, err
	}

	return item.ValueCopy(nil)
}

func (db *Database) Delete(key string) error {
	txn := db.badger.NewTransaction(true)
	defer txn.Discard()

	if err := txn.Delete([]byte(key)); err != nil {
		return err
	}

	return txn.Commit()
}

func (db *Database) Close() error {
	db.tick.Stop()
	db.done <- true
	return db.badger.Close()
}

func (db *Database) gcExec() {
	for err := error(nil); err == nil; {
		// If a GC is successful, immediately run it again.
		err = db.badger.RunValueLogGC(0.7)
	}
}

func (db *Database) gcLoop() {
	for {
		select {
		case <-db.tick.C:
			db.gcExec()
		case <-db.done:
			return
		}
	}
}
