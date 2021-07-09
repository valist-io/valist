/* eslint-disable @typescript-eslint/no-explicit-any */
/* eslint-disable @typescript-eslint/explicit-module-boundary-types */
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

const functionIDMap: Record<string, string> = {
  createOrganization: '6adea8ac-c764-45ed-a62c-e399adf7dd47',
  createRepository: '1dd94ae2-fcd1-4951-896d-a28e1445f1ae',
  voteRelease: 'db8b1576-478b-45ba-a7bf-e6cb9c2b93a4',
  voteKey: '3d2f99cf-02c7-44b6-9db7-4e672daa248b',
  voteThreshold: '9d163d2f-d1d7-4958-af5f-e2859cdad3f8',
  setOrgMeta: '152a40ff-a730-4549-8657-745bc6121e27',
  setRepoMeta: '5e460491-1d76-4bfb-9a2c-6f00f0be5ada',
  clearPendingRelease: '33c9cc4d-d17b-4c4f-a833-ead141de83ee',
  clearPendingKey: '7b0750fa-a4fc-4874-841a-274287027063',
  clearPendingThreshold: '4b3d868b-f759-4c48-904b-f49938157edc',
};

// pass the networkId to get contract addresses
const getContractAddresses = async (networkId: number) => {
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
const getBiconomyForwarderConfig = async (networkId: number) => {
  // get trusted forwarder contract address from network id
  const contractAddresses = await getContractAddresses(networkId);
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
const buildForwardTxRequest = async ({
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
const getDataToSignForEIP712 = async (web3: any, request: any, networkId: number) => {
  const contractAddresses = await getContractAddresses(networkId);
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
// eslint-disable-next-line
const getDomainSeperator = async (web3: any, networkId: number) => {
  const contractAddresses = await getContractAddresses(networkId);
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

export {
  functionIDMap,
  getDomainSeperator,
  getDataToSignForEIP712,
  buildForwardTxRequest,
  getBiconomyForwarderConfig,
};
