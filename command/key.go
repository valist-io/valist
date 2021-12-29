package command

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/valist-io/valist/contract/valist"
	"github.com/valist-io/valist/core/client"
)

// KeyAdd adds an organization or repository key.
func KeyAdd(ctx context.Context, rpath, addr string) error {
	return keyOperation(ctx, rpath, addr, client.ADD_KEY)
}

// KeyRevoke revokes an organization or repository key.
func KeyRevoke(ctx context.Context, rpath, addr string) error {
	return keyOperation(ctx, rpath, addr, client.REVOKE_KEY)
}

// KeyRotate rotates an organization or repository key.
func KeyRotate(ctx context.Context, rpath, addr string) error {
	return keyOperation(ctx, rpath, addr, client.ROTATE_KEY)
}

func keyOperation(ctx context.Context, rpath string, addr string, op common.Hash) error {
	client := ctx.Value(ClientKey).(*client.Client)

	res, err := client.ResolvePath(ctx, rpath)
	if err != nil {
		return err
	}

	if !common.IsHexAddress(addr) {
		return fmt.Errorf("Invalid address: %s", addr)
	}

	var vote *valist.ValistVoteKeyEvent
	switch {
	case res.Repository != nil:
		vote, err = client.VoteRepoDev(ctx, res.OrgID, res.RepoName, op, common.HexToAddress(addr))
	case res.Organization != nil:
		vote, err = client.VoteOrganizationAdmin(ctx, res.OrgID, op, common.HexToAddress(addr))
	default:
		return fmt.Errorf("invalid path")
	}

	if err != nil {
		return err
	}

	if big.NewInt(1).Cmp(vote.Threshold) == -1 && vote.SigCount.Cmp(vote.Threshold) == -1 {
		logger.Info("Pending %d/%d votes", vote.SigCount, vote.Threshold)
	} else {
		logger.Info("Approved!")
	}

	return nil
}
