import * as sigUtil from 'eth-sig-util';
/* eslint-disable @typescript-eslint/no-explicit-any */
/* eslint-disable @typescript-eslint/explicit-module-boundary-types */
/* eslint-disable no-underscore-dangle */
const helperAttributes: any = {};
helperAttributes.ZERO_ADDRESS = '0x0000000000000000000000000000000000000000';
helperAttributes.baseURL = 'https://api.biconomy.io';

// eslint-disable-next-line
helperAttributes.biconomyForwarderAbi = [{'inputs':[{'internalType':'address','name':'_owner','type':'address'}],'stateMutability':'nonpayable','type':'constructor'},{'anonymous':false,'inputs':[{'indexed':true,'internalType':'bytes32','name':'domainSeparator','type':'bytes32'},{'indexed':false,'internalType':'bytes','name':'domainValue','type':'bytes'}],'name':'DomainRegistered','type':'event'},{'anonymous':false,'inputs':[{'indexed':true,'internalType':'address','name':'previousOwner','type':'address'},{'indexed':true,'internalType':'address','name':'newOwner','type':'address'}],'name':'OwnershipTransferred','type':'event'},{'inputs':[],'name':'EIP712_DOMAIN_TYPE','outputs':[{'internalType':'string','name':'','type':'string'}],'stateMutability':'view','type':'function'},{'inputs':[],'name':'REQUEST_TYPEHASH','outputs':[{'internalType':'bytes32','name':'','type':'bytes32'}],'stateMutability':'view','type':'function'},{'inputs':[{'internalType':'bytes32','name':'','type':'bytes32'}],'name':'domains','outputs':[{'internalType':'bool','name':'','type':'bool'}],'stateMutability':'view','type':'function'},{'inputs':[{'components':[{'internalType':'address','name':'from','type':'address'},{'internalType':'address','name':'to','type':'address'},{'internalType':'address','name':'token','type':'address'},{'internalType':'uint256','name':'txGas','type':'uint256'},{'internalType':'uint256','name':'tokenGasPrice','type':'uint256'},{'internalType':'uint256','name':'batchId','type':'uint256'},{'internalType':'uint256','name':'batchNonce','type':'uint256'},{'internalType':'uint256','name':'deadline','type':'uint256'},{'internalType':'bytes','name':'data','type':'bytes'}],'internalType':'structERC20ForwardRequestTypes.ERC20ForwardRequest','name':'req','type':'tuple'},{'internalType':'bytes32','name':'domainSeparator','type':'bytes32'},{'internalType':'bytes','name':'sig','type':'bytes'}],'name':'executeEIP712','outputs':[{'internalType':'bool','name':'success','type':'bool'},{'internalType':'bytes','name':'ret','type':'bytes'}],'stateMutability':'nonpayable','type':'function'},{'inputs':[{'components':[{'internalType':'address','name':'from','type':'address'},{'internalType':'address','name':'to','type':'address'},{'internalType':'address','name':'token','type':'address'},{'internalType':'uint256','name':'txGas','type':'uint256'},{'internalType':'uint256','name':'tokenGasPrice','type':'uint256'},{'internalType':'uint256','name':'batchId','type':'uint256'},{'internalType':'uint256','name':'batchNonce','type':'uint256'},{'internalType':'uint256','name':'deadline','type':'uint256'},{'internalType':'bytes','name':'data','type':'bytes'}],'internalType':'structERC20ForwardRequestTypes.ERC20ForwardRequest','name':'req','type':'tuple'},{'internalType':'bytes','name':'sig','type':'bytes'}],'name':'executePersonalSign','outputs':[{'internalType':'bool','name':'success','type':'bool'},{'internalType':'bytes','name':'ret','type':'bytes'}],'stateMutability':'nonpayable','type':'function'},{'inputs':[{'internalType':'address','name':'from','type':'address'},{'internalType':'uint256','name':'batchId','type':'uint256'}],'name':'getNonce','outputs':[{'internalType':'uint256','name':'','type':'uint256'}],'stateMutability':'view','type':'function'},{'inputs':[],'name':'isOwner','outputs':[{'internalType':'bool','name':'','type':'bool'}],'stateMutability':'view','type':'function'},{'inputs':[],'name':'owner','outputs':[{'internalType':'address','name':'','type':'address'}],'stateMutability':'view','type':'function'},{'inputs':[{'internalType':'string','name':'name','type':'string'},{'internalType':'string','name':'version','type':'string'}],'name':'registerDomainSeparator','outputs':[],'stateMutability':'nonpayable','type':'function'},{'inputs':[],'name':'renounceOwnership','outputs':[],'stateMutability':'nonpayable','type':'function'},{'inputs':[{'internalType':'address','name':'newOwner','type':'address'}],'name':'transferOwnership','outputs':[],'stateMutability':'nonpayable','type':'function'},{'inputs':[{'components':[{'internalType':'address','name':'from','type':'address'},{'internalType':'address','name':'to','type':'address'},{'internalType':'address','name':'token','type':'address'},{'internalType':'uint256','name':'txGas','type':'uint256'},{'internalType':'uint256','name':'tokenGasPrice','type':'uint256'},{'internalType':'uint256','name':'batchId','type':'uint256'},{'internalType':'uint256','name':'batchNonce','type':'uint256'},{'internalType':'uint256','name':'deadline','type':'uint256'},{'internalType':'bytes','name':'data','type':'bytes'}],'internalType':'structERC20ForwardRequestTypes.ERC20ForwardRequest','name':'req','type':'tuple'},{'internalType':'bytes32','name':'domainSeparator','type':'bytes32'},{'internalType':'bytes','name':'sig','type':'bytes'}],'name':'verifyEIP712','outputs':[],'stateMutability':'view','type':'function'},{'inputs':[{'components':[{'internalType':'address','name':'from','type':'address'},{'internalType':'address','name':'to','type':'address'},{'internalType':'address','name':'token','type':'address'},{'internalType':'uint256','name':'txGas','type':'uint256'},{'internalType':'uint256','name':'tokenGasPrice','type':'uint256'},{'internalType':'uint256','name':'batchId','type':'uint256'},{'internalType':'uint256','name':'batchNonce','type':'uint256'},{'internalType':'uint256','name':'deadline','type':'uint256'},{'internalType':'bytes','name':'data','type':'bytes'}],'internalType':'structERC20ForwardRequestTypes.ERC20ForwardRequest','name':'req','type':'tuple'},{'internalType':'bytes','name':'sig','type':'bytes'}],'name':'verifyPersonalSign','outputs':[],'stateMutability':'view','type':'function'}];

