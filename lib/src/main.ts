import Web3 from 'web3';
import { provider, Account } from 'web3-core/types';
// @ts-ignore
import ipfsClient from 'ipfs-http-client'

import ValistABI from './abis/Valist.json';
import ValistOrganizationABI from './abis/ValistOrganization.json';
import ValistRepositoryABI from './abis/ValistRepository.json';

const getContractInstance = async (web3: Web3, abi: any, address: string) => {
  // create the instance
  return new web3.eth.Contract(abi, address);
}

const getValistContract = async (web3: Web3) => {
  // get network ID and the deployed address
  const networkId: number = await web3.eth.net.getId();
  // @ts-ignore
  const deployedAddress: string = ValistABI.networks[networkId].address;

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
  ipfs: ipfsClient;

  constructor(web3Provider: provider, ipfsEnabled: boolean) {
    this.web3 = new Web3(web3Provider);
    if (ipfsEnabled) {
      this.ipfs = ipfsClient({ host: 'ipfs.infura.io', port: '5001', apiPath: '/api/v0/', protocol: 'https' });
    }
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

  // returns organization contract instance
  async getRepository(orgName: string, repoName: string) {
    const org = await this.getOrganization(orgName);
    const repoAddress = await org.methods.repos(repoName).call();
    const repo = await getValistRepositoryContract(this.web3, repoAddress);
    return repo;
  }

  async getOrganizationMeta(orgName: string) {
    const org = await this.getOrganization(orgName);
    const orgMeta = await org.methods.orgMeta().call();
    return orgMeta;
  }

  // returns repository contract instance
  async getRepoFromOrganization(orgName: string, repoName: string) {
    const org = await this.getOrganization(orgName);
    const repoAddress = await org.methods.repos(repoName).call();
    const repo = await getValistRepositoryContract(this.web3, repoAddress);
    return repo;
  }

  async getRepoMeta(orgName: string, repoName: string) {
    const repo = await this.getRepoFromOrganization(orgName, repoName);
    const repoMeta = await repo.methods.repoMeta().call();
    return repoMeta;
  }

  async getLatestReleaseFromRepo(orgName: string, repoName: string) {
    const repo = await this.getRepoFromOrganization(orgName, repoName);
    const release = await repo.methods.latestRelease().call();
    return release;
  }

  async createOrganization(orgName: string, orgMeta: string, account: any) {
    const result = await this.valist.methods.createOrganization(orgName, orgMeta).send({ from: account });
    return result;
  }

  async createRepository(orgName: string, repoName: string, repoMeta: string, account: any) {
    const org = await this.getOrganization(orgName);
    const result = await org.methods.createRepository(repoName, repoMeta).send({ from: account });
    return result;
  }

  async publishRelease(orgName: string, repoName: string, release: { tag: string, hash: string, meta: string }, account: any) {
    const repo = await this.getRepository(orgName, repoName);
    const result = await repo.methods.publishRelease(release.tag, release.hash, release.meta).send({ from: account });
    return result;
  }

  async addFileIpfs(data: any){
    const file = Buffer.from(JSON.stringify(data));
    try {
      const result = await this.ipfs.add(file);
      return result
    } catch (err) {
      console.error('Error', err);
    }
  }

}

export const Web3Providers = Web3.providers;

export default Valist;
