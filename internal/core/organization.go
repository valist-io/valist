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
	Meta          *OrganizationMeta
	MetaCID       cid.Cid
	RepoNames     []string
}

// GerOrganizationID returns the ID of the organization with the given name.
func (client *Client) GetOrganizationID(ctx context.Context, name string) (common.Hash, error) {
	if orgID, ok := client.orgs[name]; ok {
		return orgID, nil
	}

	callopts := bind.CallOpts{Context: ctx}
	orgID, err := client.registryContract.NameToID(&callopts, name)
	if err != nil {
		return emptyHash, fmt.Errorf("Failed to get organization id: %v", err)
	}

	if bytes.Equal(orgID[:], emptyHash.Bytes()) {
		return emptyHash, ErrOrganizationNotExist
	}

	client.orgs[name] = orgID
	return orgID, nil
}

// GetOrganizationByName returns the organization with the given name.
func (client *Client) GetOrganizationByName(ctx context.Context, name string) (*Organization, error) {
	id, err := client.GetOrganizationID(ctx, name)
	if err != nil {
		return nil, err
	}

	return client.GetOrganizationByID(ctx, id)
}

// GetOrganizationByName returns the organization with the given ID.
func (client *Client) GetOrganizationByID(ctx context.Context, id common.Hash) (*Organization, error) {
	callopts := bind.CallOpts{Context: ctx}

	org, err := client.valistContract.Orgs(&callopts, id)
	if err != nil {
		return nil, fmt.Errorf("Failed to get organization id: %v", err)
	}

	metaCID, err := cid.Decode(org.MetaCID)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse organization meta CID: %v", err)
	}

	meta, err := client.GetOrganizationMeta(ctx, metaCID)
	if err != nil {
		return nil, err
	}

	repoNames, err := client.valistContract.GetRepoNames(&callopts, id, big.NewInt(1), big.NewInt(10))
	if err != nil {
		return nil, err
	}

	return &Organization{
		ID:            id,
		Threshold:     org.Threshold,
		ThresholdDate: org.ThresholdDate,
		MetaCID:       metaCID,
		Meta:          meta,
		RepoNames:     repoNames,
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
