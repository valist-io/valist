## Why Valist

Valist is a web3-native software distribution system, with built-in package registry support. It's powered by smart contracts (on any EVM), IPFS, and Filecoin. It's open to developers to extend and integrate into almost any system, traditional or decentralized. No need for expensive and complex PKIs or manual code signing processes!

You can point popular package managers at a Valist relay, which will ensure the integrity of the packages and act as a universal cache. This means that as you develop software, your dependencies will be stored on your local IPFS node and made available to other peers you're connected to.

This can save on build times, bandwidth, and prevents downtime if upstream sources are temporarily unavailable.

You can think of Valist as a trustless [Artifactory](https://jfrog.com/artifactory/), or a universal [Verdaccio](https://verdaccio.org/), with far more powerful access control and data integrity features. This includes multi-factor releases (M of N keys need to sign on some firmware before release), as well as the ability to use any hardware wallet to sign code.

Valist is designed to be highly extensible and interoperable with many package managers that developers are familiar with, such as NPM, Pip, Docker, APT, and Cargo.

Currently, Valist supports:

* Executable binaries

* NPM packages

* Docker images

## Architecture

Valist provides a simple web frontend and HTTP relay that can be deployed locally or on a server, with a shared Valist-SDK library used for the CLI (and other clients). It's CI/CD friendly, enabling automatic publishing of cryptographically verifiable software.

The Valist-SDK reference implementation is currently written in TypeScript, using minimal dependencies.

The following diagram is a visualization of the web implementation:

![Valist Architecture](img/current-implementation.png){width=300px}
