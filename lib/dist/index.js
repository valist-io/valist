"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.Web3Providers = void 0;
const web3_1 = __importDefault(require("web3"));
// @ts-ignore
const ipfs_http_client_1 = __importDefault(require("ipfs-http-client"));
const Valist_json_1 = __importDefault(require("./abis/Valist.json"));
const ValistOrganization_json_1 = __importDefault(require("./abis/ValistOrganization.json"));
const ValistRepository_json_1 = __importDefault(require("./abis/ValistRepository.json"));
// node-fetch polyfill
const fetch = require('node-fetch');
if (!globalThis.fetch) {
    globalThis.fetch = fetch;
}
// keccak256 hashes of each role
const ORG_ADMIN_ROLE = "0x123b642491709420c2370bb98c4e7de2b1bc05c5f9fd95ac4111e12683553c62";
const REPO_ADMIN_ROLE = "0xff7d2294a3c189284afb74beb7d578b566cf69863d5cb16db08773c21bea56c9";
const REPO_DEV_ROLE = "0x069bf569f27d389f2c70410107860b2e82ff561283b097a89e897daa5e34b1b6";
const shortnameFilterRegex = /[^A-z-]/;
const getContractInstance = (web3, abi, address) => __awaiter(void 0, void 0, void 0, function* () {
    // create the instance
    return new web3.eth.Contract(abi, address);
});
const getValistContract = (web3) => __awaiter(void 0, void 0, void 0, function* () {
    // get network ID and the deployed address
    const networkId = yield web3.eth.net.getId();
    // @ts-ignore
    const deployedAddress = Valist_json_1.default.networks[networkId].address;
    return yield getContractInstance(web3, Valist_json_1.default.abi, deployedAddress);
});
const getValistOrganizationContract = (web3, address) => __awaiter(void 0, void 0, void 0, function* () {
    // create the instance
    return yield getContractInstance(web3, ValistOrganization_json_1.default.abi, address);
});
const getValistRepositoryContract = (web3, address) => __awaiter(void 0, void 0, void 0, function* () {
    // get network ID and the deployed address
    return yield getContractInstance(web3, ValistRepository_json_1.default.abi, address);
});
class Valist {
    constructor(web3Provider, ipfsEnabled) {
        this.web3 = new web3_1.default(web3Provider);
        this.defaultAccount = "0x0";
        if (ipfsEnabled) {
            this.ipfs = ipfs_http_client_1.default({ host: `ipfs.infura.io`, port: `5001`, apiPath: `/api/v0/`, protocol: `https` });
        }
    }
    // initialize main valist contract instance for future calls
    connect() {
        return __awaiter(this, void 0, void 0, function* () {
            try {
                this.valist = yield getValistContract(this.web3);
            }
            catch (e) {
                const msg = `Could not connect to Valist registry contract`;
                console.error(msg, e);
                throw e;
            }
            try {
                const accounts = yield this.web3.eth.getAccounts();
                this.defaultAccount = accounts[0] || "0x0";
            }
            catch (e) {
                const msg = `Could not set default account`;
                console.error(msg, e);
                throw e;
            }
        });
    }
    // returns organization contract instance
    getOrganization(orgName) {
        return __awaiter(this, void 0, void 0, function* () {
            try {
                const orgAddress = yield this.valist.methods.orgs(orgName).call();
                const org = yield getValistOrganizationContract(this.web3, orgAddress);
                return org;
            }
            catch (e) {
                const msg = `Could not get organization contract`;
                console.error(msg, e);
                throw e;
            }
        });
    }
    getOrganizationMeta(orgName) {
        return __awaiter(this, void 0, void 0, function* () {
            try {
                const org = yield this.getOrganization(orgName);
                const orgMeta = yield org.methods.orgMeta().call();
                const json = yield this.fetchJSONfromIPFS(orgMeta);
                return json;
            }
            catch (e) {
                const msg = `Could not get organization metadata`;
                console.error(msg, e);
                throw e;
            }
        });
    }
    getCreatedOrganizations() {
        return __awaiter(this, void 0, void 0, function* () {
            try {
                const organizations = yield this.valist.getPastEvents('OrganizationCreated', { fromBlock: 0, toBlock: 'latest' });
                return organizations;
            }
            catch (e) {
                const msg = `Could not get created organizations`;
                console.error(msg, e);
                throw e;
            }
        });
    }
    getDeletedOrganizations() {
        return __awaiter(this, void 0, void 0, function* () {
            try {
                const organizations = yield this.valist.getPastEvents('OrganizationDeleted', { fromBlock: 0, toBlock: 'latest' });
                return organizations;
            }
            catch (e) {
                const msg = `Could not get deleted organizations`;
                console.error(msg, e);
                throw e;
            }
        });
    }
    setOrgMeta(orgName, orgMeta, account) {
        return __awaiter(this, void 0, void 0, function* () {
            try {
                const org = yield this.getOrganization(orgName);
                const hash = yield this.addJSONtoIPFS(orgMeta);
                yield org.methods.updateOrgMeta(hash).send({ from: account });
            }
            catch (e) {
                const msg = `Could not set organization metadata`;
                console.error(msg, e);
                throw e;
            }
        });
    }
    // returns repository contract instance
    getRepository(orgName, repoName) {
        return __awaiter(this, void 0, void 0, function* () {
            try {
                const org = yield this.getOrganization(orgName);
                const repoAddress = yield org.methods.repos(repoName).call();
                const repo = yield getValistRepositoryContract(this.web3, repoAddress);
                return repo;
            }
            catch (e) {
                const msg = `Could not get repository contract`;
                console.error(msg, e);
                throw e;
            }
        });
    }
    getReposFromOrganization(orgName) {
        return __awaiter(this, void 0, void 0, function* () {
            try {
                const org = yield this.getOrganization(orgName);
                const repos = yield org.getPastEvents('RepositoryCreated', { fromBlock: 0, toBlock: 'latest' });
                return repos;
            }
            catch (e) {
                const msg = `Could not get repositories from organization`;
                console.error(msg, e);
                throw e;
            }
        });
    }
    getRepoMeta(orgName, repoName) {
        return __awaiter(this, void 0, void 0, function* () {
            try {
                const repo = yield this.getRepository(orgName, repoName);
                const repoMeta = yield repo.methods.repoMeta().call();
                const json = yield this.fetchJSONfromIPFS(repoMeta);
                return json;
            }
            catch (e) {
                const msg = `Could not get repository metadata`;
                console.error(msg, e);
                throw e;
            }
        });
    }
    setRepoMeta(orgName, repoName, repoMeta, account) {
        return __awaiter(this, void 0, void 0, function* () {
            try {
                const repo = yield this.getRepository(orgName, repoName);
                const hash = yield this.addJSONtoIPFS(repoMeta);
                yield repo.methods.updateRepoMeta(hash).send({ from: account });
            }
            catch (e) {
                const msg = `Could not set repository metadata`;
                console.error(msg, e);
                throw e;
            }
        });
    }
    getLatestTagFromRepo(orgName, repoName) {
        return __awaiter(this, void 0, void 0, function* () {
            try {
                const repo = yield this.getRepository(orgName, repoName);
                const tag = yield repo.methods.tag().call();
                return tag;
            }
            catch (e) {
                const msg = `Could not get latest tag from repo`;
                console.error(msg, e);
                throw e;
            }
        });
    }
    getLatestReleaseFromRepo(orgName, repoName) {
        return __awaiter(this, void 0, void 0, function* () {
            try {
                const repo = yield this.getRepository(orgName, repoName);
                const release = yield repo.methods.latestRelease().call();
                return release;
            }
            catch (e) {
                const msg = `Could not get latest release from repo`;
                console.error(msg, e);
                throw e;
            }
        });
    }
    getLatestReleaseMetaFromRepo(orgName, repoName) {
        return __awaiter(this, void 0, void 0, function* () {
            try {
                const repo = yield this.getRepository(orgName, repoName);
                const release = yield repo.methods.releaseMeta().call();
                return release;
            }
            catch (e) {
                const msg = `Could not get latest release metadata from repo`;
                console.error(msg, e);
                throw e;
            }
        });
    }
    getReleasesFromRepo(orgName, repoName) {
        return __awaiter(this, void 0, void 0, function* () {
            try {
                const repo = yield this.getRepository(orgName, repoName);
                return yield repo.getPastEvents('Release', { fromBlock: 0, toBlock: 'latest' });
            }
            catch (e) {
                const msg = `Could not get releases from repo`;
                console.error(msg, e);
                throw e;
            }
        });
    }
    getReleaseByTag(orgName, repoName, tag) {
        return __awaiter(this, void 0, void 0, function* () {
            try {
                const events = yield this.getReleasesFromRepo(orgName, repoName);
                // @TODO make this more efficient later
                for (let i = 0; i < events.length; i++) {
                    if (events[i].returnValues.tag == tag) {
                        const { tag, release, releaseMeta } = events[i].returnValues;
                        return { tag, release, releaseMeta };
                    }
                }
                return;
            }
            catch (e) {
                const msg = `Could not get release by tag`;
                console.error(msg, e);
                throw e;
            }
        });
    }
    isOrgAdmin(orgName, account) {
        return __awaiter(this, void 0, void 0, function* () {
            try {
                const org = yield this.getOrganization(orgName);
                return yield org.methods.hasRole(ORG_ADMIN_ROLE, account).call();
            }
            catch (e) {
                const msg = `Could not check if user has ORG_ADMIN_ROLE`;
                console.error(msg, e);
                throw e;
            }
        });
    }
    isRepoAdmin(orgName, repoName, account) {
        return __awaiter(this, void 0, void 0, function* () {
            try {
                const repo = yield this.getRepository(orgName, repoName);
                return yield repo.methods.hasRole(REPO_ADMIN_ROLE, account).call();
            }
            catch (e) {
                const msg = `Could not check if user has REPO_ADMIN_ROLE`;
                console.error(msg, e);
                throw e;
            }
        });
    }
    isRepoDev(orgName, repoName, account) {
        return __awaiter(this, void 0, void 0, function* () {
            try {
                const repo = yield this.getRepository(orgName, repoName);
                return yield repo.methods.hasRole(REPO_DEV_ROLE, account).call();
            }
            catch (e) {
                const msg = `Could not check if user has REPO_DEV_ROLE`;
                console.error(msg, e);
                throw e;
            }
        });
    }
    grantOrgAdmin(orgName, granter, grantee) {
        return __awaiter(this, void 0, void 0, function* () {
            try {
                const org = yield this.getOrganization(orgName);
                yield org.methods.grantRole(ORG_ADMIN_ROLE, grantee).send({ from: granter });
            }
            catch (e) {
                const msg = `Could not grant ORG_ADMIN_ROLE`;
                console.error(msg, e);
                throw e;
            }
        });
    }
    revokeOrgAdmin(orgName, revoker, revokee) {
        return __awaiter(this, void 0, void 0, function* () {
            try {
                const org = yield this.getOrganization(orgName);
                yield org.methods.revokeRole(ORG_ADMIN_ROLE, revokee).send({ from: revoker });
            }
            catch (e) {
                const msg = `Could not revoke ORG_ADMIN_ROLE`;
                console.error(msg, e);
                throw e;
            }
        });
    }
    grantRepoAdmin(orgName, repoName, granter, grantee) {
        return __awaiter(this, void 0, void 0, function* () {
            try {
                const repo = yield this.getRepository(orgName, repoName);
                return yield repo.methods.grantRole(REPO_ADMIN_ROLE, grantee).send({ from: granter });
            }
            catch (e) {
                const msg = `Could not grant REPO_ADMIN_ROLE`;
                console.error(msg, e);
                throw e;
            }
        });
    }
    revokeRepoAdmin(orgName, repoName, revoker, revokee) {
        return __awaiter(this, void 0, void 0, function* () {
            try {
                const repo = yield this.getRepository(orgName, repoName);
                return yield repo.methods.revokeRole(REPO_ADMIN_ROLE, revokee).send({ from: revoker });
            }
            catch (e) {
                const msg = `Could not revoke REPO_ADMIN_ROLE`;
                console.error(msg, e);
                throw e;
            }
        });
    }
    grantRepoDev(orgName, repoName, granter, grantee) {
        return __awaiter(this, void 0, void 0, function* () {
            try {
                const repo = yield this.getRepository(orgName, repoName);
                return yield repo.methods.grantRole(REPO_DEV_ROLE, grantee).send({ from: granter });
            }
            catch (e) {
                const msg = `Could not grant REPO_DEV_ROLE`;
                console.error(msg, e);
                throw e;
            }
        });
    }
    revokeRepoDev(orgName, repoName, revoker, revokee) {
        return __awaiter(this, void 0, void 0, function* () {
            try {
                const repo = yield this.getRepository(orgName, repoName);
                return yield repo.methods.revokeRole(REPO_DEV_ROLE, revokee).send({ from: revoker });
            }
            catch (e) {
                const msg = `Could not revoke REPO_DEV_ROLE`;
                console.error(msg, e);
                throw e;
            }
        });
    }
    getOrgAdmins(orgName) {
        return __awaiter(this, void 0, void 0, function* () {
            try {
                const org = yield this.getOrganization(orgName);
                const adminCount = yield org.methods.getRoleMemberCount(ORG_ADMIN_ROLE).call();
                const members = [];
                for (let i = 0; i < adminCount; ++i) {
                    members.push(yield org.methods.getRoleMember(ORG_ADMIN_ROLE, i).call());
                }
                return members;
            }
            catch (e) {
                const msg = `Could not get users that have REPO_ADMIN_ROLE`;
                console.error(msg, e);
                throw e;
            }
        });
    }
    getRepoAdmins(orgName, repoName) {
        return __awaiter(this, void 0, void 0, function* () {
            try {
                const repo = yield this.getRepository(orgName, repoName);
                const adminCount = yield repo.methods.getRoleMemberCount(REPO_ADMIN_ROLE).call();
                console.log(adminCount);
                const members = [];
                for (let i = 0; i < adminCount; ++i) {
                    members.push(yield repo.methods.getRoleMember(REPO_ADMIN_ROLE, i).call());
                }
                return members;
            }
            catch (e) {
                const msg = `Could not get users that have REPO_ADMIN_ROLE`;
                console.error(msg, e);
                throw e;
            }
        });
    }
    getRepoDevs(orgName, repoName) {
        return __awaiter(this, void 0, void 0, function* () {
            try {
                const repo = yield this.getRepository(orgName, repoName);
                const devCount = yield repo.methods.getRoleMemberCount(REPO_DEV_ROLE).call();
                const members = [];
                for (let i = 0; i < devCount; ++i) {
                    members.push(yield repo.methods.getRoleMember(REPO_DEV_ROLE, i).call());
                }
                return members;
            }
            catch (e) {
                const msg = `Could not get users that have REPO_DEV_ROLE`;
                console.error(msg, e);
                throw e;
            }
        });
    }
    createOrganization(orgName, orgMeta, account) {
        return __awaiter(this, void 0, void 0, function* () {
            try {
                const metaFile = yield this.addJSONtoIPFS(orgMeta);
                const result = yield this.valist.methods.createOrganization(orgName.toLowerCase().replace(shortnameFilterRegex, ""), metaFile).send({ from: account });
                return result;
            }
            catch (e) {
                const msg = `Could not create organization`;
                console.error(msg, e);
                throw e;
            }
        });
    }
    createRepository(orgName, repoName, repoMeta, account) {
        return __awaiter(this, void 0, void 0, function* () {
            try {
                const org = yield this.getOrganization(orgName);
                const metaFile = yield this.addJSONtoIPFS(repoMeta);
                const result = yield org.methods.createRepository(repoName.toLowerCase().replace(shortnameFilterRegex, ""), metaFile).send({ from: account });
                return result;
            }
            catch (e) {
                const msg = `Could not create repository`;
                console.error(msg, e);
                throw e;
            }
        });
    }
    publishRelease(orgName, repoName, release, account) {
        return __awaiter(this, void 0, void 0, function* () {
            try {
                const repo = yield this.getRepository(orgName, repoName);
                const result = yield repo.methods.publishRelease(release.tag, release.hash, release.meta).send({ from: account });
                return result;
            }
            catch (e) {
                const msg = `Could not publish release`;
                console.error(msg, e);
                throw e;
            }
        });
    }
    addJSONtoIPFS(data) {
        return __awaiter(this, void 0, void 0, function* () {
            try {
                const file = Buffer.from(JSON.stringify(data));
                const result = yield this.ipfs.add(file);
                return result["path"];
            }
            catch (e) {
                const msg = `Could not add JSON to IPFS`;
                console.error(msg, e);
                throw e;
            }
        });
    }
    addFileToIPFS(data) {
        return __awaiter(this, void 0, void 0, function* () {
            try {
                const file = Buffer.from(data);
                const result = yield this.ipfs.add(file);
                return result["path"];
            }
            catch (e) {
                const msg = `Could not add file to IPFS`;
                console.error(msg, e);
                throw e;
            }
        });
    }
    fetchJSONfromIPFS(ipfsHash) {
        return __awaiter(this, void 0, void 0, function* () {
            try {
                const response = yield fetch(`https://cloudflare-ipfs.com/ipfs/${ipfsHash}`);
                const json = yield response.json();
                console.log(`JSON Fetched from IPFS`, json);
                return json;
            }
            catch (e) {
                const msg = `Could not fetch JSON from IPFS`;
                console.error(msg, e);
                throw e;
            }
        });
    }
}
exports.Web3Providers = web3_1.default.providers;
exports.default = Valist;
//# sourceMappingURL=index.js.map