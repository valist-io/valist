package client

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/ipfs/go-cid"
	files "github.com/ipfs/go-ipfs-files"
	"github.com/ipfs/interface-go-ipfs-core/path"
	"github.com/ipfs/interface-go-ipfs-core/options"
)

func (client *Client) WriteFile(ctx context.Context, data []byte) (cid.Cid, error) {
	path, err := client.ipfs.Unixfs().Add(ctx, files.NewBytesFile(data), options.Unixfs.Pin(true))
	if err != nil {
		return cid.Cid{}, err
	}

	return path.Cid(), nil
}

func (client *Client) ReadFile(ctx context.Context, id cid.Cid) ([]byte, error) {
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

func (client *Client) WriteFilePath(ctx context.Context, fpath string) (cid.Cid, error) {
	info, err := os.Stat(fpath)
	if err != nil {
		return cid.Cid{}, err
	}

	node, err := files.NewSerialFile(fpath, false, info)
	if err != nil {
		return cid.Cid{}, err
	}

	path, err := client.ipfs.Unixfs().Add(ctx, node, options.Unixfs.Pin(true))
	if err != nil {
		return cid.Cid{}, err
	}

	return path.Cid(), nil
}

func (client *Client) WriteDirEntries(ctx context.Context, base string, fpaths []string) (cid.Cid, error) {
	entries := make(map[string]files.Node)
	for _, fpath := range fpaths {
		info, err := os.Stat(fpath)
		if err != nil {
			return cid.Cid{}, err
		}

		node, err := files.NewSerialFile(fpath, false, info)
		if err != nil {
			return cid.Cid{}, err
		}

		rel, err := filepath.Rel(base, fpath)
		if err == nil {
			fpath = rel
		}

		entries[fpath] = node
	}

	path, err := client.ipfs.Unixfs().Add(ctx, files.NewMapDirectory(entries), options.Unixfs.Pin(true))
	if err != nil {
		return cid.Cid{}, err
	}

	return path.Cid(), nil
}
