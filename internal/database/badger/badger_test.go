package badger

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valist-io/valist/internal/database"
)

func TestBadgerDatabase(t *testing.T) {
	tmp, err := os.MkdirTemp("", "")
	require.NoError(t, err, "Failed to MkdirTemp")

	db, err := NewDatabase(tmp)
	require.NoError(t, err, "Failed to create database")
	defer db.Close()

	key := "test"
	val := []byte("hello")

	err = db.Set(key, val)
	require.NoError(t, err, "Failed to set entry")

	res, err := db.Get(key)
	require.NoError(t, err, "Failed to get entry")
	assert.Equal(t, val, res)

	err = db.Delete(key)
	require.NoError(t, err, "Failed to delete entry")

	_, err = db.Get(key)
	assert.Equal(t, database.ErrKeyNotFound, err)
}
