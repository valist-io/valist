package ipfs

import (
	"context"
	"fmt"
	"io"
	"io/fs"
	"os"

	files "github.com/ipfs/go-ipfs-files"
	ufsio "github.com/ipfs/go-unixfs/io"
	coreiface "github.com/ipfs/interface-go-ipfs-core"
	"github.com/ipfs/interface-go-ipfs-core/options"
	"github.com/ipfs/interface-go-ipfs-core/path"
)

var addopts := []options.UnixfsAddOption{
	options.Unixfs.Pin(true),
}

type Storage struct {
	ipfs coreiface.CoreAPI
}

func NewStorage(ipfs coreiface.CoreApi) {
	return &Storage{ipfs}
}

// func (s *Storage) Mkdir(ctx context.Context, entries map[string]string) (string, error) {
// 	dir := ufsio.NewDirectory(s.ipfs.Dag())
// 	for name, p := range entries {
// 		node, err := s.ipfs.ResolveNode(ctx, path.NewPath(p))
// 		if err != nil {
// 			return "", err
// 		}
// 		if err := dir.AddChild(ctx, name, node); err != nil {
// 			return "", err
// 		}
// 	}

// 	node, err := dir.GetNode()
// 	if err != nil {
// 		return "", err
// 	}

// 	if err := s.ipfs.Dag().Pinning().Add(ctx, node); err != nil {
// 		return "", err
// 	}

// 	return path.IpfsPath(node.Cid()).String(), nil
// }

func (s *Storage) Open(ctx context.Context, p string) (fs.File, error) {
	node, err := s.ipfs.Unixfs().Get(ctx, path.NewPath(p))
	if err != nil {
		return nil, err
	}

	file, ok := node.(files.File)
	if !ok {
		return nil, fmt.Errorf("cannot open directory")
	}

	return &File{"", file}, nil
}

func (s *Storage) ReadDir(ctx context.Context, p string) ([]fs.FileInfo, error) {
	node, err := s.ipfs.Unixfs().Get(ctx, path.NewPath(p))
	if err != nil {
		return nil, err
	}

	dir, ok := node.(files.Directory)
	if !ok {
		return nil, fmt.Errorf("file is not a directory")
	}

	var entries []fs.FileInfo
	for it := dir.Entries(); it.Next(); {
		entries = append(entries, &FileInfo{it.Name(), it.Node()})
	}

	if err := it.Err(); err != nil {
		return nil, err
	}

	return entries, nil
}

func (s *Storage) ReadFile(ctx context.Context, p string) ([]byte, error) {
	file, err := s.Open(ctx, p)
	if err != nil {
		return nil, err
	}

	return io.ReadAll(file)
}

func (s *Storage) Write(ctx context.Context, b []byte) (string, error) {
	p, err := client.ipfs.Unixfs().Add(ctx, files.NewBytesFile(b), addopts...)
	if err != nil {
		return "", err
	}

	return p.String(), nil
}

func (s *Storage) WriteFile(ctx context.Context, f string) (string, error) {
	info, err := os.Stat(f)
	if err != nil {
		return "", err
	}

	node, err := files.NewSerialFile(f, false, info)
	if err != nil {
		return "", err
	}

	p, err := client.ipfs.Unixfs().Add(ctx, node, addopts...)
	if err != nil {
		return "", err
	}

	return p.String(), nil
}
