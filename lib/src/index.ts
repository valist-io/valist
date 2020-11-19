import Web3 from 'web3';
import { provider } from 'web3-core/types';
// @ts-ignore
import ipfsClient from 'ipfs-http-client'

import ValistABI from './abis/Valist.json';
import ValistOrganizationABI from './abis/ValistOrganization.json';
import ValistRepositoryABI from './abis/ValistRepository.json';

// monkey patch node-fetch
const fetch = require('node-fetch');
if (!globalThis.fetch) {
    globalThis.fetch = fetch;
}

// keccak256 hashes of each role
const ORG_ADMIN_ROLE = "0x123b642491709420c2370bb98c4e7de2b1bc05c5f9fd95ac4111e12683553c62";
const REPO_ADMIN_ROLE = "0xff7d2294a3c189284afb74beb7d578b566cf69863d5cb16db08773c21bea56c9";
const REPO_DEV_ROLE = "0x069bf569f27d389f2c70410107860b2e82ff561283b097a89e897daa5e34b1b6";

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
  defaultAccount: string;

  constructor(web3Provider: provider, ipfsEnabled?: boolean) {
    this.web3 = new Web3(web3Provider);
    this.defaultAccount = "0x0";
    if (ipfsEnabled) {
      this.ipfs = ipfsClient({ host: `ipfs.infura.io`, port: `5001`, apiPath: `/api/v0/`, protocol: `https` });
    }
  }

  // initialize main valist contract instance for future calls
  async connect() {
    try {
      this.valist = await getValistContract(this.web3);
    } catch (e) {
      const msg = `Could not connect to Valist registry contract`;
      console.error(msg, e);
      throw e;
    }

    try {
      const accounts = await this.web3.eth.getAccounts();
      this.defaultAccount = accounts[0] || "0x0";
    } catch (e) {
      const msg = `Could not set default account`;
      console.error(msg, e);
      throw e;
    }
  }

  // returns organization contract instance
  async getOrganization(orgName: string) {
    try {
      const orgAddress = await this.valist.methods.orgs(orgName).call();
      const org = await getValistOrganizationContract(this.web3, orgAddress);
      return org;
    } catch (e) {
      const msg = `Could not get organization contract`;
      console.error(msg, e);
      throw e;
    }
  }

  async getOrganizationMeta(orgName: string) {
    try {
      const org = await this.getOrganization(orgName);
      const orgMeta = await org.methods.orgMeta().call();
      const json = await this.fetchJSONfromIPFS(orgMeta);

      return json;

    } catch (e) {
      const msg = `Could not get organization metadata`;
      console.error(msg, e);
      throw e;
    }
  }

  async getCreatedOrganizations() {
    try {
      const organizations = await this.valist.getPastEvents('OrganizationCreated', {fromBlock: 0, toBlock: 'latest'});
      return organizations;
    } catch (e) {
      const msg = `Could not get created organizations`;
      console.error(msg, e);
      throw e;
    }
  }

  async getDeletedOrganizations() {
    try {
      const organizations = await this.valist.getPastEvents('OrganizationDeleted', {fromBlock: 0, toBlock: 'latest'});
      return organizations;
    } catch (e) {
      const msg = `Could not get deleted organizations`;
      console.error(msg, e);
      throw e;
    }
  }

  async setOrgMeta(orgName: string, orgMeta: any, account: string) {
    try {
      const org = await this.getOrganization(orgName);
      const hash = await this.addJSONtoIPFS(orgMeta);
      await org.methods.updateOrgMeta(hash).send({ from: account });
    } catch (e) {
      const msg = `Could not set organization metadata`;
      console.error(msg, e);
      throw e;
    }
  }

  // returns repository contract instance
  async getRepository(orgName: string, repoName: string) {
    try {
      const org = await this.getOrganization(orgName);
      const repoAddress = await org.methods.repos(repoName).call();
      const repo = await getValistRepositoryContract(this.web3, repoAddress);
      return repo;
    } catch (e) {
      const msg = `Could not get repository contract`;
      console.error(msg, e);
      throw e;
    }
  }

  async getReposFromOrganization(orgName: string) {
    try {
      const org = await this.getOrganization(orgName);
      const repos = await org.getPastEvents('RepositoryCreated', {fromBlock: 0, toBlock: 'latest'});
      return repos;
    } catch (e) {
      const msg = `Could not get repositories from organization`;
      console.error(msg, e);
      throw e;
    }
  }

  async getRepoMeta(orgName: string, repoName: string) {
    try {
      const repo = await this.getRepository(orgName, repoName);
      const repoMeta = await repo.methods.repoMeta().call();
      const json = await this.fetchJSONfromIPFS(repoMeta);

      return json;

    } catch (e) {
      const msg = `Could not get repository metadata`;
      console.error(msg, e);
      throw e;
    }
  }

  async setRepoMeta(orgName: string, repoName: string, repoMeta: any, account:string) {
    try {
      const repo = await this.getRepository(orgName, repoName);
      const hash = await this.addJSONtoIPFS(repoMeta);
      await repo.methods.updateRepoMeta(hash).send({ from: account });
    } catch (e) {
      const msg = `Could not set repository metadata`;
      console.error(msg, e);
      throw e;
    }
  }

  async getLatestTagFromRepo(orgName: string, repoName: string) {
    try {
      const repo = await this.getRepository(orgName, repoName);
      const tag = await repo.methods.tag().call();
      return tag;
    } catch (e) {
      const msg = `Could not get latest tag from repo`;
      console.error(msg, e);
      throw e;
    }
  }

  async getLatestReleaseFromRepo(orgName: string, repoName: string) {
    try {
      const repo = await this.getRepository(orgName, repoName);
      const release = await repo.methods.latestRelease().call();
      return release;
    } catch (e) {
      const msg = `Could not get latest release from repo`;
      console.error(msg, e);
      throw e;
    }
  }

  async getLatestReleaseMetaFromRepo(orgName: string, repoName: string) {
    try {
      const repo = await this.getRepository(orgName, repoName);
      const release = await repo.methods.releaseMeta().call();
      return release;
    } catch (e) {
      const msg = `Could not get latest release metadata from repo`;
      console.error(msg, e);
      throw e;
    }
  }

  async getReleasesFromRepo(orgName: string, repoName: string) {
    try {
      const repo = await this.getRepository(orgName, repoName);
      return await repo.getPastEvents('Release', {fromBlock: 0, toBlock: 'latest'});
    } catch (e) {
      const msg = `Could not get releases from repo`;
      console.error(msg, e);
      throw e;
    }
  }

  async getReleaseByTag(orgName: string, repoName: string, tag: string) {
    try {
      const events = await this.getReleasesFromRepo(orgName, repoName);

      // @TODO make this more efficient later
      for (let i = 0; i < events.length; i++) {
        if (events[i].returnValues.tag == tag) {
            const { tag, release, releaseMeta } = events[i].returnValues;
            return { tag, release, releaseMeta }
        }
      }

      return;

    } catch (e) {
      const msg = `Could not get release by tag`;
      console.error(msg, e);
      throw e;
    }
  }

  async isOrgAdmin(orgName: string, account: string) {
    try {
      const org = await this.getOrganization(orgName);
      return await org.methods.hasRole(ORG_ADMIN_ROLE, account).call();
    } catch (e) {
      const msg = `Could not check if user has ORG_ADMIN_ROLE`;
      console.error(msg, e);
      throw e;
    }
  }

  async isRepoAdmin(orgName: string, repoName: string, account: string) {
    try {
      const repo = await this.getRepository(orgName, repoName);
      return await repo.methods.hasRole(REPO_ADMIN_ROLE, account).call();
    } catch (e) {
      const msg = `Could not check if user has REPO_ADMIN_ROLE`;
      console.error(msg, e);
      throw e;
    }
  }

  async isRepoDev(orgName: string, repoName: string, account: string) {
    try {
      const repo = await this.getRepository(orgName, repoName);
      return await repo.methods.hasRole(REPO_DEV_ROLE, account).call();
    } catch (e) {
      const msg = `Could not check if user has REPO_DEV_ROLE`;
      console.error(msg, e);
      throw e;
    }
  }

  async grantOrgAdmin(orgName: string, granter: string, grantee: string) {
    try {
      const org = await this.getOrganization(orgName);
      await org.methods.grantRole(ORG_ADMIN_ROLE, grantee).send({ from: granter });
    } catch (e) {
      const msg = `Could not grant ORG_ADMIN_ROLE`;
      console.error(msg, e);
      throw e;
    }
  }

  async revokeOrgAdmin(orgName: string, revoker: string, revokee: string) {
    try {
      const org = await this.getOrganization(orgName);
      await org.methods.revokeRole(ORG_ADMIN_ROLE, revokee).send({ from: revoker });
    } catch (e) {
      const msg = `Could not revoke ORG_ADMIN_ROLE`;
      console.error(msg, e);
      throw e;
    }
  }

  async grantRepoAdmin(orgName: string, repoName: string, granter: string, grantee: string) {
    try {
      const repo = await this.getRepository(orgName, repoName);
      return await repo.methods.grantRole(REPO_ADMIN_ROLE, grantee).send( { from: granter });
    } catch (e) {
      const msg = `Could not grant REPO_ADMIN_ROLE`;
      console.error(msg, e);
      throw e;
    }
  }

  async revokeRepoAdmin(orgName: string, repoName: string, revoker: string, revokee: string) {
    try {
      const repo = await this.getRepository(orgName, repoName);
      return await repo.methods.revokeRole(REPO_ADMIN_ROLE, revokee).send( { from: revoker });
    } catch (e) {
      const msg = `Could not revoke REPO_ADMIN_ROLE`;
      console.error(msg, e);
      throw e;
    }
  }

  async grantRepoDev(orgName: string, repoName: string, granter: string, grantee: string) {
    try {
      const repo = await this.getRepository(orgName, repoName);
      return await repo.methods.grantRole(REPO_DEV_ROLE, grantee).send( { from: granter });
    } catch (e) {
      const msg = `Could not grant REPO_DEV_ROLE`;
      console.error(msg, e);
      throw e;
    }
  }

  async revokeRepoDev(orgName: string, repoName: string, revoker: string, revokee: string) {
    try {
      const repo = await this.getRepository(orgName, repoName);
      return await repo.methods.revokeRole(REPO_DEV_ROLE, revokee).send( { from: revoker });
    } catch (e) {
      const msg = `Could not revoke REPO_DEV_ROLE`;
      console.error(msg, e);
      throw e;
    }
  }

  async getOrgAdmins(orgName: string) {
    try {
      const org = await this.getOrganization(orgName);
      const adminCount = await org.methods.getRoleMemberCount(ORG_ADMIN_ROLE).call();

      const members = [];
      for (let i = 0; i < adminCount; ++i) {
          members.push(await org.methods.getRoleMember(ORG_ADMIN_ROLE, i).call());
      }

      return members;
    } catch (e) {
      const msg = `Could not get users that have REPO_ADMIN_ROLE`;
      console.error(msg, e);
      throw e;
    }
  }

  async getRepoAdmins(orgName: string, repoName: string) {
    try {
      const repo = await this.getRepository(orgName, repoName);
      const adminCount = await repo.methods.getRoleMemberCount(REPO_ADMIN_ROLE).call();
      console.log(adminCount)

      const members = [];
      for (let i = 0; i < adminCount; ++i) {
          members.push(await repo.methods.getRoleMember(REPO_ADMIN_ROLE, i).call());
      }

      return members;
    } catch (e) {
      const msg = `Could not get users that have REPO_ADMIN_ROLE`;
      console.error(msg, e);
      throw e;
    }
  }

  async getRepoDevs(orgName: string, repoName: string) {
    try {
      const repo = await this.getRepository(orgName, repoName);
      const devCount = await repo.methods.getRoleMemberCount(REPO_DEV_ROLE).call();

      const members = [];
      for (let i = 0; i < devCount; ++i) {
          members.push(await repo.methods.getRoleMember(REPO_DEV_ROLE, i).call());
      }

      return members;
    } catch (e) {
      const msg = `Could not get users that have REPO_DEV_ROLE`;
      console.error(msg, e);
      throw e;
    }
  }

  async createOrganization(orgName: string, orgMeta: {name: string, description: string}, account: string) {
    try {
      const metaFile: string = await this.addJSONtoIPFS(orgMeta);
      const result = await this.valist.methods.createOrganization(orgName.toLowerCase(), metaFile).send({ from: account });
      return result;
    } catch (e) {
      const msg = `Could not create organization`;
      console.error(msg, e);
      throw e;
    }
  }

  async createRepository(orgName: string, repoName: string, repoMeta: {name: string, description: string, projectType: string, homepage: string, github: string}, account: string) {
    try {
      const org = await this.getOrganization(orgName);
      const metaFile = await this.addJSONtoIPFS(repoMeta);
      const result = await org.methods.createRepository(repoName, metaFile).send({ from: account });
      return result;
    } catch(e) {
      const msg = `Could not create repository`;
      console.error(msg, e);
      throw e;
    }
  }

  async publishRelease(orgName: string, repoName: string, release: { tag: string, hash: string, meta: string }, account: string) {
    try {
      const repo = await this.getRepository(orgName, repoName);
      const result = await repo.methods.publishRelease(release.tag, release.hash, release.meta).send({ from: account });
      return result;
    } catch (e) {
      const msg = `Could not publish release`;
      console.error(msg, e);
      throw e;
    }
  }

  async addJSONtoIPFS(data: any) {
    try {
      const file = Buffer.from(JSON.stringify(data));
      const result = await this.ipfs.add(file);
      return result["path"];
    } catch (e) {
      const msg = `Could not add JSON to IPFS`;
      console.error(msg, e);
      throw e;
    }
  }

  async addFileToIPFS(data: any) {
    try {
      const file = Buffer.from(data);
      const result = await this.ipfs.add(file);
      return result["path"];
    } catch (e) {
      const msg = `Could not add file to IPFS`;
      console.error(msg, e);
      throw e;
    }
  }

  async fetchJSONfromIPFS(ipfsHash: string) {
    try {
      const response = await fetch(`https://cloudflare-ipfs.com/ipfs/${ipfsHash}`);
      const json = await response.json();
      console.log(`JSON Fetched from IPFS`, json);
      return json;
    } catch (e) {
      const msg = `Could not fetch JSON from IPFS`;
      console.error(msg, e);
      throw e;
    }
  }
}

export const Web3Providers = Web3.providers;

export default Valist;
