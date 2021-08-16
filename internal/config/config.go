package config

import (
	"encoding/json"
	"math/big"
	"os"
	"path/filepath"

	"github.com/ethereum/go-ethereum/common"
)

const (
	rootDir    = ".valist"
	configFile = "config"
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

type Signer struct {
	// AdvancedMode allows warning instead of rejecting.
	AdvancedMode bool `json:"advanced_mode"`
	// LightKDF enables faster KDF for low power devices.
	LightKDF bool `json:"light_kdf"`
	// NoUSB disables usb signer devices.
	NoUSB bool `json:"no_usb"`
	// SmartCardPath enables smart card signing.
	SmartCardPath string `json:"smart_card_path"`
	// KeyStorePath is the path to the key store.
	KeyStorePath string `json:"key_store_path"`
	// IPCAddress is the signer ipc address.
	IPCAddress string `json:"ipc_address"`
}

type Accounts struct {
	Pinned  []common.Address `json:"pinned,omitempty"`
	Default common.Address   `json:"default,omitempty"`
}

type Config struct {
	rootPath string
	Accounts Accounts `json:"accounts"`
	Ethereum Ethereum `json:"ethereum"`
	IPFS     IPFS     `json:"ipfs"`
	Signer   Signer   `json:"signer"`
}

// Default returns a config with default settings.
func Default(rootPath string) Config {
	return Config{
		rootPath,
		Accounts{},
		Ethereum{
			RPC:     "https://rpc.valist.io",
			ChainID: big.NewInt(80001),
			Contracts: map[string]common.Address{
				"valist":    common.HexToAddress("0xA7E4124aDBBc50CF402e4Cad47de906a14daa0f6"),
				"registry":  common.HexToAddress("0x2Be6D782dBA2C52Cd0a41c6052e914dCaBcCD78e"),
				"forwarder": common.HexToAddress("0x9399BB24DBB5C4b782C70c2969F58716Ebbd6a3b"),
			},
		},
		IPFS{
			API: "/dns/pin.valist.io",
			Peers: []string{
				"/ip4/107.191.98.233/tcp/4001/p2p/QmasbWJE9C7PVFVj1CVQLX617CrDQijCxMv6ajkRfaTi98",
			},
		},
		Signer{
			AdvancedMode:  false,
			LightKDF:      false,
			NoUSB:         false,
			SmartCardPath: "",
			KeyStorePath:  filepath.Join(rootPath, "keystore"),
			IPCAddress:    filepath.Join(rootPath, "signer.ipc"),
		},
	}
}

// Exists returns true of the config root exists.
func Exists(path string) (bool, error) {
	info, err := os.Stat(filepath.Join(path, rootDir))
	if os.IsNotExist(err) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return info.IsDir(), nil
}

// Init initializes a config with a default account.
func Init(path string) error {
	rootPath := filepath.Join(path, rootDir)
	if err := os.Mkdir(rootPath, 0755); err != nil {
		return err
	}

	config := Default(rootPath)
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
