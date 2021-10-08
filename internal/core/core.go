package core

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/valist-io/valist/internal/contract"
	"github.com/valist-io/valist/internal/core/client"
	"github.com/valist-io/valist/internal/core/client/basetx"
	"github.com/valist-io/valist/internal/core/client/metatx"
	"github.com/valist-io/valist/internal/core/config"
	"github.com/valist-io/valist/internal/signer"
	"github.com/valist-io/valist/internal/storage"
	"github.com/valist-io/valist/internal/storage/estuary"
	"github.com/valist-io/valist/internal/storage/ipfs"
)

type contextKey string

const (
	ClientKey = contextKey("client")
	ConfigKey = contextKey("config")
)

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

	signer, err := signer.NewSigner(chainID, cfg.KeyStore())
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
	estuary := estuary.NewProvider("https://pin-proxy-rkl5i.ondigitalocean.app", "")
	storage, err := storage.NewStorage(estuary, ipfs)
	if err != nil {
		return nil, err
	}

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
		Storage:    storage,
		Ethereum:   eth,
		Valist:     valist,
		Registry:   registry,
		Signer:     signer,
		Transactor: transactor,
	})
}
