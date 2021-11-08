---
title: NPM Registry
---

## NPM Proxy Cache

Valist uses IPFS to store data on your local machine. You can use Valist to cache all of your NPM dependencies across every project, similar to Verdaccio.

Since IPFS is a networked filesystem, this enables easily transferring cached dependencies between your environments.

To enable the proxy cache, first, ensure the `valist daemon` is running:

```sh
valist daemon
```

Then, set your NPM registry config to use your local Valist relay:

```sh
npm config set registry http://localhost:9000/proxy/npm
```

Then, continue using and installing packages via `npm` as you normally would, and enjoy the benefits of IPFS-powered NPM packages!

!!! note
    When publishing to the upstream NPMjs.com registry, you'll need to run `npm publish --registry=https://registry.npmjs.org` since the proxy will not forward your NPM credentials.

## Web3-native NPM Registry

Every **npm** repo published on Valist is accessible via the `/api/npm` endpoint. This will resolve packages in a completely web3-native way, versus simply caching packages from NPM onto your local IPFS node.

First, ensure your `valist daemon` is running:

```sh
valist daemon
```

To install an npm package, you'll need to link your organization to the Valist registry, then install:

```bash
echo @acme-co:registry=http://localhost:9000/api/npm >> .npmrc

npm i @acme-co/npm-example
```

This is necessary to separate the namespace from NPMjs.com's package names. This prevents supply-chain attacks where a malicious package is published at the same name as a commonly used package (i.e., react). Explicitly overriding the desired organization or package name allows developers to specify where the namespace is resolved.

In the near future, different endpoints will be provided for different networks. This enables federation of packages to specific networks, while still providing global resolution across each network. Web3 projects will be able to use packages from multiple ecosystems without additional overhead.
