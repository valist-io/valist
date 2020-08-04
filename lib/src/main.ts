import Web3 from 'web3';
import axios from 'axios'
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
  ipfs: any;
  fileBuffer: any;
  ipfsEnabled: boolean;

  constructor(provider: any, ipfsEnabled:boolean) {
    this.web3 = new Web3(provider);
    this.ipfsEnabled = true;
  }

  // initialize main valist contract instance for future calls
  async connect() {
    this.valist = await getValistContract(this.web3);
    if (this.ipfsEnabled === true){
      this.ipfs = ipfsClient({ host: 'https://ipfs.infura.io', port: '5001', apiPath: '/api/v0/' })
      this.fileBuffer = (data: any) => this.ipfs.types.Buffer.from(JSON.stringify(data))
    }
  }

  // returns organization contract instance
  async getOrganization(orgName: string) {
    const orgAddress = await this.valist.methods.orgs(orgName).call();
    const org = await getValistOrganizationContract(this.web3, orgAddress);
    return org;
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

  async getLatestReleaseFromRepo(orgName: string, repoName: string) {
    const repo = await this.getRepoFromOrganization(orgName, repoName);
    const release = await repo.methods.latestRelease().call();
    return release;
  }

  async createOrganization(orgName: string, orgMeta: string, account: any) {
    const result = await this.valist.methods.createOrganization(orgName, orgMeta).send({ from: account });
    return result;
  }

  async addFileIpfs(data: any){
    const file = this.fileBuffer(data)
    try {
      const result = await this.ipfs.add(file)
      return result[0].hash
    } catch (err) {
      console.error('Error', err)
    }
  }

  async getFileIpfs(hash: any){
    const endpoint = `"https://ipfs.infura.io:5001/api/v0/block/get?arg=/ipfs/${hash}`
    try {
      const { data } = await axios.get(endpoint)
      return data
    } catch (err) {
      console.error('Error', err)
    }
  }
}

export = Valist;
