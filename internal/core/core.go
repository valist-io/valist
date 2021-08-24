package core

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/external"
	"github.com/ethereum/go-ethereum/ethclient"
	httpapi "github.com/ipfs/go-ipfs-http-client"
	"github.com/libp2p/go-libp2p-core/peer"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/valist-io/gasless"
	"github.com/valist-io/gasless/mexa"

	"github.com/valist-io/registry/internal/contract"
	"github.com/valist-io/registry/internal/core/client"
	"github.com/valist-io/registry/internal/core/client/basetx"
	"github.com/valist-io/registry/internal/core/client/metatx"
	"github.com/valist-io/registry/internal/core/config"
	"github.com/valist-io/registry/internal/signer"
)

// NewClient builds a client based on the given config.
func NewClient(ctx context.Context, cfg *config.Config, account accounts.Account) (*client.Client, error) {
	var onClose []client.Close

	listener, _, err := signer.StartIPCEndpoint(cfg)
	if err != nil {
		return nil, err
	}
	onClose = append(onClose, listener.Close)

	wallet, err := external.NewExternalSigner(cfg.Signer.IPCAddress)
	if err != nil {
		return nil, err
	}

	eth, err := ethclient.Dial(cfg.Ethereum.RPC)
	if err != nil {
		return nil, err
	}

	valist, err := contract.NewValist(cfg.Ethereum.Contracts["valist"], eth)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize valist contract: %v", err)
	}

	registry, err := contract.NewRegistry(cfg.Ethereum.Contracts["registry"], eth)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize registry contract: %v", err)
	}

	// TODO redirects do not work
	// ipfsAPI, err := ma.NewMultiaddr(cfg.IPFS.API)
	// if err != nil {
	// 	return nil, err
	// }

	// ipfs, err := httpapi.NewApi(ipfsAPI)
	// if err != nil {
	// 	return nil, err
	// }

	ipfs, err := httpapi.NewLocalApi()
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

		go ipfs.Swarm().Connect(ctx, *peerInfo) //nolint:errcheck
	}

	opts := &client.Options{
		IPFS:         ipfs,
		Ethereum:     eth,
		ChainID:      cfg.Ethereum.ChainID,
		Valist:       valist,
		Registry:     registry,
		Account:      account,
		Wallet:       wallet,
		TransactOpts: basetx.TransactOpts,
		Transactor:   basetx.NewTransactor(valist, registry),
		OnClose:      onClose,
	}

	if !cfg.Ethereum.MetaTx {
		return client.NewClient(opts)
	}

	meta, err := mexa.NewMexa(ctx, eth, cfg.Ethereum.BiconomyApiKey)
	if err != nil {
		return nil, err
	}

	signer := gasless.NewWalletSigner(opts.Account, opts.Wallet)
	opts.TransactOpts = metatx.TransactOpts
	opts.Transactor = metatx.NewTransactor(opts.Transactor, meta, signer)

	return client.NewClient(opts)
}
