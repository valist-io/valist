# Access Control

## Roles

### Organization Admin

An organization admin is the default role for managing an organization on valist. They are responsible for adding new developers to a repository as well as keeping the repository metadata up to date.

### Repository Developer

A repository developer is the default role for managing the releases under a repository. Together repository developers are able to vote on the next release of a piece of software.

## Key Management

### CLI

The valist CLI leverages the [keytar library](https://www.npmjs.com/package/keytar) for secure storage of keys in the devices keystore.

### Web

In browser, the valist relay frontend interacts with `MetaMask` for key management.
