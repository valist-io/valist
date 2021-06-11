const { expect } = require("chai");

describe("Valist Contract", function() {
  let Valist, valist, addr1, addr2, ORG_OWNER;

  beforeEach(async() => {
    Valist = await ethers.getContractFactory("Valist");
    valist = await Valist.deploy();
    ORG_OWNER = "123b642491709420c2370bb98c4e7de2b1bc05c5f9fd95ac4111e12683553c62";
    [addr1, addr2, _] = await ethers.getSigners();
  });

  describe('Deployment', () => {
    it("Should return the contract address", async function() {
      console.log(valist.address)
    });
  });

  describe('Publish Release', () => {
    it("Should create testOrg organization", async function() {
      await valist.createOrganization("testOrg", "Qmc5gCcjYypU7y28oCALwfSvxCBskLuPKWpK4qpterKC7z");
    });

    it("Creator should be the organization owner", async function() {
      await valist.isOrgOwner("testOrg", addr1.address);
    });

    it("Creator should also have admin privs", async function() {
      await valist.isOrgAdmin("testOrg", addr1.address);
    });

    it("Should create a repo under testOrg", async function() {
      await valist.createRepository("testOrg", "testRepo", "Qmc5gCcjYypU7y28oCALwfSvxCBskLuPKWpK4qpterKC7z");
    });

    it("Should propose a pending release under testOrg/testRepo", async function() {
      await valist.proposeRelease(
        "testOrg", 
        "testRepo",
        "0.0.1", 
        "Qmc5gCcjYypU7y28oCALwfSvxCBskLuPKWpK4qpterKC7z",
        "Qmc5gCcjYypU7y28oCALwfSvxCBskLuPKWpK4qpterKC7z"
      )
    });

    it("Should sign a pending release under testOrg/testRepo", async function() {
      await valist.signRelease("testOrg", "testRepo", "0.0.1");
    });

    it("Should finalize the signed pending release under testOrg/testRepo", async function() {
      await valist.finalizeRelease("testOrg", "testRepo", "0.0.1");
    });
  });
});