package core

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ipfs/go-cid"

	"github.com/valist-io/registry/internal/contract/valist"
)

type OrganizationAPI interface {
	GetOrganization(context.Context, common.Hash) (*Organization, error)
	GetOrganizationMeta(context.Context, cid.Cid) (*OrganizationMeta, error)
	CreateOrganization(context.Context, *OrganizationMeta) (*valist.ValistOrgCreated, error)
}

type OrganizationTransactorAPI interface {
	CreateOrganizationTx(context.Context, *OrganizationMeta) (*types.Transaction, error)
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
