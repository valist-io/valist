package signer

import (
	"encoding/hex"
	"net"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/signer/core"
	"github.com/ethereum/go-ethereum/signer/fourbyte"
	"github.com/ethereum/go-ethereum/signer/storage"

	"github.com/valist-io/registry/internal/contract/registry"
	"github.com/valist-io/registry/internal/contract/valist"
	"github.com/valist-io/registry/internal/core/config"
)

func NewSigner(cfg *config.Config) (*core.SignerAPI, *accounts.Manager, error) {
	validator, err := fourbyte.New()
	if err != nil {
		return nil, nil, err
	}

	for fourbytes, signature := range valist.ValistFuncSigs {
		bytes, err := hex.DecodeString(fourbytes)
		if err != nil {
			return nil, nil, err
		}

		validator.AddSelector(signature, bytes)
		if err != nil {
			return nil, nil, err
		}
	}

	for fourbytes, signature := range registry.ValistRegistryFuncSigs {
		bytes, err := hex.DecodeString(fourbytes)
		if err != nil {
			return nil, nil, err
		}

		validator.AddSelector(signature, bytes)
		if err != nil {
			return nil, nil, err
		}
	}

	ksLocation := cfg.Signer.KeyStorePath
	noUSB := cfg.Signer.NoUSB
	lightKDF := cfg.Signer.LightKDF
	scpath := cfg.Signer.SmartCardPath
	chainID := cfg.Ethereum.ChainID.Int64()
	advancedMode := cfg.Signer.AdvancedMode
	credentials := &storage.NoStorage{}

	ui := core.NewCommandlineUI()
	manager := core.StartClefAccountManager(ksLocation, noUSB, lightKDF, scpath)
	api := core.NewSignerAPI(manager, chainID, noUSB, ui, validator, advancedMode, credentials)

	return api, manager, nil
}

func StartIPCEndpoint(cfg *config.Config) (net.Listener, *rpc.Server, error) {
	api, _, err := NewSigner(cfg)
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
