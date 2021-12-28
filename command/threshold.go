package command

import (
	"context"
	"fmt"
	"math/big"

	"github.com/valist-io/valist/contract/valist"
	"github.com/valist-io/valist/core/client"
)

// Threshold updates the signature threshold for an organization or repository.
func Threshold(ctx context.Context, rpath string, threshold int64) error {
	client := ctx.Value(ClientKey).(*client.Client)

	res, err := client.ResolvePath(ctx, rpath)
	if err != nil {
		return err
	}

	var vote *valist.ValistVoteThresholdEvent
	switch {
	case res.Repository != nil:
		vote, err = client.VoteRepositoryThreshold(ctx, res.OrgID, res.RepoName, big.NewInt(threshold))
	case res.Organization != nil:
		vote, err = client.VoteOrganizationThreshold(ctx, res.OrgID, big.NewInt(threshold))
	default:
		return fmt.Errorf("invalid path")
	}

	if err != nil {
		return err
	}

	if big.NewInt(1).Cmp(vote.Threshold) == -1 && vote.SigCount.Cmp(vote.Threshold) == -1 {
		logger.Info("Voted to set threshold %d/%d", vote.SigCount, threshold)
	} else {
		logger.Info("Approved threshold %d", threshold)
	}

	return nil
}
