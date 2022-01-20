package core

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const (
	// ContractTypeEVM enables EVM contracts and accounts.
	ContractTypeEVM = "evm"
	// StorageTypeIPFS enables IPFS storage.
	StorageTypeIPFS = "ipfs"
)

type Config struct {
	// rootPath is the path the config is loaded from
	// this is used as the root for other directories
	// as well as the path for saving changes to
	rootPath string

	// ApiAddress is the api server address to use
	ApiAddress string `json:"api_address"`
	// ContractType is the type of contract provider.
	ContractType string `json:"contract_type"`
	// StorageType is the type of storage provider.
	StorageType string `json:"storage_type"`
	// DefaultAccounts is a mapping of contract types to accounts.
	DefaultAccounts map[string]string `json:"default_accounts"`

	// EthereumRPC is the ethereum rpc url.
	EthereumRPC string `json:"ethereum_rpc"`
	// EthereumMetaTx enables meta transactions.
	EthereumMetaTx bool `json:"ethereum_meta_tx"`
	// EthereumBiconomyApiKey is the mexa public api key.
	EthereumBiconomyApiKey string `json:"ethereum_biconomy_api_key"`
	// EthereumContracts is a mapping of contract addresses.
	EthereumContracts map[string]string `json:"ethereum_contracts"`

	// IpfsApiAddress is the address of a remote IPFS node to connect to.
	IpfsApiAddress string `json:"ipfs_api_address"`
	// IpfsBootstrapPeers is a list of peers to connect to.
	IpfsBootstrapPeers []string `json:"ipfs_bootstrap_peers"`
}

// NewConfig returns a config with default settings.
func NewConfig(path string) *Config {
	return &Config{
		rootPath:               filepath.Join(path, ".valist"),
		ApiAddress:             "localhost:9000",
		ContractType:           ContractTypeEVM,
		StorageType:            StorageTypeIPFS,
		DefaultAccounts:        make(map[string]string),
		EthereumRPC:            "https://rpc.valist.io",
		EthereumMetaTx:         true,
		EthereumBiconomyApiKey: "qLW9TRUjQ.f77d2f86-c76a-4b9c-b1ee-0453d0ead878",
		EthereumContracts: map[string]string{
			"valist": "0xA7E4124aDBBc50CF402e4Cad47de906a14daa0f6",
		},
		IpfsBootstrapPeers: []string{
			"/ip4/107.191.98.233/tcp/4001/p2p/QmasbWJE9C7PVFVj1CVQLX617CrDQijCxMv6ajkRfaTi98",
			"/ip4/107.191.98.233/udp/4001/quic/p2p/QmasbWJE9C7PVFVj1CVQLX617CrDQijCxMv6ajkRfaTi98",
		},
	}
}

// Init creates the required directories
// and saves the config if it does not exist.
func (c *Config) Init() error {
	// create default directories
	err := os.MkdirAll(c.rootPath, 0755)
	if err != nil {
		return err
	}
	// save config if it does not exist
	_, err = os.Stat(c.Path())
	if err == nil || !os.IsNotExist(err) {
		return err
	}
	return c.Save()
}

// Load loads the config from the root path.
func (c *Config) Load() error {
	data, err := os.ReadFile(c.Path())
	if err != nil {
		return err
	}
	return json.Unmarshal(data, c)
}

// Save writes the config to the root path.
func (c *Config) Save() error {
	data, err := json.MarshalIndent(c, "", "\t")
	if err != nil {
		return err
	}
	return os.WriteFile(c.Path(), data, 0666)
}

// Path returns the config file path.
func (c *Config) Path() string {
	return filepath.Join(c.rootPath, "config")
}

// InstallPath returns the path to install binaries.
func (c *Config) InstallPath() string {
	return filepath.Join(c.rootPath, "bin")
}

// KeyStorePath returns the keystore directory path.
func (c *Config) KeyStorePath() string {
	return filepath.Join(c.rootPath, c.ContractType, "keystore")
}

// StoragePath returns the storage directory path.
func (c *Config) StoragePath() string {
	return filepath.Join(c.rootPath, c.StorageType, "storage")
}

// SetDefaultAccount sets the default account for the current contract type.
func (c *Config) SetDefaultAccount(account string) {
	c.DefaultAccounts[c.ContractType] = account
}

// GetDefaultAccount gets the default account for the current contract type.
func (c *Config) GetDefaultAccount() string {
	return c.DefaultAccounts[c.ContractType]
}
