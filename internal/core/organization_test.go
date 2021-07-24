package core

import (
	"context"
	"math/big"
)

func (s *CoreSuite) TestCreateOrganization() {
	ctx := context.Background()

	orgName := "valist"
	orgMeta := &OrganizationMeta{
		Name:        "Valist, Inc.",
		Description: "Accelerating the transition to web3.",
	}

	txc1, err := s.client.CreateOrganization(ctx, orgMeta)
	s.Require().NoError(err, "Failed to create organization")
	s.backend.Commit()

	res1 := <-txc1
	s.Require().NoError(res1.Err, "Failed to create organization")
	orgID := res1.Log.OrgID

	txc2, err := s.client.LinkOrganizationName(ctx, orgID, orgName)
	s.Require().NoError(err, "Failed to link organization name")
	s.backend.Commit()

	res2 := <-txc2
	s.Require().NoError(res2.Err, "Failed to link organization name")

	id, err := s.client.GetOrganizationID(ctx, orgName)
	s.Require().NoError(err, "Failed to get organization id")
	s.Assert().Equal(orgID[:], id.Bytes())

	org, err := s.client.GetOrganization(ctx, orgID)
	s.Require().NoError(err, "Failed to get organization")
	s.Assert().Equal(big.NewInt(0).Cmp(org.Threshold), 0)

	meta, err := s.client.GetOrganizationMeta(ctx, org.MetaCID)
	s.Require().NoError(err, "Failed to get organization meta")
	s.Assert().Equal(orgMeta.Name, meta.Name)
	s.Assert().Equal(orgMeta.Description, meta.Description)
}
