package client

import (
	"context"
	"io"
	"os"

	files "github.com/ipfs/go-ipfs-files"
	"github.com/ipfs/interface-go-ipfs-core/options"
	"github.com/ipfs/interface-go-ipfs-core/path"

	"github.com/valist-io/valist/ipfs"
)

func (client *Client) ReadFile(ctx context.Context, fpath string) ([]byte, error) {
	node, err := client.ipfs.Unixfs().Get(ctx, path.New(fpath))
	if err != nil {
		return nil, err
	}

	file, ok := node.(files.File)
	if !ok {
		return nil, os.ErrNotExist
	}

	return io.ReadAll(file)
}

func (client *Client) WriteFile(ctx context.Context, data []byte) (string, error) {
	p, err := client.ipfs.Unixfs().Add(ctx, files.NewBytesFile(data), options.Unixfs.Pin(true))
	if err != nil {
		return "", err
	}

	if err := ipfs.ExportCAR(ctx, client.ipfs.Dag(), p.Cid()); err != nil {
		return "", err
	}

	return p.String(), nil
}
