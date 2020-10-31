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
    getCreatedRepositories(): Promise<any>;
    getDeletedOrganizations(): Promise<any>;
    getRepository(orgName: string, repoName: string): Promise<import("web3-eth-contract").Contract>;
    getRepoFromOrganization(orgName: string, repoName: string): Promise<import("web3-eth-contract").Contract>;
    getReposFromOrganization(orgName: string): Promise<import("web3-eth-contract").EventData[]>;
    getRepoMeta(orgName: string, repoName: string): Promise<any>;
    getLatestTagFromRepo(orgName: string, repoName: string): Promise<any>;
    getLatestReleaseFromRepo(orgName: string, repoName: string): Promise<any>;
    getReleasesFromRepo(orgName: string, repoName: string): Promise<import("web3-eth-contract").EventData[]>;
    getReleaseByTag(orgName: string, repoName: string, tag: string): Promise<{
        tag: any;
        release: any;
        releaseMeta: any;
    } | undefined>;
    createOrganization(orgName: string, orgMeta: {
        name: string;
        description: string;
    }, account: any): Promise<any>;
    createRepository(orgName: string, repoName: string, repoMeta: {
        name: string;
        description: string;
    }, account: any): Promise<any>;
    publishRelease(orgName: string, repoName: string, release: {
        tag: string;
        hash: string;
        meta: string;
    }, account: any): Promise<any>;
    addJSONtoIPFS(data: any): Promise<any>;
    addFileToIPFS(data: any): Promise<any>;
}
export declare const Web3Providers: import("web3-core").Providers;
export default Valist;
