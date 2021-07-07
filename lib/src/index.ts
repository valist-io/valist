/* eslint-disable @typescript-eslint/ban-ts-comment */
import Web3 from 'web3';
// @ts-ignore ipfs client types are finicky
import ipfsClient from 'ipfs-http-client';

import { provider, EventLog } from 'web3-core/types';

import * as sigUtil from 'eth-sig-util';

import {
  Biconomy,
  // @ts-ignore mexa doesn't support typescript yet
} from '@biconomy/mexa';

import {
  getDomainSeperator,
  getDataToSignForEIP712,
  buildForwardTxRequest,
  getBiconomyForwarderConfig,
} from './helpers/biconomy';

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
  VoteKeyEvent,
  VoteThresholdEvent,
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
      console.log(orgID, tx);
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

  private async getEvents(event: string, filter?: any): Promise<EventLog[]> {
    try {
      const events = this.contract.getPastEvents(event, {
        fromBlock: await this.web3.eth.getBlockNumber() - 99999,
        toBlock: 'latest',
        filter,
      });
      return events;
    } catch (e) {
      const msg = `Could not get ${event}`;
      console.error(msg, e);
      throw e;
    }
  }

  async getOrgVoteKeyEvents(orgName: string): Promise<VoteKeyEvent[]> {
    const orgID = await this.getOrgIDFromName(orgName);
    const pending = await this.getPendingOrgAdmins(orgName);

    // TODO what is the empty value for _repoName
    const eventFilter = { _key: pending, _orgID: orgID };
    const eventLogs = await this.getEvents('VoteKeyEvent', eventFilter);

    // get unique votes by key and operation
    const unique = new Map<string, VoteKeyEvent>();
    for (let i = 0; i < eventLogs.length; i++) {
      const voteEvent: VoteKeyEvent = eventLogs[i].returnValues;
      // eslint-disable-next-line no-underscore-dangle
      const voteID = `${voteEvent._key}${voteEvent._operation}`;
      unique.set(voteID, voteEvent);
    }

    return Array.from(unique.values());
  }

  async getOrgVoteThresholdEvent(orgName: string): Promise<VoteThresholdEvent[]> {
    const orgID = await this.getOrgIDFromName(orgName);
    const pending = await this.getPendingOrgThresholds(orgName);

    // TODO what is the empty value for _repoName
    const eventFilter = { _pendingThreshold: pending, _orgID: orgID };
    const eventLogs = await this.getEvents('VoteThresholdEvent', eventFilter);

    // get unique votes by pending threshold
    const unique = new Map<string, VoteThresholdEvent>();
    for (let i = 0; i < eventLogs.length; i++) {
      const voteEvent: VoteThresholdEvent = eventLogs[i].returnValues;
      // eslint-disable-next-line no-underscore-dangle
      const voteID = `${voteEvent._pendingThreshold}`;
      unique.set(voteID, voteEvent);
    }

    return Array.from(unique.values());
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
      throw new Error(`Invalid key operation ${operation}`);
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
      throw new Error(`Invalid key operation ${operation}`);
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

  async getPendingOrgThresholds(orgName: string): Promise<number[]> {
    try {
      const orgID = await this.getOrgIDFromName(orgName);
      const pendingCount = await this.contract.methods.getThresholdRequestCount(orgID).call();
      const requests = [];
      for (let i = 0; i < pendingCount; ++i) {
        // eslint-disable-next-line no-await-in-loop
        requests.push(Number(await this.contract.methods.pendingThresholdRequests(orgID, i).call()));
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

  async getPendingRepoThresholds(orgName: string, repoName: string): Promise<number[]> {
    try {
      const orgID = await this.getOrgIDFromName(orgName);
      const repoSelector = this.web3.utils.keccak256(this.web3.utils.encodePacked(orgID, repoName) || '');
      const pendingCount = await this.contract.methods.getThresholdRequestCount(repoSelector).call();
      const requests = [];
      for (let i = 0; i < pendingCount; ++i) {
        // eslint-disable-next-line no-await-in-loop
        requests.push(Number(await this.contract.methods.pendingThresholdRequests(repoSelector, i).call()));
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
      throw new Error(`Invalid key operation ${operation}`);
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
      throw new Error(`Invalid key operation ${operation}`);
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
    const gasLimit = await functionCall.estimateGas({ from: account });

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
        const networkID = await this.web3.eth.net.getId();
        const forwarder = await getBiconomyForwarderConfig(networkID);
        const forwarderContract = new this.web3.eth.Contract(forwarder.abi, forwarder.address);
        const batchNonce = await forwarderContract.methods.getNonce(account, 0).call();
        const to = this.contractAddress;
        const functionSignature = functionCall.encodeABI();

        const request = await buildForwardTxRequest({
          account,
          to,
          gasLimitNum: gasLimit,
          batchId: 0,
          batchNonce,
          data: functionSignature,
          deadline: '',
        });

        const domainSeparator = await getDomainSeperator(this.web3, networkID);
        const dataToSign = await getDataToSignForEIP712(this.web3, request, networkID);
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

        // eslint-disable-next-line no-underscore-dangle
        const functionName: string = functionCall._method.name;
        const functionIDMap: Record<string, string> = {
          createOrganization: '27bad85d-6821-43dd-aa64-924a327ef6bc',
          createRepository: 'f09f84a4-0364-4a1e-8011-deea92527823',
          voteRelease: '2b11f0d2-9eee-48cd-ba92-ad4083930142',
          voteKey: '838a8e89-3a92-481d-a708-87f62459f02e',
          voteThreshold: '8e930fd0-d781-4c2a-a7da-db961348dda3',
          setOrgMeta: 'f8b42e31-cf5c-4e16-be73-f2f94467641f',
          setRepoMeta: 'ec68cb11-e7a6-4bbf-966d-065c2a0a233c',
          clearPendingRelease: '24d70971-c3b5-4728-ad2c-310a9c717b31',
          clearPendingKey: '8e6b6d8d-123a-42ef-8751-5aea5a5d6e04',
          clearPendingThreshold: 'd98d4649-c5d6-4aab-978b-6af29afcbc5a',
        };

        const resp = await fetch('https://api.biconomy.io/api/v2/meta-tx/native', {
          method: 'POST',
          headers: {
            'x-api-key': this.biconomy.apiKey,
            'Content-Type': 'application/json;charset=utf-8',
          },
          body: JSON.stringify({
            to: this.contractAddress,
            apiId: functionIDMap[functionName],
            params: [request, domainSeparator, signed],
            from: account,
            signatureType: this.biconomy.EIP712_SIGN,
          }),
        });

        let txHash = await resp.json();
        txHash = txHash.txHash;

        const getTransactionReceiptMined = async () => {
          const transactionReceiptAsync = async (resolve: any, reject: any) => {
            const receipt = await this.web3.eth.getTransactionReceipt(txHash);
            if (receipt == null) {
              setTimeout(() => transactionReceiptAsync(resolve, reject), 500);
            } else {
              resolve(receipt);
            }
          };

          await new Promise(transactionReceiptAsync);
          return txHash;
        };

        await getTransactionReceiptMined();
        return { transactionHash: txHash };
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
