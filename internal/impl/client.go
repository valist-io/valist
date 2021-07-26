package impl

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	httpapi "github.com/ipfs/go-ipfs-http-client"
	coreiface "github.com/ipfs/interface-go-ipfs-core"
	"github.com/libp2p/go-libp2p-core/peer"
	ma "github.com/multiformats/go-multiaddr"

	"github.com/valist-io/registry/internal/contract"
	"github.com/valist-io/registry/internal/contract/registry"
	"github.com/valist-io/registry/internal/contract/valist"
	"github.com/valist-io/registry/internal/core"
)

var _ core.CoreAPI = (*Client)(nil)

var (
	chainID           = big.NewInt(80001)
	ethereumRPC       = "https://rpc.valist.io"
	ipfsAPI           = ma.StringCast("/dns/pin.valist.io")
	valistPeerAddress = ma.StringCast("/ip4/107.191.98.233/tcp/4001/p2p/QmasbWJE9C7PVFVj1CVQLX617CrDQijCxMv6ajkRfaTi98")
	valistAddress     = common.HexToAddress("0xA7E4124aDBBc50CF402e4Cad47de906a14daa0f6")
	registryAddress   = common.HexToAddress("0x2Be6D782dBA2C52Cd0a41c6052e914dCaBcCD78e")
	emptyHash         = common.HexToHash("0x0")
	emptyAddress      = common.HexToAddress("0x0")
)

// Client is a Valist SDK client.
type Client struct {
	eth      bind.DeployBackend
	ipfs     coreiface.CoreAPI
	orgs     map[string]common.Hash
	valist   *valist.Valist
	registry *registry.ValistRegistry
	chainID  *big.Int
	private  *ecdsa.PrivateKey
}

// NewClient returns a Client with default settings.
func NewClient(ctx context.Context) (core.CoreAPI, error) {
	ipfs, err := httpapi.NewApi(ipfsAPI)
	if err != nil {
		return nil, err
	}

	peerInfo, err := peer.AddrInfoFromP2pAddr(valistPeerAddress)
	if err != nil {
		return nil, err
	}

	// attempt to connect to valist gateway peer
	go ipfs.Swarm().Connect(ctx, *peerInfo)

	eth, err := ethclient.Dial(ethereumRPC)
	if err != nil {
		return nil, err
	}

	valist, err := contract.NewValist(valistAddress, eth)
	if err != nil {
		return nil, err
	}

	registry, err := contract.NewRegistry(registryAddress, eth)
	if err != nil {
		return nil, err
	}

	return &Client{
		eth:      eth,
		ipfs:     ipfs,
		orgs:     make(map[string]common.Hash),
		valist:   valist,
		registry: registry,
		chainID:  chainID,
	}, nil
}
