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
const fetch = require('node-fetch');
if (!globalThis.fetch) {
    globalThis.fetch = fetch;
}
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
        if (ipfsEnabled) {
            this.ipfs = ipfs_http_client_1.default({ host: 'ipfs.infura.io', port: '5001', apiPath: '/api/v0/', protocol: 'https' });
        }
    }
    // initialize main valist contract instance for future calls
    connect() {
        return __awaiter(this, void 0, void 0, function* () {
            this.valist = yield getValistContract(this.web3);
        });
    }
    // returns organization contract instance
    getOrganization(orgName) {
        return __awaiter(this, void 0, void 0, function* () {
            const orgAddress = yield this.valist.methods.orgs(orgName).call();
            const org = yield getValistOrganizationContract(this.web3, orgAddress);
            return org;
        });
    }
    getOrganizationMeta(orgName) {
        return __awaiter(this, void 0, void 0, function* () {
            const org = yield this.getOrganization(orgName);
            const orgMeta = yield org.methods.orgMeta().call();
            return orgMeta;
        });
    }
    getCreatedOrganizations() {
        return __awaiter(this, void 0, void 0, function* () {
            const organizations = yield this.valist.getPastEvents('OrganizationCreated', { fromBlock: 0, toBlock: 'latest' });
            return organizations;
        });
    }
    getDeletedOrganizations() {
        return __awaiter(this, void 0, void 0, function* () {
            const organizations = yield this.valist.getPastEvents('OrganizationDeleted', { fromBlock: 0, toBlock: 'latest' });
            return organizations;
        });
    }
    // returns organization contract instance
    getRepository(orgName, repoName) {
        return __awaiter(this, void 0, void 0, function* () {
            const org = yield this.getOrganization(orgName);
            const repoAddress = yield org.methods.repos(repoName).call();
            const repo = yield getValistRepositoryContract(this.web3, repoAddress);
            return repo;
        });
    }
    // returns repository contract instance
    getRepoFromOrganization(orgName, repoName) {
        return __awaiter(this, void 0, void 0, function* () {
            const org = yield this.getOrganization(orgName);
            const repoAddress = yield org.methods.repos(repoName).call();
            const repo = yield getValistRepositoryContract(this.web3, repoAddress);
            return repo;
        });
    }
    getReposFromOrganization(orgName) {
        return __awaiter(this, void 0, void 0, function* () {
            const org = yield this.getOrganization(orgName);
            const repos = yield org.getPastEvents('RepositoryCreated', { fromBlock: 0, toBlock: 'latest' });
            return repos;
        });
    }
    getRepoMeta(orgName, repoName) {
        return __awaiter(this, void 0, void 0, function* () {
            const repo = yield this.getRepoFromOrganization(orgName, repoName);
            const repoMeta = yield repo.methods.repoMeta().call();
            return repoMeta;
        });
    }
    getLatestTagFromRepo(orgName, repoName) {
        return __awaiter(this, void 0, void 0, function* () {
            const repo = yield this.getRepoFromOrganization(orgName, repoName);
            const tag = yield repo.methods.tag().call();
            return tag;
        });
    }
    getLatestReleaseFromRepo(orgName, repoName) {
        return __awaiter(this, void 0, void 0, function* () {
            const repo = yield this.getRepoFromOrganization(orgName, repoName);
            const release = yield repo.methods.latestRelease().call();
            return release;
        });
    }
    getLatestReleaseMetaFromRepo(orgName, repoName) {
        return __awaiter(this, void 0, void 0, function* () {
            const repo = yield this.getRepoFromOrganization(orgName, repoName);
            const release = yield repo.methods.releaseMeta().call();
            return release;
        });
    }
    getReleasesFromRepo(orgName, repoName) {
        return __awaiter(this, void 0, void 0, function* () {
            const repo = yield this.getRepoFromOrganization(orgName, repoName);
            return yield repo.getPastEvents('Release', { fromBlock: 0, toBlock: 'latest' });
        });
    }
    getReleaseByTag(orgName, repoName, tag) {
        return __awaiter(this, void 0, void 0, function* () {
            const repo = yield this.getRepoFromOrganization(orgName, repoName);
            const events = yield repo.getPastEvents('Release', { fromBlock: 0, toBlock: 'latest' });
            // @TODO make this more efficient later
            for (let i = 0; i < events.length; i++) {
                if (events[i].returnValues.tag == tag) {
                    const { tag, release, releaseMeta } = events[i].returnValues;
                    return { tag, release, releaseMeta };
                }
            }
            return;
        });
    }
    createOrganization(orgName, orgMeta, account) {
        return __awaiter(this, void 0, void 0, function* () {
            const metaFile = yield this.addJSONtoIPFS(orgMeta);
            const result = yield this.valist.methods.createOrganization(orgName, metaFile).send({ from: account });
            return result;
        });
    }
    createRepository(orgName, repoName, repoMeta, account) {
        return __awaiter(this, void 0, void 0, function* () {
            try {
                const org = yield this.getOrganization(orgName);
                const metaFile = yield this.addJSONtoIPFS(repoMeta);
                const result = yield org.methods.createRepository(repoName, metaFile).send({ from: account });
                return result;
            }
            catch (err) {
                console.log(err);
                return err;
            }
        });
    }
    publishRelease(orgName, repoName, release, account) {
        return __awaiter(this, void 0, void 0, function* () {
            const repo = yield this.getRepository(orgName, repoName);
            const result = yield repo.methods.publishRelease(release.tag, release.hash, release.meta).send({ from: account });
            return result;
        });
    }
    addJSONtoIPFS(data) {
        return __awaiter(this, void 0, void 0, function* () {
            const file = Buffer.from(JSON.stringify(data));
            try {
                const result = yield this.ipfs.add(file);
                return result["path"];
            }
            catch (err) {
                console.error('Error', err);
            }
        });
    }
    addFileToIPFS(data) {
        return __awaiter(this, void 0, void 0, function* () {
            console.log(data);
            const file = Buffer.from(data);
            try {
                const result = yield this.ipfs.add(file);
                return result["path"];
            }
            catch (err) {
                console.error('Error', err);
            }
        });
    }
    fetchJSONfromIPFS(ipfsHash) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield fetch(`https://cloudflare-ipfs.com/ipfs/${ipfsHash}`);
            const json = yield response.json();
            console.log("JSON Fetched from IPFS", json);
            return json;
        });
    }
}
exports.Web3Providers = web3_1.default.providers;
exports.default = Valist;
//# sourceMappingURL=index.js.map