package client

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/valist-io/valist/contract/valist"
	"github.com/valist-io/valist/core/types"
)

// GetOrganization returns the organization with the given orgID.
func (client *Client) GetOrganization(ctx context.Context, id common.Hash) (*types.Organization, error) {
	callopts := bind.CallOpts{
		Context: ctx,
		From:    client.signer.Account().Address,
	}

	org, err := client.valist.Orgs(&callopts, id)
	if err != nil {
		return nil, fmt.Errorf("Failed to get organization id: %v", err)
	}

	// TODO there's no way to check if an org exists

	return &types.Organization{
		ID:            id,
		Threshold:     org.Threshold,
		ThresholdDate: org.ThresholdDate,
		MetaCID:       org.MetaCID,
	}, nil
}

// GetOrganizationMeta returns the organization metadata from the given path.
func (client *Client) GetOrganizationMeta(ctx context.Context, p string) (*types.OrganizationMeta, error) {
	data, err := client.ReadFile(ctx, p)
	if err != nil {
		return nil, fmt.Errorf("Failed to get organization meta: %v", err)
	}

	var meta types.OrganizationMeta
	if err := json.Unmarshal(data, &meta); err != nil {
		return nil, fmt.Errorf("Failed to parse organization meta: %v", err)
	}

	return &meta, nil
}

// CreateOrganization creates a new organization using the given metadata.
func (client *Client) CreateOrganization(ctx context.Context, meta *types.OrganizationMeta) (*valist.ValistOrgCreated, error) {
	data, err := json.Marshal(meta)
	if err != nil {
		return nil, err
	}

	metaCID, err := client.WriteFile(ctx, data)
	if err != nil {
		return nil, err
	}

	txopts := client.signer.NewTransactor()
	txopts.Context = ctx

	tx, err := client.transactor.CreateOrganizationTx(txopts, metaCID)
	if err != nil {
		return nil, err
	}

	logs, err := waitMined(ctx, client.eth, tx)
	if err != nil {
		return nil, err
	}

	return client.valist.ParseOrgCreated(*logs[0])
}

// SetOrganizationMeta updates the metadata of the organization with the given orgID.
func (client *Client) SetOrganizationMeta(ctx context.Context, orgID common.Hash, meta *types.OrganizationMeta) (*valist.ValistMetaUpdate, error) {
	data, err := json.Marshal(meta)
	if err != nil {
		return nil, err
	}

	metaCID, err := client.WriteFile(ctx, data)
	if err != nil {
		return nil, err
	}

	txopts := client.signer.NewTransactor()
	txopts.Context = ctx

	tx, err := client.transactor.SetOrganizationMetaTx(txopts, orgID, metaCID)
	if err != nil {
		return nil, err
	}

	logs, err := waitMined(ctx, client.eth, tx)
	if err != nil {
		return nil, err
	}

	return client.valist.ParseMetaUpdate(*logs[0])
}

// VoteOrganizationAdmin votes to add, revoke, or rotate an organization admin key.
func (client *Client) VoteOrganizationAdmin(ctx context.Context, orgID common.Hash, operation common.Hash, address common.Address) (*valist.ValistVoteKeyEvent, error) {
	txopts := client.signer.NewTransactor()
	txopts.Context = ctx

	tx, err := client.transactor.VoteKeyTx(txopts, orgID, "", operation, address)
	if err != nil {
		return nil, err
	}

	logs, err := waitMined(ctx, client.eth, tx)
	if err != nil {
		return nil, err
	}

	return client.valist.ParseVoteKeyEvent(*logs[0])
}

// VoteOrganizationThreshold votes to set a multifactor voting threshold for the organization.
func (client *Client) VoteOrganizationThreshold(ctx context.Context, orgID common.Hash, threshold *big.Int) (*valist.ValistVoteThresholdEvent, error) {
	txopts := client.signer.NewTransactor()
	txopts.Context = ctx

	tx, err := client.transactor.VoteOrganizationThresholdTx(txopts, orgID, threshold)
	if err != nil {
		return nil, err
	}

	logs, err := waitMined(ctx, client.eth, tx)
	if err != nil {
		return nil, err
	}

	return client.valist.ParseVoteThresholdEvent(*logs[0])
}
