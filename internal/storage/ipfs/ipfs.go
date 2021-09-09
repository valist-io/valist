package ipfs

import (
	"context"
	"fmt"
	"io"
	"io/fs"
	"os"
	"strings"

	files "github.com/ipfs/go-ipfs-files"
	coreiface "github.com/ipfs/interface-go-ipfs-core"
	"github.com/ipfs/interface-go-ipfs-core/options"
	"github.com/ipfs/interface-go-ipfs-core/path"

	"github.com/valist-io/valist/internal/storage"
)

var addopts = []options.UnixfsAddOption{
	options.Unixfs.Pin(true),
}

type Storage struct {
	ipfs coreiface.CoreAPI
}

func NewStorage(ipfs coreiface.CoreAPI) *Storage {
	return &Storage{ipfs}
}

func (s *Storage) Mkdir() storage.Directory {
	return &dir{s.ipfs, emptyDirPath}
}

func (s *Storage) Open(ctx context.Context, p string) (storage.File, error) {
	node, err := s.ipfs.Unixfs().Get(ctx, path.New(p))
	if IsNotExist(err) {
		return nil, os.ErrNotExist
	}

	if err != nil {
		return nil, err
	}

	f, ok := node.(files.File)
	if !ok {
		return nil, fmt.Errorf("cannot open directory: %s", p)
	}

	return &file{"", f}, nil
}

func (s *Storage) ReadDir(ctx context.Context, p string) ([]fs.FileInfo, error) {
	node, err := s.ipfs.Unixfs().Get(ctx, path.New(p))
	if IsNotExist(err) {
		return nil, os.ErrNotExist
	}

	if err != nil {
		return nil, err
	}

	dir, ok := node.(files.Directory)
	if !ok {
		return nil, fmt.Errorf("file is not a directory")
	}
	it := dir.Entries()

	var entries []fs.FileInfo
	for it.Next() {
		entries = append(entries, &fileInfo{it.Name(), it.Node()})
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
	p, err := s.ipfs.Unixfs().Add(ctx, files.NewBytesFile(b), addopts...)
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

	p, err := s.ipfs.Unixfs().Add(ctx, node, addopts...)
	if err != nil {
		return "", err
	}

	return p.String(), nil
}

// IsNotExist returns true if the error is not exists.
func IsNotExist(err error) bool {
	if err == nil {
		return false
	}

	return strings.HasPrefix(err.Error(), "no link named")
}
