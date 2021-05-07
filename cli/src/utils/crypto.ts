import keytar from 'keytar';
import { randomBytes } from 'crypto';

const HDWalletProvider = require('@truffle/hdwallet-provider');

export const getSignerKey = async (): Promise<string> => {
  const key = process.env.VALIST_SIGNER_KEY || await keytar.getPassword('VALIST', 'SIGNER');
  if (!key) throw new Error('🔎 No key found!');
  return key;
};

export const getWeb3Provider = async (signer?: string): Promise<any> => {
  const key = signer || await getSignerKey();

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

// @TODO: Confirm if overwriting another key
export const createSignerKey = async (): Promise<string> => {
  const key = randomBytes(32).toString('hex');
  await keytar.setPassword('VALIST', 'SIGNER', key);

  const provider = await getWeb3Provider();

  return provider.addresses[0];
};
