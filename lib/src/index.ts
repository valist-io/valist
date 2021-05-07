import Web3 from 'web3';
// @ts-ignore ipfs client types are finicky
import ipfsClient from 'ipfs-http-client';
// @ts-ignore mexa doesn't support typescript yet
import Biconomy from '@biconomy/mexa';

import { provider } from 'web3-core/types';

import * as sigUtil from 'eth-sig-util';

import { OrgMeta, Release, RepoMeta } from './types';
import {
  domainData,
  domainType,
  getSignatureParameters,
  getValistContract,
  metaTransactionType,
  shortnameFilterRegex,
} from './utils';

// node-fetch polyfill
const fetch = require('node-fetch');

if (!globalThis.fetch) {
  globalThis.fetch = fetch;
}
class Valist {
  web3: Web3;

  valist: any;

  ipfs: any;

  biconomy: any;

  signer?: string;

  defaultAccount: string;

  metaTxEnabled: boolean = false;

  metaTxReady: boolean = false;

  contractAddress: string | undefined;

  constructor({
    web3Provider,
    metaTx = true,
    ipfsHost = 'https://pin.valist.io',
    contractAddress,
  }: {
    web3Provider: provider,
    metaTx?: boolean | string,
    ipfsHost?: string,
    contractAddress?: string,
  }) {
    if (metaTx === true || metaTx === 'true') {
      this.metaTxEnabled = true;

      // setup biconomy instance with public api key
      this.biconomy = new Biconomy(
        web3Provider,
        { apiKey: 'qLW9TRUjQ.f77d2f86-c76a-4b9c-b1ee-0453d0ead878', strictMode: true },
      );

      this.web3 = new Web3(this.biconomy);

      this.biconomy.onEvent(this.biconomy.READY, () => {
        this.metaTxReady = true;
        console.log('👻 MetaTransactions Enabled');
      });
    } else {
      this.web3 = new Web3(web3Provider);
    }

    this.defaultAccount = '0x0';
    this.ipfs = ipfsClient(ipfsHost);
    this.contractAddress = contractAddress;
  }

  // initialize main valist contract instance for future calls
  async connect(waitForMetaTx?: boolean) {
    try {
      this.valist = await getValistContract(this.web3, this.contractAddress);
    } catch (e) {
      const msg = 'Could not connect to Valist registry contract';
      console.error(msg, e);
      throw e;
    }

    try {
      const accounts = await this.web3.eth.getAccounts();
      this.defaultAccount = accounts[0] || '0x0';
    } catch (e) {
      const msg = 'Could not set default account';
      console.error(msg, e);
      throw e;
    }

    if (waitForMetaTx && this.biconomy) {
      await new Promise((resolve) => {
        this.biconomy.onEvent(this.biconomy.READY, () => {
          this.metaTxReady = true;
          resolve(true);
        });
      });
    }
  }

  // returns organization meta and release tags
  async getOrganization(orgName: string) {
    try {
      const org = await this.valist.methods.getOrganization(orgName).call();

      let json: any = {};

      try { json = await this.fetchJSONfromIPFS(org[0]); } catch (e) {
        // noop, just return empty object if failed
      }

      return { meta: json, repoNames: org[1] };
    } catch (e) {
      const msg = 'Could not get organization';
      console.error(msg, e);
      throw e;
    }
  }

  async getOrganizationMeta(orgName: string) {
    try {
      const orgMeta = await this.valist.methods.getOrgMeta(orgName).call();

      let json: any = {};

      try { json = await this.fetchJSONfromIPFS(orgMeta); } catch (e) {
        // noop, just return empty object if failed
      }

      return json;
    } catch (e) {
      const msg = 'Could not get organization metadata';
      console.error(msg, e);
      throw e;
    }
  }

  async getOrganizationNames() {
    try {
      const orgs = await this.valist.methods.getOrgNames().call();
      return orgs;
    } catch (e) {
      const msg = 'Could not get organization names';
      console.error(msg, e);
      throw e;
    }
  }

  async setOrgMeta(orgName: string, orgMeta: any, account: string) {
    try {
      const hash = await this.addJSONtoIPFS(orgMeta);
      return await this.sendTransaction(this.valist.methods.setOrgMeta(orgName, hash), account);
    } catch (e) {
      const msg = 'Could not set organization metadata';
      console.error(msg, e);
      throw e;
    }
  }

