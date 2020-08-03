import Web3 from 'web3';

import ValistABI from './abis/Valist.json';
import ValistOrganizationABI from './abis/ValistOrganization.json';
import ValistRepositoryABI from './abis/ValistRepository.json';

const getContractInstance = async (web3: Web3, abi: any, address: string) => {
  // create the instance
  return new web3.eth.Contract(abi, address);
}

const getValistContract = async (web3: Web3) => {
  // get network ID and the deployed address
  const networkId = await web3.eth.net.getId();
  // @ts-ignore
  const deployedAddress: any = ValistABI.networks[networkId].address

  return await getContractInstance(web3, ValistABI.abi, deployedAddress);
}

const getValistOrganizationContract = async (web3: Web3, address: string) => {
  // create the instance
  return await getContractInstance(web3, ValistOrganizationABI.abi, address);
}

const getValistRepositoryContract = async (web3: Web3, address: string) => {
  // get network ID and the deployed address
  return await getContractInstance(web3, ValistRepositoryABI.abi, address);
}

class Valist {

  web3: Web3;
  valist: any;

  constructor(provider: any) {
    this.web3 = new Web3(provider);
  }

  // initialize main valist contract instance for future calls
  async connect() {
    this.valist = await getValistContract(this.web3);
  }

  // returns organization contract instance
  async getOrganization(orgName: string) {
    const orgAddress = await this.valist.methods.orgs(orgName).call();
    const org = await getValistOrganizationContract(this.web3, orgAddress);
    return org;
  }

  // returns repository contract instance
  async getRepoFromOrganization(orgName: string, repoName: string) {
    const org = await this.getOrganization(orgName);
    const repoAddress = await org.methods.repos(repoName).call();
    const repo = await getValistRepositoryContract(this.web3, repoAddress);
    return repo;
  }

  async getLatestReleaseFromRepo(orgName: string, repoName: string) {
    const repo = await this.getRepoFromOrganization(orgName, repoName);
    const release = await repo.methods.latestRelease().call();
    return release;
  }

  async createOrganization(orgName: string, account: any) {
    const result = await this.valist.methods.createOrganization(orgName).send({ from: account });
    return result;
  }

}

module.exports = Valist;
