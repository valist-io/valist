# Publishing Software

## Binary Executables

### CLI

To publish a binary from the Valist CLI create a new `valist.yml` within your project folder containing one of the supported binary project types `binary` or `go`. Or, run `valist init` to generate a new `valist.yaml`.

```yaml
name: acme-co/go-example
tag: 0.1.6-rc.0
build: go build -o ./dist/hello src/main.go 
artifacts:
  hello: dist/hello
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
