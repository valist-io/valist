import { expect } from 'chai';
import { getContractInstance } from '../src/index';
import Web3 from 'web3';

// @ts-ignore
import ValistABI from '../dist/abis/Valist.json';

// @ts-ignore
describe('Get Valist Contract Instance', () => {
	const contract_address = "0xa1517c736b9810cC6F108e7CC9b66fA92d3F290F"
	const web3 = new Web3(new Web3.providers.HttpProvider('http://localhost:9545'));

	// @ts-ignore
	it('getContractInstance', () => {
		expect(getContractInstance(web3, ValistABI.abi, contract_address).methods.createOrganization).to.exist
	})
})