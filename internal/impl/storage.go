package impl

import (
	"context"
	"fmt"
	"io"

	"github.com/ipfs/go-cid"
	files "github.com/ipfs/go-ipfs-files"
	"github.com/ipfs/interface-go-ipfs-core/path"
)

func (client *Client) AddFile(ctx context.Context, data []byte) (cid.Cid, error) {
	path, err := client.ipfs.Unixfs().Add(ctx, files.NewBytesFile(data))
	if err != nil {
		return cid.Cid{}, err
	}

	return path.Cid(), nil
}

func (client *Client) GetFile(ctx context.Context, id cid.Cid) ([]byte, error) {
	node, err := client.ipfs.Unixfs().Get(ctx, path.IpfsPath(id))
	if err != nil {
		return nil, err
	}

	file, ok := node.(files.File)
	if !ok {
		return nil, fmt.Errorf("Failed to parse organization meta")
	}

	return io.ReadAll(file)
}
