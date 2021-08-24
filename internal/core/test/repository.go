package test

import (
	"context"
	"math/big"

	"github.com/valist-io/registry/internal/core/types"
)

func (s *CoreSuite) TestGetRepository() {
	ctx := context.Background()

	_, err := s.client.GetRepository(ctx, emptyHash, "empty")
	s.Assert().Equal(types.ErrRepositoryNotExist, err)
}

func (s *CoreSuite) TestCreateRepository() {
	ctx := context.Background()

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

	orgCreatedEvent, err := s.client.CreateOrganization(ctx, orgMeta)
	s.Require().NoError(err, "Failed to create organization")

	_, err = s.client.CreateRepository(ctx, orgCreatedEvent.OrgID, repoName, repoMeta)
	s.Require().NoError(err, "Failed to create repository")

	repo, err := s.client.GetRepository(ctx, orgCreatedEvent.OrgID, repoName)
	s.Require().NoError(err, "Failed to get repository")
	s.Assert().Equal(big.NewInt(0).Cmp(repo.Threshold), 0)

	meta, err := s.client.GetRepositoryMeta(ctx, repo.MetaCID)
	s.Require().NoError(err, "Failed to get repository meta")
	s.Assert().Equal(repoMeta.Name, meta.Name)
	s.Assert().Equal(repoMeta.Description, meta.Description)
	s.Assert().Equal(repoMeta.ProjectType, meta.ProjectType)
	s.Assert().Equal(repoMeta.Homepage, meta.Homepage)
}
