package core

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ipfs/go-cid"
)

type OrganizationAPI interface {
	GetOrganization(context.Context, common.Hash) (*Organization, error)
	GetOrganizationMeta(context.Context, cid.Cid) (*OrganizationMeta, error)
	CreateOrganization(context.Context, *OrganizationMeta) (<-chan CreateOrgResult, error)
}

type Organization struct {
	ID            common.Hash
	Threshold     *big.Int
	ThresholdDate *big.Int
	MetaCID       cid.Cid
}

type OrganizationMeta struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateOrgResult struct {
	OrgID   common.Hash
	MetaCID cid.Cid
	Err     error
}
