package config

import (
	"encoding/json"
	"math/big"
	"os"
	"path/filepath"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
)

const (
	rootDir     = ".valist"
	configFile  = "config"
	keystoreDir = "keystore"
)

const (
	scryptN = keystore.StandardScryptN
	scryptP = keystore.StandardScryptP
)

type Ethereum struct {
	// RPC is the ethereum rpc address.
	RPC string `json:"rpc"`
	// Contracts is a mapping of contract addresses.
	Contracts map[string]common.Address `json:"contracts"`
	// ChainID is the unique id of the ethereum chain.
	ChainID *big.Int `json:"chain_id"`
}

type IPFS struct {
	// API is the IPFS api address.
	API string `json:"api"`
	// Peers is a mapping of peer addresses to connect to.
	Peers []string `json:"peers"`
}

type Config struct {
	rootPath string
	// Ethereum contains ethereum config.
	Ethereum Ethereum `json:"ethereum"`
	// IPFS contains ipfs config.
	IPFS IPFS `json:"ipfs"`
	// Accounts is a mapping of names to addresses.
	Accounts map[string]common.Address `json:"accounts"`
}

// Default returns a config with default settings.
func Default() Config {
	return Config{
		Ethereum: Ethereum{
			RPC:     "https://rpc.valist.io",
			ChainID: big.NewInt(80001),
			Contracts: map[string]common.Address{
				"valist":   common.HexToAddress("0xA7E4124aDBBc50CF402e4Cad47de906a14daa0f6"),
				"registry": common.HexToAddress("0x2Be6D782dBA2C52Cd0a41c6052e914dCaBcCD78e"),
			},
		},
		IPFS: IPFS{
			API: "/dns/pin.valist.io",
			Peers: []string{
				"/ip4/107.191.98.233/tcp/4001/p2p/QmasbWJE9C7PVFVj1CVQLX617CrDQijCxMv6ajkRfaTi98",
			},
		},
		Accounts: make(map[string]common.Address),
	}
}

// Exists returns true of the config root exists.
func Exists(path string) (bool, error) {
	rootPath := filepath.Join(path, rootDir)

	info, err := os.Stat(rootPath)
	if os.IsNotExist(err) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return info.IsDir(), nil
}

// Init initializes a config with a default account.
func Init(path, password string) error {
	rootPath := filepath.Join(path, rootDir)
	keystorePath := filepath.Join(rootPath, keystoreDir)

	if err := os.Mkdir(rootPath, 0755); err != nil {
		return err
	}

	account, err := keystore.StoreKey(keystorePath, password, scryptN, scryptP)
	if err != nil {
		return err
	}

	config := Default()
	config.Accounts["default"] = account.Address
	config.rootPath = rootPath

	return config.Save()
}

// Load loads the config from the given root path.
func Load(path string) (*Config, error) {
	rootPath := filepath.Join(path, rootDir)
	configPath := filepath.Join(rootPath, configFile)

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	config.rootPath = rootPath
	return &config, nil
}

// Save writes the config to the root path.
func (c Config) Save() error {
	data, err := json.MarshalIndent(c, "", "\t")
	if err != nil {
		return err
	}

	configPath := filepath.Join(c.rootPath, configFile)
	return os.WriteFile(configPath, data, 0666)
}

// KeyStore returns the keystore from the root path.
func (c Config) KeyStore() *keystore.KeyStore {
	dir := filepath.Join(c.rootPath, keystoreDir)
	return keystore.NewKeyStore(dir, scryptN, scryptP)
}
