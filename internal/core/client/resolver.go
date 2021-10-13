package client

import (
	"context"
	"fmt"
	"strings"

	"github.com/valist-io/valist/internal/core/types"
)

func (client *Client) ResolvePath(ctx context.Context, raw string) (*types.ResolvedPath, error) {
	clean := strings.TrimLeft(raw, "/@")
	parts := strings.Split(clean, "/")

	if len(parts) == 0 || len(parts) > 3 {
		return nil, fmt.Errorf("invalid path")
	}

	orgID, err := client.GetOrganizationID(ctx, parts[0])
	if err != nil {
		return nil, err
	}

	res := types.ResolvedPath{
		OrgID:   orgID,
		OrgName: parts[0],
	}

	res.Organization, err = client.GetOrganization(ctx, res.OrgID)
	if err != nil {
		return &res, err
	}

	if len(parts) < 2 {
		return &res, nil
	}

	res.RepoName = parts[1]
	res.Repository, err = client.GetRepository(ctx, orgID, res.RepoName)
	if err != nil {
		return &res, err
	}

	if len(parts) < 3 {
		return &res, nil
	}

	res.ReleaseTag = parts[2]
	res.Release, err = client.GetRelease(ctx, orgID, res.RepoName, res.ReleaseTag)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
