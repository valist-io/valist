import Web3 from "web3";
import { expect } from "chai";
import Valist, { getContractInstance, ProjectType } from "../src/index";

import ValistABI from "../src/abis/Valist.json";
import { describe, before, it } from "mocha";

console.error = () => {}; // mute console errors

const ganache = require("ganache-core");
let web3Provider = ganache.provider();
let contractInstance: any;
let valistInstance: Valist;

const orgShortName = "secureco";
const projectName = "firmware";

const meta = {
	name: "Secure Firmware Company",
	description: "We are a secure firmware company."
};

const repoMeta = {
	name: projectName,
	description: "A secure firmware.",
	projectType: ProjectType.BINARY,
	homepage: "https://pugsandhugs.party",
	repository: "https://github.com/pugsandhugs.party"
};

const release = {
	tag: "0.0.1",
	hash: "QmU2PN4NVcAP2wCGKNpmjoEaM2xjRzjjjy9YUh25TEUPta",
	meta: "QmNTDYaHbB88ezsQuYpugXMx1X8NP2xj9S8HtSTzmKQ5XS"
};

const deployContract = async (provider: any) =>{
	const web3 = new Web3(provider);
	let accounts = await web3.eth.getAccounts();

	// @ts-ignore
	let valistContract = await new web3.eth.Contract(ValistABI.abi)
		.deploy({data: ValistABI.bytecode, arguments: []})
		.send({from: accounts[0], gas: 3333333});

		// console.log(`Contract Address: ${valistContract.options.address}`);
		return valistContract;
};

describe("Test Valist Lib", async () => {
	before("Deploy Valist Contract", async () =>{
		contractInstance = await deployContract(web3Provider);
	});

	describe("Call getContractInstance", async () => {
		it("Should be an object", async () => {
			const web3Instance = new Web3(web3Provider);
			const valistContract = await getContractInstance(web3Instance, ValistABI.abi, contractInstance.options.address);
			expect(valistContract).to.be.an("object");
		});
	});

	describe("Create new Valist Instance", async () => {
		before(() => {
			valistInstance = new Valist({
				web3Provider: web3Provider, 
				metaTx: false,
				contractAddress: contractInstance.options.address
			});
		});
	
		it("Return a Valist Object", async () => {
			expect(valistInstance).to.have.property("web3");
			expect(valistInstance).to.have.property("ipfs");
			expect(valistInstance).to.have.property("defaultAccount");
			expect(valistInstance).to.have.property("metaTxEnabled");
			expect(valistInstance).to.have.property("metaTxReady");
			expect(valistInstance).to.have.property("contractAddress");
		});

		it("Call Valist Connect", async () => {
			await valistInstance.connect();
			expect(valistInstance).to.have.property("valist");
		});
	});

	describe("Create an Organization", async () => {
		it("Should return transaction response", async () => {
			const transactionResponse = await valistInstance.createOrganization(orgShortName, meta);
			expect(transactionResponse).to.have.property("transactionHash");
			expect(transactionResponse).to.have.property("blockHash");
			expect(transactionResponse).to.have.property("blockNumber");
		});

		it("Should store orgName in list of orgNames", async () => {
			const orgNames = await valistInstance.getOrganizationNames();
			expect(orgNames[0]).to.equal("secureco");
		});
	});

	describe("Create a Project", async () => {
		it("Should return transaction response", async () => {
			const transactionResponse = await valistInstance.createRepository(orgShortName, projectName, repoMeta);
			expect(transactionResponse).to.have.property("transactionHash");
			expect(transactionResponse).to.have.property("blockHash");
			expect(transactionResponse).to.have.property("blockNumber");
		});

		it("Should fail when trying to create the same organization twice", async () => {
			try {
				await valistInstance.createOrganization(orgShortName, meta);
			} catch (e) {
				expect(e.name).to.equal("RuntimeError");
				expect(e.toString()).to.contain("Organization exists");
			}
		});
		
	});

	describe("Publish a Release", async () => {
		it("Should return transaction response", async () => {
			const transactionResponse = await valistInstance.publishRelease(orgShortName, projectName, release);
			expect(transactionResponse).to.have.property("transactionHash");
			expect(transactionResponse).to.have.property("blockHash");
			expect(transactionResponse).to.have.property("blockNumber");
		});
	});

	describe("Get Release Tags From Repo", async () => {
		it("Should return used release tag", async () => {
			const response = await valistInstance.getReleaseTagsFromRepo(orgShortName, projectName);
			expect(response).to.include.members([release.tag])
		});
	});

	describe("Get a Release from Project by Tag", async () => {
		it("Should return release and meta CID", async () => {
			const response = await valistInstance.getReleaseByTag(orgShortName, projectName, release.tag);
			expect(response).to.have.property("releaseCID");
			expect(response).to.have.property("metaCID");
		});
	});

	describe("Get the latest Release from Project", async () => {
		it("Should return release and meta CID", async () => {
			const response = await valistInstance.getLatestReleaseFromRepo(orgShortName, projectName);
			expect(response).to.have.property("releaseCID");
			expect(response).to.have.property("metaCID");
		});
	});
});
