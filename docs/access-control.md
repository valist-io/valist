# Access Control

## Roles

### Organization Admin

An organization admin is the default role for managing an organization on Valist. They are responsible for adding new developers to a repository as well as keeping the repository metadata up to date.

### Repository Developer

A repository developer is the default role for managing the releases under a repository. Together repository developers are able to vote on the next release of a piece of software.

## Key Management

### CLI

The Valist CLI leverages the [keytar](https://www.npmjs.com/package/keytar) library for secure, cross-platform storage of keys in the device's keystore.

### Web

In browser, the Valist relay frontend interacts with `MetaMask`, `WalletConnect` and/or `Magic` for key management.
