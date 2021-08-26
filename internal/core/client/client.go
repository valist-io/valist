package client

import (
	"context"
	"fmt"
	"math/big"
	"net"
	"net/http"

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
	"github.com/valist-io/registry/internal/signer"
)

var _ core.CoreAPI = (*Client)(nil)

var emptyHash = common.HexToHash("0x0")

// Client is a Valist SDK client.
type Client struct {
	eth  bind.DeployBackend
	ipfs coreiface.CoreAPI

	orgs map[string]common.Hash

	valist   *valist.Valist
	registry *registry.ValistRegistry

	chainID *big.Int
	metaTx  bool

	wallet  accounts.Wallet
	account accounts.Account

	transactor core.TransactorAPI
	listener   net.Listener
}

// NewClient create a client with a base transactor.
func NewClient(ctx context.Context, cfg *config.Config, account accounts.Account) (*Client, error) {
	listener, _, err := signer.StartIPCEndpoint(cfg)
	if err != nil {
		return nil, err
	}

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

	ipfs, err := httpapi.NewURLApiWithClient(cfg.IPFS.API, &http.Client{})
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
		wallet:     wallet,
		account:    account,
		transactor: transactor,
		listener:   listener,
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

// Close releases client resources.
func (client *Client) Close() {
	if client.listener != nil {
		client.listener.Close()
	}

	if eth, ok := client.eth.(*ethclient.Client); ok {
		eth.Close()
	}
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
