import Web3 from 'web3';
import ValistABI from './abis/contracts/Valist.sol/Valist.json';
import ValistRegistryABI from './abis/contracts/ValistRegistry.sol/ValistRegistry.json';

import { sendMetaTx } from './helpers/biconomy';

import { InvalidNetworkError } from './errors';

export const shortnameFilterRegex = /[^A-z0-9-]/;

export const getContractInstance = (web3: Web3, abi: any, address: string) => new web3.eth.Contract(abi, address);

export const getValistContract = async (web3: Web3, chainID: number, address?: string) => {
  const networkContractMap = {
    80001: '0x092039fCdc6a18Cd4e76261F30Fc88FAeC035E40',
  };

  if (!address && !Object.keys(networkContractMap).includes(chainID.toString())) {
    throw new InvalidNetworkError('Valist not found on network');
  }

  const deployedAddress: string = address || networkContractMap[80001];
  return getContractInstance(web3, ValistABI.abi, deployedAddress);
};

export const getValistRegistry = async (web3: Web3, chainID: number, address?: string) => {
  const networkContractMap = {
    80001: '0x092039fCdc6a18Cd4e76261F30Fc88FAeC035E40',
  };

  if (!address && !Object.keys(networkContractMap).includes(chainID.toString())) {
    throw new InvalidNetworkError('Valist Registry not found on network');
  }

  const deployedAddress: string = address || networkContractMap[80001];
  return getContractInstance(web3, ValistRegistryABI.abi, deployedAddress);
};

export const sendMetaTransaction = sendMetaTx;

export const Web3Providers = Web3.providers;
