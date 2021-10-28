# Publishing Software

## Binary

> [Example Binary Project](https://github.com/valist-io/example-projects/tree/main/cli-publish-go-project)

Start by creating a `valist.yml` in your project root. The organization and repository will be created if they do not exist.

```bash
valist init acme-co/go-example
```

Add your release artifacts to the `valist.yml`. Any README or LICENSE files will be included automatically.

```yaml
name: acme-co/go-example
tag: 0.1.6-rc.0
artifacts:
  main: out/main
```

To publish a new release, update the `tag` in your `valist.yml`, and run the following from your project root.

```bash
valist publish
```

### Platforms

If your project has artifacts for multiple platforms, add an entry for each supported platform to your `valist.yml`. 

```yaml
artifacts:
  linux/amd64: path_to_bin
  linux/arm64: path_to_bin
  darwin/amd64: path_to_bin
  darwin/arm64: path_to_bin
  windows/amd64: path_to_bin
```

## NPM

> [Example NPM Project](https://github.com/valist-io/example-projects/tree/main/cli-publish-npm-package)

Start by creating a repository for your NPM project.

```bash
valist create acme-co/npm-example
```

Make sure the `name` in your `package.json` matches your repository name.

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

To publish a new release, update the `version` in your `package.json`, and run the following from your project root.

```bash
# In a new terminal window run the valist daemon
valist daemon

# A prompt will appear in the daemon terminal window
npm publish --registry=http://localhost:9000/api/npm
```

## Git

> Git support is experimental.

Start by creating a repository for your Git project.

```bash
valist create acme-co/git-example
```

To publish a new release, create a Git tag, and run the following from your project root.

```bash
# In a new terminal window run the valist daemon
valist daemon

# The git tag will be your release tag
git tag -a v0.0.1 -m "release v0.0.1"

# A prompt will appear in the daemon terminal window
git push http://localhost:9000/api/git/acme-co/git-example tags/v0.0.1
```

## Docker

> Docker support is experimental.

Start by creating a repository for your Docker project.

```bash
valist create acme-co/docker-example
```

Add an entry to `/etc/hosts` for your local valist node.

```text
127.0.0.1 valist.local
```

To publish a new release, create a Docker tag, and run the following from your project root.

```bash
# In a new terminal window run the valist daemon
valist daemon

# The docker tag must match your valist host entry
docker build --tag valist.local:9000/acme-co/docker-example .

# A prompt will appear in the daemon terminal window
docker push valist.local:9000/acme-co/docker-example
```

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
