import Web3 from 'web3';
import { expect } from 'chai';
import { describe, before, it } from 'mocha';
import Valist from '../dist/';
import { getContractInstance } from '../src/utils';

import ValistABI from '../src/abis/contracts/Valist.sol/Valist.json';

console.error = () => {}; // mute console errors

const ganache = require('ganache-core');

const web3Provider = ganache.provider();
let contractInstance: any;
let valist: Valist;
let accounts: string[];

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
    .deploy({ data: ValistABI.bytecode, arguments: ['0x9399BB24DBB5C4b782C70c2969F58716Ebbd6a3b'] })
    .send({ from: accounts[0], gas: 4333333 });

  return valistContract;
};

describe('Test Valist Lib', async () => {
  before('Deploy Valist Contract', async () => {
    contractInstance = await deployContract(web3Provider);
  });

  describe('Call getContractInstance', async () => {
    it('Should be an object', async () => {
      const web3Instance = new Web3(web3Provider);
      const valistContract = getContractInstance(web3Instance, ValistABI.abi, contractInstance.options.address);
      expect(valistContract).to.be.an('object');
    });
  });

  describe('Create new Valist Instance', async () => {
    before(() => {
      valist = new Valist({
        web3Provider,
        metaTx: false,
        contractAddress: contractInstance.options.address,
      });
    });

    it('Return a Valist Object', async () => {
      expect(valist).to.have.property('web3');
      expect(valist).to.have.property('ipfs');
      expect(valist).to.have.property('defaultAccount');
      expect(valist).to.have.property('metaTxEnabled');
      expect(valist).to.have.property('metaTxReady');
      expect(valist).to.have.property('contractAddress');
    });

    it('Call Valist Connect', async () => {
      await valist.connect();
      expect(valist).to.have.property('contract');
    });
  });

  describe('Create an Organization', async () => {
    it('Should return transaction response', async () => {
      const transactionResponse = await valist.createOrganization(orgShortName, meta);
      expect(transactionResponse).to.have.property('transactionHash');
      expect(transactionResponse).to.have.property('blockHash');
      expect(transactionResponse).to.have.property('blockNumber');
    });

    it('Should store orgName in list of orgNames', async () => {
      const orgNames = await valist.getOrganizationNames();
      expect(orgNames[0]).to.equal('secureco');
    });

    it('Should fetch organization', async () => {
      const org = await valist.getOrganization(orgShortName);
      expect(org.orgID).to.equal('0xcc69885fda6bcc1a4ace058b4a62bf5e179ea78fd58a1ccd71c22cc9b688792f');
      expect(org.metaCID).to.equal('bafkreiacinnkuxv46nybpqjtxizecpytoskdeukd7scunuu4aqovjbrvqy');
    });
  });

  describe('Create a Repository', async () => {
    it('Should return transaction response', async () => {
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

    it('Should fail when trying to create the same organization twice', async () => {
      try {
        await valist.createOrganization(orgShortName, meta);
      } catch (e) {
        expect(e.name).to.equal('RuntimeError');
        expect(e.toString()).to.contain('Org exists');
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
        expect(e.toString()).to.contain('User does not have permission to publish release');
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

});
