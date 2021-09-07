import Web3 from 'web3';
import { expect } from 'chai';
import { describe, before, it } from 'mocha';
import Valist from '../dist/';
import { ADD_KEY, REVOKE_KEY } from '../dist/constants';

import ValistABI from '../src/abis/contracts/Valist.sol/Valist.json';
import ValistRegistryABI from '../src/abis/contracts/ValistRegistry.sol/ValistRegistry.json';

console.error = () => {}; // mute console errors

const ganache = require('ganache-core');

const web3Provider = ganache.provider();
let contractInstance: any;
let registryInstance: any;
let valist: Valist;
let accounts: string[];

const metaTxRelay = '0x9399BB24DBB5C4b782C70c2969F58716Ebbd6a3b';

let orgID: string;
const orgShortName = 'secureco';
const repoName = 'firmware';

const meta = {
  name: 'Secure Firmware Company',
  description: 'We are a secure firmware company.',
};

const repoMeta: {
  name: string;
  description: string;
  projectType: 'binary';
  homepage: string;
  repository: string;
} = {
  name: repoName,
  description: 'A secure firmware.',
  projectType: 'binary',
  homepage: 'https://pugsandhugs.party',
  repository: 'https://github.com/pugsandhugs.party',
};

const release = {
  tag: '0.0.1',
  releaseCID: 'QmU2PN4NVcAP2wCGKNpmjoEaM2xjRzjjjy9YUh25TEUPta',
  metaCID: 'QmNTDYaHbB88ezsQuYpugXMx1X8NP2xj9S8HtSTzmKQ5XS',
};

const deployContract = async (provider: any) => {
  const web3 = new Web3(provider);
  accounts = await web3.eth.getAccounts();

  const valistContract = await new web3.eth.Contract(ValistABI.abi as any)
    .deploy({ data: ValistABI.bytecode, arguments: [metaTxRelay] })
    .send({ from: accounts[0], gas: 4333333 });

  return valistContract;
};

const deployRegistry = async (provider: any) => {
  const web3 = new Web3(provider);
  accounts = await web3.eth.getAccounts();

  const valistRegistry = await new web3.eth.Contract(ValistRegistryABI.abi as any)
    .deploy({ data: ValistRegistryABI.bytecode, arguments: [metaTxRelay] })
    .send({ from: accounts[0], gas: 4333333 });

  return valistRegistry;
};

