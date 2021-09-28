# Fetching & Installing Software

## Web

To download a release artifact from the Web UI, navigate to the target repository's profile page (https://app.valist.io/`<orgName>`/`<repoName>`), and click `versions`, then choose your desired release from the release list.

![valist-release-page](img/valist-release-page.png){width="600px"}

## SDK

Artifacts can be downloaded using the Valist SDK by filling in a web3 provider and running the following code:

```javascript
const Valist = require('@valist/sdk');

(async () => {
  const valist = new Valist({ web3Provider: YOUR_WEB3_PROVIDER, metaTx: false });
  await valist.connect();

  const releases = await valist.getReleases('valist', 'sdk');

  const latest = await valist.getLatestRelease('valist', 'sdk');

  console.log(releases);
})();
```

[Example Node SDK Project](https://github.com/valist-io/example-projects/tree/main/sdk-node)

## NPM Registry

Every **npm** repo type published on Valist is accessible via the relay API at [https://valist.io/api/npm](https://valist.io/api/npm).

To install an npm package, you'll need to link your organization to the Valist registry, then install!

```bash
echo @acme-co:registry=https://valist.io/api/npm >> .npmrc

npm i @acme-co/npm-example
```

Or set a Valist relay as your default **NPM registry**:

```bash
npm config set registry https://valist.io/api/npm
```
