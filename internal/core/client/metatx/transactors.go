package metatx

import (
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ipfs/go-cid"

	"github.com/valist-io/gasless"
	"github.com/valist-io/registry/internal/contract/registry"
	"github.com/valist-io/registry/internal/contract/valist"
	"github.com/valist-io/registry/internal/core/config"
	"github.com/valist-io/registry/internal/core/types"
)

func (t *Transactor) CreateOrganizationTx(txopts *bind.TransactOpts, metaCID cid.Cid) (*ethtypes.Transaction, error) {
	parsedABI, err := abi.JSON(strings.NewReader(valist.ValistABI))
	if err != nil {
		return nil, err
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	builder := gasless.NewMessageBuilder(parsedABI, config.Default(home).Ethereum.Contracts["valist"], t.eth)
	msg, err := builder.Message(txopts.Context, "createOrganization", metaCID.String())
	if err != nil {
		return nil, err
	}

	return t.meta.Transact(txopts.Context, msg, t.signer, createOrganizationBFID)
}

func (t *Transactor) LinkOrganizationNameTx(txopts *bind.TransactOpts, orgID common.Hash, name string) (*ethtypes.Transaction, error) {
	parsedABI, err := abi.JSON(strings.NewReader(registry.ValistRegistryABI))
	if err != nil {
		return nil, err
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	builder := gasless.NewMessageBuilder(parsedABI, config.Default(home).Ethereum.Contracts["registry"], t.eth)
	msg, err := builder.Message(txopts.Context, "linkNameToID", orgID, name)
	if err != nil {
		return nil, err
	}

	return t.meta.Transact(txopts.Context, msg, t.signer, linkNameToIDBFID)
}

func (t *Transactor) CreateRepositoryTx(txopts *bind.TransactOpts, orgID common.Hash, repoName string, repoMeta string) (*ethtypes.Transaction, error) {
	parsedABI, err := abi.JSON(strings.NewReader(valist.ValistABI))
	if err != nil {
		return nil, err
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	builder := gasless.NewMessageBuilder(parsedABI, config.Default(home).Ethereum.Contracts["valist"], t.eth)
	msg, err := builder.Message(txopts.Context, "createRepository", orgID, repoName, repoMeta)
	if err != nil {
		return nil, err
	}

	return t.meta.Transact(txopts.Context, msg, t.signer, createRepositoryBFID)
}

func (t *Transactor) VoteReleaseTx(txopts *bind.TransactOpts, orgID common.Hash, repoName string, release *types.Release) (*ethtypes.Transaction, error) {
	parsedABI, err := abi.JSON(strings.NewReader(valist.ValistABI))
	if err != nil {
		return nil, err
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	builder := gasless.NewMessageBuilder(parsedABI, config.Default(home).Ethereum.Contracts["valist"], t.eth)
	msg, err := builder.Message(txopts.Context, "voteRelease", orgID, repoName, release.ReleaseCID, release.MetaCID)
	if err != nil {
		return nil, err
	}

	return t.meta.Transact(txopts.Context, msg, t.signer, voteReleaseBFID)
}

func (t *Transactor) VoteKeyTx(txopts *bind.TransactOpts, orgID common.Hash, repoName string, operation common.Hash, address common.Address) (*ethtypes.Transaction, error) {
	parsedABI, err := abi.JSON(strings.NewReader(valist.ValistABI))
	if err != nil {
		return nil, err
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	builder := gasless.NewMessageBuilder(parsedABI, config.Default(home).Ethereum.Contracts["valist"], t.eth)
	msg, err := builder.Message(txopts.Context, "voteKey", orgID, repoName, operation, address)
	if err != nil {
		return nil, err
	}

	return t.meta.Transact(txopts.Context, msg, t.signer, voteKeyBFID)
}

func (t *Transactor) SetRepositoryMetaTx(txopts *bind.TransactOpts, orgID common.Hash, repoName string, repoMeta string) (*ethtypes.Transaction, error) {
	parsedABI, err := abi.JSON(strings.NewReader(valist.ValistABI))
	if err != nil {
		return nil, err
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	builder := gasless.NewMessageBuilder(parsedABI, config.Default(home).Ethereum.Contracts["valist"], t.eth)
	msg, err := builder.Message(txopts.Context, "setRepoMeta", orgID, repoName, repoMeta)
	if err != nil {
		return nil, err
	}
	return t.meta.Transact(txopts.Context, msg, t.signer, setRepoMetaBFID)
}

func (t *Transactor) VoteRepositoryThresholdTx(txopts *bind.TransactOpts, orgID common.Hash, repoName string, threshold *big.Int) (*ethtypes.Transaction, error) {
	parsedABI, err := abi.JSON(strings.NewReader(valist.ValistABI))
	if err != nil {
		return nil, err
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	builder := gasless.NewMessageBuilder(parsedABI, config.Default(home).Ethereum.Contracts["valist"], t.eth)
	msg, err := builder.Message(txopts.Context, "voteThreshold", orgID, repoName, threshold)
	if err != nil {
		return nil, err
	}

	return t.meta.Transact(txopts.Context, msg, t.signer, voteThresholdBFID)
}

func (t *Transactor) VoteOrganizationThresholdTx(txopts *bind.TransactOpts, orgID common.Hash, threshold *big.Int) (*ethtypes.Transaction, error) {
	parsedABI, err := abi.JSON(strings.NewReader(valist.ValistABI))
	if err != nil {
		return nil, err
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	builder := gasless.NewMessageBuilder(parsedABI, config.Default(home).Ethereum.Contracts["valist"], t.eth)
	msg, err := builder.Message(txopts.Context, "voteThreshold", orgID, "", threshold)
	if err != nil {
		return nil, err
	}

	return t.meta.Transact(txopts.Context, msg, t.signer, voteThresholdBFID)

}
