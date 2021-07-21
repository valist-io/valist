package core

import (
	"context"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	httpapi "github.com/ipfs/go-ipfs-http-client"
	"github.com/libp2p/go-libp2p-core/peer"
	ma "github.com/multiformats/go-multiaddr"

	"github.com/valist-io/registry/internal/contract/valist"
	"github.com/valist-io/registry/internal/contract/valist/registry"
)

const defaultRPC = "https://rpc.valist.io"

var (
	emptyHash                     = common.HexToHash("0x0000000000000000000000000000000000000000")
	valistPeerAddress             = ma.StringCast("/ip4/107.191.98.233/tcp/4001/p2p/QmasbWJE9C7PVFVj1CVQLX617CrDQijCxMv6ajkRfaTi98")
	valistContractAddress         = common.HexToAddress("0xA7E4124aDBBc50CF402e4Cad47de906a14daa0f6")
	valistRegistryContractAddress = common.HexToAddress("0x2Be6D782dBA2C52Cd0a41c6052e914dCaBcCD78e")
)

var (
	ErrOrganizationNotExist = errors.New("Organization does not exist")
	ErrRepositoryNotExist   = errors.New("Repository does not exist")
	ErrReleaseNotExist      = errors.New("Release does not exist")
)

// Client is a Valist SDK client.
type Client struct {
	eth              *ethclient.Client
	ipfs             *httpapi.HttpApi
	orgs             map[string]common.Hash
	valistContract   *valist.Valist
	registryContract *registry.ValistRegistry
}

// NewClient returns a Client with default settings.
func NewClient() (*Client, error) {
	ipfs, err := httpapi.NewLocalApi()
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to ipfs: %v", err)
	}

	// TODO move this to bootstrap peer list
	peerInfo, err := peer.AddrInfoFromP2pAddr(valistPeerAddress)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse valist peer: %v", err)
	}

	if err := ipfs.Swarm().Connect(context.TODO(), *peerInfo); err != nil {
		return nil, err
	}

	eth, err := ethclient.Dial(defaultRPC)
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to rpc: %v", err)
	}

	valistContract, err := valist.NewValist(valistContractAddress, eth)
	if err != nil {
		return nil, fmt.Errorf("Failed to instantiate valist contract: %v", err)
	}

	registryContract, err := registry.NewValistRegistry(valistRegistryContractAddress, eth)
	if err != nil {
		return nil, fmt.Errorf("Failed to instantiate valist registry contract: %v", err)
	}

	return &Client{
		eth:              eth,
		ipfs:             ipfs,
		orgs:             make(map[string]common.Hash),
		valistContract:   valistContract,
		registryContract: registryContract,
	}, nil
}
