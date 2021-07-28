package impl

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ipfs/go-cid"

	"github.com/valist-io/registry/internal/core"
)

// GetOrganization returns the organization with the given ID.
func (client *Client) GetOrganization(ctx context.Context, id common.Hash) (*core.Organization, error) {
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

	return &core.Organization{
		ID:            id,
		Threshold:     org.Threshold,
		ThresholdDate: org.ThresholdDate,
		MetaCID:       metaCID,
	}, nil
}

// GetOrganizationMeta returns the organization meta with the given CID.
func (client *Client) GetOrganizationMeta(ctx context.Context, id cid.Cid) (*core.OrganizationMeta, error) {
	data, err := client.GetFile(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("Failed to get organization meta: %v", err)
	}

	var meta core.OrganizationMeta
	if err := json.Unmarshal(data, &meta); err != nil {
		return nil, fmt.Errorf("Failed to parse organization meta: %v", err)
	}

	return &meta, nil
}

// CreateOrganization creates a new organization with the given meta and returns the orgID.
func (client *Client) CreateOrganization(ctx context.Context, meta *core.OrganizationMeta) (<-chan core.CreateOrgResult, error) {
	if client.transact == nil {
		return nil, ErrNoTransactor
	}

	data, err := json.Marshal(meta)
	if err != nil {
		return nil, err
	}

	metaCID, err := client.AddFile(ctx, data)
	if err != nil {
		return nil, err
	}

	txopts, err := client.transact()
	if err != nil {
		return nil, err
	}
	txopts.Context = ctx

	tx, err := client.valist.CreateOrganization(txopts, metaCID.String())
	if err != nil {
		return nil, err
	}

	result := make(chan core.CreateOrgResult, 1)
	go client.createOrganization(ctx, tx, result)

	return result, nil
}

func (client *Client) createOrganization(ctx context.Context, tx *types.Transaction, result chan<- core.CreateOrgResult) {
	defer close(result)

	receipt, err := bind.WaitMined(ctx, client.eth, tx)
	if err != nil {
		result <- core.CreateOrgResult{Err: err}
		return
	}

	if len(receipt.Logs) == 0 {
		result <- core.CreateOrgResult{Err: fmt.Errorf("Failed to parse log")}
		return
	}

	log, err := client.valist.ParseOrgCreated(*receipt.Logs[0])
	if err != nil {
		result <- core.CreateOrgResult{Err: err}
		return
	}

	metaCID, err := cid.Decode(log.MetaCID)
	if err != nil {
		result <- core.CreateOrgResult{Err: err}
		return
	}

	result <- core.CreateOrgResult{
		OrgID:   log.OrgID,
		MetaCID: metaCID,
	}
}
