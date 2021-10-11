package client

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/valist-io/valist/internal/contract/valist"
	"github.com/valist-io/valist/internal/core/types"
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
		From:    client.signer.Account().Address,
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

// GetReleaseMeta returns the release meta with the given CID.
func (client *Client) GetReleaseMeta(ctx context.Context, p string) (*types.ReleaseMeta, error) {
	releaseData, err := client.Storage().ReadFile(ctx, p)
	if err != nil {
		return nil, err
	}

	var meta types.ReleaseMeta
	if err := json.Unmarshal(releaseData, &meta); err != nil {
		return nil, err
	}

	return &meta, nil
}

// GetLatestRelease returns the latest release from the repository with the given name and orgID.
func (client *Client) GetLatestRelease(ctx context.Context, orgID common.Hash, repoName string) (*types.Release, error) {
	callopts := bind.CallOpts{
		Context: ctx,
		From:    client.signer.Account().Address,
	}

	tag, releaseCID, metaCID, signers, err := client.valist.GetLatestRelease(&callopts, orgID, repoName)
	if err != nil && err.Error() == "execution reverted" {
		return nil, types.ErrReleaseNotExist
	}

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
	txopts := client.signer.NewTransactor()
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
