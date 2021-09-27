import Web3 from 'web3';
import ValistABI from './abis/contracts/Valist.sol/Valist.json';
import ValistRegistryABI from './abis/contracts/ValistRegistry.sol/ValistRegistry.json';

import { sendMetaTx } from './helpers/biconomy';

import { InvalidNetworkError } from './errors';

export const shortnameFilterRegex = /[^A-z0-9-]/g;

export const getContractInstance = (web3: Web3, abi: any, address: string) => new web3.eth.Contract(abi, address);

export const getValistContract = async (web3: Web3, chainID: number, address?: string) => {
  const networkContractMap = {
    80001: '0xA7E4124aDBBc50CF402e4Cad47de906a14daa0f6',
  };

  if (!address && !Object.keys(networkContractMap).includes(chainID.toString())) {
    throw new InvalidNetworkError('Valist not found on network');
  }

  const deployedAddress: string = address || networkContractMap[80001];
  return getContractInstance(web3, ValistABI.abi, deployedAddress);
};

export const getValistRegistry = async (web3: Web3, chainID: number, address?: string) => {
  const networkContractMap = {
    80001: '0x2Be6D782dBA2C52Cd0a41c6052e914dCaBcCD78e',
  };

  if (!address && !Object.keys(networkContractMap).includes(chainID.toString())) {
    throw new InvalidNetworkError('Valist Registry not found on network');
  }

  const deployedAddress: string = address || networkContractMap[80001];
  return getContractInstance(web3, ValistRegistryABI.abi, deployedAddress);
};

export const sendMetaTransaction = sendMetaTx;

export const Web3Providers = Web3.providers;

// strip first occurrence of parent folder from absolute path.
// allows for sub-folders to be the same name as the parent
// e.g. '/example-projects/cli-publish-binary/dist/dist/dist/nested' => '/dist/dist/nested'
// also works with 'C:\\example-projects\\cli-publish-binary\\dist\\dist\\dist\\nested' => '\dist\dist\nested'
export const stripParentFolderFromPath = (filePath: string, parent: string): string => filePath.replace(
  new RegExp(`^.+?${parent.replace(/\/|\\$/, '')}`),
  '',
);

export const parseCID = (url: string) => url.replace('/ipfs/', '');