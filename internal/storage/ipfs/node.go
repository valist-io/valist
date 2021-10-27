package ipfs

import (
	"context"
	"fmt"
	"io/ioutil"

	config "github.com/ipfs/go-ipfs-config"
	httpapi "github.com/ipfs/go-ipfs-http-client"
	"github.com/ipfs/go-ipfs/core"
	"github.com/ipfs/go-ipfs/core/coreapi"
	"github.com/ipfs/go-ipfs/core/node/libp2p"
	"github.com/ipfs/go-ipfs/plugin/loader"
	"github.com/ipfs/go-ipfs/repo/fsrepo"
	coreiface "github.com/ipfs/interface-go-ipfs-core"
)

// NewCoreAPI returns an IPFS CoreAPI. If a local IPFS istance is running
// a local connection will be attempted, otherwise a new instance is started.
func NewCoreAPI(ctx context.Context, repoPath string) (coreiface.CoreAPI, error) {
	local, err := httpapi.NewLocalApi()
	if err == nil {
		return local, nil
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

	// create fsrepo if not initialized
	if _, err = fsrepo.ConfigAt(repoPath); err != nil {
		repoCfg, err := config.Init(ioutil.Discard, 2048)
		if err != nil {
			return nil, err
		}

		err = fsrepo.Init(repoPath, repoCfg)
		if err != nil {
			return nil, fmt.Errorf("failed to init fsrepo: %s", err)
		}
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

	return coreapi.NewCoreAPI(node)
}
