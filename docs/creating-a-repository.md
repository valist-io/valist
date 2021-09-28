# Creating a Repository

After you have created an organization, you are ready to create a repository. Creating a repository allows you to mange, publish, and retrieve your software releases on Valist.

## CLI

To create a new repository on from the CLI, call `valist repo create` and pass it your `organization` and your `repository name`.

```bash
valist repo create [org-name] [repo-name]
```

## SDK

To create a new repository using the Valist SDK, create a new javascript file containing the following:

```javascript
  import Valist from '@valist/sdk';
  const HDWalletProvider = require('@truffle/hdwallet-provider');

  const key = "<key>";
  const orgName = "<orgName>";
  const repoName = "<repoName>";
  const metaData = {
      name: 'Awesome Project',
      description: 'The coolest project in the world',
      homepage: 'https://cool.project',
      repository: 'https://github.com/',
      projectType: 'binary',
    };

  const web3Provider = new HDWalletProvider({
    privateKeys: [key],
    providerOrUrl: 'https://rpc.valist.io',
  });

  const valist = new Valist({ web3Provider });

  (async () => {
    const { transactionHash } = await valist.createRepository(orgName, repoName, metaData);

    console.log(transactionHash);
  })();
```
