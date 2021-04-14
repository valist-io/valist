#!/usr/bin/env node
import * as yargs from 'yargs';
import * as fs from 'fs';
import Valist from 'valist';
import { getWeb3Provider, createSignerKey, getSignerKey } from './utils/crypto';
import { npmPack } from './utils/npm';
import { parseValistConfig } from './utils/config';

const initValist = async () => {
  try {
    const valist = new Valist({ web3Provider: await getWeb3Provider() });

    const waitForMetaTx: boolean = true;

    await valist.connect(waitForMetaTx);

    valist.signer = await getSignerKey();

    console.log('Account:', valist.defaultAccount);

    return valist;
  } catch (e) {
    const msg = 'Could not connect to Valist';
    console.error(msg, e);
    throw e;
  }
};

yargs.command('create signer', 'Create a new signer key', () => {}, async () => {
  console.log('Generating new signer key...');
  const address = await createSignerKey();
  console.log('Successfully stored in keychain/keyring');
  console.log('Your new signer address is', address);
  process.exit(0);
});

yargs.command('publish', 'Publish an NPM package to Valist', () => {}, async () => {
  console.log('Connecting to Valist...');
  const valist = await initValist();
  console.log('Connected');

  const {
    project,
    org,
    tag,
    meta,
    type,
  }: {
    project: string,
    org: string,
    tag: string,
    meta: string,
    type: string,
  } = parseValistConfig();

  if (type === 'npm') {
    console.log('Packing NPM Package');
    const tarballName = await npmPack();
    console.log('Packed:', tarballName);

    const releaseFile = fs.createReadStream(tarballName);
    const metaFile = fs.createReadStream(meta);

    console.log('Preparing release on IPFS');
    const releaseObject = await valist.prepareRelease(tag, releaseFile, metaFile);
    console.log('Release Object:', releaseObject);

    console.log('Publishing Release to Valist');
    const { transactionHash } = await valist.publishRelease(org, project, releaseObject);

    console.log(`Successfully Released ${project} ${tag}!`);
    console.log('Transaction Hash:', transactionHash);
  }
  process.exit(0);
});

yargs.parse();
