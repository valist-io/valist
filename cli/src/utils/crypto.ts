import { randomBytes } from 'crypto';
import { MissingKeyError } from './errors';

// eslint-disable-next-line @typescript-eslint/no-var-requires
const HDWalletProvider = require('@truffle/hdwallet-provider');

let keytar: any;
if (process.env.CI) {
  if (!process.env.VALIST_SIGNER_KEY) {
    throw new Error('VALIST_SIGNER_KEY needed in CI environment');
  }
} else {
  // eslint-disable-next-line global-require
  keytar = require('keytar');
}

export const getSignerKey = async (): Promise<string> => {
  const key = process.env.VALIST_SIGNER_KEY || await keytar.getPassword('VALIST', 'SIGNER');
  if (!key) throw new MissingKeyError();
  return key;
};

export const getWeb3Provider = async (signer?: string): Promise<any> => {
  const key = signer || await getSignerKey();

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
