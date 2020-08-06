import Web3 from 'web3';
import { provider } from 'web3-core/types';
import ipfsClient from 'ipfs-http-client';
declare class Valist {
    web3: Web3;
    valist: any;
    ipfs: ipfsClient;
    constructor(web3Provider: provider, ipfsEnabled: boolean);
    connect(): Promise<void>;
    getOrganization(orgName: string): Promise<import("web3-eth-contract").Contract>;
    getRepository(orgName: string, repoName: string): Promise<import("web3-eth-contract").Contract>;
    getOrganizationMeta(orgName: string): Promise<any>;
    getRepoFromOrganization(orgName: string, repoName: string): Promise<import("web3-eth-contract").Contract>;
    getRepoMeta(orgName: string, repoName: string): Promise<any>;
    getLatestReleaseFromRepo(orgName: string, repoName: string): Promise<any>;
    createOrganization(orgName: string, orgMeta: string, account: any): Promise<any>;
    createRepository(orgName: string, repoName: string, repoMeta: string, account: any): Promise<any>;
    publishRelease(orgName: string, repoName: string, release: {
        tag: string;
        hash: string;
        meta: string;
    }, account: any): Promise<any>;
    addFileIpfs(data: any): Promise<any>;
}
export declare const Web3Providers: import("web3-core").Providers;
export default Valist;
