package client

import (
	"bytes"
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/valist-io/valist/contract/registry"
	"github.com/valist-io/valist/core/types"
)

// GetOrganizationID returns the orgID for the given orgName.
func (client *Client) GetOrganizationID(ctx context.Context, name string) (common.Hash, error) {
	if orgID, ok := client.orgs[name]; ok {
		return orgID, nil
	}

	callopts := bind.CallOpts{
		Context: ctx,
		From:    client.signer.Account().Address,
	}

	orgID, err := client.registry.NameToID(&callopts, name)
	if err != nil {
		return emptyHash, fmt.Errorf("Failed to get organization id: %v", err)
	}

	if bytes.Equal(orgID[:], emptyHash.Bytes()) {
		return emptyHash, types.ErrOrgNotExist
	}

	client.orgs[name] = orgID
	return orgID, nil
}

// LinkOrganizationName links the given orgID to the orgName.
func (client *Client) LinkOrganizationName(ctx context.Context, orgID common.Hash, name string) (*registry.ValistRegistryMappingEvent, error) {
	txopts := client.signer.NewTransactor()
	txopts.Context = ctx

	tx, err := client.transactor.LinkOrganizationNameTx(txopts, orgID, name)
	if err != nil {
		return nil, err
	}

	logs, err := waitMined(ctx, client.eth, tx)
	if err != nil {
		return nil, err
	}

	return client.registry.ParseMappingEvent(*logs[0])
}
