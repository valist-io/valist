# Valist Meta

This repo contains all services required to run Valist on your machine.

## Development

This repo uses submodules to manage dependencies.

To start the docker-compose run the following.

```bash
git clone --recurse-submodules https://github.com/valist-io/valist-meta

cd valist-meta

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
