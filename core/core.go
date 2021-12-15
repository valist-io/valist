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
	var opt client.Options

	eth, err := ethclient.Dial(cfg.Ethereum.RPC)
	if err != nil {
		return nil, err
	}
	opt.Ethereum = eth

	chainID, err := eth.ChainID(ctx)
	if err != nil {
		return nil, err
	}

	kstore := keystore.NewKeyStore(cfg.KeyStorePath(), keystore.StandardScryptN, keystore.StandardScryptP)
	opt.Signer, err = signer.NewSigner(chainID, kstore)
	if err != nil {
		return nil, err
	}

	valistAddress := cfg.Ethereum.Contracts["valist"]
	opt.Valist, err = contract.NewValist(valistAddress, eth)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize valist contract: %v", err)
	}

	registryAddress := cfg.Ethereum.Contracts["registry"]
	opt.Registry, err = contract.NewRegistry(registryAddress, eth)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize registry contract: %v", err)
	}

	opt.IPFS, err = ipfs.NewCoreAPI(ctx, cfg.StoragePath())
	if err != nil {
		return nil, err
	}
	ipfs.Bootstrap(ctx, opt.IPFS)

	if cfg.Ethereum.MetaTx {
		opt.Transactor, err = metatx.NewTransactor(eth, valistAddress, registryAddress, cfg.Ethereum.BiconomyApiKey)
	} else {
		opt.Transactor, err = basetx.NewTransactor(eth, valistAddress, registryAddress)
	}

	if err != nil {
		return nil, err
	}

	return client.NewClient(opt)
}
