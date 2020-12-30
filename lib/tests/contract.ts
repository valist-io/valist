import Web3 from 'web3';
import { expect } from 'chai';
import { getValistContract, getContractInstance } from '../src/index';

// @ts-ignore
import ValistABI from '../dist/abis/Valist.json';
import Valist from '../src/index';

/*
npm install -g typings
typings install dt~mocha --global --save
typings install npm~chai --save
*/

describe('Call getContractInstance', () => {
	const contract_address = "0xa1517c736b9810cC6F108e7CC9b66fA92d3F290F"
	const web3 = new Web3(new Web3.providers.HttpProvider('http://localhost:9545'));
	const contractInstance = getContractInstance(web3, ValistABI.abi, contract_address)

	it('Should be an object', () => {
		expect(contractInstance).to.be.an('object');
	});

	it('Should contain createOrganization', () => {
		expect(contractInstance.methods.createOrganization).to.exist;
	}); 
});

describe('Call getValistContract', () => {
	const web3 = new Web3(new Web3.providers.HttpProvider('http://localhost:9545'));

	it('Should be an object', async () => {
		let valistContract = await getValistContract(web3)
		expect(valistContract).to.be.an('object');
	});

	it('Should contain createOrganization', async () => {
		let valistContract = await getValistContract(web3)
		expect(valistContract.methods.createOrganization).to.exist;
	}); 
});

describe('Create new Valist Instance', () => {
	const web3Provider = new Web3.providers.HttpProvider('http://localhost:9545');

	it('Should be an object', () => {
		expect(new Valist(web3Provider, false)).to.be.an('object');
	});

	it('Should connect to valist contract and return valist attribute as object', async () => {
		let valist = new Valist(web3Provider, false)
		await valist.connect()

		expect(valist.valist).to.be.an('object');
	});
});

