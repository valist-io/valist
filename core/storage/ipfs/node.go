package ipfs

import (
	"context"
	"io/ioutil"
	"net/http"
	"os"
	"sync"

	config "github.com/ipfs/go-ipfs-config"
	httpapi "github.com/ipfs/go-ipfs-http-client"
	"github.com/ipfs/go-ipfs/core"
	"github.com/ipfs/go-ipfs/core/coreapi"
	"github.com/ipfs/go-ipfs/core/node/libp2p"
	"github.com/ipfs/go-ipfs/plugin/loader"
	"github.com/ipfs/go-ipfs/repo"
	"github.com/ipfs/go-ipfs/repo/fsrepo"
	coreiface "github.com/ipfs/interface-go-ipfs-core"
	"github.com/libp2p/go-libp2p-core/peer"
	multiaddr "github.com/multiformats/go-multiaddr"

	"github.com/valist-io/valist/log"
)

var logger = log.New()

// once is used to ensure plugins are only initialized once.
var once sync.Once

// newCoreAPI returns an IPFS CoreAPI. If a local IPFS istance is running
// a local connection will be attempted, otherwise a new instance is started.
func newCoreAPI(ctx context.Context, repoPath, apiAddr string) (coreiface.CoreAPI, error) {
	var api coreiface.CoreAPI
	var err error

	if env := os.Getenv("VALIST_IPFS_ADDR"); env != "" {
		api, err = httpapi.NewURLApiWithClient(env, &http.Client{})
	} else if apiAddr != "" {
		api, err = httpapi.NewURLApiWithClient(apiAddr, &http.Client{})
	} else {
		api, err = httpapi.NewLocalApi()
	}

	// if we have a local or remote connection return
	if err == nil {
		return api, nil
	}

	logger.Warn("Local IPFS node not found, starting embedded node.")
	logger.Warn("Use a persistent node for a better experience.")

	// make sure plugins are loaded once
	once.Do(setupPlugins)

	repo, err := setupRepo(repoPath)
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

// bootstrap attempts to connect to bootstrap peers.
func bootstrap(ctx context.Context, ipfs coreiface.CoreAPI, peers []string) {
	var wg sync.WaitGroup
	for _, peerString := range peers {
		peerAddr, err := multiaddr.NewMultiaddr(peerString)
		if err != nil {
			logger.Warn("Failed to parse bootstrap peer addr %s", peerString)
			continue
		}

		peerInfo, err := peer.AddrInfoFromP2pAddr(peerAddr)
		if err != nil {
			logger.Warn("Failed to parse bootstrap peer info %s", peerString)
			continue
		}

		wg.Add(1)
		go func(info peer.AddrInfo) {
			defer wg.Done()
			if err := ipfs.Swarm().Connect(ctx, info); err != nil {
				logger.Warn("Failed to bootstrap %s %v", peerInfo.ID, err)
			}
		}(*peerInfo)
	}
	wg.Wait()
}

// setupPlugins initializes the IPFS plugins once.
func setupPlugins() {
	plugins, err := loader.NewPluginLoader("")
	if err != nil {
		panic(err)
	}
	if err := plugins.Initialize(); err != nil {
		panic(err)
	}
	if err := plugins.Inject(); err != nil {
		panic(err)
	}
}

// setupRepo creates the repo if it does not exist.
func setupRepo(repoPath string) (repo.Repo, error) {
	_, err := fsrepo.ConfigAt(repoPath)
	if err == nil {
		return fsrepo.Open(repoPath)
	}
	repoCfg, err := config.Init(ioutil.Discard, 2048)
	if err != nil {
		return nil, err
	}
	err = fsrepo.Init(repoPath, repoCfg)
	if err != nil {
		return nil, err
	}
	return fsrepo.Open(repoPath)
}
