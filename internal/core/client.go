package core

import (
	"crypto/ecdsa"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	coreiface "github.com/ipfs/interface-go-ipfs-core"

	"github.com/valist-io/registry/internal/contract/registry"
	"github.com/valist-io/registry/internal/contract/valist"
)

var (
	emptyHash    = common.HexToHash("0x0")
	emptyAddress = common.HexToAddress("0x0")
)

var (
	ErrOrganizationNotExist = errors.New("Organization does not exist")
	ErrRepositoryNotExist   = errors.New("Repository does not exist")
	ErrReleaseNotExist      = errors.New("Release does not exist")
)

// Client is a Valist SDK client.
type Client struct {
	eth      bind.DeployBackend
	ipfs     coreiface.CoreAPI
	orgs     map[string]common.Hash
	valist   *valist.Valist
	registry *registry.ValistRegistry
	private  *ecdsa.PrivateKey
	chainID  *big.Int
}

// NewClient returns a Client with default settings.
func NewClient(
	eth bind.DeployBackend,
	ipfs coreiface.CoreAPI,
	valist *valist.Valist,
	registry *registry.ValistRegistry,
	private *ecdsa.PrivateKey,
	chainID *big.Int) *Client {
	// TODO build this from a config file
	return &Client{
		eth:      eth,
		ipfs:     ipfs,
		orgs:     make(map[string]common.Hash),
		valist:   valist,
		registry: registry,
		private:  private,
		chainID:  chainID,
	}
}
