import keytar from 'keytar';
import { randomBytes } from 'crypto';

const HDWalletProvider = require('@truffle/hdwallet-provider');

export const getSignerKey = async () => {
  const key = process.env.VALIST_SIGNER_KEY || await keytar.getPassword('VALIST', 'SIGNER');
  if (!key) throw new Error('No key found!');
  return key;
};

export const getWeb3Provider = async () => {
  const key = await getSignerKey();

  const web3Provider = new HDWalletProvider({
    privateKeys: [key],
    providerOrUrl: process.env.WEB3_PROVIDER || 'https://rpc-mumbai.matic.today',
  });

  return web3Provider;
};

export const createSignerKey = async (): Promise<string> => {
  const key = randomBytes(32).toString('hex');
  await keytar.setPassword('VALIST', 'SIGNER', key);

  const provider = await getWeb3Provider();

  return provider.addresses[0];
};
