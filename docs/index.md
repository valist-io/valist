# Introduction

## Overview

Valist is a software/firmware/binary data notary system, similar to the concept that Apple uses to digitally sign and secure applications, but open to developers to extend and integrate into almost any system, traditional or decentralized.

### Signed Binaries, Firmware, and universal package distribution

Valist is designed to be highly extensible and interoperable with many package managers that developers are familiar with, such as NPM, Pip, Docker, APT, and Cargo.

Currently, the Valist API supports:

* Arbitrary Binaries, with automatic code-signing

* NPM packages

* Pip packages

* Docker images

To achieve this, Valist uses the following semantics for organizing releases:

https://`relay`/`organization`/`project`/`tag`

This allows for arbitrary versioning, and better compatibility with the majority of package managers that exist today.

## Motivation

Secure (and simplified) software updating is a common problem within IoT, traditional systems, and many cases now, dApps. Typically, it is necessary to roll your own upgrading solution, or be stuck with a centralized app store acceptance and delivery process. The former is ineffective, as constant re-implementation of a process that should be as secure as possible dramatically increases risk, while the latter is ineffective since you are tied to a central entity that manages your distribution on your behalf (i.e., requires permission).

The idea is to leverage Ethereum, IPFS and/or Filecoin to create a public "base" layer for a simplified binary repository that both integrates with traditional systems and is built upon decentralized protocols. Smart contracts on Ethereum manage the latest source of truth for binary data stored in another layer such as IPFS and/or Filecoin. Clients can then query the software notary for the latest version of some software and be pointed to a verifiable, decentralized store.

Imagine the following scenario:

* A developer wants to distribute a new firmware version for a hardware wallet (or some other arbitrary software) they have been building in a secure, verifiable way.

* Using a simple frontend, the developer registers their credentials (one or more public keys, perhaps leveraging ERC-725) to the software notary dApp. This can also be organization-level credentials, with individual developer/team access control.

* The developer then signs the firmware with a private key associated with the public identity.

* The developer uploads the firmware to the binary store (IPFS/Filecoin) using the simple frontend, and can set any relevant metadata such as version number, update notes, etc.

* The registry (on Ethereum, and potentially other blockchains in the future) is updated with the latest verified version.

    Clients with the software installed automatically detect the change and proceed to notify and/or trigger an auto-update.

Here's a visual of what this flow looks like:

![Signed release](img/signed-release-flow.png)

## Architecture

To start, we are providing a simple web frontend and HTTP relay that can be deployed locally or on a server, and a shared library that will be used in a future CLI (and other clients). The intent is to be CI/CD friendly, enabling automatic publishing of cryptographically verifiable software.

The Valist core lib reference implementation is currently written in TypeScript, using minimal dependencies.

The core lib is used to build a REST API layer to ensure backwards compatibility with traditional package managers and HTTP compatibility.

The following diagram is a visualization of the current implementation:

![Valist Architecture](img/architecture.svg)
