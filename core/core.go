package core

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/valist-io/valist/contract"
	"github.com/valist-io/valist/core/client"
	"github.com/valist-io/valist/core/client/basetx"
	"github.com/valist-io/valist/core/client/metatx"
	"github.com/valist-io/valist/core/config"
	"github.com/valist-io/valist/ipfs"
	"github.com/valist-io/valist/signer"
)

// NewClient creates a new valist client using the given config.
func NewClient(ctx context.Context, cfg *config.Config) (*client.Client, error) {
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

	valistAddress := cfg.Ethereum.Contracts["valist"]
	valist, err := contract.NewValist(valistAddress, eth)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize valist contract: %v", err)
	}

	registryAddress := cfg.Ethereum.Contracts["registry"]
	registry, err := contract.NewRegistry(registryAddress, eth)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize registry contract: %v", err)
	}

	coreAPI, err := ipfs.NewCoreAPI(ctx, cfg.StoragePath())
	if err != nil {
		return nil, err
	}
	ipfs.Bootstrap(ctx, coreAPI)

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
		Ethereum:   eth,
		IPFS:       coreAPI,
		Signer:     signer,
		Valist:     valist,
		Registry:   registry,
		Transactor: transactor,
	})
}
