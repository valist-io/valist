package client

import (
	"context"
	"fmt"
	"strings"

	"github.com/ipfs/interface-go-ipfs-core/path"

	"github.com/valist-io/registry/internal/core/types"
)

func (client *Client) ResolvePath(ctx context.Context, raw string) (*types.ResolvedPath, error) {
	clean := strings.TrimLeft(raw, "/@")
	parts := strings.SplitN(clean, "/", 4)
	if len(parts) == 0 {
		return nil, fmt.Errorf("invalid path")
	}

	orgID, err := client.GetOrganizationID(ctx, parts[0])
	if err != nil {
		return nil, err
	}

	var res types.ResolvedPath
	res.Organization, err = client.GetOrganization(ctx, orgID)
	if err != nil {
		return nil, err
	}

	if len(parts) < 2 {
		return &res, nil
	}

	res.Repository, err = client.GetRepository(ctx, orgID, parts[1])
	if err != nil {
		return nil, err
	}

	if len(parts) < 3 {
		return &res, nil
	}

	res.Release, err = client.GetRelease(ctx, orgID, parts[1], parts[2])
	if err != nil {
		return nil, err
	}

	if len(parts) < 4 {
		return &res, nil
	}

	var p path.Path
	p = path.IpfsPath(res.Release.ReleaseCID)
	p = path.Join(p, parts[3])

	res.Node, err = client.ipfs.ResolveNode(ctx, p)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
