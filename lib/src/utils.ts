import Web3 from 'web3';
import ValistABI from './abis/contracts/Valist.sol/Valist.json';

import { InvalidNetworkError } from './errors';

export const shortnameFilterRegex = /[^A-z0-9-]/;

export const getContractInstance = (web3: Web3, abi: any, address: string) => new web3.eth.Contract(abi, address);

export const getValistContract = async (web3: Web3, address?: string) => {
  const networkContractMap = {
    80001: '0xB068CB8e72186Af3D7Be8ba97B16DC5E4Ce4C538',
  };
  // get network ID to fetch deployed address
  const networkId: number = await web3.eth.net.getId();

  if (!address && !Object.keys(networkContractMap).includes(networkId.toString())) {
    throw new InvalidNetworkError('Valist not found on network');
  }

  const deployedAddress: string = address || networkContractMap[80001];
  return getContractInstance(web3, ValistABI.abi, deployedAddress);
};

export const Web3Providers = Web3.providers;
