[![Logo](./assets/Logo.png)](https://valist.io)

[![Discord](https://img.shields.io/discord/785535462311591976)](https://discord.com/channels/785535462311591976)
[![Valist](https://img.shields.io/badge/valist-published-blue)](https://app.valist.io/valist)

A trustless universal package repository enabling you to digitally sign and distribute software in just a few steps.

## Features

Valist supports the following package types natively (with many more in the pipeline).

* Binaries (software, firmware, you name it)
* NPM (native `npm install` support)
* Pip
* Docker

## Documentation

Documentation for how to get started with Valist can be found at [https://docs.valist.io](https://docs.valist.io).

## Packages

| Directory            | Description                   |
| -------------------- | ----------------------------- |
| [cli](./cli)         | Valist command line interface |
| [hardhat](./hardhat) | Ethereum smart contracts      |
| [lib](./lib)         | Valist core library           |
| [relay](./relay)     | Valist relay web application  |
| [test](./test)       | Acceptance tests              |

## Building

Install the following requirements:

- node `14.17`
- npm `6.14`

To start the development server:

```bash
make dev
```

To build the contracts, lib, and relay:

```bash
make all
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
