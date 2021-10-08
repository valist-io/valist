package ipfs

import (
	"context"
	"fmt"
	"io"
	"io/fs"
	"os"
	"strings"

	files "github.com/ipfs/go-ipfs-files"
	httpapi "github.com/ipfs/go-ipfs-http-client"
	"github.com/ipfs/go-ipfs/core"
	"github.com/ipfs/go-ipfs/core/coreapi"
	"github.com/ipfs/go-ipfs/core/node/libp2p"
	"github.com/ipfs/go-ipfs/plugin/loader"
	"github.com/ipfs/go-ipfs/repo/fsrepo"
	coreiface "github.com/ipfs/interface-go-ipfs-core"
	"github.com/ipfs/interface-go-ipfs-core/options"
	"github.com/ipfs/interface-go-ipfs-core/path"

	"github.com/valist-io/valist/internal/storage"
)

var addopts = []options.UnixfsAddOption{
	options.Unixfs.Pin(true),
}

type Provider struct {
	ipfs coreiface.CoreAPI
}

func NewProvider(ctx context.Context, repoPath string) (*Provider, error) {
	local, err := httpapi.NewLocalApi()
	if err == nil {
		return &Provider{local}, nil
	}

	plugins, err := loader.NewPluginLoader("")
	if err != nil {
		return nil, err
	}

	if err := plugins.Initialize(); err != nil {
		return nil, err
	}

	if err := plugins.Inject(); err != nil {
		return nil, err
	}

	repo, err := fsrepo.Open(repoPath)
	if err != nil {
		return nil, err
	}

	cfg := &core.BuildCfg{
		Online:  true,
		Routing: libp2p.DHTOption,
		Repo:    repo,
	}

	node, err := core.NewNode(ctx, cfg)
	if err != nil {
		return nil, err
	}

	ipfs, err := coreapi.NewCoreAPI(node)
	if err != nil {
		return nil, err
	}

	return &Provider{ipfs}, nil
}

func (prov *Provider) Prefix() string {
	return "ipfs"
}

func (prov *Provider) Open(ctx context.Context, fpath string) (storage.File, error) {
	node, err := prov.ipfs.Unixfs().Get(ctx, path.New(fpath))
	if IsNotExist(err) {
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

// IsNotExist returns true if the error is not exists.
func IsNotExist(err error) bool {
	if err == nil {
		return false
	}

	return strings.HasPrefix(err.Error(), "no link named")
}
