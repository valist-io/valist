package core

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/valist-io/valist/contract"
	"github.com/valist-io/valist/core/client"
	"github.com/valist-io/valist/core/client/basetx"
	"github.com/valist-io/valist/core/client/metatx"
	"github.com/valist-io/valist/core/config"
	"github.com/valist-io/valist/signer"
	"github.com/valist-io/valist/storage"
	"github.com/valist-io/valist/storage/estuary"
	"github.com/valist-io/valist/storage/ipfs"
)

// NewClient creates a new valist client using the given config.
func NewClient(ctx context.Context, cfg *config.Config) (*client.Client, error) {
	valistAddress := cfg.Ethereum.Contracts["valist"]
	registryAddress := cfg.Ethereum.Contracts["registry"]

	eth, err := ethclient.Dial(cfg.Ethereum.RPC)
	if err != nil {
		return nil, err
	}

	chainID, err := eth.ChainID(ctx)
	if err != nil {
		return nil, err
	}

	kstore := keystore.NewKeyStore(cfg.KeyStorePath(), keystore.StandardScryptN, keystore.StandardScryptP)
	signer, err := signer.NewSigner(chainID, kstore)
	if err != nil {
		return nil, err
	}

	valist, err := contract.NewValist(valistAddress, eth)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize valist contract: %v", err)
	}

	registry, err := contract.NewRegistry(registryAddress, eth)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize registry contract: %v", err)
	}

	ipfs, err := ipfs.NewProvider(ctx, cfg.StoragePath())
	if err != nil {
		return nil, err
	}

	// TODO move to config once URL is proxied
	var provider storage.Provider
	provider = estuary.NewProvider("https://pin-proxy-rkl5i.ondigitalocean.app", "", ipfs)
	provider = storage.WithTimeout(provider, 5*time.Minute, 5*time.Minute)

	var transactor client.TransactorAPI
	if cfg.Ethereum.MetaTx {
		transactor, err = metatx.NewTransactor(eth, valistAddress, registryAddress, cfg.Ethereum.BiconomyApiKey)
	} else {
		transactor, err = basetx.NewTransactor(eth, valistAddress, registryAddress)
	}

	if err != nil {
		return nil, err
	}

	return client.NewClient(client.Options{
		Storage:    provider,
		Ethereum:   eth,
		Valist:     valist,
		Registry:   registry,
		Signer:     signer,
		Transactor: transactor,
	})
}
