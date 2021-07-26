package impl

import (
	"context"
	"fmt"
	"io"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ipfs/go-cid"

	"github.com/valist-io/registry/internal/core"
)

// GetRelease returns the release with the given tag in the repository with the given name and orgID.
func (client *Client) GetRelease(ctx context.Context, orgID common.Hash, repoName, tag string) (*core.Release, error) {
	if tag == "latest" {
		return client.GetLatestRelease(ctx, orgID, repoName)
	}

	packed, err := packedEncoding(orgID, repoName, tag)
	if err != nil {
		return nil, fmt.Errorf("Failed to get release: %v", err)
	}

	callopts := bind.CallOpts{Context: ctx}
	selector := crypto.Keccak256Hash(packed)

	release, err := client.valist.Releases(&callopts, selector)
	if err != nil {
		return nil, err
	}

	if release.ReleaseCID == "" {
		return nil, core.ErrReleaseNotExist
	}

	releaseCID, err := cid.Decode(release.ReleaseCID)
	if err != nil {
		return nil, fmt.Errorf("Failed to get release: %v", err)
	}

	metaCID, err := cid.Decode(release.MetaCID)
	if err != nil {
		return nil, fmt.Errorf("Failed to get release: %v", err)
	}

	return &core.Release{
		Tag:        tag,
		ReleaseCID: releaseCID,
		MetaCID:    metaCID,
	}, nil
}

// GetLatestRelease returns the latest release from the repository with the given name and orgID.
func (client *Client) GetLatestRelease(ctx context.Context, orgID common.Hash, repoName string) (*core.Release, error) {
	callopts := bind.CallOpts{Context: ctx}

	tag, release, meta, signers, err := client.valist.GetLatestRelease(&callopts, orgID, repoName)
	if err != nil {
		return nil, err
	}

	releaseCID, err := cid.Decode(release)
	if err != nil {
		return nil, err
	}

	metaCID, err := cid.Decode(meta)
	if err != nil {
		return nil, err
	}

	return &core.Release{
		Tag:        tag,
		ReleaseCID: releaseCID,
		MetaCID:    metaCID,
		Signers:    signers,
	}, nil
}

// VoteRelease votes on a release in the given organization's repository with the given release and meta CIDs.
func (client *Client) VoteRelease(ctx context.Context, orgID common.Hash, repoName string, release *core.Release) (<-chan core.VoteReleaseResult, error) {
	txopts, err := bind.NewKeyedTransactorWithChainID(client.private, client.chainID)
	if err != nil {
		return nil, err
	}
	txopts.Context = ctx

	releaseCID := release.ReleaseCID.String()
	metaCID := release.MetaCID.String()

	tx, err := client.valist.VoteRelease(txopts, orgID, repoName, release.Tag, releaseCID, metaCID)
	if err != nil {
		return nil, err
	}

	result := make(chan core.VoteReleaseResult, 1)
	go client.voteRelease(ctx, tx, result)

	return result, nil
}

func (client *Client) voteRelease(ctx context.Context, tx *types.Transaction, result chan<- core.VoteReleaseResult) {
	defer close(result)

	receipt, err := bind.WaitMined(ctx, client.eth, tx)
	if err != nil {
		result <- core.VoteReleaseResult{Err: err}
		return
	}

	if len(receipt.Logs) == 0 {
		result <- core.VoteReleaseResult{Err: fmt.Errorf("Failed to parse log")}
		return
	}

	log, err := client.valist.ParseVoteReleaseEvent(*receipt.Logs[0])
	if err != nil {
		result <- core.VoteReleaseResult{Err: err}
		return
	}

	releaseCID, err := cid.Decode(log.ReleaseCID)
	if err != nil {
		result <- core.VoteReleaseResult{Err: err}
		return
	}

	metaCID, err := cid.Decode(log.MetaCID)
	if err != nil {
		result <- core.VoteReleaseResult{Err: err}
		return
	}

	result <- core.VoteReleaseResult{
		OrgID:      log.OrgID,
		RepoName:   log.RepoName,
		Tag:        log.Tag,
		ReleaseCID: releaseCID,
		MetaCID:    metaCID,
		Signer:     log.Signer,
		SigCount:   log.SigCount,
		Threshold:  log.Threshold,
	}
}

// ReleaseTagIterator is used to iterate release tags.
type ReleaseTagIterator struct {
	client   *Client
	orgID    common.Hash
	repoName string
	tags     []string
	page     *big.Int
	limit    *big.Int
}

// ListReleaseTags returns a new ReleaseTagIterator.
func (client *Client) ListReleaseTags(orgID common.Hash, repoName string, page, limit *big.Int) core.ReleaseTagIterator {
	return &ReleaseTagIterator{
		client:   client,
		orgID:    orgID,
		repoName: repoName,
		page:     page,
		limit:    limit,
	}
}

func (it *ReleaseTagIterator) paginate(ctx context.Context) error {
	if len(it.tags) != 0 {
		return nil
	}

	callopts := bind.CallOpts{Context: ctx}
	selector := crypto.Keccak256Hash(it.orgID[:], []byte(it.repoName))

	tags, err := it.client.valist.GetReleaseTags(&callopts, selector, it.page, it.limit)
	if err != nil {
		return err
	}

	it.tags = tags
	it.page.Add(it.page, big.NewInt(1))
	return nil
}

// Next returns the next tag in the iterator.
func (it *ReleaseTagIterator) Next(ctx context.Context) (string, error) {
	if err := it.paginate(ctx); err != nil {
		return "", err
	}

	if it.tags[0] == "" {
		return "", io.EOF
	}

	tag := it.tags[0]
	it.tags = it.tags[1:]

	return tag, nil
}

// ReleaseIterator is used to iterate releases.
type ReleaseIterator struct {
	client   *Client
	orgID    common.Hash
	repoName string
	tags     core.ReleaseTagIterator
}

// ListReleases returns a new ReleaseIterator.
func (client *Client) ListReleases(orgID common.Hash, repoName string, page, limit *big.Int) core.ReleaseIterator {
	return &ReleaseIterator{
		client:   client,
		orgID:    orgID,
		repoName: repoName,
		tags:     client.ListReleaseTags(orgID, repoName, page, limit),
	}
}

// Next returns the next release from the iterator.
// Returns EOF when no releases are left.
func (it *ReleaseIterator) Next(ctx context.Context) (*core.Release, error) {
	tag, err := it.tags.Next(ctx)
	if err != nil {
		return nil, err
	}

	release, err := it.client.GetRelease(ctx, it.orgID, it.repoName, tag)
	if err != nil {
		return nil, err
	}

	return release, nil
}

// ForEach calls the given callback for each release.
func (it *ReleaseIterator) ForEach(ctx context.Context, cb func(*core.Release)) error {
	for {
		release, err := it.Next(ctx)
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		cb(release)
	}
}
