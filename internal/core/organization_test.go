package core

import (
	"context"
	"math/big"
)

func (s *CoreSuite) TestOrganization() {
	ctx := context.Background()

	orgName := "valist"
	orgMeta := &OrganizationMeta{
		Name:        "Valist, Inc.",
		Description: "Accelerating the transition to web3.",
	}

	orgID, err := s.client.CreateOrganization(ctx, orgMeta)
	s.Require().NoError(err, "Failed to create organization")

	err = s.client.LinkOrganizationName(ctx, orgID, orgName)
	s.Require().NoError(err, "Failed to link organization name")

	id, err := s.client.GetOrganizationID(ctx, orgName)
	s.Require().NoError(err, "Failed to link organization name")
	s.Assert().Equal(orgID, id)

	org, err := s.client.GetOrganization(ctx, orgName)
	s.Require().NoError(err, "Failed to get organization")
	s.Assert().Equal(big.NewInt(0).Cmp(org.Threshold), 0)

	meta, err := s.client.GetOrganizationMeta(ctx, org.MetaCID)
	s.Require().NoError(err, "Failed to get organization meta")
	s.Assert().Equal(orgMeta.Name, meta.Name)
	s.Assert().Equal(orgMeta.Description, meta.Description)
}
