# Getting Started

## Fetching a Binary from the Valist relay

To fetch a release from a project, you can run the following `curl` command:

```bash
curl -L -o example-bin https://app.valist.io/api/exampleorg/binary/latest
```

Once downloaded, you can set the binary as executable:

```bash
chmod +x example-bin
```

You can also download the binary by its release version/tag using by changing `latest` to the `tag` you wish to request:

```bash
curl -L -o example-bin https://app.valist.io/api/exampleorg/binary/0.0.1
```

## Valist for NodeJS Projects

To add a package stored on Valist to your Node project, run the following command:

```bash
npm install https://app.valist.io/api/exampleorg/npm/latest
```

This will install the NPM release from an IPFS relay through the valist.io API.

A basic implementation of the NPM registry API is available. You can use it by setting your registry to:

```bash
npm config set registry https://valist.io/api/npm
```

Then, you can install your package by using the NPM `@organization`/`package` format:

```bash
npm install @exampleorg/npm
```

This is generated entirely from the Valist smart contracts, and are enforced using the Access Control you can configure using the [valist dashboard](https://app.valist.io), or using the core lib SDK.

> Note: Switching to the Valist relay at this time will fetch directly from the smart contracts only. In the future, Valist <> npmjs.com proxy support will allow you to cache from upstream and use Valist as a universal cache.

Next, you can import the `Valist` SDK and pass in a `Web3Provider`. Valist exposes a set of `Web3Providers`, including HTTP:

```typescript
import Valist from 'valist';
import { Web3Providers } from 'valist/dist/utils';

const web3Provider = new Web3Providers.HttpProvider(process.env.WEB3_PROVIDER);

const valist = new Valist({ web3Provider });

await valist.connect();

valist.defaultAccount;

```

Here, Valist is exposing the `Web3Providers` object which allows us to pass in a provider from an environment variable.

Then, we initialize the Valist object using `await valist.connect()`, and check the default account.

## Using the Valist SDK in a React App

The Valist SDK can be imported into any React 16 and above project. If you are using NextJS you would mostly likely perform the following within the `_app.tsx` entry-point.

```typescript
import Valist from 'valist';
```

The `valist` object can be mapped to the component state for ease of use.

```typescript
const [valist, setValist] = useState<Valist>();
```

To instantiate a new `valist object` from a `web3 provider` you will need to do so within an `async` function within a react `useEffect`.

```typescript
  useEffect(() => {
    // self executing async function enables async support within useEffect
    (async function () {

    })();

  }, []);
```

Inside the `useEffect` you can then setup a `try catch` with any custom web3 error handling. You will also want to check if `window.ethereum` is defined if you are using an extension like [MetaMask](https://metamask.io).

```typescript
  useEffect(() => {
    (async function () {
        try {
          if (window.ethereum) {

          }
        } catch (error) {
            console.log(error);
        }
    })();
  }, []);
```

Now that you've defined basic asynchronous error handing for your application to communicate with `Ethereum` and `IPFS`. If your application is using `web3` through `MetaMask` you can enable your web3 provider by calling `window.ethereum.enable()`. A new valist instance can then be created with `new Valist(window.ethereum, true)` and the enabled web3 provider can be passed in from `window.ethereum`.

After the valist object is instantiated the initial wallet connection can be established by awaiting the `valist.connect()` method. Once the wallet connection is established you can then set the valist state value using your setValist.

```typescript
  // initialize web3 and valist object on document load (this effect is only triggered once)
  useEffect(() => {
    (async function () {
        try {
          if (window.ethereum) {
            window.ethereum.enable();
            let valist = new Valist({ web3Provider: window.ethereum });
            await valist.connect();
            setValist(valist);
          }
        } catch (error) {
            console.log(error);
        }
    })();
  }, []);
```

### Using React Context

You can set the Valist object as a context in React, and import it from any component using `useContext`.

For example:

```typescript
import ValistContext from '../ValistContext/ValistContext';

const ExampleComponent:FunctionComponent<any> = () => {
    const valist = useContext(ValistContext);

    return (
            <div>valist.defaultAccount</div>
    )
}
```

The valist state object can now be passed into the `ValistContext` as the valist value prop.

```typescript
    <ValistContext.Provider value={valist}>
```

## Next.js API support

You can easily import Valist into a Next.js backend!

Here's an example for how the `getLatestReleaseFromRepo` API call is implementated in Valist using Next.js:

```typescript
import { NextApiRequest, NextApiResponse } from 'next'
import Valist from 'valist';
import { Web3Providers } from 'valist/dist/utils';

export default async function getLatestReleaseFromRepo(req: NextApiRequest, res: NextApiResponse) {

  // set .env.local to your local chain or set in production deployment
  if (process.env.WEB3_PROVIDER) {
    const web3Provider = new Web3Providers.HttpProvider(process.env.WEB3_PROVIDER);

    const valist = new Valist({ web3Provider });
    await valist.connect();

    const {
      query: { orgName, repoName },
    } = req;

    const latestRelease = await valist.getLatestReleaseFromRepo(orgName.toString(), repoName.toString());

    if (latestRelease) {
      //return res.status(200).json({latestRelease});
      return res.redirect(`https://cloudflare-ipfs.com/ipfs/${latestRelease}`);
    } else {
      return res.status(404).json({statusCode: 404, message: "No release found!"});
    }

  } else {
    return res.status(500).json({statusCode: 500, message: "No Web3 Provider!"});
  }
}
```
