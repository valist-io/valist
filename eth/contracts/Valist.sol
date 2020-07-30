// SPDX-License-Identifier: MIT
pragma solidity >=0.4.21 <0.7.0;

import "./ValistOrganization.sol";

contract Valist {

    // map an organization/username handle to an Organization contract
    // [valist.io]/[organization]/[repository]
    mapping(string => ValistOrganization) public orgs;

    event Update(string orgName, string repoName, string meta, string changelog, string release);

    function getOrganizationMeta(string memory orgName) public view returns(string memory meta) {
        return orgs[orgName].meta();
    }

    function getRepoMetaByOrganization(string memory orgName, string memory repoName) public view returns(string memory meta) {
        return orgs[orgName].repos(repoName).meta();
    }

    function getRepoLatestReleaseByOrganization(string memory orgName, string memory repoName) public view returns(string memory meta) {
        return orgs[orgName].repos(repoName).latestRelease();
    }

    // register organization/username to the global valist namespace
    function createOrganization(string memory orgName, string memory meta) public {
        require(address(orgs[orgName]) == address(0), "Organization already exists!");
        orgs[orgName] = new ValistOrganization(msg.sender, meta);
    }

}
