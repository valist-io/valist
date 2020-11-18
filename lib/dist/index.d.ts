import Web3 from 'web3';
import { provider } from 'web3-core/types';
import ipfsClient from 'ipfs-http-client';
declare class Valist {
    web3: Web3;
    valist: any;
    ipfs: ipfsClient;
    constructor(web3Provider: provider, ipfsEnabled?: boolean);
    connect(): Promise<void>;
    getOrganization(orgName: string): Promise<import("web3-eth-contract").Contract>;
    getOrganizationMeta(orgName: string): Promise<any>;
    getCreatedOrganizations(): Promise<any>;
    getDeletedOrganizations(): Promise<any>;
    getRepository(orgName: string, repoName: string): Promise<import("web3-eth-contract").Contract>;
    getReposFromOrganization(orgName: string): Promise<import("web3-eth-contract").EventData[]>;
    getRepoMeta(orgName: string, repoName: string): Promise<any>;
    getLatestTagFromRepo(orgName: string, repoName: string): Promise<any>;
    getLatestReleaseFromRepo(orgName: string, repoName: string): Promise<any>;
    getLatestReleaseMetaFromRepo(orgName: string, repoName: string): Promise<any>;
    getReleasesFromRepo(orgName: string, repoName: string): Promise<import("web3-eth-contract").EventData[]>;
    getReleaseByTag(orgName: string, repoName: string, tag: string): Promise<{
        tag: any;
        release: any;
        releaseMeta: any;
    } | undefined>;
    isOrgAdmin(orgName: string, account: any): Promise<any>;
    grantOrgAdmin(orgName: string, account: any): Promise<void>;
    revokeOrgAdmin(orgName: string, account: any): Promise<void>;
    isRepoAdmin(orgName: string, repoName: string, account: any): Promise<any>;
    isRepoDev(orgName: string, repoName: string, account: any): Promise<any>;
    getOrgAdmins(orgName: string, repoName: string): Promise<any[]>;
    getRepoAdmins(orgName: string, repoName: string): Promise<any[]>;
    getRepoDevs(orgName: string, repoName: string): Promise<any[]>;
    grantRepoAdmin(orgName: string, repoName: string, account: any): Promise<any>;
    revokeRepoAdmin(orgName: string, repoName: string, account: any): Promise<any>;
    grantRepoDev(orgName: string, repoName: string, account: any): Promise<any>;
    revokeRepoDev(orgName: string, repoName: string, account: any): Promise<any>;
    createOrganization(orgName: string, orgMeta: {
        name: string;
        description: string;
    }, account: any): Promise<any>;
    createRepository(orgName: string, repoName: string, repoMeta: {
        name: string;
        description: string;
        projectType: string;
        homepage: string;
        github: string;
    }, account: any): Promise<any>;
    publishRelease(orgName: string, repoName: string, release: {
        tag: string;
        hash: string;
        meta: string;
    }, account: any): Promise<any>;
    addJSONtoIPFS(data: any): Promise<any>;
    addFileToIPFS(data: any): Promise<any>;
    fetchJSONfromIPFS(ipfsHash: string): Promise<any>;
}
export declare const Web3Providers: import("web3-core").Providers;
export default Valist;
