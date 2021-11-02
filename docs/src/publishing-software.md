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

## Websites

Any static websites can be published and hosted through Valist.

Start by creating a repository for your web project.

```bash
valist create acme-co/website-example
```

Next, add your website files to the `valist.yml`.

```yaml
name: acme-co/website-example
tag: 0.0.1
artifacts:
  web/index: index.html
  logo: logo.png
```

To publish a new release, update the `tag` in your `valist.yml`, and run the following from your project root.

```bash
valist publish
```

## NFTs

Valist can be used to publish and pin metadata, images or any other web3 assets required for minting an nft.

Start by creating a repository for your NFT project.

```bash
valist create acme-co/nft-example
```

Next, add your nft assets to the `valist.yml`.

```yaml
name: acme-co/nft-example
tag: 0.0.1
artifacts:
  meta: data/meta.json
  image: data/ape.png
```

To publish a new release, update the `tag` in your `valist.yml`, and run the following from your project root.

```bash
valist publish
```

> **Note:** In a future upgrade to the valist system, all Valist repos will implement the `ERC-721` interface to allow repositories to be managed, owned, and traded. If you are interested in contributing or following the status of these features, checkout the current implementation and planning doc on [hackmd](https://hackmd.io/YF5CsRv_QZWk7o7ZzgRxDg?both)
