// SPDX-License-Identifier: MIT
pragma solidity >=0.4.21 <0.7.0;

import "@openzeppelin/contracts/access/AccessControl.sol";

contract Valist is AccessControl {

    bytes32 constant REPO_OWNER = keccak256("REPO_OWNER_ROLE");
    bytes32 constant ORG_OWNER = keccak256("ORG_OWNER_ROLE");

    bytes32 constant REPO_ADMIN = keccak256("REPO_ADMIN_ROLE");
    bytes32 constant ORG_ADMIN = keccak256("ORG_ADMIN_ROLE");

    bytes32 constant APP_DEV = keccak256("APP_DEV_ROLE");

    struct Repository {
        mapping(bytes32 => address[]) usersByRole;
        string[] changelog; // list of IPFS uris for any changelogs (also emitted as an event during update)
        string[] releases; // list of previous release ipfs hashes
        string meta; // ipfs URI for metadata (image, description, etc)
        string latest; // latest release hash (should be the latest push to releases array)
        bool active;
    }

    struct Organization {
        mapping(bytes32 => address[]) usersByRole;
        mapping(string => Repository) repos;
        string meta; // org metadata (full name, image, description, etc)
        bool active;
    }

    event Update(string orgName, string repoName, string meta, string changelog, string release);

    // [valist.io]/[organization]/[repository]

    // map an organization handle to an Organization struct
    mapping(string => Organization) public orgs;

    function getRepositoryMetaByOrganization(string memory orgName, string memory repoName) public view returns(string memory meta) {
        return orgs[orgName].repos[repoName].meta;
    }

    function getRepositoryLatestReleaseByOrganization(string memory orgName, string memory repoName) public view returns(string memory meta) {
        return orgs[orgName].repos[repoName].latest;
    }

    function createOrganization(string memory orgName, string memory meta) public {
        require(orgs[orgName].active == false, "Organization already exists!");
        orgs[orgName].meta = meta;
        orgs[orgName].usersByRole[ORG_OWNER].push(msg.sender); // add user to org owner role
    }

    function createRepository(string memory orgName, string memory repoName, string memory meta) public {
        require(orgs[orgName].repos[repoName].active == false, "Repository already exists!");
        // @TODO require that user is of admin or owner role
        // (and likely change the way we store the roles, since OZ uses roles globally, when we need them per-org and per-repo)
        orgs[orgName].repos[repoName].meta = meta;
        orgs[orgName].repos[repoName].usersByRole[REPO_OWNER].push(msg.sender);
    }

}
