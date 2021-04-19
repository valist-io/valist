#!/usr/bin/env node
import * as yargs from 'yargs';
import * as fs from 'fs';
import * as path from 'path';
import { initValist, parseValistConfig } from './utils/config';
import { createSignerKey } from './utils/crypto';
import { npmPack } from './utils/npm';

yargs.command('create signer', 'Create a new signer key', () => {}, async () => {
  console.log('🛠 Generating new signer key...');
  const address = await createSignerKey();
  console.log('🔒 Successfully stored in keychain/keyring');
  console.log('📇 Your new signer address is:', address);
  process.exit(0);
});

yargs.command('publish', 'Publish package to Valist', () => {}, async () => {
  console.log('📡 Connecting to Valist...');
  const valist = await initValist();
  console.log('⚡️ Connected!');

  const {
    project,
    org,
    tag,
    meta,
    type,
    artifact,
  }: {
    project: string,
    org: string,
    tag: string,
    meta: string,
    type: string,
    artifact?: string,
  } = parseValistConfig();

  let releaseFile: fs.ReadStream;
  let metaFile: fs.ReadStream;

  if (type === 'npm') {
    console.log('🛠  Packing NPM Package...');
    const tarballName = await npmPack();
    console.log('💼 Packed:', tarballName);
    releaseFile = fs.createReadStream(path.join(process.cwd(), tarballName));
    metaFile = fs.createReadStream(path.join(process.cwd(), meta));
  } else if (type === 'binary') {
    if (!artifact) {
      console.error('👻 No build artifact found!');
      process.exit(1);
    }

    releaseFile = fs.createReadStream(path.join(process.cwd(), artifact));
    metaFile = fs.createReadStream(path.join(process.cwd(), meta));
  } else {
    console.error('🚧 Project type not supported!');
    process.exit(1);
  }

  console.log('🪐 Preparing release on IPFS...');
  const releaseObject = await valist.prepareRelease(tag, releaseFile, metaFile);
  console.log('📦 Release Object:', releaseObject);

  console.log('⚡️ Publishing Release to Valist...');
  const { transactionHash } = await valist.publishRelease(org, project, releaseObject);

  console.log(`✅ Successfully Released ${project} ${tag}!`);
  console.log('📖 IPFS address of release:', `ipfs://${releaseObject.releaseCID}`);
  console.log('🔗 Transaction Hash:', transactionHash);

  // cleanup generated tarball/build artifact
  if (type === 'npm') {
    fs.unlinkSync(releaseFile.path);
  }

  process.exit(0);
});

yargs.parse();
