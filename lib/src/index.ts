import Web3 from 'web3';
import { provider } from 'web3-core/types';
// @ts-ignore
import ipfsClient from 'ipfs-http-client';

import ValistABI from './abis/Valist.json';

// node-fetch polyfill
const fetch = require('node-fetch');
if (!globalThis.fetch) {
    globalThis.fetch = fetch;
}

export const shortnameFilterRegex = /[^A-z0-9-]/;

export type ProjectType = "binary" | "npm" | "pip" | "docker";

const getContractInstance = (web3: Web3, abi: any, address: string) => {
  // create the instance
  return new web3.eth.Contract(abi, address);
}

const getValistContract = async (web3: Web3) => {
  // get network ID and the deployed address
  const networkId: number = await web3.eth.net.getId();
  // @ts-ignore
  const deployedAddress: string = ValistABI.networks[networkId].address;

  return getContractInstance(web3, ValistABI.abi, deployedAddress);
}
class Valist {

  web3: Web3;
  valist: any;
  ipfs: any;
  defaultAccount: string;

  constructor({ web3Provider, metaTx = false, ipfsHost = `ipfs.infura.io`}: { web3Provider: provider, metaTx?: boolean, ipfsHost?: string }) {
    if (metaTx) {
      this.web3 = new Web3(web3Provider);
    } else {
      this.web3 = new Web3(web3Provider);
    }
    this.defaultAccount = "0x0";
    this.ipfs = ipfsClient({ host: ipfsHost, port: 5001, apiPath: `/api/v0/`, protocol: `https` });
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

  // returns organization meta and release tags
  async getOrganization(orgName: string) {
    try {
      const org = await this.valist.methods.getOrganization(orgName).call();

      let json: any = {};

      try { json = await this.fetchJSONfromIPFS(org[0]) } catch (e) {}

      return { meta: json, repoNames: org[1] };

    } catch (e) {
      const msg = `Could not get organization`;
      console.error(msg, e);
      throw e;
    }
  }

  async getOrganizationMeta(orgName: string) {
    try {
      const orgMeta = await this.valist.methods.getOrgMeta(orgName).call();

      let json: any = {};

      try { json = await this.fetchJSONfromIPFS(orgMeta) } catch (e) {}

      return json;

    } catch (e) {
      const msg = `Could not get organization metadata`;
      console.error(msg, e);
      throw e;
    }
  }

  async getOrganizationNames() {
    try {
      const orgs = await this.valist.methods.getOrgNames().call();
      return orgs;
    } catch (e) {
      const msg = `Could not get organization names`;
      console.error(msg, e);
      throw e;
    }
  }

  async setOrgMeta(orgName: string, orgMeta: any, account: string) {
    try {
      const hash = await this.addJSONtoIPFS(orgMeta);
      await this.valist.methods.setOrgMeta(orgName, hash).send({ from: account });
    } catch (e) {
      const msg = `Could not set organization metadata`;
      console.error(msg, e);
      throw e;
    }
  }

  // returns repository
  async getRepository(orgName: string, repoName: string) {
    try {
      const repo = await this.valist.methods.getRepository(orgName, repoName).call();

      let json: any = {};

      try { json = await this.fetchJSONfromIPFS(repo[0]) } catch (e) {}

      return { meta: json, tags: repo[1] };

    } catch (e) {
      const msg = `Could not get repository contract`;
      console.error(msg, e);
      throw e;
    }
  }

  async getReposFromOrganization(orgName: string) {
    try {
      const repos = await this.valist.methods.getRepoNames(orgName).call();
      return repos;
    } catch (e) {
      const msg = `Could not get repositories from organization`;
      console.error(msg, e);
      throw e;
    }
  }

  async getRepoMeta(orgName: string, repoName: string) {
    try {
      const repoMeta = await this.valist.methods.getRepoMeta(orgName, repoName).call();

      let json: any = {};

      try { json = await this.fetchJSONfromIPFS(repoMeta) } catch (e) {};

      return json;

    } catch (e) {
      const msg = `Could not get repository metadata`;
      console.error(msg, e);
      throw e;
    }
  }

  async setRepoMeta(orgName: string, repoName: string, repoMeta: any, account: string) {
    try {
      const hash = await this.addJSONtoIPFS(repoMeta);
      await this.valist.methods.setRepoMeta(orgName, repoName, hash).send({ from: account });
    } catch (e) {
      const msg = `Could not set repository metadata`;
      console.error(msg, e);
      throw e;
    }
  }

  async getLatestReleaseFromRepo(orgName: string, repoName: string) {
    try {
      const release = await this.valist.methods.getLatestRelease(orgName, repoName).call();
      return release;
    } catch (e) {
      const msg = `Could not get latest release from repo`;
      console.error(msg, e);
      throw e;
    }
  }

  async getLatestTagFromRepo(orgName: string, repoName: string) {
    try {
      const tag = await this.valist.methods.getLatestTag(orgName, repoName).call();
      return tag;
    } catch (e) {
      const msg = `Could not get latest tag from repo`;
      console.error(msg, e);
      throw e;
    }
  }

  async getReleaseTagsFromRepo(orgName: string, repoName: string) {
    try {
      const tags = await this.valist.methods.getReleaseTags(orgName, repoName).call();
      return tags;
    } catch (e) {
      const msg = `Could not get release tags from repo`;
      console.error(msg, e);
      throw e;
    }
  }

  async getReleasesFromRepo(orgName: string, repoName: string) {
    try {
      const tags = await this.valist.methods.getReleaseTags(orgName, repoName).call();
      const releases: any[] = [];

      for (let i = 0; i < tags.length; i++) {
        const release = await this.valist.methods.getRelease(orgName, repoName, tags[i]).call();
        releases.push({...release, tag: tags[i]});
      };

      return releases;
    } catch (e) {
      const msg = `Could not get releases from repo`;
      console.error(msg, e);
      throw e;
    }
  }

  async getReleaseByTag(orgName: string, repoName: string, tag: string) {
    try {
      const release = await this.valist.methods.getRelease(orgName, repoName, tag).call();

      return release;
    } catch (e) {
      const msg = `Could not get release by tag`;
      console.error(msg, e);
      throw e;
    }
  }

  async isOrgOwner(orgName: string, account: string) {
    try {
      return await this.valist.methods.isOrgOwner(orgName, account).call();
    } catch (e) {
      const msg = `Could not check if user has ORG_ADMIN_ROLE`;
      console.error(msg, e);
      throw e;
    }
  }

  async isOrgAdmin(orgName: string, account: string) {
    try {
      return await this.valist.methods.isOrgAdmin(orgName, account).call();
    } catch (e) {
      const msg = `Could not check if user has ORG_ADMIN_ROLE`;
      console.error(msg, e);
      throw e;
    }
  }

  async isRepoAdmin(orgName: string, repoName: string, account: string) {
    try {
      return await this.valist.methods.isRepoAdmin(orgName, repoName, account).call();
    } catch (e) {
      const msg = `Could not check if user has REPO_ADMIN_ROLE`;
      console.error(msg, e);
      throw e;
    }
  }

  async isRepoDev(orgName: string, repoName: string, account: string) {
    try {
      return await this.valist.methods.isRepoDev(orgName, repoName, account).call();
    } catch (e) {
      const msg = `Could not check if user has REPO_DEV_ROLE`;
      console.error(msg, e);
      throw e;
    }
  }

  async grantOrgAdmin(orgName: string, granter: string, grantee: string) {
    try {
      await this.valist.methods.grantOrgAdmin(orgName, grantee).send({ from: granter });
    } catch (e) {
      const msg = `Could not grant ORG_ADMIN_ROLE`;
      console.error(msg, e);
      throw e;
    }
  }

  async revokeOrgAdmin(orgName: string, revoker: string, revokee: string) {
    try {
      await this.valist.methods.revokeOrgAdmin(orgName, revokee).send({ from: revoker });
    } catch (e) {
      const msg = `Could not revoke ORG_ADMIN_ROLE`;
      console.error(msg, e);
      throw e;
    }
  }

  async grantRepoAdmin(orgName: string, repoName: string, granter: string, grantee: string) {
    try {
      await this.valist.methods.grantRepoAdmin(orgName, repoName, grantee).send( { from: granter });
    } catch (e) {
      const msg = `Could not grant REPO_ADMIN_ROLE`;
      console.error(msg, e);
      throw e;
    }
  }

  async revokeRepoAdmin(orgName: string, repoName: string, revoker: string, revokee: string) {
    try {
      await this.valist.methods.revokeRepoAdmin(orgName, repoName, revokee).send( { from: revoker });
    } catch (e) {
      const msg = `Could not revoke REPO_ADMIN_ROLE`;
      console.error(msg, e);
      throw e;
    }
  }

  async grantRepoDev(orgName: string, repoName: string, granter: string, grantee: string) {
    try {
      await this.valist.methods.grantRepoDev(orgName, repoName, grantee).send( { from: granter });
    } catch (e) {
      const msg = `Could not grant REPO_DEV_ROLE`;
      console.error(msg, e);
      throw e;
    }
  }

  async revokeRepoDev(orgName: string, repoName: string, revoker: string, revokee: string) {
    try {
      await this.valist.methods.revokeRepoDev(orgName, repoName, revokee).send( { from: revoker });
    } catch (e) {
      const msg = `Could not revoke REPO_DEV_ROLE`;
      console.error(msg, e);
      throw e;
    }
  }

  async getOrgOwners(orgName: string) {
    try {
      const members = await this.valist.methods.getOrgOwners(orgName).call();

      return members;
    } catch (e) {
      const msg = `Could not get users that have ORG_OWNER_ROLE`;
      console.error(msg, e);
      throw e;
    }
  }

  async getOrgAdmins(orgName: string) {
    try {
      const members = await this.valist.methods.getOrgAdmins(orgName).call();

      return members;
    } catch (e) {
      const msg = `Could not get users that have ORG_ADMIN_ROLE`;
      console.error(msg, e);
      throw e;
    }
  }

  async getRepoAdmins(orgName: string, repoName: string) {
    try {
      const members = await this.valist.methods.getRepoAdmins(orgName, repoName).call();

      return members;
    } catch (e) {
      const msg = `Could not get users that have REPO_ADMIN_ROLE`;
      console.error(msg, e);
      throw e;
    }
  }

  async getRepoDevs(orgName: string, repoName: string) {
    try {
      const members = await this.valist.methods.getRepoDevs(orgName, repoName).call();

      return members;
    } catch (e) {
      const msg = `Could not get users that have REPO_DEV_ROLE`;
      console.error(msg, e);
      throw e;
    }
  }

  async createOrganization(orgName: string, orgMeta: { name: string, description: string }, account: string) {
    try {
      const metaFile: string = await this.addJSONtoIPFS(orgMeta);
      const block = await this.web3.eth.getBlock("latest");
      const result = await this.valist.methods.createOrganization(orgName.toLowerCase().replace(shortnameFilterRegex, ""), metaFile).send({ from: account, gasLimit: block.gasLimit });
      return result;
    } catch (e) {
      const msg = `Could not create organization`;
      console.error(msg, e);
      throw e;
    }
  }

  async createRepository(orgName: string, repoName: string, repoMeta: { name: string, description: string, projectType: ProjectType, homepage: string, repository: string }, account: string) {
    try {
      const metaFile = await this.addJSONtoIPFS(repoMeta);
      const block = await this.web3.eth.getBlock("latest");
      const result = await this.valist.methods.createRepository(orgName, repoName.toLowerCase().replace(shortnameFilterRegex, ""), metaFile).send({ from: account, gasLimit: block.gasLimit });
      return result;
    } catch(e) {
      const msg = `Could not create repository`;
      console.error(msg, e);
      throw e;
    }
  }

  async publishRelease(orgName: string, repoName: string, release: { tag: string, hash: string, meta: string }, account: string) {
    try {
      const block = await this.web3.eth.getBlock("latest");
      await this.valist.methods.publishRelease(orgName, repoName, release.tag, release.hash, release.meta).send({ from: account, gasLimit: block.gasLimit });
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