  // returns repository
  async getRepository(orgName: string, repoName: string) {
    try {
      const repo = await this.valist.methods.getRepository(orgName, repoName).call();

      let json: any = {};

      try { json = await this.fetchJSONfromIPFS(repo[0]); } catch (e) {
        // noop, just return empty object if failed
      }

      return { meta: json, tags: repo[1] };
    } catch (e) {
      const msg = 'Could not get repository contract';
      console.error(msg, e);
      throw e;
    }
  }

  async getReposFromOrganization(orgName: string) {
    try {
      const repos = await this.valist.methods.getRepoNames(orgName).call();
      return repos;
    } catch (e) {
      const msg = 'Could not get repositories from organization';
      console.error(msg, e);
      throw e;
    }
  }

  async getRepoMeta(orgName: string, repoName: string) {
    try {
      const repoMeta = await this.valist.methods.getRepoMeta(orgName, repoName).call();

      let json: any = {};

      try { json = await this.fetchJSONfromIPFS(repoMeta); } catch (e) {
        // noop, just return empty object if failed
      }

      return json;
    } catch (e) {
      const msg = 'Could not get repository metadata';
      console.error(msg, e);
      throw e;
    }
  }

  async setRepoMeta(orgName: string, repoName: string, repoMeta: any, account: string) {
    try {
      const hash = await this.addJSONtoIPFS(repoMeta);
      return await this.sendTransaction(this.valist.methods.setRepoMeta(orgName, repoName, hash), account);
    } catch (e) {
      const msg = 'Could not set repository metadata';
      console.error(msg, e);
      throw e;
    }
  }

  async getLatestReleaseFromRepo(orgName: string, repoName: string) {
    try {
      const release = await this.valist.methods.getLatestRelease(orgName, repoName).call();
      return release;
    } catch (e) {
      const msg = 'Could not get latest release from repo';
      console.error(msg, e);
      throw e;
    }
  }

  async getLatestTagFromRepo(orgName: string, repoName: string) {
    try {
      const tag = await this.valist.methods.getLatestTag(orgName, repoName).call();
      return tag;
    } catch (e) {
      const msg = 'Could not get latest tag from repo';
      console.error(msg, e);
      throw e;
    }
  }

  async getReleaseTagsFromRepo(orgName: string, repoName: string) {
    try {
      const tags = await this.valist.methods.getReleaseTags(orgName, repoName).call();
      return tags;
    } catch (e) {
      const msg = 'Could not get release tags from repo';
      console.error(msg, e);
      throw e;
    }
  }

  async getReleasesFromRepo(orgName: string, repoName: string): Promise<Release[]> {
    try {
      const tags = await this.valist.methods.getReleaseTags(orgName, repoName).call();
      const releases: Release[] = [];

      for (let i = 0; i < tags.length; i++) {
        // eslint-disable-next-line no-await-in-loop
        const release = await this.valist.methods.getRelease(orgName, repoName, tags[i]).call();
        releases.push({ ...release, tag: tags[i] });
      }

      return releases;
    } catch (e) {
      const msg = 'Could not get releases from repo';
      console.error(msg, e);
      throw e;
    }
  }

  async getReleaseByTag(orgName: string, repoName: string, tag: string) {
    try {
      const release = await this.valist.methods.getRelease(orgName, repoName, tag).call();

      return release;
    } catch (e) {
      const msg = 'Could not get release by tag';
      console.error(msg, e);
      throw e;
    }
  }

  async isOrgOwner(orgName: string, account: string) {
    try {
      const isOrgOwner = await this.valist.methods.isOrgOwner(orgName, account).call();
      return isOrgOwner;
    } catch (e) {
      const msg = 'Could not check if user has ORG_ADMIN_ROLE';
      console.error(msg, e);
      throw e;
    }
  }

  async isOrgAdmin(orgName: string, account: string) {
    try {
      const isOrgAdmin = await this.valist.methods.isOrgAdmin(orgName, account).call();
      return isOrgAdmin;
    } catch (e) {
      const msg = 'Could not check if user has ORG_ADMIN_ROLE';
      console.error(msg, e);
      throw e;
    }
  }

  async isRepoAdmin(orgName: string, repoName: string, account: string) {
    try {
      const isRepoAdmin = await this.valist.methods.isRepoAdmin(orgName, repoName, account).call();
      return isRepoAdmin;
    } catch (e) {
      const msg = 'Could not check if user has REPO_ADMIN_ROLE';
      console.error(msg, e);
      throw e;
    }
  }

