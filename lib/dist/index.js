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
var __generator = (this && this.__generator) || function (thisArg, body) {
    var _ = { label: 0, sent: function() { if (t[0] & 1) throw t[1]; return t[1]; }, trys: [], ops: [] }, f, y, t, g;
    return g = { next: verb(0), "throw": verb(1), "return": verb(2) }, typeof Symbol === "function" && (g[Symbol.iterator] = function() { return this; }), g;
    function verb(n) { return function (v) { return step([n, v]); }; }
    function step(op) {
        if (f) throw new TypeError("Generator is already executing.");
        while (_) try {
            if (f = 1, y && (t = op[0] & 2 ? y["return"] : op[0] ? y["throw"] || ((t = y["return"]) && t.call(y), 0) : y.next) && !(t = t.call(y, op[1])).done) return t;
            if (y = 0, t) op = [op[0] & 2, t.value];
            switch (op[0]) {
                case 0: case 1: t = op; break;
                case 4: _.label++; return { value: op[1], done: false };
                case 5: _.label++; y = op[1]; op = [0]; continue;
                case 7: op = _.ops.pop(); _.trys.pop(); continue;
                default:
                    if (!(t = _.trys, t = t.length > 0 && t[t.length - 1]) && (op[0] === 6 || op[0] === 2)) { _ = 0; continue; }
                    if (op[0] === 3 && (!t || (op[1] > t[0] && op[1] < t[3]))) { _.label = op[1]; break; }
                    if (op[0] === 6 && _.label < t[1]) { _.label = t[1]; t = op; break; }
                    if (t && _.label < t[2]) { _.label = t[2]; _.ops.push(op); break; }
                    if (t[2]) _.ops.pop();
                    _.trys.pop(); continue;
            }
            op = body.call(thisArg, _);
        } catch (e) { op = [6, e]; y = 0; } finally { f = t = 0; }
        if (op[0] & 5) throw op[1]; return { value: op[0] ? op[1] : void 0, done: true };
    }
};
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.Web3Providers = void 0;
var web3_1 = __importDefault(require("web3"));
// @ts-ignore
var ipfs_http_client_1 = __importDefault(require("ipfs-http-client"));
var Valist_json_1 = __importDefault(require("./abis/Valist.json"));
var ValistOrganization_json_1 = __importDefault(require("./abis/ValistOrganization.json"));
var ValistRepository_json_1 = __importDefault(require("./abis/ValistRepository.json"));
var getContractInstance = function (web3, abi, address) { return __awaiter(void 0, void 0, void 0, function () {
    return __generator(this, function (_a) {
        // create the instance
        return [2 /*return*/, new web3.eth.Contract(abi, address)];
    });
}); };
var getValistContract = function (web3) { return __awaiter(void 0, void 0, void 0, function () {
    var networkId, deployedAddress;
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0: return [4 /*yield*/, web3.eth.net.getId()];
            case 1:
                networkId = _a.sent();
                deployedAddress = Valist_json_1.default.networks[networkId].address;
                return [4 /*yield*/, getContractInstance(web3, Valist_json_1.default.abi, deployedAddress)];
            case 2: return [2 /*return*/, _a.sent()];
        }
    });
}); };
var getValistOrganizationContract = function (web3, address) { return __awaiter(void 0, void 0, void 0, function () {
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0: return [4 /*yield*/, getContractInstance(web3, ValistOrganization_json_1.default.abi, address)];
            case 1: 
            // create the instance
            return [2 /*return*/, _a.sent()];
        }
    });
}); };
var getValistRepositoryContract = function (web3, address) { return __awaiter(void 0, void 0, void 0, function () {
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0: return [4 /*yield*/, getContractInstance(web3, ValistRepository_json_1.default.abi, address)];
            case 1: 
            // get network ID and the deployed address
            return [2 /*return*/, _a.sent()];
        }
    });
}); };
var Valist = /** @class */ (function () {
    function Valist(web3Provider, ipfsEnabled) {
        this.web3 = new web3_1.default(web3Provider);
        if (ipfsEnabled) {
            this.ipfs = ipfs_http_client_1.default({ host: 'ipfs.infura.io', port: '5001', apiPath: '/api/v0/', protocol: 'https' });
        }
    }
    // initialize main valist contract instance for future calls
    Valist.prototype.connect = function () {
        return __awaiter(this, void 0, void 0, function () {
            var _a;
            return __generator(this, function (_b) {
                switch (_b.label) {
                    case 0:
                        _a = this;
                        return [4 /*yield*/, getValistContract(this.web3)];
                    case 1:
                        _a.valist = _b.sent();
                        return [2 /*return*/];
                }
            });
        });
    };
    // returns organization contract instance
    Valist.prototype.getOrganization = function (orgName) {
        return __awaiter(this, void 0, void 0, function () {
            var orgAddress, org;
            return __generator(this, function (_a) {
                switch (_a.label) {
                    case 0: return [4 /*yield*/, this.valist.methods.orgs(orgName).call()];
                    case 1:
                        orgAddress = _a.sent();
                        return [4 /*yield*/, getValistOrganizationContract(this.web3, orgAddress)];
                    case 2:
                        org = _a.sent();
                        return [2 /*return*/, org];
                }
            });
        });
    };
    // returns organization contract instance
    Valist.prototype.getRepository = function (orgName, repoName) {
        return __awaiter(this, void 0, void 0, function () {
            var org, repoAddress, repo;
            return __generator(this, function (_a) {
                switch (_a.label) {
                    case 0: return [4 /*yield*/, this.getOrganization(orgName)];
                    case 1:
                        org = _a.sent();
                        return [4 /*yield*/, org.methods.repos(repoName).call()];
                    case 2:
                        repoAddress = _a.sent();
                        return [4 /*yield*/, getValistRepositoryContract(this.web3, repoAddress)];
                    case 3:
                        repo = _a.sent();
                        return [2 /*return*/, repo];
                }
            });
        });
    };
    Valist.prototype.getOrganizationMeta = function (orgName) {
        return __awaiter(this, void 0, void 0, function () {
            var org, orgMeta;
            return __generator(this, function (_a) {
                switch (_a.label) {
                    case 0: return [4 /*yield*/, this.getOrganization(orgName)];
                    case 1:
                        org = _a.sent();
                        return [4 /*yield*/, org.methods.orgMeta().call()];
                    case 2:
                        orgMeta = _a.sent();
                        return [2 /*return*/, orgMeta];
                }
            });
        });
    };
    // returns repository contract instance
    Valist.prototype.getRepoFromOrganization = function (orgName, repoName) {
        return __awaiter(this, void 0, void 0, function () {
            var org, repoAddress, repo;
            return __generator(this, function (_a) {
                switch (_a.label) {
                    case 0: return [4 /*yield*/, this.getOrganization(orgName)];
                    case 1:
                        org = _a.sent();
                        return [4 /*yield*/, org.methods.repos(repoName).call()];
                    case 2:
                        repoAddress = _a.sent();
                        return [4 /*yield*/, getValistRepositoryContract(this.web3, repoAddress)];
                    case 3:
                        repo = _a.sent();
                        return [2 /*return*/, repo];
                }
            });
        });
    };
    Valist.prototype.getRepoMeta = function (orgName, repoName) {
        return __awaiter(this, void 0, void 0, function () {
            var repo, repoMeta;
            return __generator(this, function (_a) {
                switch (_a.label) {
                    case 0: return [4 /*yield*/, this.getRepoFromOrganization(orgName, repoName)];
                    case 1:
                        repo = _a.sent();
                        return [4 /*yield*/, repo.methods.repoMeta().call()];
                    case 2:
                        repoMeta = _a.sent();
                        return [2 /*return*/, repoMeta];
                }
            });
        });
    };
    Valist.prototype.getLatestTagFromRepo = function (orgName, repoName) {
        return __awaiter(this, void 0, void 0, function () {
            var repo, tag;
            return __generator(this, function (_a) {
                switch (_a.label) {
                    case 0: return [4 /*yield*/, this.getRepoFromOrganization(orgName, repoName)];
                    case 1:
                        repo = _a.sent();
                        return [4 /*yield*/, repo.methods.tag().call()];
                    case 2:
                        tag = _a.sent();
                        return [2 /*return*/, tag];
                }
            });
        });
    };
    Valist.prototype.getLatestReleaseFromRepo = function (orgName, repoName) {
        return __awaiter(this, void 0, void 0, function () {
            var repo, release;
            return __generator(this, function (_a) {
                switch (_a.label) {
                    case 0: return [4 /*yield*/, this.getRepoFromOrganization(orgName, repoName)];
                    case 1:
                        repo = _a.sent();
                        return [4 /*yield*/, repo.methods.latestRelease().call()];
                    case 2:
                        release = _a.sent();
                        return [2 /*return*/, release];
                }
            });
        });
    };
    Valist.prototype.getReleasesFromRepo = function (orgName, repoName, tag) {
        return __awaiter(this, void 0, void 0, function () {
            var repo;
            return __generator(this, function (_a) {
                switch (_a.label) {
                    case 0: return [4 /*yield*/, this.getRepoFromOrganization(orgName, repoName)];
                    case 1:
                        repo = _a.sent();
                        return [2 /*return*/, repo.getPastEvents('Release', { fromBlock: 0, toBlock: 'latest' })];
                }
            });
        });
    };
    Valist.prototype.getReleaseByTag = function (orgName, repoName, tag) {
        return __awaiter(this, void 0, void 0, function () {
            var repo, topic, events, _a, release, releaseMeta;
            return __generator(this, function (_b) {
                switch (_b.label) {
                    case 0: return [4 /*yield*/, this.getRepoFromOrganization(orgName, repoName)];
                    case 1:
                        repo = _b.sent();
                        topic = this.web3.utils.sha3(tag);
                        return [4 /*yield*/, repo.getPastEvents('Release', { topics: [, topic], fromBlock: 0, toBlock: 'latest' })];
                    case 2:
                        events = _b.sent();
                        if (events[0]) {
                            _a = events[0].returnValues, release = _a.release, releaseMeta = _a.releaseMeta;
                            return [2 /*return*/, { tag: tag, release: release, releaseMeta: releaseMeta }];
                        }
                        else {
                            return [2 /*return*/];
                        }
                        return [2 /*return*/];
                }
            });
        });
    };
    Valist.prototype.createOrganization = function (orgName, orgMeta, account) {
        return __awaiter(this, void 0, void 0, function () {
            var metaFile, result;
            return __generator(this, function (_a) {
                switch (_a.label) {
                    case 0: return [4 /*yield*/, this.addJSONtoIPFS(orgMeta)];
                    case 1:
                        metaFile = _a.sent();
                        return [4 /*yield*/, this.valist.methods.createOrganization(orgName, metaFile).send({ from: account })];
                    case 2:
                        result = _a.sent();
                        return [2 /*return*/, result];
                }
            });
        });
    };
    Valist.prototype.createRepository = function (orgName, repoName, repoMeta, account) {
        return __awaiter(this, void 0, void 0, function () {
            var org, metaFile, result, err_1;
            return __generator(this, function (_a) {
                switch (_a.label) {
                    case 0:
                        _a.trys.push([0, 4, , 5]);
                        return [4 /*yield*/, this.getOrganization(orgName)];
                    case 1:
                        org = _a.sent();
                        return [4 /*yield*/, this.addJSONtoIPFS(repoMeta)];
                    case 2:
                        metaFile = _a.sent();
                        return [4 /*yield*/, org.methods.createRepository(repoName, metaFile).send({ from: account })];
                    case 3:
                        result = _a.sent();
                        return [2 /*return*/, result];
                    case 4:
                        err_1 = _a.sent();
                        console.log(err_1);
                        return [2 /*return*/, err_1];
                    case 5: return [2 /*return*/];
                }
            });
        });
    };
    Valist.prototype.publishRelease = function (orgName, repoName, release, account) {
        return __awaiter(this, void 0, void 0, function () {
            var repo, result;
            return __generator(this, function (_a) {
                switch (_a.label) {
                    case 0: return [4 /*yield*/, this.getRepository(orgName, repoName)];
                    case 1:
                        repo = _a.sent();
                        return [4 /*yield*/, repo.methods.publishRelease(release.tag, release.hash, release.meta).send({ from: account })];
                    case 2:
                        result = _a.sent();
                        return [2 /*return*/, result];
                }
            });
        });
    };
    Valist.prototype.addJSONtoIPFS = function (data) {
        return __awaiter(this, void 0, void 0, function () {
            var file, result, err_2;
            return __generator(this, function (_a) {
                switch (_a.label) {
                    case 0:
                        file = Buffer.from(JSON.stringify(data));
                        _a.label = 1;
                    case 1:
                        _a.trys.push([1, 3, , 4]);
                        return [4 /*yield*/, this.ipfs.add(file)];
                    case 2:
                        result = _a.sent();
                        return [2 /*return*/, result["path"]];
                    case 3:
                        err_2 = _a.sent();
                        console.error('Error', err_2);
                        return [3 /*break*/, 4];
                    case 4: return [2 /*return*/];
                }
            });
        });
    };
    return Valist;
}());
exports.Web3Providers = web3_1.default.providers;
exports.default = Valist;
//# sourceMappingURL=index.js.map