helperAttributes.biconomyForwarderDomainData = {
  name: 'Biconomy Forwarder',
  version: '1',
};

helperAttributes.domainType = [
  { name: 'name', type: 'string' },
  { name: 'version', type: 'string' },
  { name: 'verifyingContract', type: 'address' },
  { name: 'salt', type: 'bytes32' },
];

helperAttributes.forwardRequestType = [
  { name: 'from', type: 'address' },
  { name: 'to', type: 'address' },
  { name: 'token', type: 'address' },
  { name: 'txGas', type: 'uint256' },
  { name: 'tokenGasPrice', type: 'uint256' },
  { name: 'batchId', type: 'uint256' },
  { name: 'batchNonce', type: 'uint256' },
  { name: 'deadline', type: 'uint256' },
  { name: 'data', type: 'bytes' },
];

/* on dashboard page, this script helps scrape these values quickly:

const table = document.getElementsByTagName('table')[0]
let functionIDMap = {}
for (let i = 1; i < table.rows.length; ++i) {
    functionIDMap[table.rows[i].children[1].textContent] = table.rows[i].children[2].textContent
}

*/
const functionIDMap: Record<string, string> = {
  clearPendingKey: 'a0dfd7b2-fb2b-46da-a662-3cbb87c7b83e',
  clearPendingRelease: 'b95d7f2d-6d40-4690-b7df-ec36928aaf77',
  clearPendingThreshold: 'f154fe5a-cd81-4a31-8536-6ea999795f56',
  createOrganization: '7cb293ac-5ed6-4dd8-9956-eb5a9a236403',
  createRepository: '3b40c07a-d9dd-401a-913b-ef395648ba4d',
  setOrgMeta: '1292cba4-8b4e-4828-8989-e2583017cda7',
  setRepoMeta: '1857aa6a-b334-4b6a-bf7c-959d5581e8d4',
  voteKey: '82d84700-7a9a-44f5-865d-f34badb00852',
  voteRelease: 'c8fc037a-dc5c-4fe3-b2fd-f8c602986d72',
  voteThreshold: 'f0b640b6-4280-4cf0-afca-0d62046cee09',
  grantRole: '17ec42d7-9f19-407c-8131-3033f7dcc142',
  init: '5336e4c2-fc5c-49bd-b41d-9990dde03982',
  linkNameToID: '8fc893ff-08e1-4cda-9264-62f6467d91a8',
  overrideNameToID: '0455fbcd-4d1e-45ec-b0ce-5eaf73169b3e',
  renounceRole: '08c8a75f-e9d2-4e9d-82e9-8f6c5b2bf8a0',
  revokeRole: 'd4040355-b755-4a1a-9f16-0f0462bd56d1',
};

