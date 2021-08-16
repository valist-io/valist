package client

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/external"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	httpapi "github.com/ipfs/go-ipfs-http-client"
	coreiface "github.com/ipfs/interface-go-ipfs-core"
	"github.com/libp2p/go-libp2p-core/peer"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/valist-io/gasless"
	"github.com/valist-io/gasless/mexa"

	"github.com/valist-io/registry/internal/config"
	"github.com/valist-io/registry/internal/contract"
	"github.com/valist-io/registry/internal/contract/registry"
	"github.com/valist-io/registry/internal/contract/valist"
	"github.com/valist-io/registry/internal/core"
	"github.com/valist-io/registry/internal/core/basetx"
	"github.com/valist-io/registry/internal/core/metatx"
)

var _ core.CoreAPI = (*Client)(nil)

var emptyHash = common.HexToHash("0x0")

// Client is a Valist SDK client.
type Client struct {
	eth  bind.DeployBackend
	ipfs coreiface.CoreAPI

	orgs map[string]common.Hash

	chainID  *big.Int
	valist   *valist.Valist
	registry *registry.ValistRegistry

	metaTx bool

	wallet  accounts.Wallet
	account accounts.Account

	transactor core.TransactorAPI
}

// NewClient create a client with a base transactor.
func NewClient(ctx context.Context, cfg *config.Config, account accounts.Account) (*Client, error) {
	signer, err := external.NewExternalSigner(cfg.Signer.IPCAddress)
	if err != nil {
		return nil, err
	}

	eth, err := ethclient.Dial(cfg.Ethereum.RPC)
	if err != nil {
		return nil, err
	}

	valistAddr, ok := cfg.Ethereum.Contracts["valist"]
	if !ok {
		return nil, fmt.Errorf("Valist contract address required")
	}

	registryAddr, ok := cfg.Ethereum.Contracts["registry"]
	if !ok {
		return nil, fmt.Errorf("Registry contract address required")
	}

	valist, err := contract.NewValist(valistAddr, eth)
	if err != nil {
		return nil, err
	}

	registry, err := contract.NewRegistry(registryAddr, eth)
	if err != nil {
		return nil, err
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

	// attempt to add all IPFS peers to swarm
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

	transactor := basetx.NewTransactor(valist, registry)

	return &Client{
		eth:        eth,
		ipfs:       ipfs,
		orgs:       make(map[string]common.Hash),
		chainID:    cfg.Ethereum.ChainID,
		valist:     valist,
		registry:   registry,
		wallet:     signer,
		account:    account,
		transactor: transactor,
	}, nil
}

// NewClientWithMetaTx creates a client with a metatx transactor.
func NewClientWithMetaTx(ctx context.Context, cfg *config.Config, account accounts.Account) (*Client, error) {
	client, err := NewClient(ctx, cfg, account)
	if err != nil {
		return nil, err
	}

	eth, ok := client.eth.(*ethclient.Client)
	if !ok {
		return nil, fmt.Errorf("cannot create metatx transactor with simulated backend")
	}

	meta, err := mexa.NewMexa(ctx, eth, cfg.Ethereum.BiconomyApiKey)
	if err != nil {
		return nil, err
	}

	signer := gasless.NewWalletSigner(client.account, client.wallet)
	client.transactor = metatx.NewTransactor(client.transactor, meta, signer)

	return client, nil
}

// TransactOpts creates a transaction signer.
func (client *Client) TransactOpts() *bind.TransactOpts {
	return &bind.TransactOpts{
		From:   client.account.Address,
		NoSend: client.metaTx,
		Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			if address != client.account.Address {
				return nil, bind.ErrNotAuthorized
			}

			if client.metaTx {
				return tx, nil
			}

			return client.wallet.SignTx(client.account, tx, client.chainID)
		},
	}
}
