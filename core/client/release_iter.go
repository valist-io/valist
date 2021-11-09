package client

import (
	"context"
	"io"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/valist-io/valist/core/types"
)

// ReleaseTagIterator is used to iterate release tags.
type ReleaseTagIterator struct {
	client   *Client
	orgID    common.Hash
	repoName string
	tags     []string
	page     *big.Int
	limit    *big.Int
}

// ListReleaseTags returns an iterator for retrieving all release tags from the repo with the given repoName and orgID.
func (client *Client) ListReleaseTags(orgID common.Hash, repoName string) types.ReleaseTagIterator {
	return &ReleaseTagIterator{
		client:   client,
		orgID:    orgID,
		repoName: repoName,
		page:     big.NewInt(1),
		limit:    big.NewInt(10),
	}
}

func (it *ReleaseTagIterator) paginate(ctx context.Context) error {
	if len(it.tags) != 0 {
		return nil
	}

	selector := crypto.Keccak256Hash(it.orgID[:], []byte(it.repoName))
	callopts := bind.CallOpts{
		Context: ctx,
		From:    it.client.signer.Account().Address,
	}

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

// ForEach calls the given callback for each release tag.
func (it *ReleaseTagIterator) ForEach(ctx context.Context, cb func(string)) error {
	for {
		tag, err := it.Next(ctx)
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		cb(tag)
	}
}

// ReleaseIterator is used to iterate releases.
type ReleaseIterator struct {
	client   *Client
	orgID    common.Hash
	repoName string
	tags     types.ReleaseTagIterator
}

// ListReleases returns an iterator for retrieving all releases from the repo with the given repoName and orgID.
func (client *Client) ListReleases(orgID common.Hash, repoName string) types.ReleaseIterator {
	return &ReleaseIterator{
		client:   client,
		orgID:    orgID,
		repoName: repoName,
		tags:     client.ListReleaseTags(orgID, repoName),
	}
}

// Next returns the next release from the iterator.
// Returns EOF when no releases are left.
func (it *ReleaseIterator) Next(ctx context.Context) (*types.Release, error) {
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
func (it *ReleaseIterator) ForEach(ctx context.Context, cb func(*types.Release)) error {
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
