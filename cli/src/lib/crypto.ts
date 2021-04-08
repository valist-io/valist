const HDWalletProvider = require('@truffle/hdwallet-provider');

// const mnemonic = process.env.MNEMONIC;
const key = process.env.VALIST_SIGNING_KEY;

if (/* !mnemonic && */ !key) throw new Error('No key found!');

export const web3Provider = new HDWalletProvider({
//  mnemonic: { phrase: mnemonic },
  privateKeys: [key],
  providerOrUrl: process.env.WEB3_PROVIDER || 'https://rpc-mumbai.matic.today',
});

export const getWeb3Provider = () => web3Provider;
