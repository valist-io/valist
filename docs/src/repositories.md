# Repositories

After you have created an organization, you are ready to create a repository. Creating a repository allows you to mange, publish, and retrieve your software releases on Valist.

If you have not created an organization yet, don't worry -- creating a repository will automatically prompt you to create the parent organization if it doesn't exist.

## CLI

To create a new repository on from the CLI, call `valist create` and pass it your `organization` and `repository name`.

```bash
valist create [org-name]/[repo-name]
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
