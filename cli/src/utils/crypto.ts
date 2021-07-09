import { randomBytes } from 'crypto';
import HDWalletProvider from '@truffle/hdwallet-provider';
import { MissingKeyError } from './errors';

let keytar: any;
if (process.env.CI) {
  if (!process.env.VALIST_SIGNER_KEY) {
    throw new Error('VALIST_SIGNER_KEY needed in CI environment');
  }
} else {
  // eslint-disable-next-line global-require
  keytar = require('keytar');
}

export const getSignerKey = async (): Promise<string | null> => {
  const key = process.env.VALIST_SIGNER_KEY || await keytar.getPassword('VALIST', 'SIGNER');
  return key;
};

export const getWeb3Provider = async (signer?: string): Promise<any> => {
  const key = signer || await getSignerKey();
  if (!key) throw new MissingKeyError();

  const web3Provider = new HDWalletProvider({
    privateKeys: [key],
    providerOrUrl: process.env.WEB3_PROVIDER || 'https://rpc.valist.io',
  });

  return web3Provider;
};

export const getSignerAddress = async (): Promise<string> => {
  const provider = await getWeb3Provider();
  return provider.addresses[0];
};

export const createSignerKey = async (): Promise<void> => {
  const key = randomBytes(32).toString('hex');
  await keytar.setPassword('VALIST', 'SIGNER', key);
};
