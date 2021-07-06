/* eslint-disable @typescript-eslint/ban-ts-comment */
import Web3 from 'web3';
// @ts-ignore ipfs client types are finicky
import ipfsClient from 'ipfs-http-client';
// @ts-ignore mexa doesn't support typescript yet
import Biconomy from '@biconomy/mexa';

import { provider } from 'web3-core/types';

import * as sigUtil from 'eth-sig-util';

import {
  Organization,
  OrgID,
  OrgMeta,
  Release,
  RepoMeta,
  Repository,
  ValistCache,
} from './types';

import {
  ADD_KEY,
  REVOKE_KEY,
  ROTATE_KEY,
  ORG_ADMIN,
  REPO_DEV,
} from './constants';

import {
  domainData,
  domainType,
  getSignatureParameters,
  getValistContract,
  metaTransactionType,
  shortnameFilterRegex,
} from './utils';

// node-fetch polyfill
// eslint-disable-next-line @typescript-eslint/no-var-requires
const fetch = require('node-fetch');

if (!globalThis.fetch) {
  globalThis.fetch = fetch;
}
class Valist {
  web3: Web3;

  contract: any;

  ipfs: any;

  biconomy: any;

  signer?: string;

  defaultAccount: string;

  metaTxEnabled = false;

  metaTxReady = false;

  contractAddress: string | undefined;

