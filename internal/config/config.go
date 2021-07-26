package config

import (
	"os"
	"path/filepath"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

const (
	rootDir     = ".valist"
	keystoreDir = "keystore"
)

func GetKeyStore() (*keystore.KeyStore, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	dir := filepath.Join(home, rootDir, keystoreDir)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}

	return keystore.NewKeyStore(dir, keystore.StandardScryptN, keystore.StandardScryptP), nil
}
