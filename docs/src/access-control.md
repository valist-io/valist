# Access Control

## Roles

### Organization Admin

An organization admin is the default role for managing an organization on Valist. They are responsible for adding new developers to a repository as well as keeping the repository metadata up to date.

### Repository Developer

A repository developer is the default role for managing the releases under a repository. Together repository developers are able to vote on the next release of a piece of software.

## Key Management

### CLI

The Valist CLI leverages the [Web3 Storage Definition](https://github.com/ethereum/wiki/wiki/Web3-Secret-Storage-Definition) to store encrypted keystore files in `~/.valist/keystore`.

These keystore files are compatible with common Ethereum wallets like MetaMask, Geth/Clef, and MyCrypto.

### Web

In browser, the Valist relay frontend interacts with `MetaMask` or `WalletConnect` for key management.
