const HDWalletProvider = require('@truffle/hdwallet-provider');

const mnemonic = process.env.MNEMONIC;
const key = process.env.KEY;

export const web3Provider = new HDWalletProvider({
  mnemonic: { phrase: mnemonic },
  privateKeys: [key],
  providerOrUrl: 'https://rpc-mumbai.matic.today',
});

export const getWeb3Provider = () => web3Provider;
