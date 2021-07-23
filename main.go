package main

import (
	"context"
	"log"
	"math/big"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/libp2p/go-libp2p-core/peer"
	ma "github.com/multiformats/go-multiaddr"

	"github.com/valist-io/registry/internal/contract"
	"github.com/valist-io/registry/internal/core"
	"github.com/valist-io/registry/internal/http"
	"github.com/valist-io/registry/internal/ipfs"
)

const (
	bindAddr    = ":8080"
	ethereumRPC = "https://rpc.valist.io"
)

var (
	chainID         = big.NewInt(80001)
	peerAddress     = ma.StringCast("/ip4/107.191.98.233/tcp/4001/p2p/QmasbWJE9C7PVFVj1CVQLX617CrDQijCxMv6ajkRfaTi98")
	valistAddress   = common.HexToAddress("0xA7E4124aDBBc50CF402e4Cad47de906a14daa0f6")
	registryAddress = common.HexToAddress("0x2Be6D782dBA2C52Cd0a41c6052e914dCaBcCD78e")
)

func main() {
	ipfs, err := ipfs.NewCoreAPI(context.Background())
	if err != nil {
		log.Fatalf("Failed to connect to ipfs: %v", err)
	}

	peerInfo, err := peer.AddrInfoFromP2pAddr(peerAddress)
	if err != nil {
		log.Fatalf("Failed to get peer info: %v", err)
	}

	// attempt to connect to valist gateway peer
	go ipfs.Swarm().Connect(context.Background(), *peerInfo)

	eth, err := ethclient.Dial(ethereumRPC)
	if err != nil {
		log.Fatalf("Failed to connect to rpc: %v", err)
	}

	valist, err := contract.NewValist(valistAddress, eth)
	if err != nil {
		log.Fatalf("Failed to instantiate valist contract: %v", err)
	}

	registry, err := contract.NewRegistry(registryAddress, eth)
	if err != nil {
		log.Fatalf("Failed to instantiate registry contract: %v", err)
	}

	client := core.NewClient(eth, ipfs, valist, registry, nil, chainID)
	server := http.NewServer(client, bindAddr)

	log.Println("Server running on", bindAddr)
	go server.ListenAndServe()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("Shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	server.Shutdown(ctx)
}
