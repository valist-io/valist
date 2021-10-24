# Publishing Software

## Binary Executables

### CLI

To publish a binary from the Valist CLI create a new `valist.yml` within your project folder containing one of the supported binary project types `binary` or `go`. Or, run `valist init` to generate a new `valist.yaml`.

```yaml
type: go
org: acme-co
repo: go-example
tag: 0.1.6-rc.0
meta: README.md
build: go build -o ./dist/hello src/main.go 
out: dist/hello
```

```bash
valist publish
```

[Example Go Project](https://github.com/valist-io/example-projects/tree/main/cli-publish-go-project)

## NPM Packages

### CLI

To publish an NPM package from the Valist CLI, create a new `package.json` with the name structured as `@orgName/repoName` and the `version` as the current version to be published.


```json
{
  "name": "@acme-co/npm-example",
  "version": "0.0.1-rc.0",
  "description": "",
  "main": "dist/index.js",
  "scripts": {
    "build": "tsc",
    "test": "echo \"Error: no test specified\" && exit 1"
  },
  "author": "ACME",
  "license": "MIT",
  "dependencies": {
    "typescript": "^4.3.4"
  }
}
```

In a new terminal window start the valist daemon

```bash
valist daemon
```

Now just publish to the registry using the npm publish command and the registry flag!

```bash
npm publish --registry=http://localhost:9000/api/npm
```

[Example NPM Project](https://github.com/valist-io/example-projects/tree/main/cli-publish-npm-package)

<!-- ### SDK

To publish a binary from the Valist SDK, create and run the following javascript:

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
``` -->

<!-- ### SDK

To publish a binary from the Valist SDK create and run the following javascript:

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
``` -->
<!--
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
``` -->
