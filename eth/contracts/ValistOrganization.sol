// SPDX-License-Identifier: MIT
pragma solidity >=0.4.21 <0.7.0;

import "@openzeppelin/contracts/access/AccessControl.sol";
import "./ValistRepository.sol";

contract ValistOrganization is AccessControl {
    bytes32 constant ORG_OWNER = keccak256("ORG_OWNER_ROLE");
    bytes32 constant ORG_ADMIN = keccak256("ORG_ADMIN_ROLE");

    struct Organization {
        mapping(string => ValistRepository) repos;
        string meta; // org metadata (full name, image, description, etc)
        bool active;
    }

    Organization public org;

    constructor(string memory meta, address owner) public {
        _setupRole(ORG_OWNER, owner);
        org.meta = meta;
        org.active = true;
    }

    function isActive() public view returns(bool) {
        return org.active;
    }

    function createRepository(string memory repoName, string memory meta, address owner) public returns(address) {
        require(org.repos[repoName].isActive() == false, "Repository already exists!");
        org.repos[repoName] = new ValistRepository(meta, owner);
        return address(org.repos[repoName]);
    }

    function getRepositoryAddress(string memory repoName) public view returns(address) {
        return address(org.repos[repoName]);
    }

    function getRepository(string memory repoName) public view returns(ValistRepository) {
        return org.repos[repoName];
    }

    function updateMetadata(string memory meta) public {
        require(hasRole(ORG_OWNER, msg.sender) || hasRole(ORG_ADMIN, msg.sender), "You do not have permission to modify this organization!");
        org.meta = meta;
    }

}
