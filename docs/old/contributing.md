# Contributing

We welcome pull requests, and would love to support our early contributors with some awesome perks!

We'd also love for you to [join our Discord server](https://discord.gg/DcFb7SGDtN) and hang out with us, we're building a community around infosec, software engineers, and blockchain enthusiasts!

## Building

A `Makefile` is provided to simplify development workflow.

### Installing Dependencies

To install all dependencies, simply run the following:

```bash
make install
```

This will `cd` into the `lib`, `relay`, and `eth` folders, running an `npm install` in each.

You can also manually run `npm install` in each of the folders.

### Building the frontend

You can build the `lib` and `relay` separately as Make targets.

However, you also can build the whole frontend using:

```bash
make frontend
```

This works great inside of CI/CD tools as well. For example, `make install frontend` will work in most runners as the build command.

### Building all artifacts

To build the contracts in `eth`, the core in `lib`, and the `relay`, simply run:

```bash
make all
```

## Development

### Running the dev server

The frontend and API is powered by Next.js. To start the development server, run the following command:

```bash
make dev
```

This will build the `lib` and `relay` folders and start the local Next.js development server at `localhost:3000`. Hot reloading works for both `lib` and `relay` using `tsc -w` in `lib`, and `next dev` in `relay`. They are executed in parallel, piping output to the same shell, thanks to `make -j`.

### Migrating contracts

Truffle is managing the contract migrations and deployments. Make sure you have a development chain setup, and modify the `truffle-config.js` accordingly.

Once your Truffle environment is established, simply run:

```bash
make migrate
```

You can also run the `make deploy` alias.

This will deploy the Solidity contracts to your chain of choice. Make sure to update your contract addresses if/when necessary!
