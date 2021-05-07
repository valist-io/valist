import Web3 from 'web3';
import ValistABI from './abis/Valist.json';

import { InvalidNetworkError } from './errors';

export const shortnameFilterRegex = /[^A-z0-9-]/;

export const getContractInstance = (web3: Web3, abi: any, address: string) => new web3.eth.Contract(abi, address);

export const getValistContract = async (web3: Web3, address?: string) => {
  // get network ID and the deployed address
  const networkId: number = await web3.eth.net.getId();

  if (networkId !== 80001) {
    throw new InvalidNetworkError('Incorrect network ID must be Matic (80001)');
  }

  const deployedAddress: string = address || ValistABI.networks[networkId].address;

  return getContractInstance(web3, ValistABI.abi, deployedAddress);
};

export const getSignatureParameters = (web3: Web3, signature: string) => {
  if (!web3.utils.isHexStrict(signature)) throw new Error(`Not a valid hex string: ${signature}`);

  const r = signature.slice(0, 66);
  const s = '0x'.concat(signature.slice(66, 130));
  let v: string | number = '0x'.concat(signature.slice(130, 132));
  v = web3.utils.hexToNumber(v);
  if (![27, 28].includes(v)) v += 27;

  return { r, s, v };
};

export const domainType = [
  { name: 'name', type: 'string' },
  { name: 'version', type: 'string' },
  { name: 'chainId', type: 'uint256' },
  { name: 'verifyingContract', type: 'address' },
];

export const metaTransactionType = [
  { name: 'nonce', type: 'uint256' },
  { name: 'from', type: 'address' },
  { name: 'functionSignature', type: 'bytes' },
];

export const domainData = {
  name: 'Valist',
  version: '0',
  chainId: 80001,
  verifyingContract: ValistABI.networks[80001].address,
};

export const Web3Providers = Web3.providers;
