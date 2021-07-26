package core

import (
	"context"

	"github.com/ipfs/go-cid"
)

type StorageAPI interface {
	AddFile(context.Context, []byte) (cid.Cid, error)
	GetFile(context.Context, cid.Cid) ([]byte, error)
}
