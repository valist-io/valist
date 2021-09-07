package client

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/valist-io/registry/internal/contract/valist"
	"github.com/valist-io/registry/internal/core/types"
)

// GetRelease returns the release with the given tag in the repository with the given name and orgID.
func (client *Client) GetRelease(ctx context.Context, orgID common.Hash, repoName, tag string) (*types.Release, error) {
	if tag == "latest" {
		return client.GetLatestRelease(ctx, orgID, repoName)
	}

	packed, err := packedEncoding(orgID, repoName, tag)
	if err != nil {
		return nil, fmt.Errorf("Failed to get release: %v", err)
	}

	selector := crypto.Keccak256Hash(packed)
	callopts := bind.CallOpts{
		Context: ctx,
		From:    client.account.Address,
	}

	release, err := client.valist.Releases(&callopts, selector)
	if err != nil {
		return nil, err
	}

	if release.ReleaseCID == "" {
		return nil, types.ErrReleaseNotExist
	}

	return &types.Release{
		Tag:        tag,
		ReleaseCID: release.ReleaseCID,
		MetaCID:    release.MetaCID,
	}, nil
}

// GetLatestRelease returns the latest release from the repository with the given name and orgID.
func (client *Client) GetLatestRelease(ctx context.Context, orgID common.Hash, repoName string) (*types.Release, error) {
	callopts := bind.CallOpts{
		Context: ctx,
		From:    client.account.Address,
	}

	tag, releaseCID, metaCID, signers, err := client.valist.GetLatestRelease(&callopts, orgID, repoName)
	if err != nil {
		return nil, err
	}

	return &types.Release{
		Tag:        tag,
		ReleaseCID: releaseCID,
		MetaCID:    metaCID,
		Signers:    signers,
	}, nil
}

// VoteRelease votes on a release in the given organization's repository with the given release and meta CIDs.
func (client *Client) VoteRelease(ctx context.Context, orgID common.Hash, repoName string, release *types.Release) (*valist.ValistVoteReleaseEvent, error) {
	txopts := client.transactOpts(client.account, client.wallet, client.chainID)
	txopts.Context = ctx

	tx, err := client.transactor.VoteReleaseTx(txopts, orgID, repoName, release.Tag, release.ReleaseCID, release.MetaCID)
	if err != nil {
		return nil, err
	}

	logs, err := waitMined(ctx, client.eth, tx)
	if err != nil {
		return nil, err
	}

	return client.valist.ParseVoteReleaseEvent(*logs[0])
}
