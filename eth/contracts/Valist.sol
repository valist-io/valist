// SPDX-License-Identifier: MIT
pragma solidity >=0.4.21 <0.7.0;

import "./ValistOrganization.sol";

contract Valist {

    // map an organization/username handle to an Organization contract
    // [valist.io]/[organization]/[repository]
    mapping(string => ValistOrganization) public valist;

    event Update(string orgName, string repoName, string meta, string changelog, string release);

    function getRepositoryMetaByOrganization(string memory orgName, string memory repoName) public view returns(string memory meta) {
        return valist[orgName].getRepository(repoName).getMeta();
    }

    function getRepositoryLatestReleaseByOrganization(string memory orgName, string memory repoName) public view returns(string memory meta) {
        return valist[orgName].getRepository(repoName).getLatestRelease();
    }

    // register organization/username to the global valist namespace
    function createOrganization(string memory orgName, string memory meta) public {
        require(valist[orgName].isActive() == false, "Organization already exists!");
        valist[orgName] = new ValistOrganization(meta, msg.sender);
    }

}
