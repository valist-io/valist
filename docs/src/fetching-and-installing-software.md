# Fetching & Installing Software

## CLI

To install a binary artifact, run the following:

```bash
valist install [org-name]/[repo-name]
```

The CLI will detect your current platform and install the correct artifact into your `~/.valist/bin` directory.

You can also install packages at specific versions using:

```sh
valist install [org-name]/[repo-name]/[version-tag]
```

Finally, ensure your PATH includes the ~/.valist/bin folder:

```sh
export PATH="$PATH:$HOME/.valist/bin"
echo PATH=\"\$PATH:$HOME/.valist/bin\" >> ~/.zshrc # or .bashrc
```

## Web

To download a release artifact from the Web UI, navigate to the target repository's profile page (https://app.valist.io/`<orgName>`/`<repoName>`), and click `versions`, then choose your desired release from the release list.

![valist-release-page](./img/valist-release-page.png){width="600px"}

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
