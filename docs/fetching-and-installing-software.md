# Fetching & Installing Software

## Web

To download a release artifact from the Web UI, navigate to the target repository's profile page (https://app.valist.io/`<orgName>`/`<repoName>`), and then choose your desired release from the release list.

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

Check out our [example repo here](https://github.com/valist-io/example-projects/tree/main/cli-publish-binary) for more!

## NPM Registry

Every **node** repo type published on Valist is accessible via the relay API at [https://valist.io/api/npm](https://valist.io/api/npm).

To install a package directly from a repository simply append the `npm --registry` flag:

```bash
npm install <examplePackage> --registry=https://valist.io/api/npm
```

Or set a Valist relay as your default **NPM registry**:

```bash
npm config set registry https://valist.io/api/npm
```
