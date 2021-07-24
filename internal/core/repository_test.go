package core

import (
	"context"
	"math/big"
)

func (s *CoreSuite) TestCreateRepository() {
	ctx := context.Background()

	orgMeta := &OrganizationMeta{
		Name:        "Valist, Inc.",
		Description: "Accelerating the transition to web3.",
	}

	repoName := "sdk"
	repoMeta := &RepositoryMeta{
		Name:        "sdk",
		Description: "Valist core sdk.",
		ProjectType: "npm",
		Homepage:    "https://valist.io",
		Repository:  "https://github.com/valist-io/valist",
	}

	txc1, err := s.client.CreateOrganization(ctx, orgMeta)
	s.Require().NoError(err, "Failed to create organization")
	s.backend.Commit()

	res1 := <-txc1
	s.Require().NoError(res1.Err, "Failed to create organization")
	orgID := res1.Log.OrgID

	txc2, err := s.client.CreateRepository(ctx, orgID, repoName, repoMeta)
	s.Require().NoError(err, "Failed to create repository")
	s.backend.Commit()

	res2 := <-txc2
	s.Require().NoError(res2.Err, "Failed to create repository")

	repo, err := s.client.GetRepository(ctx, orgID, repoName)
	s.Require().NoError(err, "Failed to get repository")
	s.Assert().Equal(big.NewInt(0).Cmp(repo.Threshold), 0)

	meta, err := s.client.GetRepositoryMeta(ctx, repo.MetaCID)
	s.Require().NoError(err, "Failed to get repository meta")
	s.Assert().Equal(repoMeta.Name, meta.Name)
	s.Assert().Equal(repoMeta.Description, meta.Description)
	s.Assert().Equal(repoMeta.ProjectType, meta.ProjectType)
	s.Assert().Equal(repoMeta.Homepage, meta.Homepage)
}
