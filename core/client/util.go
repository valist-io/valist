package client

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func packedEncoding(args ...interface{}) ([]byte, error) {
	var arguments abi.Arguments

	for _, arg := range args {
		switch arg.(type) {
		case [32]byte, common.Hash:
			argtype, _ := abi.NewType("bytes32", "", nil)
			arguments = append(arguments, abi.Argument{Type: argtype})
		case string:
			argtype, _ := abi.NewType("string", "", nil)
			arguments = append(arguments, abi.Argument{Type: argtype})
		default:
			return nil, fmt.Errorf("packed encoding type not supported")
		}
	}

	return arguments.Pack(args...)
}

func waitMined(ctx context.Context, eth bind.DeployBackend, tx *types.Transaction) ([]*types.Log, error) {
	if sim, ok := eth.(*backends.SimulatedBackend); ok {
		sim.Commit()
	}

	fmt.Printf("Waiting for transaction %s", tx.Hash().Hex())
	fmt.Printf("View status: https://mumbai.polygonscan.com/tx/%s", tx.Hash().Hex())

	receipt, err := bind.WaitMined(ctx, eth, tx)
	if err != nil {
		return nil, err
	}

	return receipt.Logs, nil
}