  async isRepoDev(orgName: string, repoName: string, account: string) {
    try {
      const isRepoDev = await this.valist.methods.isRepoDev(orgName, repoName, account).call();
      return isRepoDev;
    } catch (e) {
      const msg = 'Could not check if user has REPO_DEV_ROLE';
      console.error(msg, e);
      throw e;
    }
  }

  async grantOrgAdmin(orgName: string, granter: string, grantee: string) {
    try {
      const tx = await this.sendTransaction(this.valist.methods.grantOrgAdmin(orgName, grantee), granter);
      return tx;
    } catch (e) {
      const msg = 'Could not grant ORG_ADMIN_ROLE';
      console.error(msg, e);
      throw e;
    }
  }

  async revokeOrgAdmin(orgName: string, revoker: string, revokee: string) {
    try {
      const tx = await this.sendTransaction(this.valist.methods.revokeOrgAdmin(orgName, revokee), revoker);
      return tx;
    } catch (e) {
      const msg = 'Could not revoke ORG_ADMIN_ROLE';
      console.error(msg, e);
      throw e;
    }
  }

  async grantRepoAdmin(orgName: string, repoName: string, granter: string, grantee: string) {
    try {
      const tx = await this.sendTransaction(this.valist.methods.grantRepoAdmin(orgName, repoName, grantee), granter);
      return tx;
    } catch (e) {
      const msg = 'Could not grant REPO_ADMIN_ROLE';
      console.error(msg, e);
      throw e;
    }
  }

  async revokeRepoAdmin(orgName: string, repoName: string, revoker: string, revokee: string) {
    try {
      const tx = await this.sendTransaction(this.valist.methods.revokeRepoAdmin(orgName, repoName, revokee), revoker);
      return tx;
    } catch (e) {
      const msg = 'Could not revoke REPO_ADMIN_ROLE';
      console.error(msg, e);
      throw e;
    }
  }

  async grantRepoDev(orgName: string, repoName: string, granter: string, grantee: string) {
    try {
      const tx = await this.sendTransaction(this.valist.methods.grantRepoDev(orgName, repoName, grantee), granter);
      return tx;
    } catch (e) {
      const msg = 'Could not grant REPO_DEV_ROLE';
      console.error(msg, e);
      throw e;
    }
  }

  async revokeRepoDev(orgName: string, repoName: string, revoker: string, revokee: string) {
    try {
      const tx = await this.sendTransaction(this.valist.methods.revokeRepoDev(orgName, repoName, revokee), revoker);
      return tx;
    } catch (e) {
      const msg = 'Could not revoke REPO_DEV_ROLE';
      console.error(msg, e);
      throw e;
    }
  }

  async getOrgOwners(orgName: string) {
    try {
      const members = await this.valist.methods.getOrgOwners(orgName).call();
      return members;
    } catch (e) {
      const msg = 'Could not get users that have ORG_OWNER_ROLE';
      console.error(msg, e);
      throw e;
    }
  }

  async getOrgAdmins(orgName: string) {
    try {
      const members = await this.valist.methods.getOrgAdmins(orgName).call();
      return members;
    } catch (e) {
      const msg = 'Could not get users that have ORG_ADMIN_ROLE';
      console.error(msg, e);
      throw e;
    }
  }

  async getRepoAdmins(orgName: string, repoName: string) {
    try {
      const members = await this.valist.methods.getRepoAdmins(orgName, repoName).call();
      return members;
    } catch (e) {
      const msg = 'Could not get users that have REPO_ADMIN_ROLE';
      console.error(msg, e);
      throw e;
    }
  }

  async getRepoDevs(orgName: string, repoName: string) {
    try {
      const members = await this.valist.methods.getRepoDevs(orgName, repoName).call();
      return members;
    } catch (e) {
      const msg = 'Could not get users that have REPO_DEV_ROLE';
      console.error(msg, e);
      throw e;
    }
  }

  async createOrganization(orgName: string, orgMeta: OrgMeta, account?: string) {
    try {
      const metaFile: string = await this.addJSONtoIPFS(orgMeta);
      const tx = await this.sendTransaction(
        this.valist.methods.createOrganization(
          orgName.toLowerCase().replace(shortnameFilterRegex, ''), metaFile,
        ),
        account,
      );
      return tx;
    } catch (e) {
      const msg = 'Could not create organization';
      console.error(msg, e);
      throw e;
    }
  }

