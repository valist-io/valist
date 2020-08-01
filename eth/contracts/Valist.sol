// SPDX-License-Identifier: MIT
pragma solidity >=0.4.21 <0.7.0;

import "./ValistOrganization.sol";

contract Valist {

    bytes32 constant ORG_ADMIN = keccak256("ORG_ADMIN_ROLE");
    bytes32 constant REPO_ADMIN = keccak256("REPO_ADMIN_ROLE");

    // map an organization/username handle to an Organization contract
    // [valist.io]/[organization]/[repository]
    mapping(string => ValistOrganization) public orgs;

    event OrganizationCreated(string orgName, string orgMeta);

    event OrganizationDeleted(string orgName);

    function getOrganizationMeta(string memory orgName) public view returns(string memory meta) {
        return orgs[orgName].orgMeta();
    }

    function getRepoMetaByOrganization(string memory orgName, string memory repoName) public view returns(string memory meta) {
        return orgs[orgName].repos(repoName).repoMeta();
    }

    function getRepoLatestReleaseByOrganization(string memory orgName, string memory repoName) public view returns(string memory meta) {
        return orgs[orgName].repos(repoName).latestRelease();
    }

    // register organization/username to the global valist namespace
    function createOrganization(string memory orgName, string memory orgMeta) public returns(address) {
        require(address(orgs[orgName]) == address(0), "Organization already exists!");

        orgs[orgName] = new ValistOrganization(msg.sender, orgMeta);

        emit OrganizationCreated(orgName, orgMeta);

        return address(orgs[orgName]);
    }

    function deleteOrganization(string memory orgName) public {
        require(orgs[orgName].hasRole(ORG_ADMIN, msg.sender), "You do not have permission to perform this action!");

        orgs[orgName]._deleteOrganization(msg.sender);

        delete orgs[orgName];

        emit OrganizationDeleted(orgName);
    }

}
