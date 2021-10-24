# Creating an Organization

Creating an organization is the first step to publishing software on Valist! Organizations are how users are able to manage their various software repositories and admin credentials on Valist. Each organization receives a unique `valist-ID` which is then linked to a namespace that is governed by the members (keys) of the organization.

## CLI

To create an organization from the CLI it's as simple as running `valist org create` and passing it the name of the organization that you would like to create.

```bash
  valist org create [organization-name]
```

You will then be prompted to enter some additional metadata about the organization.

![cli-create-org-metadata](../img/cli-org-create-metadata.png)

## SDK

To create a new organization using the Valist SDK, create a new javascript file containing the following:

```javascript
  import Valist from '@valist/sdk';
  const HDWalletProvider = require('@truffle/hdwallet-provider');

  const key = "<key>"; // securely store this!
  const orgName = "<orgName>";
  const metaData = {
    name: 'Awesome Developer',
    description: 'Much talent',
  };

  const web3Provider = new HDWalletProvider({
    privateKeys: [key],
    providerOrUrl: 'https://rpc.valist.io',
  });

  const valist = new Valist({ web3Provider });

  (async () => {
    const { transactionHash } = await valist.createOrganization(orgName, metaData);

    console.log(transactionHash);
  })();
```
