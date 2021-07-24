package core

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
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
