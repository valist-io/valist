const { expect, assert } = require("chai");

describe("Valist Contract", function() {
  let Valist;
  let valist;
  let signer1;
  let signer2;
  let signer3;
  let ORG_OWNER;

  const releaseCID = "Qmc5gCcjYypU7y28oCALwfSvxCBskLuPKWpK4qpterKC7z";
  const metaCID = "Qmc5gCcjYypU7y28oCALwfSvxCBskLuPKWpK4qpterKC7z";

  before(async() => {
    // Deploy Valist Contract
    Valist = await ethers.getContractFactory("Valist");
    valist = await Valist.deploy();

    // Setup Accounts and Constants
    ORG_OWNER = "123b642491709420c2370bb98c4e7de2b1bc05c5f9fd95ac4111e12683553c62"; // keccak256 hash of "ORG_OWNER" string
    [signer1, signer2, signer3, signer4, _] = await ethers.getSigners();
  });

  describe('Deployment', () => {
    it("Should return the Valist contract address", async function() {
      expect(valist.address);
    });
  });

  describe('Publish a Release with Multi-Factor Release Policy', () => {
    it("Should create testOrg organization", async function() {
      await valist.createOrganization("testOrg", metaCID);
    });

    it("Creator should be the organization owner", async function() {
      expect(await valist.isOrgOwner("testOrg", signer1.address)).to.be.true;
    });

    it("Creator should also have admin privs", async function() {
      expect(await valist.isOrgAdmin("testOrg", signer1.address)).to.be.true;
    });

    it("Should create a repo under testOrg", async function() {
      await valist.createRepository("testOrg", "testRepo", metaCID);
    });

    it("Add addr2 as repoDev under testOrg", async function() {
      await valist.voteAddRepoAdmin("testOrg", "testRepo", signer2.address);
      expect(await valist.isRepoDev("testOrg", "testRepo", signer2.address)).to.be.true;
    });

    it("Add addr3 as repoDev under testOrg", async function() {
      await valist.voteAddRepoAdmin("testOrg", "testRepo", signer3.address);
      expect(await valist.isRepoDev("testOrg", "testRepo", signer3.address)).to.be.true;
    });

    it("Should enable multi-factor release policy", async function() {
      await valist.setRepoSignerThreshold("testOrg", "testRepo", 3);
      expect(await valist.isRepoDev("testOrg", "testRepo", signer3.address)).to.be.true;
    });

    it("Should propose a pending release under testOrg/testRepo", async function() {
      await valist.voteRelease(
        "testOrg",
        "testRepo",
        "0.0.1",
        releaseCID,
        metaCID
      )
    });

    it("Should sign a pending release under testOrg/testRepo (2nd key)", async function() {
      await valist.connect(signer2).voteRelease(
        "testOrg",
        "testRepo",
        "0.0.1",
        releaseCID,
        metaCID
      );
    });

    it("Should sign a pending release under testOrg/testRepo (3rd key)", async function() {
      await valist.connect(signer3).voteRelease(
        "testOrg",
        "testRepo",
        "0.0.1",
        releaseCID,
        metaCID
      );
    });

    it("Release should be finalized after threshold has been met", async function() {
      const release = await valist.getLatestRelease("testOrg", "testRepo");
      expect(release[0]).to.equal(releaseCID);
      expect(release[1]).to.equal(metaCID);
    });
    // things that need multifactor operations:
    // publish release
    // add users/keys
    // remove users/keys
    // 
  });

  describe('Vote on adding a new repoAdminKey to testOrg/testRepo with multi-factor policy 3', () => {
    it("Vote on adding key 4 (1st key)", async function() {
      await valist.voteAddRepoAdmin("testOrg", "testRepo", signer4.address);
    });

    it("Vote on adding key 4 (2nd key)", async function() {
      await valist.connect(signer2).voteAddRepoAdmin("testOrg", "testRepo", signer4.address);
    });

    it("Vote on adding key 4 (3rd key)", async function() {
      await valist.connect(signer3).voteAddRepoAdmin("testOrg", "testRepo", signer4.address);
    });

    it("Validate that key 4 is now repo admin", async function() {
      expect(await valist.isRepoAdmin("testOrg", "testRepo", signer4.address)).to.be.true;
    });
  });

  /*
  describe('Vote on revoking key 4 as repoAdmin from testOrg/testRepo with multi-factor policy 3', () => {
    it("Vote on revoking key 4 (1st key)", async function() {
      await valist.voteRevokeRepoAdmin("testOrg", "testRepo", signer4.address);
    });
  });
  */

  describe('Read from Valist contract', () => {
    it("Should get testOrg", async function() {
      const org = await valist.getOrganization("testOrg");
      expect(org[0]).to.equal(0);
      expect(org[1]).to.equal(metaCID);
    });

    it("Should get 10 orgNames", async function() {
      const orgNames = await valist.getOrgNames(1, 10);
      expect(orgNames).to.contain('testOrg');
      expect(orgNames.length).to.equal(10);
    });

    it("Should get number of orgs", async function() {
      expect(Number(await valist.getOrgCount())).to.equal(1);
    });
  });
});