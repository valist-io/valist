#!/usr/bin/env node
import * as yargs from 'yargs';
import * as fs from 'fs';
import Valist from 'valist';
import { getWeb3Provider } from './lib/crypto';
import { npmPack } from './lib/npm';
import { parseValistConfig } from './lib/config';

const initValist = async () => {
  try {
    const valist = new Valist({ web3Provider: getWeb3Provider(), metaTx: false });

    const waitForMetaTx: boolean = false;

    await valist.connect(waitForMetaTx);

    console.log('Account:', valist.defaultAccount);

    return valist;
  } catch (e) {
    const msg = 'Could not connect to Valist';
    console.error(msg, e);
    throw e;
  }
};

yargs.command('start', 'Starts the Valist service', () => {}, async (argv) => {
  // const valist = await initValist();
  console.log('Valist service started', argv);
});

yargs.command('publish', 'Publish an NPM package to Valist', () => {}, async () => {
  console.log('Connecting to Valist...');
  const valist = await initValist();
  console.log('Connected');

  console.log('Fetching Valist Config...');
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

    console.log('Files', releaseFile, metaFile);

    console.log('Preparing release on IPFS');
    const releaseObject = await valist.prepareRelease(tag, releaseFile, metaFile);
    console.log('Release Object:', releaseObject);

    console.log('Publishing Release to Valist');
    const { transactionHash } = await valist.publishRelease(org, project, releaseObject);

    console.log(`Successfully Released ${project} ${tag}!`);
    console.log('Transaction Hash:', transactionHash);
  }
});

yargs.parse();