// pass the networkId to get GSN forwarder contract addresses
const getContractAddresses = (networkId: number) => {
  const addressMap: Record<number, string> = {
    80001: '0x9399BB24DBB5C4b782C70c2969F58716Ebbd6a3b',
    137: '0x86C80a8aa58e0A4fa09A69624c31Ab2a6CAD56b8',
    1: '0x84a0856b038eaAd1cC7E297cF34A7e72685A8693',
  };
  const contractAddresses: any = {};
  contractAddresses.biconomyForwarderAddress = addressMap[networkId || 80001];
  return contractAddresses;
};

/**
 * Returns ABI and contract address based on network Id
 * You can build biconomy forwarder contract object using above values and calculate the nonce
 * @param {*} networkId
 */
const getBiconomyForwarderConfig = (networkId: number) => {
  // get trusted forwarder contract address from network id
  const contractAddresses = getContractAddresses(networkId);
  const forwarderAddress = contractAddresses.biconomyForwarderAddress;
  return { abi: helperAttributes.biconomyForwarderAbi, address: forwarderAddress };
};

/**
 * pass the below params in any order e.g. account=<account>,batchNone=<batchNone>,...
 * @param {*}  account - from (end user's) address for this transaction
 * @param {*}  to - target recipient contract address
 * @param {*}  gasLimitNum - gas estimation of your target method in numeric format
 * @param {*}  batchId - batchId
 * @param {*}  batchNonce - batchNonce which can be verified and obtained from the biconomy forwarder
 * @param {*}  data - functionSignature of target method
 * @param {*}  deadline - optional deadline for this forward request
 */
const buildForwardTxRequest = ({
  account, to, gasLimitNum, batchId, batchNonce, data, deadline,
}: any) => {
  const req = {
    from: account,
    to,
    token: helperAttributes.ZERO_ADDRESS,
    txGas: gasLimitNum,
    tokenGasPrice: '0',
    // eslint-disable-next-line
    batchId: parseInt(batchId),
    // eslint-disable-next-line
    batchNonce: parseInt(batchNonce),
    deadline: deadline || Math.floor(Date.now() / 1000 + 3600),
    data,
  };
  return req;
};

/**
 * pass your forward request and network Id
 * use this method to build message to be signed by end user in EIP712 signature format
 * @param {*} request - forward request object
 * @param {*} networkId
 */
const getDataToSignForEIP712 = (web3: any, request: any, networkId: number) => {
  const contractAddresses = getContractAddresses(networkId);
  const forwarderAddress = contractAddresses.biconomyForwarderAddress;
  const domainData = helperAttributes.biconomyForwarderDomainData;
  domainData.salt = web3.utils.padLeft(web3.utils.toHex(web3.utils.toBN(networkId)), 64);
  domainData.verifyingContract = forwarderAddress;

  const dataToSign = JSON.stringify({
    types: {
      EIP712Domain: helperAttributes.domainType,
      ERC20ForwardRequest: helperAttributes.forwardRequestType,
    },
    domain: domainData,
    primaryType: 'ERC20ForwardRequest',
    message: request,
  });
  return dataToSign;
};

