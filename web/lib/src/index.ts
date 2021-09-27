/* eslint-disable @typescript-eslint/ban-ts-comment */
/* eslint-disable no-underscore-dangle */
import Web3 from 'web3';
// @ts-ignore ipfs client types are finicky
import ipfsClient from 'ipfs-http-client';

import { provider, EventLog } from 'web3-core/types';

import {
  Organization,
  OrgID,
  OrgMeta,
  PendingRelease,
  PendingVote,
  Release,
  RepoMeta,
  Repository,
  ValistCache,
  ValistConfig,
} from './types';

import {
  ADD_KEY,
  REVOKE_KEY,
  ROTATE_KEY,
  ORG_ADMIN,
  REPO_DEV,
} from './constants';

import {
  getValistContract,
  shortnameFilterRegex,
  sendMetaTransaction,
  getValistRegistry,
  stripParentFolderFromPath,
  parseCID,
} from './utils';

import { ValistSDKError } from './errors';

// node-fetch polyfill
// eslint-disable-next-line @typescript-eslint/no-var-requires
const fetch = require('node-fetch');

if (!globalThis.fetch) {
  globalThis.fetch = fetch;
}
class Valist {
  web3: Web3;

  contract: any;

  registry: any;

  ipfs: any;

  defaultAccount: string;

  metaTxEnabled = false;

  gatewayHost?: string;

  signer?: string;

  contractAddress?: string;

  registryAddress?: string;

  chainID?: number;

  cache: ValistCache;

  constructor({
    web3Provider,
    metaTx = true,
    ipfsHost = 'https://pin.valist.io',
    gatewayHost = 'https://gateway.valist.io',
    contractAddress,
    registryAddress,
  }: {
    web3Provider: provider,
    metaTx?: boolean,
    ipfsHost?: string,
    gatewayHost?: string,
    contractAddress?: string,
    registryAddress?: string,
  }) {
    this.web3 = new Web3(web3Provider);
    this.metaTxEnabled = metaTx;
    this.defaultAccount = '0x0';
    this.ipfs = ipfsClient(ipfsHost);
    this.gatewayHost = gatewayHost;
    this.contractAddress = contractAddress;
    this.registryAddress = registryAddress;
    this.cache = {
      orgIDs: {},
    };
  }

