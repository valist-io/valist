/* eslint-disable max-len */
import Markdown from '../Markdown';

export default function RepoReadme(): JSX.Element {
  const defaultContent = `
# Valist

A trustless universal package repository enabling you to digitally sign and distribute software via IPFS in just a few steps.


## Features

Valist supports building and distributing software from source code to the end user.

* Publish Binaries (software, firmware, you name it)
* Easy Reproducible Builds
* NPM Registry
* Docker Registry
* Multi-Factor Releases (multi-sig for publishing software)
* Native IPFS Support
* Multi-Chain Support (Ethereum-based only for now)
`;

  return (
    <div className="p-10">
      <Markdown markdown={defaultContent} />
    </div>
  );
}
