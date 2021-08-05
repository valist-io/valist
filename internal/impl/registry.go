package impl

import (
	"bytes"
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/valist-io/registry/internal/core"
)

// GetOrganizationID returns the ID of the organization with the given name.
func (client *Client) GetOrganizationID(ctx context.Context, name string) (common.Hash, error) {
	if orgID, ok := client.orgs[name]; ok {
		return orgID, nil
	}

	callopts := bind.CallOpts{
		Context: ctx,
		From:    client.account.Address,
	}

	orgID, err := client.registry.NameToID(&callopts, name)
	if err != nil {
		return emptyHash, fmt.Errorf("Failed to get organization id: %v", err)
	}

	if bytes.Equal(orgID[:], emptyHash.Bytes()) {
		return emptyHash, core.ErrOrganizationNotExist
	}

	client.orgs[name] = orgID
	return orgID, nil
}

// LinkOrganizationName creates a link from the given orgID to the given name.
func (client *Client) LinkOrganizationName(ctx context.Context, orgID common.Hash, name string) (<-chan core.LinkOrgNameResult, error) {
	txopts := bind.TransactOpts{
		Context: ctx,
		From:    client.account.Address,
		Signer:  client.Signer,
	}

	tx, err := client.registry.LinkNameToID(&txopts, orgID, name)
	if err != nil {
		return nil, err
	}

	result := make(chan core.LinkOrgNameResult, 1)
	go client.linkOrganizationName(ctx, tx, result)

	return result, err
}

func (client *Client) linkOrganizationName(ctx context.Context, tx *types.Transaction, result chan<- core.LinkOrgNameResult) {
	defer close(result)

	receipt, err := bind.WaitMined(ctx, client.eth, tx)
	if err != nil {
		result <- core.LinkOrgNameResult{Err: err}
		return
	}

	if len(receipt.Logs) == 0 {
		result <- core.LinkOrgNameResult{Err: fmt.Errorf("Failed to parse log")}
		return
	}

	log, err := client.registry.ParseMappingEvent(*receipt.Logs[0])
	if err != nil {
		result <- core.LinkOrgNameResult{Err: err}
		return
	}

	result <- core.LinkOrgNameResult{
		OrgID: log.OrgID,
		Name:  log.Name,
	}
}
