# Getting Started

## Installation

### CLI

The Valist-CLI can be globally installed by running the following commands:

```bash
# Or, run the following (based on your os/arch):
mkdir -p ~/.local/bin
curl https://gateway.valist.io/ipfs/QmebByjWtVazSsyvhmFQQGpA5aeNYkbFKFWkZPi3Sqtmw6/linux-amd64/valist --output ~/.local/bin/

# Add valist to your $PATH
export PATH="$PATH:$HOME/.local/bin"
echo PATH="$PATH:$HOME/.local/bin" >> ~/.zshrc
```

### SDK

The Valist JS-SDK can be installed by running the following command:

```bash
npm install @valist/sdk --registry=https://valist.io/api/npm
```

## Why Valist

Valist is a software/firmware/binary data notary system, similar to the concept that Apple uses to digitally sign and secure applications, but open to developers to extend and integrate into almost any system, traditional or decentralized. No need for expensive and centralized PKIs or manual code signing processes!

The goal is to point **any** software distribution system at a Valist relay, which will ensure the integrity of the packages and act as a universal cache.

You can think of Valist as a trustless [Artifactory](https://jfrog.com/artifactory/), or a universal [Verdaccio](https://verdaccio.org/), but with far more powerful access control and data integrity features. This includes multi-factor releases (M of N keys need to sign on some firmware before release), as well as the ability to use any hardware wallet to sign code.

Valist is designed to be highly extensible and interoperable with many package managers that developers are familiar with, such as NPM, Pip, Docker, APT, and Cargo.

Currently, Valist supports:

* Executable Binaries, with automatic code-signing

* NPM packages

* Python packages

* Docker images

## Architecture

Valist provides a simple web frontend and HTTP relay that can be deployed locally or on a server, with a shared Valist-SDK library used for the CLI (and other clients). The intent is to be CI/CD friendly, enabling automatic publishing of cryptographically verifiable software.

The Valist-SDK reference implementation is currently written in TypeScript, using minimal dependencies and is leveraged to build a REST API layer to ensure backwards compatibility with traditional package managers and HTTP compatibility.

The following diagram is a visualization of the current implementation:

![Valist Architecture](img/current-implementation.png){width=300px}
