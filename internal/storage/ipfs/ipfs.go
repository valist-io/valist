package ipfs

import (
	"context"
	"fmt"
	"io"
	"io/fs"
	"os"
	"strings"

	cid "github.com/ipfs/go-cid"
	files "github.com/ipfs/go-ipfs-files"
	merkledag "github.com/ipfs/go-merkledag"
	coreiface "github.com/ipfs/interface-go-ipfs-core"
	"github.com/ipfs/interface-go-ipfs-core/options"
	"github.com/ipfs/interface-go-ipfs-core/path"
	car "github.com/ipld/go-car"

	"github.com/valist-io/valist/internal/storage"
)

var addopts = []options.UnixfsAddOption{
	options.Unixfs.Pin(true),
}

type Provider struct {
	ipfs coreiface.CoreAPI
}

func NewProvider(ctx context.Context, repoPath string) (*Provider, error) {
	ipfs, err := NewCoreAPI(ctx, repoPath)
	if err != nil {
		return nil, err
	}

	return &Provider{ipfs}, nil
}

// Export returns a reader for exporting CAR data.
func (prov *Provider) Export(ctx context.Context, fpath string) (io.Reader, error) {
	res, err := prov.ipfs.ResolvePath(ctx, path.New(fpath))
	if err != nil {
		return nil, err
	}

	pr, pw := io.Pipe()
	go func() {
		ses := merkledag.NewSession(ctx, prov.ipfs.Dag())
		err := car.WriteCar(ctx, ses, []cid.Cid{res.Cid()}, pw)
		pw.CloseWithError(err) //nolint:errcheck
	}()

	return pr, nil
}

func (prov *Provider) Open(ctx context.Context, fpath string) (storage.File, error) {
	node, err := prov.ipfs.Unixfs().Get(ctx, path.New(fpath))
	if isNotExist(err) {
		return nil, os.ErrNotExist
	}

	if err != nil {
		return nil, err
	}

	f, ok := node.(files.File)
	if !ok {
		return nil, fmt.Errorf("cannot open directory: %s", fpath)
	}

	return &file{"", f}, nil
}

func (prov *Provider) ReadDir(ctx context.Context, fpath string) ([]fs.FileInfo, error) {
	node, err := prov.ipfs.Unixfs().Get(ctx, path.New(fpath))
	if isNotExist(err) {
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

func (prov *Provider) ReadFile(ctx context.Context, fpath string) ([]byte, error) {
	file, err := prov.Open(ctx, fpath)
	if err != nil {
		return nil, err
	}

	return io.ReadAll(file)
}

func (prov *Provider) WriteFile(ctx context.Context, fpath string) (string, error) {
	info, err := os.Stat(fpath)
	if err != nil {
		return "", err
	}

	node, err := files.NewSerialFile(fpath, false, info)
	if err != nil {
		return "", err
	}

	p, err := prov.ipfs.Unixfs().Add(ctx, node, addopts...)
	if err != nil {
		return "", err
	}

	return p.String(), nil
}

func (prov *Provider) Write(ctx context.Context, data []byte) (string, error) {
	p, err := prov.ipfs.Unixfs().Add(ctx, files.NewBytesFile(data), addopts...)
	if err != nil {
		return "", err
	}

	return p.String(), nil
}

// isNotExist returns true if the error is not exists.
func isNotExist(err error) bool {
	if err == nil {
		return false
	}

	return strings.HasPrefix(err.Error(), "no link named")
}
