package ipfs

import (
	"context"
	"fmt"
	"io/ioutil"
	"sync"

	config "github.com/ipfs/go-ipfs-config"
	httpapi "github.com/ipfs/go-ipfs-http-client"
	"github.com/ipfs/go-ipfs/core"
	"github.com/ipfs/go-ipfs/core/coreapi"
	"github.com/ipfs/go-ipfs/core/node/libp2p"
	"github.com/ipfs/go-ipfs/plugin/loader"
	"github.com/ipfs/go-ipfs/repo/fsrepo"
	coreiface "github.com/ipfs/interface-go-ipfs-core"
	"github.com/libp2p/go-libp2p-core/peer"
	multiaddr "github.com/multiformats/go-multiaddr"
)

var bootstrapPeers = []string{
	"/dnsaddr/gateway.valist.io/p2p/QmasbWJE9C7PVFVj1CVQLX617CrDQijCxMv6ajkRfaTi98",
}

// once is used to ensure plugins are only initialized once.
var once sync.Once

// NewCoreAPI returns a new IPFS instance bootstrapped with the default peers.
func NewCoreAPI(ctx context.Context, repoPath string) (coreiface.CoreAPI, error) {
	ipfs, err := newIPFS(ctx, repoPath)
	if err != nil {
		return nil, err
	}

	for _, peerString := range bootstrapPeers {
		peerAddr, err := multiaddr.NewMultiaddr(peerString)
		if err != nil {
			continue
		}

		peerInfo, err := peer.AddrInfoFromP2pAddr(peerAddr)
		if err != nil {
			continue
		}

		go ipfs.Swarm().Connect(ctx, *peerInfo) //nolint:errcheck
	}

	return ipfs, nil
}

// newIPFS returns an IPFS CoreAPI. If a local IPFS istance is running
// a local connection will be attempted, otherwise a new instance is started.
func newIPFS(ctx context.Context, repoPath string) (coreiface.CoreAPI, error) {
	local, err := connectToIPFS(ctx)
	if err == nil {
		return local, nil
	}

	fmt.Println("Local IPFS node not found, using embedded node instead.")
	once.Do(setupPlugins)

	if err := initRepo(repoPath); err != nil {
		return nil, fmt.Errorf("failed to init fsrepo: %s", err)
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

// connectToIPFS attempts to connect to the local IPFS API and
// makes a request to ensure the API is running.
func connectToIPFS(ctx context.Context) (coreiface.CoreAPI, error) {
	local, err := httpapi.NewLocalApi()
	if err != nil {
		return nil, err
	}

	// make a request to ensure the api is actually running
	_, err = local.Swarm().ListenAddrs(ctx)
	if err != nil {
		return nil, err
	}

	return local, nil
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

// initRepo creates the repo if it does not exist.
func initRepo(repoPath string) error {
	if _, err := fsrepo.ConfigAt(repoPath); err == nil {
		return nil
	}

	repoCfg, err := config.Init(ioutil.Discard, 2048)
	if err != nil {
		return err
	}

	return fsrepo.Init(repoPath, repoCfg)
}
