package ipfs

import (
	"context"
	"io"
	"path/filepath"

	config "github.com/ipfs/go-ipfs-config"
	httpapi "github.com/ipfs/go-ipfs-http-client"
	"github.com/ipfs/go-ipfs/core"
	"github.com/ipfs/go-ipfs/core/coreapi"
	libp2p "github.com/ipfs/go-ipfs/core/node/libp2p"
	"github.com/ipfs/go-ipfs/plugin/loader"
	"github.com/ipfs/go-ipfs/repo/fsrepo"
	icore "github.com/ipfs/interface-go-ipfs-core"
)

// NewCoreAPI returns a new IPFS core api.
func NewCoreAPI(ctx context.Context) (icore.CoreAPI, error) {
	api, err := httpapi.NewLocalApi()
	if err == nil {
		return api, nil
	}

	root, err := config.PathRoot()
	if err != nil {
		return nil, err
	}

	plugins, err := loader.NewPluginLoader(filepath.Join(root, "plugins"))
	if err != nil {
		return nil, err
	}

	if err := plugins.Initialize(); err != nil {
		return nil, err
	}

	if err := plugins.Inject(); err != nil {
		return nil, err
	}

	cfg, err := config.Init(io.Discard, 2048)
	if err != nil {
		return nil, err
	}

	if err := fsrepo.Init(root, cfg); err != nil {
		return nil, err
	}

	repo, err := fsrepo.Open(root)
	if err != nil {
		return nil, err
	}

	opts := &core.BuildCfg{
		Online:    true,
		Permanent: true,
		Routing:   libp2p.DHTOption,
		Repo:      repo,
	}

	node, err := core.NewNode(ctx, opts)
	if err != nil {
		return nil, err
	}

	return coreapi.NewCoreAPI(node)
}
