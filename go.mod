module github.com/valist-io/registry

go 1.16

require (
	github.com/ethereum/go-ethereum v1.10.6
	github.com/gin-gonic/gin v1.7.2
	github.com/ipfs/go-cid v0.0.7
	github.com/ipfs/go-ipfs v0.9.1
	github.com/ipfs/go-ipfs-files v0.0.8
	github.com/ipfs/go-ipfs-http-client v0.1.0
	github.com/ipfs/interface-go-ipfs-core v0.4.0
	github.com/libp2p/go-libp2p-core v0.8.5
	github.com/manifoldco/promptui v0.8.0
	github.com/mattn/go-runewidth v0.0.10 // indirect
	github.com/multiformats/go-multiaddr v0.3.3
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/stretchr/testify v1.7.0
	github.com/urfave/cli/v2 v2.3.0
	github.com/valist-io/gasless v0.0.1
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c
)

replace github.com/ethereum/go-ethereum => github.com/nasdf/go-ethereum v1.10.7-0.20210731182913-02804a7b22b2

replace github.com/valist-io/gasless => ../gasless
