package client

import (
	"context"
	"fmt"
	"strings"

	"github.com/valist-io/valist/core/types"
)

// ResolvePath resolves the organization, repository, release, and node from the given path.
func (client *Client) ResolvePath(ctx context.Context, raw string) (types.ResolvedPath, error) {
	var res types.ResolvedPath
	var err error

	clean := strings.TrimLeft(raw, "/@")
	parts := strings.Split(clean, "/")

	if len(parts) == 0 || len(parts) > 3 {
		return res, fmt.Errorf("invalid path")
	}

	res.OrgName = parts[0]
	res.OrgID, err = client.GetOrganizationID(ctx, res.OrgName)
	if err != nil {
		return res, err
	}

	res.Organization, err = client.GetOrganization(ctx, res.OrgID)
	if err != nil {
		return res, err
	}

	if len(parts) < 2 {
		return res, nil
	}

	res.RepoName = parts[1]
	res.Repository, err = client.GetRepository(ctx, res.OrgID, res.RepoName)
	if err != nil {
		return res, err
	}

	if len(parts) < 3 {
		return res, nil
	}

	res.ReleaseTag = parts[2]
	res.Release, err = client.GetRelease(ctx, res.OrgID, res.RepoName, res.ReleaseTag)
	if err != nil {
		return res, err
	}

	return res, nil
}
