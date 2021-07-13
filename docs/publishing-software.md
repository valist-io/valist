# Publishing Software

## Binary Executables

### CLI

### Web

### SDK

```bash
const Valist = require('@valist/sdk');

(async () => {
  const valist = new Valist({ web3Provider: YOUR_WEB3_PROVIDER, metaTx: false });
  await valist.connect();

  const releases = await valist.getReleasesFromRepo('valist', 'sdk');

  console.log(releases);
})();
```

## NPM Packages

### CLI

### Web

### SDK

## Python Packages

### CLI

### Web

### SDK

## Docker Images

### CLI

### Web

### SDK