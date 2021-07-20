[![Logo](./docs/img/logo-large-with-text.png)](https://valist.io)

[![Discord](https://img.shields.io/discord/785535462311591976)](https://discord.com/channels/785535462311591976)
[![Valist](https://img.shields.io/badge/valist-published-blue)](https://app.valist.io/valist)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fvalist-io%2Fvalist.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fvalist-io%2Fvalist?ref=badge_shield)

A trustless universal package repository enabling you to digitally sign and distribute software in just a few steps.

## Features

Valist supports the following package types natively (with many more in the pipeline).

* Binaries (software, firmware, you name it)
* NPM Packages (native `npm install` support)
* Python Packages
* Docker Images

## Documentation

Documentation for how to get started with Valist can be found at [https://docs.valist.io](https://docs.valist.io).

## Packages

| Directory            | Description                         |
| -------------------- | ----------------------------------- |
| [cli](./cli)         | Valist command line interface       |
| [hardhat](./hardhat) | Ethereum smart contracts            |
| [lib](./lib)         | Valist core library/sdk             |
| [relay](./relay)     | Valist relay web application & API  |
| [test](./test)       | Acceptance tests                    |

## Building

Install the following requirements:

* node `>= 14.17`
* npm `>= 6.14`

Install all dependencies:

```bash
make install
```

To start the development server:

```bash
make dev
```

To build the contracts, lib, and relay:

```bash
make all
```

To start the relay server in production mode:

```bash
make start
```

## Contributing

We welcome pull requests and would love to support our early contributors with some awesome perks!

Found a bug or have an idea for a feature? [Create an issue](https://github.com/valist-io/valist/issues/new).

## Maintainers

[@awantoch](https://github.com/awantoch)

[@jiyuu-jin](https://github.com/jiyuu-jin)

[@nasdf](https://github.com/nasdf)

## License

Valist is licensed under the [Mozilla Public License Version 2.0](https://www.mozilla.org/en-US/MPL/2.0/).
