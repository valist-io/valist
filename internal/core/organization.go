package core

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ipfs/go-cid"
	files "github.com/ipfs/go-ipfs-files"
	"github.com/ipfs/interface-go-ipfs-core/path"

	"github.com/valist-io/registry/internal/contract/registry"
	"github.com/valist-io/registry/internal/contract/valist"
)

type OrganizationMeta struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Organization struct {
	ID            common.Hash
	Threshold     *big.Int
	ThresholdDate *big.Int
	MetaCID       cid.Cid
}

type CreateOrgResult struct {
	Log *valist.ValistOrgCreated
	Err error
}

type LinkOrgNameResult struct {
	Log *registry.ValistRegistryMappingEvent
	Err error
}

// GerOrganizationID returns the ID of the organization with the given name.
func (client *Client) GetOrganizationID(ctx context.Context, name string) (common.Hash, error) {
	if orgID, ok := client.orgs[name]; ok {
		return orgID, nil
	}

	callopts := bind.CallOpts{Context: ctx}
	orgID, err := client.registry.NameToID(&callopts, name)
	if err != nil {
		return emptyHash, fmt.Errorf("Failed to get organization id: %v", err)
	}

	if bytes.Equal(orgID[:], emptyHash.Bytes()) {
		return emptyHash, ErrOrganizationNotExist
	}

	client.orgs[name] = orgID
	return orgID, nil
}

// GetOrganization returns the organization with the given ID.
func (client *Client) GetOrganization(ctx context.Context, id common.Hash) (*Organization, error) {
	callopts := bind.CallOpts{Context: ctx}
	org, err := client.valist.Orgs(&callopts, id)
	if err != nil {
		return nil, fmt.Errorf("Failed to get organization id: %v", err)
	}

	// TODO there's no way to check if an org exists

	metaCID, err := cid.Decode(org.MetaCID)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse organization meta CID: %v", err)
	}

	return &Organization{
		ID:            id,
		Threshold:     org.Threshold,
		ThresholdDate: org.ThresholdDate,
		MetaCID:       metaCID,
	}, nil
}

// GetOrganizationMeta returns the organization meta with the given CID.
func (client *Client) GetOrganizationMeta(ctx context.Context, id cid.Cid) (*OrganizationMeta, error) {
	node, err := client.ipfs.Unixfs().Get(ctx, path.IpfsPath(id))
	if err != nil {
		return nil, err
	}

	file, ok := node.(files.File)
	if !ok {
		return nil, fmt.Errorf("Failed to parse organization meta")
	}

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("Failed to get organization meta: %v", err)
	}

	var meta OrganizationMeta
	if err := json.Unmarshal(data, &meta); err != nil {
		return nil, fmt.Errorf("Failed to parse organization meta: %v", err)
	}

	return &meta, nil
}

// CreateOrganization creates a new organization with the given meta and returns the orgID.
func (client *Client) CreateOrganization(ctx context.Context, meta *OrganizationMeta) (<-chan CreateOrgResult, error) {
	data, err := json.Marshal(meta)
	if err != nil {
		return nil, err
	}

	path, err := client.ipfs.Unixfs().Add(ctx, files.NewBytesFile(data))
	if err != nil {
		return nil, err
	}

	txopts, err := bind.NewKeyedTransactorWithChainID(client.private, client.chainID)
	if err != nil {
		return nil, err
	}
	txopts.Context = ctx

	tx, err := client.valist.CreateOrganization(txopts, path.Cid().String())
	if err != nil {
		return nil, err
	}

	result := make(chan CreateOrgResult, 1)
	go client.createOrganization(ctx, tx, result)

	return result, nil
}

func (client *Client) createOrganization(ctx context.Context, tx *types.Transaction, result chan<- CreateOrgResult) {
	defer close(result)

	receipt, err := bind.WaitMined(ctx, client.eth, tx)
	if err != nil {
		result <- CreateOrgResult{Err: err}
		return
	}

	if len(receipt.Logs) == 0 {
		result <- CreateOrgResult{Err: fmt.Errorf("Failed to parse log")}
		return
	}

	log, err := client.valist.ParseOrgCreated(*receipt.Logs[0])
	if err != nil {
		result <- CreateOrgResult{Err: err}
		return
	}

	result <- CreateOrgResult{Log: log}
}

// LinkOrganizationName creates a link from the given orgID to the given name.
func (client *Client) LinkOrganizationName(ctx context.Context, orgID common.Hash, name string) (<-chan LinkOrgNameResult, error) {
	// TODO validate name

	txopts, err := bind.NewKeyedTransactorWithChainID(client.private, client.chainID)
	if err != nil {
		return nil, err
	}
	txopts.Context = ctx

	tx, err := client.registry.LinkNameToID(txopts, orgID, name)
	if err != nil {
		return nil, err
	}

	result := make(chan LinkOrgNameResult, 1)
	go client.linkOrganizationName(ctx, tx, result)

	return result, err
}

func (client *Client) linkOrganizationName(ctx context.Context, tx *types.Transaction, result chan<- LinkOrgNameResult) {
	defer close(result)

	receipt, err := bind.WaitMined(ctx, client.eth, tx)
	if err != nil {
		result <- LinkOrgNameResult{Err: err}
		return
	}

	if len(receipt.Logs) == 0 {
		result <- LinkOrgNameResult{Err: fmt.Errorf("Failed to parse log")}
		return
	}

	log, err := client.registry.ParseMappingEvent(*receipt.Logs[0])
	if err != nil {
		result <- LinkOrgNameResult{Err: err}
		return
	}

	result <- LinkOrgNameResult{Log: log}
}