  cache: ValistCache;

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
    this.cache = {
      orgIDs: [],
      orgNames: [],
      orgs: {},
    };
  }

  // initialize main valist contract instance for future calls
  async connect(waitForMetaTx?: boolean): Promise<void> {
    try {
      this.contract = await getValistContract(this.web3, this.contractAddress);
      // eslint-disable-next-line no-underscore-dangle
      this.contractAddress = this.contract._address;
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

  async getOrgIDFromName(orgName: string): Promise<OrgID> {
    try {
      if (!this.cache.orgs[orgName]) {
        this.cache.orgs[orgName] = {
          orgID: await this.contract.methods.orgIDByName(orgName).call(),
          threshold: 0,
          thresholdDate: 0,
          meta: {
            name: '',
            description: '',
          },
          metaCID: '',
          repoNames: [],
        };
      }
      return this.cache.orgs[orgName].orgID;
    } catch (e) {
      const msg = 'Could not get organization ID';
      console.error(msg, e);
      throw e;
    }
  }

  // returns organization meta and release tags
  async getOrganization(orgName: string, page = 1, resultsPerPage = 10): Promise<Organization> {
    try {
      const orgID: OrgID = await this.getOrgIDFromName(orgName);
      // @TODO add cache expiration on org metadata
      if (!this.cache.orgs[orgName].metaCID) {
        this.cache.orgs[orgName] = await this.contract.methods.orgs(orgID).call();
        this.cache.orgs[orgName].orgID = orgID;
      }

      this.cache.orgs[orgName].repoNames = await this.getRepoNames(orgName, page, resultsPerPage);

      try {
        this.cache.orgs[orgName].meta = await this.fetchJSONfromIPFS(this.cache.orgs[orgName].metaCID);
      } catch (e) {
        // noop, just return empty object if failed
      }

      return this.cache.orgs[orgName];
    } catch (e) {
      const msg = 'Could not get organization';
      console.error(msg, e);
      throw e;
    }
  }

  async getOrganizationNames(page = 1, resultsPerPage = 10): Promise<string[]> {
    try {
      const orgs = await this.contract.methods.getOrgNames(page, resultsPerPage).call();
      return orgs.filter(Boolean);
    } catch (e) {
      const msg = 'Could not get organization names';
      console.error(msg, e);
      throw e;
    }
  }

  async getRepoNames(orgName: string, page = 1, resultsPerPage = 10): Promise<string[]> {
    try {
      const orgID = await this.getOrgIDFromName(orgName);
      const repoNames = await this.contract.methods.getRepoNames(orgID, page, resultsPerPage).call();
      return repoNames.filter(Boolean);
    } catch (e) {
      const msg = 'Could not get organization names';
      console.error(msg, e);
      throw e;
    }
  }

  async setOrgMeta(orgName: string, orgMeta: OrgMeta, account: string = this.defaultAccount): Promise<any> {
    try {
      const orgID = await this.getOrgIDFromName(orgName);
      const hash = await this.addJSONtoIPFS(orgMeta);
      return await this.sendTransaction(this.contract.methods.setOrgMeta(orgID, hash), account);
    } catch (e) {
      const msg = 'Could not set organization metadata';
      console.error(msg, e);
      throw e;
    }
  }

  // returns repository
  async getRepository(orgName: string, repoName: string): Promise<Repository> {
    try {
      const orgID = await this.getOrgIDFromName(orgName);
      const repoSelector = this.web3.utils.keccak256(this.web3.utils.encodePacked(orgID, repoName) || '');
      const repo = await this.contract.methods.repos(repoSelector).call();

      let json: RepoMeta = {
        name: '',
        description: '',
        projectType: 'binary',
        homepage: '',
        repository: '',
      };

      try { json = await this.fetchJSONfromIPFS(repo.metaCID); } catch (e) {
        // noop, just return empty object if failed
      }

      return {
        orgID,
        meta: json,
        metaCID: repo.metaCID,
        tags: repo.tags,
        threshold: repo.threshold,
        thresholdDate: repo.thresholdDate,
      };
    } catch (e) {
      const msg = 'Could not get repository';
      console.error(msg, e);
      throw e;
    }
  }

  async setRepoMeta(
    orgName: string,
    repoName: string,
    repoMeta: RepoMeta,
    account: string = this.defaultAccount,
  ): Promise<any> {
    try {
      const orgID = await this.getOrgIDFromName(orgName);
      const hash = await this.addJSONtoIPFS(repoMeta);
      return await this.sendTransaction(this.contract.methods.setRepoMeta(orgID, repoName, hash), account);
    } catch (e) {
      const msg = 'Could not set repository metadata';
      console.error(msg, e);
      throw e;
    }
  }

  async getLatestRelease(orgName: string, repoName: string): Promise<Release> {
    try {
      const orgID = await this.getOrgIDFromName(orgName);
      const release = await this.contract.methods.getLatestRelease(orgID, repoName).call();
      return {
        tag: release[0],
        releaseCID: release[1],
        metaCID: release[2],
        signers: release[3],
      };
    } catch (e) {
      const msg = 'Could not get latest release from repo';
      console.error(msg, e);
      throw e;
    }
  }

  async getReleaseTags(orgName: string, repoName: string, page = 1, resultsPerPage = 10): Promise<string[]> {
    try {
      const orgID = await this.getOrgIDFromName(orgName);
      const repoSelector = this.web3.utils.keccak256(this.web3.utils.encodePacked(orgID, repoName) || '');
      const tags = await this.contract.methods.getReleaseTags(repoSelector, page, resultsPerPage).call();
      return tags.filter(Boolean);
    } catch (e) {
      const msg = 'Could not get release tags from repo';
      console.error(msg, e);
      throw e;
    }
  }

  async getReleases(orgName: string, repoName: string, page = 1, resultsPerPage = 10): Promise<Release[]> {
    try {
      const orgID = await this.getOrgIDFromName(orgName);
      const tags = await this.getReleaseTags(orgName, repoName, page, resultsPerPage);
      const releases: Release[] = [];

      for (let i = 0; i < tags.length; ++i) {
        const releaseSelector = this.web3.utils.keccak256(
          this.web3.eth.abi.encodeParameters(['bytes32', 'string', 'string'], [orgID, repoName, tags[i]]),
        );
        // eslint-disable-next-line no-await-in-loop
        const release = await this.contract.methods.releases(releaseSelector).call();
        releases.push({
          tag: tags[i],
          releaseCID: release.releaseCID,
          metaCID: release.metaCID,
        });
      }

      return releases;
    } catch (e) {
      const msg = 'Could not get releases from repo';
      console.error(msg, e);
      throw e;
    }
  }

  async getReleaseByTag(orgName: string, repoName: string, tag: string): Promise<Release> {
    try {
      const orgID = await this.getOrgIDFromName(orgName);
      const releaseSelector = this.web3.utils.keccak256(
        this.web3.eth.abi.encodeParameters(['bytes32', 'string', 'string'], [orgID, repoName, tag]),
      );
      const release = await this.contract.methods.releases(releaseSelector).call();
      return {
        tag,
        releaseCID: release.releaseCID,
        metaCID: release.metaCID,
        signers: release.signers,
      };
    } catch (e) {
      const msg = 'Could not get release by tag';
      console.error(msg, e);
      throw e;
    }
  }

  async isOrgAdmin(orgName: string, account: string = this.defaultAccount): Promise<boolean> {
    try {
      const orgID = await this.getOrgIDFromName(orgName);
      const isOrgAdmin = await this.contract.methods.isOrgAdmin(orgID, account).call();
      return isOrgAdmin;
    } catch (e) {
      const msg = 'Could not check if user has ORG_ADMIN_ROLE';
      console.error(msg, e);
      throw e;
    }
  }

  async isRepoDev(orgName: string, repoName: string, account: string = this.defaultAccount): Promise<boolean> {
    try {
      const orgID = await this.getOrgIDFromName(orgName);
      const isRepoDev = await this.contract.methods.isRepoDev(orgID, repoName, account).call();
      return isRepoDev;
    } catch (e) {
      const msg = 'Could not check if user has REPO_DEV_ROLE';
      console.error(msg, e);
      throw e;
    }
  }

  async voteOrgAdmin(orgName: string, key: string): Promise<any> {
    try {
      const orgID = await this.getOrgIDFromName(orgName);
      const tx = await this.sendTransaction(
        this.contract.methods.voteKey(orgID, '', ADD_KEY, key),
        this.defaultAccount,
      );
      return tx;
    } catch (e) {
      const msg = 'Could not vote to grant ORG_ADMIN_ROLE';
      console.error(msg, e);
      throw e;
    }
  }

  async revokeOrgAdmin(orgName: string, key: string): Promise<any> {
    try {
      const orgID = await this.getOrgIDFromName(orgName);
      const tx = await this.sendTransaction(
        this.contract.methods.voteKey(orgID, '', REVOKE_KEY, key),
        this.defaultAccount,
      );
      return tx;
    } catch (e) {
      const msg = 'Could not vote to revoke ORG_ADMIN_ROLE';
      console.error(msg, e);
      throw e;
    }
  }

  async rotateOrgAdmin(orgName: string, key: string): Promise<any> {
    try {
      const orgID = await this.getOrgIDFromName(orgName);
      const tx = await this.sendTransaction(
        this.contract.methods.voteKey(orgID, '', ROTATE_KEY, key),
        this.defaultAccount,
      );
      return tx;
    } catch (e) {
      const msg = 'Could not vote to revoke ORG_ADMIN_ROLE';
      console.error(msg, e);
      throw e;
    }
  }

  async voteRepoDev(orgName: string, repoName: string, key: string): Promise<any> {
    try {
      const orgID = await this.getOrgIDFromName(orgName);
      const tx = await this.sendTransaction(
        this.contract.methods.voteKey(orgID, repoName, ADD_KEY, key),
        this.defaultAccount,
      );
      return tx;
    } catch (e) {
      const msg = 'Could not vote to grant REPO_DEV_ROLE';
      console.error(msg, e);
      throw e;
    }
  }

  async revokeRepoDev(orgName: string, repoName: string, key: string): Promise<any> {
    try {
      const orgID = await this.getOrgIDFromName(orgName);
      const tx = await this.sendTransaction(
        this.contract.methods.voteKey(orgID, repoName, REVOKE_KEY, key),
        this.defaultAccount,
      );
      return tx;
    } catch (e) {
      const msg = 'Could not vote to revoke REPO_DEV_ROLE';
      console.error(msg, e);
      throw e;
    }
  }

  async rotateRepoDev(orgName: string, repoName: string, key: string): Promise<any> {
    try {
      const orgID = await this.getOrgIDFromName(orgName);
      const tx = await this.sendTransaction(
        this.contract.methods.voteKey(orgID, repoName, ROTATE_KEY, key),
        this.defaultAccount,
      );
      return tx;
    } catch (e) {
      const msg = 'Could not vote to revoke REPO_DEV_ROLE';
      console.error(msg, e);
      throw e;
    }
  }

  async getOrgAdmins(orgName: string): Promise<string[]> {
    try {
      const orgID = await this.getOrgIDFromName(orgName);
      const roleSelector = this.web3.utils.keccak256(this.web3.utils.encodePacked(orgID, ORG_ADMIN) || '');
      const members = await this.contract.methods.getRoleMembers(roleSelector).call();
      return members;
    } catch (e) {
      const msg = 'Could not get users that have ORG_ADMIN_ROLE';
      console.error(msg, e);
      throw e;
    }
  }

  async getRepoDevs(orgName: string, repoName: string): Promise<string[]> {
    try {
      const orgID = await this.getOrgIDFromName(orgName);
      const roleSelector = this.web3.utils.keccak256(this.web3.utils.encodePacked(orgID, repoName, REPO_DEV) || '');
      const members = await this.contract.methods.getRoleMembers(roleSelector).call();
      return members;
    } catch (e) {
      const msg = 'Could not get users that have REPO_DEV_ROLE';
      console.error(msg, e);
      throw e;
    }
  }

  async createOrganization(orgName: string, orgMeta: OrgMeta, account: string = this.defaultAccount): Promise<any> {
    try {
      const metaFile: string = await this.addJSONtoIPFS(orgMeta);
      const tx = await this.sendTransaction(
        this.contract.methods.createOrganization(
          orgName.toLowerCase().replace(shortnameFilterRegex, ''),
          metaFile,
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

  async createRepository(
    orgName: string,
    repoName: string,
    repoMeta: RepoMeta,
    account: string = this.defaultAccount,
  ): Promise<any> {
    try {
      const metaFile = await this.addJSONtoIPFS(repoMeta);
      const orgID = await this.getOrgIDFromName(orgName);
      const tx = await this.sendTransaction(
        this.contract.methods.createRepository(
          orgID, repoName.toLowerCase().replace(shortnameFilterRegex, ''),
          metaFile,
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

  async prepareRelease(tag: string, releaseFile: any, metaFile: any): Promise<Release> {
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

  async publishRelease(
    orgName: string,
    repoName: string,
    release: Release,
    account: string = this.defaultAccount,
  ): Promise<any> {
    try {
      const orgID = await this.getOrgIDFromName(orgName);
      const isRepoDev = await this.contract.methods.isRepoDev(orgID, repoName, account).call();
      if (!isRepoDev) throw new Error('User does not have permission to publish release');

      const tx = await this.sendTransaction(
        this.contract.methods.voteRelease(
          orgID, repoName, release.tag, release.releaseCID, release.metaCID,
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
  async fetchJSONfromIPFS(ipfsHash: string): Promise<any> {
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

  async addJSONtoIPFS(data: any, onlyHash?: boolean): Promise<string> {
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

  async addFileToIPFS(data: any, onlyHash?: boolean): Promise<string> {
    try {
      const result = await this.ipfs.add(data, { onlyHash, cidVersion: 1 });
      return result.cid.string;
    } catch (e) {
      const msg = 'Could not add file to IPFS';
      console.error(msg, e);
      throw e;
    }
  }

  async sendTransaction(functionCall: any, account: string = this.defaultAccount): Promise<any> {
    if (this.metaTxEnabled) {
      if (!this.metaTxReady) throw new Error('MetaTransactions not ready!');

      try {
        let tx;

        if (this.signer) {
          const txParams = {
            from: account,
            data: functionCall.encodeABI(),
            gasLimit: this.web3.utils.toHex(await functionCall.estimateGas({ from: account })),
          };
          const signedTx = await this.web3.eth.accounts.signTransaction(txParams, `0x${this.signer}`);
          const forwardData = await this.biconomy.getForwardRequestAndMessageToSign(signedTx.rawTransaction);
          const signature = sigUtil.signTypedData_v4(
            Buffer.from(this.signer, 'hex'),
            { data: forwardData.eip712Format },
          );
          const { rawTransaction } = signedTx;
          const data = {
            signature,
            forwardRequest: forwardData.request,
            rawTransaction,
            signatureType: this.biconomy.EIP712_SIGN,
          };
          // @ts-ignore
          tx = await this.web3.eth.sendSignedTransaction(data);
        } else {
          tx = await functionCall.send({
            from: account,
            signatureType: this.biconomy.EIP712_SIGN,
          });
        }

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
