# Publishing Software

## Binary Executables

### CLI

To publish a binary from the Valist-CLI create a new `valist.yml` within your project folder containing one of supported binary project types `binary` or `go`. Or run a `valist init` to generate a new `valist.yaml`.

```yaml
type: go
org: test
repo: testProject
tag: 0.0.1
out: dist
```

```bash
valist publish
```

### SDK

To publish a binary from the Valist-SDK create and run the following javascript:

```javascript
  import Valist from '@valist/sdk';
  const HDWalletProvider = require('@truffle/hdwallet-provider');

  const key = "<key>";
  const orgName = "<orgName";
  const repoName = "<repoName>";
  const metaData = "<orgMeta>";

  const web3Provider = new HDWalletProvider({
    privateKeys: [key],
    providerOrUrl: 'https://rpc.valist.io',
  });

  const valist = new Valist({web3Provider});

  (async () => {
    const { transactionHash } = await valist.publishRelease(orgName, repoName, releaseObject);

    console.log(transactionHash);
  })();
```

## NPM Packages

### CLI

To publish an NPM package from the Valist-CLI create a new `valist.yml` within your project with a project type of `node`. Or run a `valist init` to generate a new `valist.yaml`.

```yaml
type: node
org: test
repo: testProject
tag: 0.0.1
out: dist
```

```bash
valist publish
```

### SDK

To publish a binary from the Valist-SDK create and run the following javascript:

```javascript
  import Valist from '@valist/sdk';
  const HDWalletProvider = require('@truffle/hdwallet-provider');

  const key = "<key>";
  const orgName = "<orgName";
  const repoName = "<repoName>";
  const metaData = "<orgMeta>";

  const web3Provider = new HDWalletProvider({
    privateKeys: [key],
    providerOrUrl: 'https://rpc.valist.io',
  });

  const valist = new Valist({web3Provider});

  (async () => {
    const { transactionHash } = await valist.publishRelease(orgName, repoName, releaseObject);

    console.log(transactionHash);
  })();
```

## Python Packages

### CLI

To publish a python package from the Valist-CLI create a new `valist.yml` within your project with a project type of `python`. Or run a `valist init` to generate a new `valist.yaml`.

```yaml
type: python
org: test
repo: testProject
tag: 0.0.1
out: dist
```

```bash
valist publish
```

### SDK

To publish a binary from the Valist-SDK create and run the following javascript:

```javascript
  import Valist from '@valist/sdk';
  const HDWalletProvider = require('@truffle/hdwallet-provider');

  const key = "<key>";
  const orgName = "<orgName";
  const repoName = "<repoName>";
  const metaData = "<orgMeta>";

  const web3Provider = new HDWalletProvider({
    privateKeys: [key],
    providerOrUrl: 'https://rpc.valist.io',
  });

  const valist = new Valist({web3Provider});

  (async () => {
    const { transactionHash } = await valist.publishRelease(orgName, repoName, releaseObject);

    console.log(transactionHash);
  })();
```

## Docker Images

### CLI

To publish a docker image from the Valist-CLI create a new `valist.yml` within your project with a project type of `docker`. Or run a `valist init` to generate a new `valist.yaml`.

```yaml
type: docker
org: test
repo: testProject
tag: 0.0.1
out: dist
```

```bash
valist publish
```

### SDK

To publish a binary from the Valist-SDK create and run the following javascript:

```javascript
  import Valist from '@valist/sdk';
  const HDWalletProvider = require('@truffle/hdwallet-provider');

  const key = "<key>";
  const orgName = "<orgName";
  const repoName = "<repoName>";
  const metaData = "<orgMeta>";

  const web3Provider = new HDWalletProvider({
    privateKeys: [key],
    providerOrUrl: 'https://rpc.valist.io',
  });

  const valist = new Valist({web3Provider});

  (async () => {
    const { transactionHash } = await valist.publishRelease(orgName, repoName, releaseObject);

    console.log(transactionHash);
  })();
```