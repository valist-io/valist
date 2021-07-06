import * as keytar from 'keytar';
import { randomBytes } from 'crypto';
import { MissingKeyError } from './errors';

const HDWalletProvider = require('@truffle/hdwallet-provider');

export const getSignerKey = async (): Promise<string | null> => {
  const key = process.env.VALIST_SIGNER_KEY || await keytar.getPassword('VALIST', 'SIGNER');
  return key;
};

export const getWeb3Provider = async (signer?: string): Promise<any> => {
  const key = signer || await getSignerKey();
  if (!key) throw new MissingKeyError();

  const web3Provider = new HDWalletProvider({
    privateKeys: [key],
    providerOrUrl: process.env.WEB3_PROVIDER || 'https://matic-mumbai.chainstacklabs.com',
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
