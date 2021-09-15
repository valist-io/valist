package core

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/ethclient"
	httpapi "github.com/ipfs/go-ipfs-http-client"
	"github.com/libp2p/go-libp2p-core/peer"
	ma "github.com/multiformats/go-multiaddr"

	"github.com/valist-io/valist/internal/contract"
	"github.com/valist-io/valist/internal/core/client"
	"github.com/valist-io/valist/internal/core/client/basetx"
	"github.com/valist-io/valist/internal/core/client/metatx"
	"github.com/valist-io/valist/internal/core/config"
	"github.com/valist-io/valist/internal/core/signer"
	"github.com/valist-io/valist/internal/storage/ipfs"
)

type contextKey string

const (
	ClientKey = contextKey("client")
	ConfigKey = contextKey("config")
)

func newClientOpts(ctx context.Context, cfg *config.Config, account accounts.Account) (*client.Options, error) {
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

	signer := signer.NewSigner(chainID, cfg.KeyStore())

	valist, err := contract.NewValist(valistAddress, eth)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize valist contract: %v", err)
	}

	registry, err := contract.NewRegistry(registryAddress, eth)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize registry contract: %v", err)
	}

	ipfsapi, err := httpapi.NewURLApiWithClient(cfg.IPFS.API, &http.Client{})
	if err != nil {
		return nil, err
	}

	for _, peerString := range cfg.IPFS.Peers {
		peerAddr, err := ma.NewMultiaddr(peerString)
		if err != nil {
			continue
		}

		peerInfo, err := peer.AddrInfoFromP2pAddr(peerAddr)
		if err != nil {
			continue
		}

		go ipfsapi.Swarm().Connect(ctx, *peerInfo) //nolint:errcheck
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

	return &client.Options{
		Storage:    ipfs.NewStorage(ipfsapi),
		Ethereum:   eth,
		Valist:     valist,
		Registry:   registry,
		Account:    account,
		Signer:     signer,
		Transactor: transactor,
	}, nil
}

// NewClient builds a client based on the given config.
func NewClient(ctx context.Context, cfg *config.Config, account accounts.Account) (*client.Client, error) {
	opts, err := newClientOpts(ctx, cfg, account)
	if err != nil {
		return nil, err
	}
	return client.NewClient(opts)
}

// NewClientWithPassphrase builds a client based on the given config and unlocks the signer.
func NewClientWithPassphrase(ctx context.Context, cfg *config.Config, account accounts.Account, passphrase string) (*client.Client, error) {
	opts, err := newClientOpts(ctx, cfg, account)
	if err != nil {
		return nil, err
	}
	if err := opts.Signer.Unlock(account, passphrase); err != nil {
		return nil, err
	}
	return client.NewClient(opts)
}
