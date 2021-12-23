package ipfs

import (
	"context"
	"fmt"
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
	"github.com/ipfs/go-ipfs/repo/fsrepo"
	coreiface "github.com/ipfs/interface-go-ipfs-core"
	"github.com/libp2p/go-libp2p-core/peer"
	multiaddr "github.com/multiformats/go-multiaddr"
)

// once is used to ensure plugins are only initialized once.
var once sync.Once

// NewCoreAPI returns an IPFS CoreAPI. If a local IPFS istance is running
// a local connection will be attempted, otherwise a new instance is started.
func NewCoreAPI(ctx context.Context, repoPath, apiAddr string) (coreiface.CoreAPI, error) {
	local, err := connectToIPFS(ctx, apiAddr)
	if err == nil {
		return local, nil
	}

	fmt.Println("Local IPFS node not found, starting embedded node.")
	fmt.Println("Use a persistent node for a better experience.")

	once.Do(setupPlugins)
	if err := initRepo(repoPath); err != nil {
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

	return coreapi.NewCoreAPI(node)
}

// Bootstrap attempts to connect to bootstrap peers.
func Bootstrap(ctx context.Context, ipfs coreiface.CoreAPI, peers []string) {
	var wg sync.WaitGroup
	for _, peerString := range peers {
		peerAddr, err := multiaddr.NewMultiaddr(peerString)
		if err != nil {
			fmt.Printf("Failed to parse bootstrap peer addr %s\n", peerString)
			continue
		}

		peerInfo, err := peer.AddrInfoFromP2pAddr(peerAddr)
		if err != nil {
			fmt.Printf("Failed to parse bootstrap peer info %s\n", peerString)
			continue
		}

		wg.Add(1)
		go func(info peer.AddrInfo) {
			defer wg.Done()
			if err := ipfs.Swarm().Connect(ctx, info); err != nil {
				fmt.Printf("Failed to bootstrap %s %v\n", peerInfo.ID, err)
			}
		}(*peerInfo)
	}
	wg.Wait()
}

// connectToIPFS attempts to connect to the local IPFS API and
// makes a request to ensure the API is running.
func connectToIPFS(ctx context.Context, apiAddr string) (coreiface.CoreAPI, error) {
	var api coreiface.CoreAPI
	var err error

	// choose a host in this order
	// 1. address from environment
	// 2. address from config
	// 3. attempt local node
	if envAddr := os.Getenv("VALIST_IPFS_ADDR"); envAddr != "" {
		api, err = httpapi.NewURLApiWithClient(envAddr, &http.Client{})
	} else if apiAddr != "" {
		api, err = httpapi.NewURLApiWithClient(apiAddr, &http.Client{})
	} else {
		api, err = httpapi.NewLocalApi()
	}

	if err != nil {
		return nil, err
	}

	// sanity check that the api is active
	_, err = api.Swarm().ListenAddrs(ctx)
	if err != nil {
		return nil, err
	}

	return api, nil
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
