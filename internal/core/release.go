package core

import (
	"context"
	"fmt"
	"io"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ipfs/go-cid"
	files "github.com/ipfs/go-ipfs-files"
	"github.com/ipfs/interface-go-ipfs-core/path"
)

type Release struct {
	Tag        string
	ReleaseCID cid.Cid
	MetaCID    cid.Cid
	Signers    []common.Address
}

// GetRelease returns the release with the given tag in the repository with the given name and orgID.
func (client *Client) GetRelease(ctx context.Context, orgName, repoName, tag string) (*Release, error) {
	if tag == "latest" {
		return client.GetLatestRelease(ctx, orgName, repoName)
	}

	orgID, err := client.GetOrganizationID(ctx, orgName)
	if err != nil {
		return nil, err
	}

	packed, err := packedEncoding(orgID, repoName, tag)
	if err != nil {
		return nil, fmt.Errorf("Failed to get release: %v", err)
	}

	callopts := bind.CallOpts{Context: ctx}
	selector := crypto.Keccak256Hash(packed)

	release, err := client.valist.Releases(&callopts, selector)
	if err != nil {
		return nil, ErrReleaseNotExist
	}

	releaseCID, err := cid.Decode(release.ReleaseCID)
	if err != nil {
		return nil, fmt.Errorf("Failed to get release: %v", err)
	}

	metaCID, err := cid.Decode(release.MetaCID)
	if err != nil {
		return nil, fmt.Errorf("Failed to get release: %v", err)
	}

	return &Release{
		Tag:        tag,
		ReleaseCID: releaseCID,
		MetaCID:    metaCID,
	}, nil
}

// GetLatestRelease returns the latest release from the repository with the given name and orgID.
func (client *Client) GetLatestRelease(ctx context.Context, orgName, repoName string) (*Release, error) {
	callopts := bind.CallOpts{Context: ctx}

	orgID, err := client.GetOrganizationID(ctx, orgName)
	if err != nil {
		return nil, err
	}

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

	return &Release{
		Tag:        tag,
		ReleaseCID: releaseCID,
		MetaCID:    metaCID,
		Signers:    signers,
	}, nil
}

// GetReleaseMeta returns the release meta with the given CID.
func (client *Client) GetReleaseMeta(ctx context.Context, id cid.Cid) ([]byte, error) {
	node, err := client.ipfs.Unixfs().Get(ctx, path.IpfsPath(id))
	if err != nil {
		return nil, err
	}

	file, ok := node.(files.File)
	if !ok {
		return nil, fmt.Errorf("Failed to parse release meta")
	}

	return io.ReadAll(file)
}

// ReleaseTagIterator is used to iterate release tags.
type ReleaseTagIterator struct {
	client   *Client
	orgName  string
	repoName string
	tags     []string
	page     *big.Int
	limit    *big.Int
}

// ListReleaseTags returns a new ReleaseTagIterator.
func (client *Client) ListReleaseTags(orgName, repoName string, page, limit *big.Int) *ReleaseTagIterator {
	return &ReleaseTagIterator{
		client:   client,
		orgName:  orgName,
		repoName: repoName,
		page:     page,
		limit:    limit,
	}
}

func (it *ReleaseTagIterator) load(ctx context.Context) error {
	if len(it.tags) != 0 {
		return nil
	}

	orgID, err := it.client.GetOrganizationID(ctx, it.orgName)
	if err != nil {
		return err
	}

	callopts := bind.CallOpts{Context: ctx}
	selector := crypto.Keccak256Hash(orgID[:], []byte(it.repoName))

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
	if err := it.load(ctx); err != nil {
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
	orgName  string
	repoName string
	tags     *ReleaseTagIterator
}

// ListReleases returns a new ReleaseIterator.
func (client *Client) ListReleases(orgName, repoName string, page, limit *big.Int) *ReleaseIterator {
	return &ReleaseIterator{
		client:   client,
		orgName:  orgName,
		repoName: repoName,
		tags:     client.ListReleaseTags(orgName, repoName, page, limit),
	}
}

// Next returns the next release from the iterator.
// Returns EOF when no releases are left.
func (it *ReleaseIterator) Next(ctx context.Context) (*Release, error) {
	tag, err := it.tags.Next(ctx)
	if err != nil {
		return nil, err
	}

	release, err := it.client.GetRelease(ctx, it.orgName, it.repoName, tag)
	if err != nil {
		return nil, err
	}

	return release, nil
}

// ForEach calls the given callback for each release.
func (it *ReleaseIterator) ForEach(ctx context.Context, cb func(*Release)) error {
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
