import { expect } from 'chai';
import { ethers } from 'hardhat';
import { Signer } from 'ethers';
import { release } from 'os';

describe('Valist Contract', () => {
  let valist: any;
  let accounts: Signer[];

  const ADD_KEY = ethers.utils.keccak256(ethers.utils.solidityPack(['string'], ['ADD_KEY_OPERATION']));
  const REVOKE_KEY = ethers.utils.keccak256(ethers.utils.solidityPack(['string'], ['REVOKE_KEY_OPERATION']));
  const ROTATE_KEY = ethers.utils.keccak256(ethers.utils.solidityPack(['string'], ['ROTATE_KEY_OPERATION']));

  const orgName = 'testOrg';
  const repoName = 'testRepo';
  const releaseCID = 'bafybeig5g7gpjxl5mmkufdkf4amj4ttmy4eni5ghgi4huw5w57s6e3cf6y';
  const metaCID = 'bafybeigmfwlweiecbubdw4lq6uqngsioqepntcfohvrccr2o5f7flgydme';

  let orgID = ethers.utils.keccak256(ethers.utils.solidityPack(['uint'], [1]));
  let repoSelector = ethers.utils.keccak256(ethers.utils.solidityPack(['bytes32', 'string'], [orgID, repoName]));

  before(async () => {
    // Deploy Valist Contract
    const contractFactory = await ethers.getContractFactory('Valist');
    valist = await contractFactory.deploy('0x9399BB24DBB5C4b782C70c2969F58716Ebbd6a3b');
    // Setup Accounts and Constants
    accounts = await ethers.getSigners();
  });

  describe('Create an organization', () => {
    it('Should create testOrg organization', async () => {
      await valist.createOrganization(orgName, metaCID);
    });

    it('Should fetch orgID from orgName', async () => {
      orgID = await valist.orgIDByName(orgName);
    });

    it('Org ID should be generated using keccak256(++orgCount)', async () => {
      const expectedOrgID = ethers.utils.keccak256(ethers.utils.solidityPack(['uint'], [await valist.orgCount()]));
      expect(orgID).to.equal(expectedOrgID);
    });

    it('Creator should be an organization admin', async () => {
      expect(await valist.isOrgAdmin(orgID, accounts[0].getAddress())).to.be.true;
    });
  });

  describe('Create a repository', () => {

    it('Should create a repo under testOrg', async () => {
      await valist.createRepository(orgID, repoName, metaCID);
    });

    it('Add addr2 as repoDev under testRepo', async () => {
      // const selector = ethers.utils.keccak256(ethers.utils.solidityPack(['uint'], [await valist.orgCount()]));
      await valist.voteKey(orgID, repoName, ADD_KEY, accounts[1].getAddress());
      expect(await valist.isRepoDev(orgID, repoName, accounts[1].getAddress())).to.be.true;
    });

    it('Add addr3 as repoDev under testRepo', async () => {
      await valist.voteKey(orgID, repoName, ADD_KEY, accounts[2].getAddress());
      expect(await valist.isRepoDev(orgID, repoName, accounts[2].getAddress())).to.be.true;
    });

    it('Should publish a release under testOrg/testRepo', async () => {
      const releaseSelector = ethers.utils.keccak256(
        ethers.utils.defaultAbiCoder.encode(['bytes32', 'string', 'string'],[orgID, repoName, '0.0.0'])
      );
      await valist.voteRelease(
        orgID,
        repoName,
        '0.0.0',
        releaseCID,
        metaCID
      )
      const release = await valist.releases(releaseSelector);
      expect(release.releaseCID).to.equal(releaseCID);
      expect(release.metaCID).to.equal(metaCID);
    });

    it('Should fetch release using releaseSelector', async () => {
      const releaseSelector = ethers.utils.keccak256(
        ethers.utils.defaultAbiCoder.encode(['bytes32', 'string', 'string'],[orgID, repoName, '0.0.0'])
      );
      const release = await valist.releases(releaseSelector);
      expect(release.releaseCID).to.equal(releaseCID);
      expect(release.metaCID).to.equal(metaCID);
    });

    it('Should fetch release using getLatestRelease', async () => {
      const release = await valist.getLatestRelease(orgID, repoName);
      expect(release[0]).to.equal('0.0.0');
      expect(release[1]).to.equal(releaseCID);
      expect(release[2]).to.equal(metaCID);
    });
  });

  describe('Multi-Factor Releases', () => {

    it('Should enable multi-factor release policy on repo', async () => {
      await valist.voteThreshold(orgID, repoName, 2);
      expect(await valist.pendingThresholdRequests(repoSelector, 0)).to.equal(2);
      await valist.connect(accounts[1]).voteThreshold(orgID, repoName, 2);
      try {
        await valist.connect(accounts[2]).voteThreshold(orgID, repoName, 2);
      } catch(e) {
        expect(e.message).to.contain('Threshold set');
      }

      const repo = await valist.repos(repoSelector);
      expect(repo.threshold).to.equal(2);
    });

    it('Should throw error when trying to vote for threshold that is already set', async () => {
      try {
        await valist.connect(accounts[2]).voteThreshold(orgID, repoName, 2);
      } catch(e) {
        expect(e.message).to.contain('Threshold set');
      }
    });

    it('Should propose a pending release under testOrg/testRepo', async () => {
      await valist.voteRelease(
        orgID,
        repoName,
        '0.0.1',
        releaseCID,
        metaCID
      )
    });

    it('Should fetch pending release request', async () => {
      const pendingRelease = await valist.pendingReleaseRequests(repoSelector, 0);
      expect(pendingRelease.tag).to.equal('0.0.1');
      expect(pendingRelease.releaseCID).to.equal(releaseCID);
      expect(pendingRelease.metaCID).to.equal(metaCID);
    });

    it('Should sign a pending release under testOrg/testRepo (2nd key)', async () => {
      await valist.connect(accounts[1]).voteRelease(
        orgID,
        repoName,
        '0.0.1',
        releaseCID,
        metaCID
      );
    });

    it('Release should be finalized after threshold has been met', async () => {
      const releaseSelector = ethers.utils.keccak256(
        ethers.utils.defaultAbiCoder.encode(['bytes32', 'string', 'string'], [orgID, repoName, '0.0.1'])
      );
      const release = await valist.releases(releaseSelector);
      expect(release.releaseCID).to.equal(releaseCID);
      expect(release.metaCID).to.equal(metaCID);
    });

    it('Should fail to propose release that has been finalized', async () => {
      try {
        await valist.voteRelease(orgID, repoName, '0.0.1', releaseCID, metaCID);
      } catch (e) {
        expect(e.message).to.contain('Tag used');
      }
    });

    it('Should be able to clear old pending release that has already met threshold', async () => {
      await valist.clearPendingRelease(orgID, repoName, '0.0.1', releaseCID, metaCID, 0);
      try {
        await valist.pendingReleaseRequests(repoSelector, 0);
      } catch (e) {
        expect(e.message).to.contain('Transaction reverted without a reason');
      }
    });
  });

  describe('Vote on adding a new repoDev key to testOrg/testRepo with multi-factor threshold of 2', () => {
    it('Vote on adding key 4 (1st key)', async () => {
      await valist.voteKey(orgID, repoName, ADD_KEY, accounts[3].getAddress());
    });

    it('Vote on adding key 4 (2nd key)', async () => {
      await valist.connect(accounts[1]).voteKey(orgID, repoName, ADD_KEY, accounts[3].getAddress());
    });

    it('Should fail to vote when key is already added', async () => {
      try {
        await valist.voteKey(orgID, repoName, ADD_KEY, accounts[3].getAddress());
      } catch (e) {
        expect(e.message).to.contain('Key exists');
      }
    });

    it('Validate that key 4 is now repo dev', async () => {
      expect(await valist.isRepoDev(orgID, repoName, accounts[3].getAddress())).to.be.true;
    });
  });

  /*
  describe('Vote on revoking key 4 as repoAdmin from testOrg/testRepo with multi-factor policy 3', () => {
    it('Vote on revoking key 4 (1st key)', async () => {
      await valist.voteRevokeRepoAdmin(orgID, repoName, signer4.address);
    });
  });
  */

  describe('Read from Valist contract', () => {
    it('Should get testOrg metadata', async () => {
      const org = await valist.orgs(orgID);
      expect(org.metaCID).to.equal(metaCID);
    });

    it('Should get 10 orgNames', async () => {
      const orgNames = await valist.getOrgNames(1, 10);
      expect(orgNames).to.contain('testOrg');
      expect(orgNames.length).to.equal(10);
    });

    it('Should get number of orgs', async () => {
      expect(Number(await valist.orgCount())).to.equal(1);
    });
  });
});
