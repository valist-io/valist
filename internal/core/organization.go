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
	"github.com/ipfs/go-cid"
	files "github.com/ipfs/go-ipfs-files"
	"github.com/ipfs/interface-go-ipfs-core/path"
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

// GetOrganizationByName returns the organization with the given ID.
func (client *Client) GetOrganization(ctx context.Context, name string) (*Organization, error) {
	id, err := client.GetOrganizationID(ctx, name)
	if err != nil {
		return nil, err
	}

	callopts := bind.CallOpts{Context: ctx}
	org, err := client.valist.Orgs(&callopts, id)
	if err != nil {
		return nil, fmt.Errorf("Failed to get organization id: %v", err)
	}

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
func (client *Client) CreateOrganization(ctx context.Context, meta *OrganizationMeta) (common.Hash, error) {
	data, err := json.Marshal(meta)
	if err != nil {
		return emptyHash, err
	}

	path, err := client.ipfs.Unixfs().Add(ctx, files.NewBytesFile(data))
	if err != nil {
		return emptyHash, err
	}

	txopts, err := bind.NewKeyedTransactorWithChainID(client.private, client.chainID)
	if err != nil {
		return emptyHash, err
	}
	txopts.Context = ctx

	tx, err := client.valist.CreateOrganization(txopts, path.Cid().String())
	if err != nil {
		return emptyHash, err
	}

	receipt, err := bind.WaitMined(ctx, client.eth, tx)
	if err != nil {
		return emptyHash, err
	}

	if len(receipt.Logs) < 1 || len(receipt.Logs[0].Topics) < 2 {
		return emptyHash, fmt.Errorf("Failed to get create organization log")
	}

	return receipt.Logs[0].Topics[1], nil
}

// LinkOrganizationName creates a link from the given orgID to the given name.
func (client *Client) LinkOrganizationName(ctx context.Context, orgID common.Hash, name string) error {
	// TODO validate name

	txopts, err := bind.NewKeyedTransactorWithChainID(client.private, client.chainID)
	if err != nil {
		return err
	}
	txopts.Context = ctx

	tx, err := client.registry.LinkNameToID(txopts, orgID, name)
	if err != nil {
		return err
	}

	_, err = bind.WaitMined(ctx, client.eth, tx)
	return err
}
