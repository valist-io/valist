import Web3 from 'web3';
declare class Valist {
    web3: Web3;
    valist: any;
    constructor(provider: any);
    connect(): Promise<void>;
    getOrganization(orgName: string): Promise<import("web3-eth-contract").Contract>;
    getRepoFromOrganization(orgName: string, repoName: string): Promise<import("web3-eth-contract").Contract>;
    getLatestReleaseFromRepo(orgName: string, repoName: string): Promise<any>;
    createOrganization(orgName: string, account: any): Promise<any>;
}
export = Valist;