describe('Test Valist Lib', async () => {
  before('Deploy Valist Contracts', async () => {
    contractInstance = await deployContract(web3Provider);
    registryInstance = await deployRegistry(web3Provider);
  });

  describe('Create new Valist Instance', async () => {
    before(() => {
      valist = new Valist({
        web3Provider,
        metaTx: false,
        contractAddress: contractInstance.options.address,
        registryAddress: registryInstance.options.address,
      });
    });

    it('Return a Valist Object', async () => {
      expect(valist).to.have.property('web3');
      expect(valist).to.have.property('ipfs');
      expect(valist).to.have.property('defaultAccount');
      expect(valist).to.have.property('metaTxEnabled');
      expect(valist).to.have.property('contractAddress');
      expect(valist).to.have.property('registryAddress');
    });

    it('Call Valist Connect', async () => {
      await valist.connect();
      expect(valist).to.have.property('contract');
      expect(valist).to.have.property('registry');
    });
  });

  describe('Create an Organization', async () => {

    it('Should get empty orgID by name', async () => {
      try {
        await valist.getOrgIDFromName(orgShortName);
      } catch (e) {
        expect(e.message).to.contain('orgID not found');
      }
    });

    it('Should create organization', async () => {
      const transactionResponse = await valist.createOrganization(orgShortName, meta);
      orgID = transactionResponse.orgID;
      expect(orgID).to.equal('0xcc69885fda6bcc1a4ace058b4a62bf5e179ea78fd58a1ccd71c22cc9b688792f')
      expect(transactionResponse).to.have.property('transactionHash');
      expect(transactionResponse).to.have.property('blockHash');
      expect(transactionResponse).to.have.property('blockNumber');
    });

    it('Should link orgID to namespace', async () => {
      const transactionResponse = await valist.linkNameToID(orgShortName, orgID);
      expect(transactionResponse).to.have.property('transactionHash');
      expect(transactionResponse).to.have.property('blockHash');
      expect(transactionResponse).to.have.property('blockNumber');
    });

    it('Should get orgID by name', async () => {
      const id = await valist.getOrgIDFromName(orgShortName);
      expect(id).to.equal(orgID);
    });

    it('Should store orgName in list of orgNames', async () => {
      const orgNames = await valist.getOrganizationNames();
      expect(orgNames[0]).to.equal('secureco');
    });

    it('Should fetch organization', async () => {
      const org = await valist.getOrganization(orgShortName);
      expect(org.orgID).to.equal(orgID);
      expect(org.repoNames.length).to.equal(0);
      expect(org.meta.name).to.equal(meta.name);
      expect(org.meta.description).to.equal(meta.description);
      expect(org.threshold).to.equal('0');
      expect(org.thresholdDate).to.equal('0');
      expect(org.metaCID).to.equal('bafkreiacinnkuxv46nybpqjtxizecpytoskdeukd7scunuu4aqovjbrvqy');
    });
  });

  describe('Create a Repository', async () => {
    it('Should create repository', async () => {
      const transactionResponse = await valist.createRepository(orgShortName, repoName, repoMeta);
      expect(transactionResponse).to.have.property('transactionHash');
      expect(transactionResponse).to.have.property('blockHash');
      expect(transactionResponse).to.have.property('blockNumber');
    });

    it('Org should now include repo in repoNames', async () => {
      const org = await valist.getOrganization(orgShortName);
      expect(org.repoNames).to.include(repoName);
    });

    it('Should get list of repoNames', async () => {
      const repoNames = await valist.getRepoNames(orgShortName);
      expect(repoNames).to.include(repoName);
    });

    it('Should get repository', async () => {
      const repo = await valist.getRepository(orgShortName, repoName);
      expect(repo.orgID).to.equal('0xcc69885fda6bcc1a4ace058b4a62bf5e179ea78fd58a1ccd71c22cc9b688792f');
      expect(repo.meta).to.deep.equal(repoMeta);
    });

    it('Should fail when trying to create the same organization twice', async () => {
      try {
        await valist.createRepository(orgShortName, repoName, repoMeta);
      } catch (e) {
        expect(e.name).to.equal('RuntimeError');
        expect(e.toString()).to.contain('Repo exists');
      }
    });
  });

  describe('Publish a Release', async () => {
    it('Should return transaction response', async () => {
      const transactionResponse = await valist.publishRelease(orgShortName, repoName, release);
      expect(transactionResponse).to.have.property('transactionHash');
      expect(transactionResponse).to.have.property('blockHash');
      expect(transactionResponse).to.have.property('blockNumber');
    });

    it('Should fetch latest release and meta CID', async () => {
      const resp = await valist.getLatestRelease(orgShortName, repoName);
      expect(resp.releaseCID).to.equal(release.releaseCID);
      expect(resp.metaCID).to.equal(release.metaCID);
      expect(resp.tag).to.equal(release.tag);
    });

    it('Should fail when user does not have permission', async () => {
      try {
        await valist.publishRelease(orgShortName, repoName, { ...release, tag: '0.0.1' }, accounts[1]);
      } catch (e) {
        expect(e.toString()).to.contain('User does not have permission to publish release');
      }
    });

    it ('Should fail when org does not exist', async () => {
      try {
        await valist.publishRelease('', repoName, release);
      } catch (e) {
        expect(e.toString()).to.contain('orgID not found');
      }
    });

    it ('Should fail when repo does not exist', async () => {
      try {
        await valist.publishRelease(orgShortName, '', { ...release, tag: '0.0.1' });
      } catch (e) {
        expect(e.toString()).to.contain('User does not have permission to publish release');
      }
    });
  });

  describe('Get Release Tags From Repo', async () => {
    it('Should return release tag', async () => {
      const response = await valist.getReleaseTags(orgShortName, repoName);
      expect(response).to.include.members([release.tag]);
    });
  });

  describe('Get a Release from Project by Tag', async () => {
    it('Should return release and meta CID', async () => {
      const resp = await valist.getReleaseByTag(orgShortName, repoName, release.tag);
      expect(resp.releaseCID).to.equal(release.releaseCID);
      expect(resp.metaCID).to.equal(release.metaCID);
      expect(resp.tag).to.equal(release.tag);
    });
  });

  describe('Multi-Factor Org Setup', async () => {
    it('Should add key2 as orgAdmin', async () => {
      await valist.voteOrgAdmin(orgShortName, accounts[1]);
    });

    it('Should add key3 as orgAdmin', async () => {
      await valist.voteOrgAdmin(orgShortName, accounts[2]);
    });

    it('Should vote for org threshold', async () => {
      await valist.voteOrgThreshold(orgShortName, 2);
    });

    it('Should fail to vote for org threshold twice with same key', async () => {
      try {
        await valist.voteOrgThreshold(orgShortName, 2);
      } catch (e) {
        expect(e.message).to.contain('User voted');
      }
    });

    it('Should fetch pending org threshold requests', async () => {
      const requests = await valist.getPendingOrgThresholds(orgShortName);
      expect(Number(requests[0])).to.equal(2);
    });

    it('Should vote for org threshold with key2', async () => {
      valist.defaultAccount = accounts[1];
      await valist.voteOrgThreshold(orgShortName, 2);
      valist.defaultAccount = accounts[0];
    });

    it('Vote should pass and org threshold set', async () => {
      const org = await valist.getOrganization(orgShortName);
      expect(Number(org.threshold)).to.equal(2);
    });

    it('Should clear pending org threshold', async () => {
      await valist.clearPendingOrgThreshold(orgShortName, 2, 0);
      const requests = await valist.getPendingOrgThresholds(orgShortName);
      const pendingVote = await valist.getPendingOrgThresholdVotes(orgShortName, 2);
      expect(requests.length).to.equal(0);
      expect(Number(pendingVote.expiration)).to.equal(0);
      expect(pendingVote.signers.length).to.equal(0);
    });

    it('Should vote to add key4 as a orgAdmin from key1', async () => {
      await valist.voteOrgAdmin(orgShortName, accounts[3]);
    });

    it('Should fetch pending orgAdmin key', async () => {
      const pendingOrgAdmins = await valist.getPendingOrgAdmins(orgShortName);
      expect(pendingOrgAdmins[0]).to.equal(accounts[3]);
    });

    it('Should fetch pending orgAdmin votes for key4', async () => {
      const pendingVote = await valist.getPendingOrgAdminVotes(orgShortName, ADD_KEY, accounts[3]);
      expect(pendingVote.signers[0]).to.equal(accounts[0]);
    });

    it('Should vote to add key4 as an orgAdmin from key2', async () => {
      valist.defaultAccount = accounts[1];
      await valist.voteOrgAdmin(orgShortName, accounts[3]);
      valist.defaultAccount = accounts[0];
    });

    it('Vote should pass and key4 should be orgAdmin', async () => {
      const isOrgAdmin = await valist.isOrgAdmin(orgShortName, accounts[3]);
      expect(isOrgAdmin).to.be.true;
    });

    it('Should clear pending orgAdmin key', async () => {
      await valist.clearPendingOrgKey(orgShortName, ADD_KEY, accounts[3], 0);
      const pendingVote = await valist.getPendingOrgAdminVotes(orgShortName, ADD_KEY, accounts[3]);
      const pendingOrgAdmins = await valist.getPendingOrgAdmins(orgShortName);
      expect(pendingOrgAdmins.length).to.equal(0);
      expect(Number(pendingVote.expiration)).to.equal(0);
      expect(pendingVote.signers.length).to.equal(0);
    });
  });

  describe('Multi-Factor Org Remove Key', async () => {
    it('Should vote to revoke key2 as orgAdmin', async () => {
      await valist.revokeOrgAdmin(orgShortName, accounts[1]);
    });

    it('Should fail to vote for revocation twice with same key', async () => {
      try {
        await valist.revokeOrgAdmin(orgShortName, accounts[1]);
      } catch (e) {
        expect(e.message).to.contain('User voted');
      }
    });

    it('Should fetch pending orgAdmin key', async () => {
      const pendingOrgAdmins = await valist.getPendingOrgAdmins(orgShortName);
      expect(pendingOrgAdmins[0]).to.equal(accounts[1]);
    });

    it('Should fetch pending orgAdmin revocation votes for key2', async () => {
      const pendingVote = await valist.getPendingOrgAdminVotes(orgShortName, REVOKE_KEY, accounts[1]);
      expect(pendingVote.signers[0]).to.equal(accounts[0]);
    });

    it('Should vote from key3 to revoke key2 as orgAdmin', async () => {
      valist.defaultAccount = accounts[2];
      await valist.revokeOrgAdmin(orgShortName, accounts[1]);
      valist.defaultAccount = accounts[0];
    });

    it('Vote should pass and key2 should be orgAdmin', async () => {
      const isOrgAdmin = await valist.isOrgAdmin(orgShortName, accounts[1]);
      expect(isOrgAdmin).to.be.false;
    });

    it('Should clear pending orgAdmin key', async () => {
      await valist.clearPendingOrgKey(orgShortName, REVOKE_KEY, accounts[1], 0);
      const pendingVote = await valist.getPendingOrgAdminVotes(orgShortName, REVOKE_KEY, accounts[1]);
      const pendingOrgAdmins = await valist.getPendingOrgAdmins(orgShortName);
      expect(pendingOrgAdmins.length).to.equal(0);
      expect(Number(pendingVote.expiration)).to.equal(0);
      expect(pendingVote.signers.length).to.equal(0);
    });

    it('Threshold should be one less after revocation', async () => {
      const org = await valist.getOrganization(orgShortName);
      expect(Number(org.threshold)).to.equal(1);
    });
  });

  describe('Multi-Factor Repo Setup', async () => {
    it('Should add key2 as repoDev', async () => {
      await valist.voteRepoDev(orgShortName, repoName, accounts[1]);
    });

    it('Should add key3 as repoDev', async () => {
      await valist.voteRepoDev(orgShortName, repoName, accounts[2]);
    });

    it('Should vote for repo threshold', async () => {
      await valist.voteRepoThreshold(orgShortName, repoName, 2);
    });

    it('Should fail to vote for repo threshold twice with same key', async () => {
      try {
        await valist.voteRepoThreshold(orgShortName, repoName, 2);
      } catch (e) {
        expect(e.message).to.contain('User voted');
      }
    });

    it('Should fetch pending repo threshold requests', async () => {
      const requests = await valist.getPendingRepoThresholds(orgShortName, repoName);
      expect(Number(requests[0])).to.equal(2);
    });

    it('Should vote for repo threshold with key2', async () => {
      valist.defaultAccount = accounts[1];
      await valist.voteRepoThreshold(orgShortName, repoName, 2);
      valist.defaultAccount = accounts[0];
    });

    it('Vote should pass and repo threshold set', async () => {
      const repo = await valist.getRepository(orgShortName, repoName);
      expect(repo.threshold).to.equal(2);
    });

    it('Should clear pending repo threshold', async () => {
      await valist.clearPendingRepoThreshold(orgShortName, repoName, 2, 0);
      const requests = await valist.getPendingRepoThresholds(orgShortName, repoName);
      const pendingVote = await valist.getPendingRepoThresholdVotes(orgShortName, repoName, 2);
      expect(requests.length).to.equal(0);
      expect(Number(pendingVote.expiration)).to.equal(0);
      expect(pendingVote.signers.length).to.equal(0);
    });

    it('Should vote to add key4 as a repoDev from key1', async () => {
      await valist.voteRepoDev(orgShortName, repoName, accounts[3]);
    });

    it('Should fetch pending repo dev key', async () => {
      const pendingRepoDevs = await valist.getPendingRepoDevs(orgShortName, repoName);
      expect(pendingRepoDevs[0]).to.equal(accounts[3]);
    });

    it('Should fetch pending repo dev votes for key4', async () => {
      const pendingVote = await valist.getPendingRepoDevVotes(orgShortName, repoName, ADD_KEY, accounts[3]);
      expect(pendingVote.signers[0]).to.equal(accounts[0]);
    });

    it('Should vote to add key4 as a repoDev from key2', async () => {
      valist.defaultAccount = accounts[1];
      await valist.voteRepoDev(orgShortName, repoName, accounts[3]);
      valist.defaultAccount = accounts[0];
    });

    it('Vote should pass and key4 should be repoDev', async () => {
      const isRepoDev = await valist.isRepoDev(orgShortName, repoName, accounts[3]);
      expect(isRepoDev).to.be.true;
    });

    it('Should clear pending repoDev key', async () => {
      await valist.clearPendingRepoKey(orgShortName, repoName, ADD_KEY, accounts[3], 0);
      const pendingVote = await valist.getPendingRepoDevVotes(orgShortName, repoName, ADD_KEY, accounts[3]);
      const pendingRepoDevs = await valist.getPendingRepoDevs(orgShortName, repoName);
      expect(pendingRepoDevs.length).to.equal(0);
      expect(Number(pendingVote.expiration)).to.equal(0);
      expect(pendingVote.signers.length).to.equal(0);
    });
  });

  describe('Multi-Factor Release', async () => {
    it('Should propose new release', async () => {
      release.tag = '0.0.2';
      const transactionResponse = await valist.publishRelease(orgShortName, repoName, release);
      expect(transactionResponse).to.have.property('transactionHash');
      expect(transactionResponse).to.have.property('blockHash');
      expect(transactionResponse).to.have.property('blockNumber');
    });

    it('Should fetch pending release requests', async () => {
      const pendingReleases = await valist.getPendingReleases(orgShortName, repoName);
      expect(pendingReleases[0].tag).to.equal(release.tag);
      expect(pendingReleases[0].releaseCID).to.equal(release.releaseCID);
      expect(pendingReleases[0].metaCID).to.equal(release.metaCID);
    });

    it('Should fetch pending release votes', async () => {
      const pendingVote = await valist.getPendingReleaseVotes(orgShortName, repoName, release);
      expect(Number(pendingVote.expiration)).to.be.greaterThan(0);
      expect(pendingVote.signers[0]).to.equal(accounts[0]);
    });

    // it('Should fetch VoteReleaseEvent', async () => {
    //   const events = await valist.getVoteReleaseEvents();
    //   expect(events[events.length - 1].returnValues._sigCount).to.be.equal('1');
    //   expect(events[events.length - 1].returnValues._threshold).to.be.equal('2');
    // });

    it('Should finalize vote on new release', async () => {
      valist.defaultAccount = accounts[1];
      const transactionResponse = await valist.publishRelease(orgShortName, repoName, release);
      valist.defaultAccount = accounts[0];
      expect(transactionResponse).to.have.property('transactionHash');
      expect(transactionResponse).to.have.property('blockHash');
      expect(transactionResponse).to.have.property('blockNumber');
    });

    // it('Should fetch VoteReleaseEvent', async () => {
    //   const events = await valist.getVoteReleaseEvents();
    //   expect(events[events.length - 1].returnValues._sigCount).to.be.equal('2');
    //   expect(events[events.length - 1].returnValues._threshold).to.be.equal('2');
    // });

    it('Should clear pending release', async () => {
      await valist.clearPendingRelease(orgShortName, repoName, release, 0);
      const pendingReleases = await valist.getPendingReleases(orgShortName, repoName);
      expect(pendingReleases.length).to.equal(0);
    });
  });

});
