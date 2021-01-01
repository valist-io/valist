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
exports.Web3Providers = exports.shortnameFilterRegex = void 0;
const web3_1 = __importDefault(require("web3"));
// @ts-ignore
const ipfs_http_client_1 = __importDefault(require("ipfs-http-client"));
const Valist_json_1 = __importDefault(require("./abis/Valist.json"));
// node-fetch polyfill
const fetch = require('node-fetch');
if (!globalThis.fetch) {
    globalThis.fetch = fetch;
}
exports.shortnameFilterRegex = /[^A-z0-9-]/;
const getContractInstance = (web3, abi, address) => {
    // create the instance
    return new web3.eth.Contract(abi, address);
};
const getValistContract = (web3) => __awaiter(void 0, void 0, void 0, function* () {
    // get network ID and the deployed address
    const networkId = yield web3.eth.net.getId();
    // @ts-ignore
    const deployedAddress = Valist_json_1.default.networks[networkId].address;
    return getContractInstance(web3, Valist_json_1.default.abi, deployedAddress);
});
class Valist {
    constructor({ web3Provider, metaTx = false, ipfsHost = `ipfs.infura.io` }) {
        if (metaTx) {
            this.web3 = new web3_1.default(web3Provider);
        }
        else {
            this.web3 = new web3_1.default(web3Provider);
        }
        this.defaultAccount = "0x0";
        this.ipfs = ipfs_http_client_1.default({ host: ipfsHost, port: 5001, apiPath: `/api/v0/`, protocol: `https` });
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
    // returns organization meta and release tags
    getOrganization(orgName) {
        return __awaiter(this, void 0, void 0, function* () {
            try {
                const org = yield this.valist.methods.getOrganization(orgName).call();
                let json = {};
                try {
                    json = yield this.fetchJSONfromIPFS(org[0]);
                }
                catch (e) { }
                return { meta: json, repoNames: org[1] };
            }
            catch (e) {
                const msg = `Could not get organization`;
                console.error(msg, e);
                throw e;
            }
        });
    }
    getOrganizationMeta(orgName) {
        return __awaiter(this, void 0, void 0, function* () {
            try {
                const orgMeta = yield this.valist.methods.getOrgMeta(orgName).call();
                let json = {};
                try {
                    json = yield this.fetchJSONfromIPFS(orgMeta);
                }
                catch (e) { }
                return json;
            }
            catch (e) {
                const msg = `Could not get organization metadata`;
                console.error(msg, e);
                throw e;
            }
        });
    }
    getOrganizationNames() {
        return __awaiter(this, void 0, void 0, function* () {
            try {
                const orgs = yield this.valist.methods.getOrgNames().call();
                return orgs;
            }
            catch (e) {
                const msg = `Could not get organization names`;
                console.error(msg, e);
                throw e;
            }
        });
    }
    setOrgMeta(orgName, orgMeta, account) {
        return __awaiter(this, void 0, void 0, function* () {
            try {
                const hash = yield this.addJSONtoIPFS(orgMeta);
                yield this.valist.methods.setOrgMeta(orgName, hash).send({ from: account });
            }
            catch (e) {
                const msg = `Could not set organization metadata`;
                console.error(msg, e);
                throw e;
            }
        });
    }
    // returns repository
    getRepository(orgName, repoName) {
        return __awaiter(this, void 0, void 0, function* () {
            try {
                const repo = yield this.valist.methods.getRepository(orgName, repoName).call();
                let json = {};
                try {
                    json = yield this.fetchJSONfromIPFS(repo[0]);
                }
                catch (e) { }
                return { meta: json, tags: repo[1] };
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
                const repos = yield this.valist.methods.getRepoNames(orgName).call();
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
                const repoMeta = yield this.valist.methods.getRepoMeta(orgName, repoName).call();
                let json = {};
                try {
                    json = yield this.fetchJSONfromIPFS(repoMeta);
                }
                catch (e) { }
                ;
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
                const hash = yield this.addJSONtoIPFS(repoMeta);
                yield this.valist.methods.setRepoMeta(orgName, repoName, hash).send({ from: account });
            }
            catch (e) {
                const msg = `Could not set repository metadata`;
                console.error(msg, e);
                throw e;
            }
        });
    }
    getLatestReleaseFromRepo(orgName, repoName) {
        return __awaiter(this, void 0, void 0, function* () {
            try {
                const release = yield this.valist.methods.getLatestRelease(orgName, repoName).call();
                return release;
            }
            catch (e) {
                const msg = `Could not get latest release from repo`;
                console.error(msg, e);
                throw e;
            }
        });
    }
    getLatestTagFromRepo(orgName, repoName) {
        return __awaiter(this, void 0, void 0, function* () {
            try {
                const tag = yield this.valist.methods.getLatestTag(orgName, repoName).call();
                return tag;
            }
            catch (e) {
                const msg = `Could not get latest tag from repo`;
                console.error(msg, e);
                throw e;
            }
        });
    }
    getReleaseTagsFromRepo(orgName, repoName) {
        return __awaiter(this, void 0, void 0, function* () {
            try {
                const tags = yield this.valist.methods.getReleaseTags(orgName, repoName).call();
                return tags;
            }
            catch (e) {
                const msg = `Could not get release tags from repo`;
                console.error(msg, e);
                throw e;
            }
        });
    }
    getReleasesFromRepo(orgName, repoName) {
        return __awaiter(this, void 0, void 0, function* () {
            try {
                const tags = yield this.valist.methods.getReleaseTags(orgName, repoName).call();
                const releases = [];
                for (let i = 0; i < tags.length; i++) {
                    const release = yield this.valist.methods.getRelease(orgName, repoName, tags[i]).call();
                    releases.push(Object.assign(Object.assign({}, release), { tag: tags[i] }));
                }
                ;
                return releases;
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
                const release = yield this.valist.methods.getRelease(orgName, repoName, tag).call();
                return release;
            }
            catch (e) {
                const msg = `Could not get release by tag`;
                console.error(msg, e);
                throw e;
            }
        });
    }
    isOrgOwner(orgName, account) {
        return __awaiter(this, void 0, void 0, function* () {
            try {
                return yield this.valist.methods.isOrgOwner(orgName, account).call();
            }
            catch (e) {
                const msg = `Could not check if user has ORG_ADMIN_ROLE`;
                console.error(msg, e);
                throw e;
            }
        });
    }
    isOrgAdmin(orgName, account) {
        return __awaiter(this, void 0, void 0, function* () {
            try {
                return yield this.valist.methods.isOrgAdmin(orgName, account).call();
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
                return yield this.valist.methods.isRepoAdmin(orgName, repoName, account).call();
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
                return yield this.valist.methods.isRepoDev(orgName, repoName, account).call();
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
                yield this.valist.methods.grantOrgAdmin(orgName, grantee).send({ from: granter });
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
                yield this.valist.methods.revokeOrgAdmin(orgName, revokee).send({ from: revoker });
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
                yield this.valist.methods.grantRepoAdmin(orgName, repoName, grantee).send({ from: granter });
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
                yield this.valist.methods.revokeRepoAdmin(orgName, repoName, revokee).send({ from: revoker });
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
                yield this.valist.methods.grantRepoDev(orgName, repoName, grantee).send({ from: granter });
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
                yield this.valist.methods.revokeRepoDev(orgName, repoName, revokee).send({ from: revoker });
            }
            catch (e) {
                const msg = `Could not revoke REPO_DEV_ROLE`;
                console.error(msg, e);
                throw e;
            }
        });
    }
    getOrgOwners(orgName) {
        return __awaiter(this, void 0, void 0, function* () {
            try {
                const members = yield this.valist.methods.getOrgOwners(orgName).call();
                return members;
            }
            catch (e) {
                const msg = `Could not get users that have ORG_OWNER_ROLE`;
                console.error(msg, e);
                throw e;
            }
        });
    }
    getOrgAdmins(orgName) {
        return __awaiter(this, void 0, void 0, function* () {
            try {
                const members = yield this.valist.methods.getOrgAdmins(orgName).call();
                return members;
            }
            catch (e) {
                const msg = `Could not get users that have ORG_ADMIN_ROLE`;
                console.error(msg, e);
                throw e;
            }
        });
    }
    getRepoAdmins(orgName, repoName) {
        return __awaiter(this, void 0, void 0, function* () {
            try {
                const members = yield this.valist.methods.getRepoAdmins(orgName, repoName).call();
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
                const members = yield this.valist.methods.getRepoDevs(orgName, repoName).call();
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
                const block = yield this.web3.eth.getBlock("latest");
                const result = yield this.valist.methods.createOrganization(orgName.toLowerCase().replace(exports.shortnameFilterRegex, ""), metaFile).send({ from: account, gasLimit: block.gasLimit });
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
                const metaFile = yield this.addJSONtoIPFS(repoMeta);
                const block = yield this.web3.eth.getBlock("latest");
                const result = yield this.valist.methods.createRepository(orgName, repoName.toLowerCase().replace(exports.shortnameFilterRegex, ""), metaFile).send({ from: account, gasLimit: block.gasLimit });
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
                const block = yield this.web3.eth.getBlock("latest");
                yield this.valist.methods.publishRelease(orgName, repoName, release.tag, release.hash, release.meta).send({ from: account, gasLimit: block.gasLimit });
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