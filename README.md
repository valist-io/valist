[![Logo](./img/logo-header.png)](https://valist.io)

[![Go Reference](https://pkg.go.dev/badge/github.com/valist-io/valist.svg)](https://pkg.go.dev/github.com/valist-io/valist)
[![Discord](https://img.shields.io/discord/785535462311591976)](https://discord.com/channels/785535462311591976)
[![Valist](https://img.shields.io/badge/valist-published-blue)](https://app.valist.io/valist)
[![Gitcoin](https://img.shields.io/badge/gitcoin-grant-brightgreen)](https://gitcoin.co/grants/1776/valist)

Web3-native software distribution.

## Features

* Publish Applications and Assets (EXEs, WASM, Docker images, NFTs, etc.)
* Multi-Factor Releases (multi-sig for publishing assets)
* Multi-Chain Support (EVM based only for now, NEAR [coming soon](https://github.com/filecoin-project/devgrants/pull/368))

## Documentation

Documentation for how to get started with Valist can be found at [https://docs.valist.io](https://docs.valist.io).

## Building

This repo contains all services required to run Valist on your machine. However, you can also clone each repo and build components separately.

## Development

Start by cloning the repo and updating submodules.

```bash
git clone https://github.com/valist-io/valist-meta

cd valist-meta

bash update-submodules.sh
```

To start the docker-compose run the following.

```bash
make dev
```

To deploy the contracts and subgraphs to the local ganache run the following.

```bash
make deploy
```

## Contributing

We welcome pull requests and would love to support our early contributors with some awesome perks!

Found a bug or have an idea for a feature? [Create an issue](https://github.com/valist-io/valist-meta/issues/new).

## Maintainers

[@awantoch](https://github.com/awantoch)

[@jiyuu-jin](https://github.com/jiyuu-jin)

[@nasdf](https://github.com/nasdf)

## License

Valist is licensed under the [Mozilla Public License Version 2.0](https://www.mozilla.org/en-US/MPL/2.0/).
