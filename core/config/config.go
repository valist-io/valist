package config

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/ethereum/go-ethereum/common"
)

const (
	rootDir        = ".valist"
	installDir     = "bin"
	configFile     = "config"
	keystoreDir    = "keystore"
	storageDir     = "storage"
	TelemetryAllow = "allow"
	TelemetryDeny  = "deny"
	TelemetryNone  = "none"
)

type IPFS struct {
	// ApiAddr is the address of the IPFS node
	ApiAddr string `json:"api_address"`
	// BootstrapPeers is a list of peers to connect to
	BootstrapPeers []string `json:"bootstrap_peers"`
}

type Ethereum struct {
	// BiconomyApiKey is the mexa public api key.
	BiconomyApiKey string `json:"biconomy_api_key"`
	// Contracts is a mapping of contract addresses.
	Contracts map[string]common.Address `json:"contracts"`
	// MetaTx enables meta transactions.
	MetaTx bool `json:"meta_tx"`
	// RPC is the ethereum rpc address.
	RPC string `json:"rpc"`
}

type Accounts struct {
	// Pinned is a list of all accounts.
	Pinned []common.Address `json:"pinned,omitempty"`
	// Default is the default account.
	Default common.Address `json:"default,omitempty"`
}

type HTTP struct {
	// ApiAddr is the api server address to use
	ApiAddr string `json:"api_address"`
	// WebAddr is the static web server address to use
	WebAddr string `json:"web_address"`
}

type Config struct {
	rootPath  string
	Accounts  Accounts `json:"accounts"`
	Ethereum  Ethereum `json:"ethereum"`
	IPFS      IPFS     `json:"ipfs"`
	HTTP      HTTP     `json:"http"`
	Telemetry string   `json:"telemetry"`
}

// NewConfig returns a config with default settings.
func NewConfig(path string) *Config {
	return &Config{
		rootPath: filepath.Join(path, rootDir),
		Accounts: Accounts{},
		IPFS: IPFS{
			BootstrapPeers: []string{
				"/ip4/107.191.98.233/tcp/4001/p2p/QmasbWJE9C7PVFVj1CVQLX617CrDQijCxMv6ajkRfaTi98",
				"/ip4/107.191.98.233/udp/4001/quic/p2p/QmasbWJE9C7PVFVj1CVQLX617CrDQijCxMv6ajkRfaTi98",
			},
		},
		Ethereum: Ethereum{
			BiconomyApiKey: "qLW9TRUjQ.f77d2f86-c76a-4b9c-b1ee-0453d0ead878",
			Contracts: map[string]common.Address{
				"valist":   common.HexToAddress("0xA7E4124aDBBc50CF402e4Cad47de906a14daa0f6"),
				"registry": common.HexToAddress("0x2Be6D782dBA2C52Cd0a41c6052e914dCaBcCD78e"),
			},
			MetaTx: true,
			RPC:    "https://rpc.valist.io",
		},
		HTTP: HTTP{
			ApiAddr: "localhost:9000",
		},
		Telemetry: TelemetryNone,
	}
}

// Initialize creates a config with default settings if one does not exist.
func Initialize(path string) error {
	rootPath := filepath.Join(path, rootDir)
	confPath := filepath.Join(rootPath, configFile)

	_, err := os.Stat(confPath)
	if err == nil || !os.IsNotExist(err) {
		return err
	}

	if err := os.MkdirAll(rootPath, 0755); err != nil {
		return err
	}

	return NewConfig(path).Save()
}

// Load loads the config from the root path.
func (c *Config) Load() error {
	path := filepath.Join(c.rootPath, configFile)

	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, c)
}

// Save writes the config to the root path.
func (c *Config) Save() error {
	path := filepath.Join(c.rootPath, configFile)

	data, err := json.MarshalIndent(c, "", "\t")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0666)
}

// KeyStorePath returns the keystore directory path.
func (c *Config) KeyStorePath() string {
	return filepath.Join(c.rootPath, keystoreDir)
}

// InstallPath returns the path to install binaries.
func (c *Config) InstallPath() string {
	return filepath.Join(c.rootPath, installDir)
}

// StoragePath is the path to the IPFS repo storage.
func (c *Config) StoragePath() string {
	return filepath.Join(c.rootPath, storageDir)
}
