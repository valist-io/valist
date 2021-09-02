package test

import (
	"context"

	"github.com/valist-io/registry/internal/core/types"
)

func (s *CoreSuite) TestResolvePath() {
	ctx := context.Background()

	orgName := "valist"
	orgMeta := &types.OrganizationMeta{
		Name:        "Valist, Inc.",
		Description: "Accelerating the transition to web3.",
	}

	repoName := "sdk"
	repoMeta := &types.RepositoryMeta{
		Name:        "sdk",
		Description: "Valist core sdk.",
		ProjectType: "npm",
		Homepage:    "https://valist.io",
		Repository:  "https://github.com/valist-io/valist",
	}

	metaCID, err := s.client.Storage().Write(ctx, []byte("hello"))
	s.Require().NoError(err, "Failed to add meta file")

	releaseCID, err := s.client.Storage().Write(ctx, []byte("world"))
	s.Require().NoError(err, "Failed to add release file")

	release := &types.Release{
		Tag:        "v0.0.1",
		ReleaseCID: releaseCID,
		MetaCID:    metaCID,
	}

	orgCreatedEvent, err := s.client.CreateOrganization(ctx, orgMeta)
	s.Require().NoError(err, "Failed to create organization")
	orgID := orgCreatedEvent.OrgID

	_, err = s.client.LinkOrganizationName(ctx, orgID, orgName)
	s.Require().NoError(err, "Failed to link organization name")

	_, err = s.client.CreateRepository(ctx, orgID, repoName, repoMeta)
	s.Require().NoError(err, "Failed to create repository")

	_, err = s.client.VoteRelease(ctx, orgID, repoName, release)
	s.Require().NoError(err, "Failed to vote release")

	res, err := s.client.ResolvePath(ctx, "valist/sdk/v0.0.1")
	s.Require().NoError(err, "Failed to resolve path")
	s.Require().NotNil(res.Organization)
	s.Require().NotNil(res.Repository)
	s.Require().NotNil(res.Release)
}