  async createRepository(orgName: string, repoName: string, repoMeta: RepoMeta, account?: string) {
    try {
      const metaFile = await this.addJSONtoIPFS(repoMeta);
      const tx = await this.sendTransaction(
        this.valist.methods.createRepository(
          orgName, repoName.toLowerCase().replace(shortnameFilterRegex, ''), metaFile,
        ),
        account,
      );
      return tx;
    } catch (e) {
      const msg = 'Could not create repository';
      console.error(msg, e);
      throw e;
    }
  }

  async prepareRelease(tag: string, releaseFile: any, metaFile: any) {
    try {
      const releaseCID: string = await this.addFileToIPFS(releaseFile);
      const metaCID: string = await this.addFileToIPFS(metaFile);
      return { tag, releaseCID, metaCID };
    } catch (e) {
      const msg = 'Could not publish release';
      console.error(msg, e);
      throw e;
    }
  }

  async publishRelease(orgName: string, repoName: string, release: Release, account?: string) {
    try {
      const tx = await this.sendTransaction(
        this.valist.methods.publishRelease(
          orgName, repoName, release.tag, release.releaseCID, release.metaCID,
        ),
        account,
      );
      return tx;
    } catch (e) {
      const msg = 'Could not publish release';
      console.error(msg, e);
      throw e;
    }
  }

  // eslint-disable-next-line class-methods-use-this
  async fetchJSONfromIPFS(ipfsHash: string) {
    try {
      const response = await fetch(`https://gateway.valist.io/ipfs/${ipfsHash}`);
      const json = await response.json();
      return json;
    } catch (e) {
      const msg = 'Could not fetch JSON from IPFS';
      console.error(msg, e);
      throw e;
    }
  }

  async addJSONtoIPFS(data: any, onlyHash?: boolean) {
    try {
      const file = Buffer.from(JSON.stringify(data));
      const result = await this.ipfs.add(file, { onlyHash, cidVersion: 1 });
      return result.cid.string;
    } catch (e) {
      const msg = 'Could not add JSON to IPFS';
      console.error(msg, e);
      throw e;
    }
  }

  async addFileToIPFS(data: any, onlyHash?: boolean) {
    try {
      const result = await this.ipfs.add(data, { onlyHash, cidVersion: 1 });
      return result.cid.string;
    } catch (e) {
      const msg = 'Could not add file to IPFS';
      console.error(msg, e);
      throw e;
    }
  }

  async sendTransaction(functionCall: any, account: string = this.defaultAccount) {
    if (this.metaTxEnabled) {
      if (!this.metaTxReady) throw new Error('MetaTransactions not ready!');

      const sendAsync = (params: any): any => new Promise((resolve, reject) => {
        // @ts-ignore sendAsync is conflicting with the Magic RPCProvider type
        this.web3.currentProvider.sendAsync(params, async (e: any, signed: any) => {
          if (e) reject(e);
          resolve(signed);
        });
      });

      try {
        const nonce = await this.valist.methods.getNonce(account).call();
        const functionSignature = functionCall.encodeABI();

        const message = {
          nonce: Number(nonce),
          from: account,
          functionSignature,
        };

        const dataToSign = JSON.stringify({
          types: {
            EIP712Domain: domainType,
            MetaTransaction: metaTransactionType,
          },
          domain: domainData,
          primaryType: 'MetaTransaction',
          message,
        });

        let signed;

        if (this.signer) {
          signed = sigUtil.signTypedData_v4(Buffer.from(this.signer, 'hex'), { data: JSON.parse(dataToSign) });
        } else {
          const sig = await sendAsync({
            jsonrpc: '2.0',
            id: new Date().getTime(),
            method: 'eth_signTypedData_v4',
            params: [account, dataToSign],
          });
          signed = sig.result;
        }

        const { r, s, v } = getSignatureParameters(this.web3, signed);

        // console.log("R", r, "S", s, "V", v, "Function signature", functionSignature, account);

        const gasLimit = await this.valist.methods
          .executeMetaTransaction(account, functionSignature, r, s, v)
          .estimateGas({ from: account });

        const gasPrice = await this.web3.eth.getGasPrice();

        const tx = await this.valist.methods
          .executeMetaTransaction(account, functionSignature, r, s, v)
          .send({
            from: account,
            gasPrice,
            gasLimit,
          });
        return tx;
      } catch (e) {
        const msg = 'Could not send meta transaction';
        console.error(msg, e);
        throw e;
      }
    } else {
      const gasLimit = await functionCall.estimateGas({ from: account });
      const tx = await functionCall.send({ from: account, gasLimit });
      return tx;
    }
  }
}

export = Valist;
