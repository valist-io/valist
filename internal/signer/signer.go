package signer

import (
	"net"

	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/signer/core"
	"github.com/ethereum/go-ethereum/signer/fourbyte"
	"github.com/ethereum/go-ethereum/signer/storage"

	"github.com/valist-io/registry/internal/config"
)

func NewSignerAPI(cfg *config.Config) (*core.SignerAPI, error) {
	validator, err := fourbyte.New()
	if err != nil {
		return nil, err
	}

	ksLocation := cfg.Signer.KeyStorePath
	noUSB := cfg.Signer.NoUSB
	lightKDF := cfg.Signer.LightKDF
	scpath := cfg.Signer.SmartCardPath
	chainID := cfg.Ethereum.ChainID.Int64()
	advancedMode := cfg.Signer.AdvancedMode
	credentials := &storage.NoStorage{}

	ui := core.NewCommandlineUI()
	am := core.StartClefAccountManager(ksLocation, noUSB, lightKDF, scpath)
	return core.NewSignerAPI(am, chainID, noUSB, ui, validator, advancedMode, credentials), nil
}

func StartIPCEndpoint(cfg *config.Config) (net.Listener, *rpc.Server, error) {
	api, err := NewSignerAPI(cfg)
	if err != nil {
		return nil, nil, err
	}

	rpcAPI := []rpc.API{{
		Namespace: "account",
		Public:    true,
		Service:   api,
		Version:   "1.0",
	}}

	return rpc.StartIPCEndpoint(cfg.Signer.IPCAddress, rpcAPI)
}
