import Web3 from 'web3';
import { provider } from 'web3-core/types';
import ipfsClient from 'ipfs-http-client';
export declare type ProjectType = "binary" | "npm" | "pip" | "docker";
declare class Valist {
    web3: Web3;
    valist: any;
    ipfs: ipfsClient;
    defaultAccount: string;
    constructor(web3Provider: provider, ipfsEnabled?: boolean);
    connect(): Promise<void>;
    getOrganization(orgName: string): Promise<import("web3-eth-contract").Contract>;
    getOrganizationMeta(orgName: string): Promise<any>;
    getCreatedOrganizations(): Promise<any>;
    getDeletedOrganizations(): Promise<any>;
    setOrgMeta(orgName: string, orgMeta: any, account: string): Promise<void>;
    getRepository(orgName: string, repoName: string): Promise<import("web3-eth-contract").Contract>;
    getReposFromOrganization(orgName: string): Promise<import("web3-eth-contract").EventData[]>;
    getRepoMeta(orgName: string, repoName: string): Promise<any>;
    setRepoMeta(orgName: string, repoName: string, repoMeta: any, account: string): Promise<void>;
    getLatestTagFromRepo(orgName: string, repoName: string): Promise<any>;
    getLatestReleaseFromRepo(orgName: string, repoName: string): Promise<any>;
    getLatestReleaseMetaFromRepo(orgName: string, repoName: string): Promise<any>;
    getReleasesFromRepo(orgName: string, repoName: string): Promise<import("web3-eth-contract").EventData[]>;
    getReleaseByTag(orgName: string, repoName: string, tag: string): Promise<{
        tag: any;
        release: any;
        releaseMeta: any;
    } | undefined>;
    isOrgAdmin(orgName: string, account: string): Promise<any>;
    isRepoAdmin(orgName: string, repoName: string, account: string): Promise<any>;
    isRepoDev(orgName: string, repoName: string, account: string): Promise<any>;
    grantOrgAdmin(orgName: string, granter: string, grantee: string): Promise<void>;
    revokeOrgAdmin(orgName: string, revoker: string, revokee: string): Promise<void>;
    grantRepoAdmin(orgName: string, repoName: string, granter: string, grantee: string): Promise<any>;
    revokeRepoAdmin(orgName: string, repoName: string, revoker: string, revokee: string): Promise<any>;
    grantRepoDev(orgName: string, repoName: string, granter: string, grantee: string): Promise<any>;
    revokeRepoDev(orgName: string, repoName: string, revoker: string, revokee: string): Promise<any>;
    getOrgAdmins(orgName: string): Promise<any[]>;
    getRepoAdmins(orgName: string, repoName: string): Promise<any[]>;
    getRepoDevs(orgName: string, repoName: string): Promise<any[]>;
    createOrganization(orgName: string, orgMeta: {
        name: string;
        description: string;
    }, account: string): Promise<any>;
    createRepository(orgName: string, repoName: string, repoMeta: {
        name: string;
        description: string;
        projectType: ProjectType;
        homepage: string;
        repository: string;
    }, account: string): Promise<any>;
    publishRelease(orgName: string, repoName: string, release: {
        tag: string;
        hash: string;
        meta: string;
    }, account: string): Promise<any>;
    addJSONtoIPFS(data: any): Promise<any>;
    addFileToIPFS(data: any): Promise<any>;
    fetchJSONfromIPFS(ipfsHash: string): Promise<any>;
}
export declare const Web3Providers: import("web3-core").Providers;
export default Valist;