/**
 * get the domain seperator that needs to be passed while using EIP712 signature type
 * @param {*} networkId
 */
const getDomainSeperator = (web3: any, networkId: number) => {
  const contractAddresses = getContractAddresses(networkId);
  const forwarderAddress = contractAddresses.biconomyForwarderAddress;
  const domainData = helperAttributes.biconomyForwarderDomainData;
  domainData.salt = web3.utils.padLeft(web3.utils.toHex(web3.utils.toBN(networkId)), 64);
  domainData.verifyingContract = forwarderAddress;

  const domainSeparator = web3.utils.keccak256(
    web3.eth.abi.encodeParameters([
      'bytes32',
      'bytes32',
      'bytes32',
      'address',
      'bytes32',
    ], [
      web3.utils.keccak256('EIP712Domain(string name,string version,address verifyingContract,bytes32 salt)'),
      web3.utils.keccak256(domainData.name),
      web3.utils.keccak256(domainData.version),
      domainData.verifyingContract,
      domainData.salt,
    ]),
  );
  return domainSeparator;
};

const sendAsync = (web3: any, params: any): any => new Promise((resolve, reject) => {
  // eslint-disable-next-line @typescript-eslint/ban-ts-comment
  // @ts-ignore sendAsync is conflicting with the RPCProvider type
  web3.currentProvider.sendAsync(params, async (e: any, signed: any) => {
    if (e) reject(e);
    resolve(signed);
  });
});

const sendMetaTx = async (
  web3: any,
  networkID: number,
  functionCall: any,
  account: string,
  gasLimit: string | number,
  signer?: string,
) => {
  const forwarder = getBiconomyForwarderConfig(networkID);
  const forwarderContract = new web3.eth.Contract(forwarder.abi, forwarder.address);
  const batchNonce = await forwarderContract.methods.getNonce(account, 0).call();
  const functionSignature = functionCall.encodeABI();

  const functionName: string = functionCall._method.name;
  const contractAddress: string = functionCall._parent._address;

  const request = buildForwardTxRequest({
    account,
    to: contractAddress,
    gasLimitNum: gasLimit,
    batchId: 0,
    batchNonce,
    data: functionSignature,
    deadline: '',
  });

  const domainSeparator = getDomainSeperator(web3, networkID);
  const dataToSign = getDataToSignForEIP712(web3, request, networkID);
  let signed;

  if (signer) {
    signed = sigUtil.signTypedData_v4(Buffer.from(signer, 'hex'), { data: JSON.parse(dataToSign) });
  } else {
    const sig = await sendAsync(web3, {
      jsonrpc: '2.0',
      id: new Date().getTime(),
      method: 'eth_signTypedData_v4',
      params: [account, dataToSign],
    });
    signed = sig.result;
  }

  const resp = await fetch('https://api.biconomy.io/api/v2/meta-tx/native', {
    method: 'POST',
    headers: {
      // public biconomy key
      'x-api-key': 'qLW9TRUjQ.f77d2f86-c76a-4b9c-b1ee-0453d0ead878',
      'Content-Type': 'application/json;charset=utf-8',
    },
    body: JSON.stringify({
      to: contractAddress,
      apiId: functionIDMap[functionName],
      params: [request, domainSeparator, signed],
      from: account,
      signatureType: 'EIP712_SIGN',
    }),
  });

  let txHash = await resp.json();
  txHash = txHash.txHash;

  const getTransactionReceiptMined = async () => {
    const transactionReceiptAsync = async (resolve: any, reject: any) => {
      const receipt = await web3.eth.getTransactionReceipt(txHash);
      if (receipt == null) {
        setTimeout(() => transactionReceiptAsync(resolve, reject), 500);
      } else {
        resolve(receipt);
      }
    };

    const receipt = await new Promise(transactionReceiptAsync);
    return receipt;
  };
  const receipt = await getTransactionReceiptMined();
  return receipt;
};

export {
  sendMetaTx,
};