  // initialize main valist contract instance for future calls
  async connect(): Promise<void> {
    this.chainID = await this.web3.eth.getChainId();
    try {
      this.contract = await getValistContract(this.web3, this.chainID, this.contractAddress);
      this.contractAddress = this.contract._address;
      if (!this.contractAddress) throw new ValistSDKError('Could not get Valist contract address');
    } catch (e) {
      const msg = 'Could not connect to Valist registry contract';
      console.error(msg, e);
      throw e;
    }

    try {
      this.registry = await getValistRegistry(this.web3, this.chainID, this.registryAddress);
      this.registryAddress = this.registry._address;
      if (!this.registryAddress) throw new ValistSDKError('Could not get Valist registry address');
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

    if (this.metaTxEnabled) console.log('👻 MetaTransactions Enabled');
  }

  async createOrganization(orgName: string, orgMeta: OrgMeta): Promise<any> {
    try {
      await this.getOrgIDFromName(orgName);
      throw new ValistSDKError('Namespace already taken, please choose another name');
    } catch (e) {
      if (e.message !== 'orgID not found') {
        throw e;
      }
    }
    try {
      const metaFile: string = await this.addJSONtoIPFS(orgMeta);
      const tx = await this.sendTransaction(this.contract.methods.createOrganization(metaFile));
      let orgID: string;
      try {
        // eslint-disable-next-line prefer-destructuring
        orgID = this.metaTxEnabled ? tx.logs[0].topics[1] : tx.events.OrgCreated.returnValues._orgID;
      } catch (e) {
        throw new ValistSDKError('Could not parse new orgID from transaction');
      }
      return { ...tx, orgID };
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
      const orgID = await this.getOrgIDFromName(orgName);
      const metaFile = await this.addJSONtoIPFS(repoMeta);
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

  async prepareRelease(
    config: ValistConfig,
    releaseFiles: any,
    metaFile: any,
    parentFolder = '',
  ): Promise<Release> {
    try {
      let releaseCID: string;

      if (releaseFiles.length > 1) {
        releaseCID = await this.addFolderToIPFS(releaseFiles, parentFolder);
      } else {
        releaseCID = await this.addFileToIPFS(releaseFiles[0]);
      }

      const metaCID: string = await this.addFileToIPFS(metaFile);
      const { tag } = config;
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
      const isRepoDev = await this.isRepoDev(orgName, repoName, account);
      if (!isRepoDev) throw new ValistSDKError('User does not have permission to publish release');

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

  async repoHasReleased(orgName: string, repoName: string):Promise<boolean> {
    try {
      await this.getLatestRelease(orgName, repoName);
      return true;
    } catch (e) {
      return false;
    }
  }

  async linkNameToID(name: string, orgID: string): Promise<any> {
    try {
      await this.getOrgIDFromName(name);
      throw new ValistSDKError('Namespace already taken, please choose another name');
    } catch (e) {
      if (e.message !== 'orgID not found') {
        throw e;
      }
    }
    try {
      const tx = await this.sendTransaction(
        this.registry.methods.linkNameToID(orgID, name.toLowerCase().replace(shortnameFilterRegex, '')),
      );
      return tx;
    } catch (e) {
      const msg = 'Could not link namespace to orgID';
      console.error(msg, e);
      throw e;
    }
  }

  async setOrgMeta(orgName: string, orgMeta: OrgMeta, account: string = this.defaultAccount): Promise<any> {
    try {
      if (!orgMeta.name) throw new Error('orgMeta.name not found');
      if (!orgMeta.description) throw new Error('orgMeta.description not found');
      const orgID = await this.getOrgIDFromName(orgName);
      const hash = await this.addJSONtoIPFS(orgMeta);
      return await this.sendTransaction(this.contract.methods.setOrgMeta(orgID, hash), account);
    } catch (e) {
      const msg = 'Could not set organization metadata';
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
      if (!repoMeta.name) throw new Error('repoMeta.name is empty');
      if (!repoMeta.description) throw new Error('repoMeta.description is empty');
      if (!repoMeta.projectType) throw new Error('repoMeta.projectType is empty');
      const orgID = await this.getOrgIDFromName(orgName);
      const hash = await this.addJSONtoIPFS(repoMeta);
      const tx = await this.sendTransaction(this.contract.methods.setRepoMeta(orgID, repoName, hash), account);
      return tx;
    } catch (e) {
      const msg = 'Could not set repository metadata';
      console.error(msg, e);
      throw e;
    }
  }

  async getOrgIDFromName(orgName: string): Promise<OrgID> {
    const name = orgName.toLowerCase().replace(shortnameFilterRegex, '');
    if (!this.cache.orgIDs[name]) {
      this.cache.orgIDs[name] = await this.registry.methods.nameToID(name).call();
    }
    if (this.cache.orgIDs[name] === '0x0000000000000000000000000000000000000000000000000000000000000000') {
      this.cache.orgIDs[name] = '';
      throw new ValistSDKError('orgID not found');
    }
    return this.cache.orgIDs[name];
  }

  // returns organization meta and release tags
  async getOrganization(orgName: string, page = 1, resultsPerPage = 10): Promise<Organization> {
    try {
      const orgID: OrgID = await this.getOrgIDFromName(orgName);
      const org = await this.contract.methods.orgs(orgID).call();

      const repoNames = await this.getRepoNames(orgName, page, resultsPerPage);

      let meta = { name: '', description: '' };
      try {
        meta = await this.fetchJSONfromIPFS(org.metaCID);
      } catch (e) {
        // noop, just return empty object if failed
      }

      return {
        orgID,
        threshold: org.threshold,
        thresholdDate: org.thresholdDate,
        meta,
        metaCID: org.metaCID,
        repoNames,
      };
    } catch (e) {
      const msg = 'Could not get organization';
      console.error(msg, e);
      throw e;
    }
  }

  async getOrganizationNames(page = 1, resultsPerPage = 10): Promise<string[]> {
    try {
      const orgs = await this.registry.methods.getNames(page, resultsPerPage).call();
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
        threshold: Number(repo.threshold),
        thresholdDate: repo.thresholdDate,
      };
    } catch (e) {
      const msg = 'Could not get repository';
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
          tag: tags[i] || '',
          releaseCID: release.releaseCID || '',
          metaCID: release.metaCID || '',
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
        releaseCID: release.releaseCID  || '',
        metaCID: release.metaCID  || '',
        signers: release.signers  || '',
      };
    } catch (e) {
      const msg = 'Could not get release by tag';
      console.error(msg, e);
      throw e;
    }
  }

  async getPendingReleases(orgName: string, repoName: string): Promise<PendingRelease[]> {
    try {
      const orgID = await this.getOrgIDFromName(orgName);
      const repoSelector = this.web3.utils.keccak256(this.web3.utils.encodePacked(orgID, repoName) || '');
      const pendingCount = await this.contract.methods.getPendingReleaseCount(repoSelector).call();
      const requests: PendingRelease[] = [];
      for (let i = 0; i < pendingCount; ++i) {
        // eslint-disable-next-line no-await-in-loop
        requests.push(await this.contract.methods.pendingReleaseRequests(repoSelector, i).call());
      }
      return requests;
    } catch (e) {
      const msg = 'Could not get pending releases';
      console.error(msg, e);
      throw e;
    }
  }

  async getPendingReleaseVotes(orgName: string, repoName: string, release: Release): Promise<PendingVote> {
    try {
      const orgID = await this.getOrgIDFromName(orgName);
      const voteSelector = this.web3.utils.keccak256(
        this.web3.eth.abi.encodeParameters(
          ['bytes32', 'string', 'string', 'string', 'string'],
          [orgID, repoName, release.tag, release.releaseCID, release.metaCID],
        ),
      );
      const votes = await this.getPendingVotes(voteSelector);
      return votes;
    } catch (e) {
      const msg = 'Could not get pending release votes';
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
      if (e.message.includes('orgID not found')) {
        return false;
      }
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
      if (e.message.includes('No repo') || e.message.includes('orgID not found')) {
        return false;
      }
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
      const msg = 'Could not vote to rotate ORG_ADMIN_ROLE';
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

  private async getPendingVotes(voteSelector: string): Promise<PendingVote> {
    try {
      const votes = await this.contract.methods.getPendingVotes(voteSelector).call();
      return {
        expiration: votes[0],
        signers: votes[1],
      };
    } catch (e) {
      const msg = 'Could not get pending votes';
      console.error(msg, e);
      throw e;
    }
  }

  private async getEvents(topics?: any[]): Promise<EventLog[]> {
    try {
      const fromBlock = await this.web3.eth.getBlockNumber() - 99990;
      const options = { fromBlock, toBlock: 'latest', topics };
      const events = await this.contract.getPastEvents('allEvents', options);
      return events;
    } catch (e) {
      const msg = 'Could not get allEvents';
      console.error(msg, e);
      throw e;
    }
  }

  async getOrgEvents(orgName: string): Promise<EventLog[]> {
    const orgID = await this.getOrgIDFromName(orgName);
    // this should be the empty repo keccak256
    const repoTopic = '0xc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470';
    const eventLogs = await this.getEvents([null, orgID, repoTopic]);
    return eventLogs.reverse();
  }

  async getRepoEvents(orgName: string, repoName: string): Promise<EventLog[]> {
    const orgID = await this.getOrgIDFromName(orgName);
    const repoTopic = this.web3.utils.keccak256(repoName);
    const eventLogs = await this.getEvents([null, orgID, repoTopic]);
    return eventLogs.reverse();
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

  async getPendingOrgAdmins(orgName: string): Promise<string[]> {
    try {
      const orgID = await this.getOrgIDFromName(orgName);
      const pendingCount = await this.contract.methods.getRoleRequestCount(orgID).call();
      const requests = [];
      for (let i = 0; i < pendingCount; ++i) {
        // eslint-disable-next-line no-await-in-loop
        requests.push(await this.contract.methods.pendingRoleRequests(orgID, i).call());
      }
      return requests;
    } catch (e) {
      const msg = 'Could not get pending orgAdmin requests';
      console.error(msg, e);
      throw e;
    }
  }

  async getPendingOrgAdminVotes(orgName: string, operation: string, address: string): Promise<PendingVote> {
    if (![ADD_KEY, REVOKE_KEY].includes(operation)) {
      throw new ValistSDKError(`Invalid key operation ${operation}`);
    }
    try {
      const orgID = await this.getOrgIDFromName(orgName);
      // keccak256(abi.encodePacked(orgID, ORG_ADMIN, OPERATION, pendingOrgAdminAddress))
      const voteSelector = this.web3.utils.keccak256(
        this.web3.utils.encodePacked(orgID, ORG_ADMIN, operation, address) || '',
      );
      const votes = await this.getPendingVotes(voteSelector);
      return votes;
    } catch (e) {
      const msg = 'Could not get pending orgAdmin votes';
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

  async getPendingRepoDevs(orgName: string, repoName: string): Promise<string[]> {
    try {
      const orgID = await this.getOrgIDFromName(orgName);
      const roleSelector = this.web3.utils.keccak256(this.web3.utils.encodePacked(orgID, repoName) || '');
      const pendingCount = await this.contract.methods.getRoleRequestCount(roleSelector).call();
      const requests = [];
      for (let i = 0; i < pendingCount; ++i) {
        // eslint-disable-next-line no-await-in-loop
        requests.push(await this.contract.methods.pendingRoleRequests(roleSelector, i).call());
      }
      return requests;
    } catch (e) {
      const msg = 'Could not get users that have REPO_DEV_ROLE';
      console.error(msg, e);
      throw e;
    }
  }

  async getPendingRepoDevVotes(
    orgName: string,
    repoName: string,
    operation: string,
    address: string,
  ): Promise<PendingVote> {
    if (![ADD_KEY, REVOKE_KEY].includes(operation)) {
      throw new ValistSDKError(`Invalid key operation ${operation}`);
    }
    try {
      const orgID = await this.getOrgIDFromName(orgName);
      // keccak256(abi.encodePacked(orgID, repoName, REPO_DEV, OPERATION, pendingRepoDevAddress))
      const voteSelector = this.web3.utils.keccak256(
        this.web3.utils.encodePacked(orgID, repoName, REPO_DEV, operation, address) || '',
      );
      const votes = await this.getPendingVotes(voteSelector);
      return votes;
    } catch (e) {
      const msg = 'Could not get pending repoDev votes';
      console.error(msg, e);
      throw e;
    }
  }

  async voteOrgThreshold(orgName: string, threshold: number): Promise<any> {
    try {
      const orgID = await this.getOrgIDFromName(orgName);
      const tx = await this.sendTransaction(
        this.contract.methods.voteThreshold(orgID, '', threshold),
        this.defaultAccount,
      );
      return tx;
    } catch (e) {
      const msg = 'Could not vote for org threshold';
      console.error(msg, e);
      throw e;
    }
  }

  async getPendingOrgThresholds(orgName: string): Promise<string[]> {
    try {
      const orgID = await this.getOrgIDFromName(orgName);
      const pendingCount = await this.contract.methods.getThresholdRequestCount(orgID).call();
      const requests = [];
      for (let i = 0; i < pendingCount; ++i) {
        // eslint-disable-next-line no-await-in-loop
        requests.push(await this.contract.methods.pendingThresholdRequests(orgID, i).call());
      }
      return requests;
    } catch (e) {
      const msg = 'Could not get pending org thresholds';
      console.error(msg, e);
      throw e;
    }
  }

  async getPendingOrgThresholdVotes(orgName: string, threshold: number): Promise<PendingVote> {
    try {
      const orgID = await this.getOrgIDFromName(orgName);
      // keccak256(abi.encodePacked(orgID, repoName, pendingRepoThreshold))
      const voteSelector = this.web3.utils.keccak256(
        this.web3.utils.encodePacked(orgID, '', threshold) || '',
      );
      const votes = await this.getPendingVotes(voteSelector);
      return votes;
    } catch (e) {
      const msg = 'Could not get pending repo threshold votes';
      console.error(msg, e);
      throw e;
    }
  }

  async voteRepoThreshold(orgName: string, repoName: string, threshold: number): Promise<any> {
    try {
      const orgID = await this.getOrgIDFromName(orgName);
      const tx = await this.sendTransaction(
        this.contract.methods.voteThreshold(orgID, repoName, threshold),
        this.defaultAccount,
      );
      return tx;
    } catch (e) {
      const msg = 'Could not vote for repo threshold';
      console.error(msg, e);
      throw e;
    }
  }

  async getPendingRepoThresholds(orgName: string, repoName: string): Promise<string[]> {
    try {
      const orgID = await this.getOrgIDFromName(orgName);
      const repoSelector = this.web3.utils.keccak256(this.web3.utils.encodePacked(orgID, repoName) || '');
      const pendingCount = await this.contract.methods.getThresholdRequestCount(repoSelector).call();
      const requests = [];
      for (let i = 0; i < pendingCount; ++i) {
        // eslint-disable-next-line no-await-in-loop
        requests.push(await this.contract.methods.pendingThresholdRequests(repoSelector, i).call());
      }
      return requests;
    } catch (e) {
      const msg = 'Could not get pending repo thresholds';
      console.error(msg, e);
      throw e;
    }
  }

  async getPendingRepoThresholdVotes(orgName: string, repoName: string, threshold: number): Promise<PendingVote> {
    try {
      const orgID = await this.getOrgIDFromName(orgName);
      // keccak256(abi.encodePacked(orgID, repoName, pendingRepoThreshold))
      const voteSelector = this.web3.utils.keccak256(
        this.web3.utils.encodePacked(orgID, repoName, threshold) || '',
      );
      const votes = await this.getPendingVotes(voteSelector);
      return votes;
    } catch (e) {
      const msg = 'Could not get pending repo threshold votes';
      console.error(msg, e);
      throw e;
    }
  }

  async clearPendingOrgKey(orgName: string, operation: string, key: string, index: number): Promise<any> {
    if (![ADD_KEY, REVOKE_KEY].includes(operation)) {
      throw new ValistSDKError(`Invalid key operation ${operation}`);
    }
    try {
      const orgID = await this.getOrgIDFromName(orgName);
      const tx = await this.sendTransaction(
        this.contract.methods.clearPendingKey(orgID, '', operation, key, index),
        this.defaultAccount,
      );
      return tx;
    } catch (e) {
      const msg = 'Could not clear pending role request';
      console.error(msg, e);
      throw e;
    }
  }

  async clearPendingRepoKey(
    orgName: string,
    repoName: string,
    operation: string,
    key: string,
    index: number,
  ): Promise<any> {
    if (![ADD_KEY, REVOKE_KEY].includes(operation)) {
      throw new ValistSDKError(`Invalid key operation ${operation}`);
    }
    try {
      const orgID = await this.getOrgIDFromName(orgName);
      const tx = await this.sendTransaction(
        this.contract.methods.clearPendingKey(orgID, repoName, operation, key, index),
        this.defaultAccount,
      );
      return tx;
    } catch (e) {
      const msg = 'Could not clear pending role request';
      console.error(msg, e);
      throw e;
    }
  }

  async clearPendingOrgThreshold(orgName: string, threshold: number, index: number): Promise<any> {
    try {
      const orgID = await this.getOrgIDFromName(orgName);
      const tx = await this.sendTransaction(
        this.contract.methods.clearPendingThreshold(orgID, '', threshold, index),
        this.defaultAccount,
      );
      return tx;
    } catch (e) {
      const msg = 'Could not clear pending org threshold';
      console.error(msg, e);
      throw e;
    }
  }

  async clearPendingRepoThreshold(orgName: string, repoName: string, threshold: number, index: number): Promise<any> {
    try {
      const orgID = await this.getOrgIDFromName(orgName);
      const tx = await this.sendTransaction(
        this.contract.methods.clearPendingThreshold(orgID, repoName, threshold, index),
        this.defaultAccount,
      );
      return tx;
    } catch (e) {
      const msg = 'Could not clear pending repo threshold';
      console.error(msg, e);
      throw e;
    }
  }

  async clearPendingRelease(orgName: string, repoName: string, release: Release, index: number): Promise<any> {
    try {
      const orgID = await this.getOrgIDFromName(orgName);
      const tx = await this.sendTransaction(
        this.contract.methods.clearPendingRelease(
          orgID,
          repoName,
          release.tag,
          release.releaseCID,
          release.metaCID,
          index,
        ),
        this.defaultAccount,
      );
      return tx;
    } catch (e) {
      const msg = 'Could not clear pending release';
      console.error(msg, e);
      throw e;
    }
  }

  // eslint-disable-next-line class-methods-use-this
  async fetchJSONfromIPFS(ipfsHash: string): Promise<any> {
    try {
      const response = await fetch(`${this.gatewayHost}/ipfs/${parseCID(ipfsHash)}`);
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

  async addFolderToIPFS(files: any[], parent = ''): Promise<string> {
    try {
      const fileObjects = files.map((file) => {
        const fileObject = {
          path: stripParentFolderFromPath(file.path, parent),
          content: file,
        };
        return fileObject;
      });

      let cid;
      // eslint-disable-next-line no-restricted-syntax
      for await (const result of this.ipfs.addAll(fileObjects, { wrapWithDirectory: true, cidVersion: 1 })) {
        cid = result.cid.string;
      }

      return cid;
    } catch (e) {
      const msg = 'Could not add file to IPFS';
      console.error(msg, e);
      throw e;
    }
  }

  async sendTransaction(functionCall: any, account: string = this.defaultAccount): Promise<any> {
    const gasLimit = await functionCall.estimateGas({ from: account });

    if (this.metaTxEnabled) {
      try {
        const tx = await sendMetaTransaction(
          this.web3,
          this.chainID as number,
          functionCall,
          account,
          gasLimit,
          this.signer,
        );
        return tx;
      } catch (e) {
        const msg = 'Could not send meta transaction';
        console.error(msg, e);
        throw e;
      }
    } else {
      const tx = await functionCall.send({ from: account, gasLimit });
      return tx;
    }
  }
}

export = Valist;
