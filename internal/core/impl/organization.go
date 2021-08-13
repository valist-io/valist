package impl

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ipfs/go-cid"

	"github.com/valist-io/registry/internal/contract/valist"
	"github.com/valist-io/registry/internal/core"
)

// GetOrganization returns the organization with the given ID.
func (client *Client) GetOrganization(ctx context.Context, id common.Hash) (*core.Organization, error) {
	callopts := bind.CallOpts{
		Context: ctx,
		From:    client.account.Address,
	}

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

func (client *Client) CreateOrganization(ctx context.Context, txopts *bind.TransactOpts, meta *core.OrganizationMeta) (*valist.ValistOrgCreated, error) {

	data, err := json.Marshal(meta)
	if err != nil {
		return nil, err
	}

	metaCID, err := client.AddFile(ctx, data)
	if err != nil {
		return nil, err
	}

	tx, err := client.transactor.CreateOrganizationTx(ctx, txopts, metaCID)
	if err != nil {
		return nil, err
	}

	logs, err := getTxLogs(ctx, client.eth, tx)
	if err != nil {
		return nil, err
	}

	return client.valist.ParseOrgCreated(*logs[0])
}
