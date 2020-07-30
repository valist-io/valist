// SPDX-License-Identifier: MIT
pragma solidity >=0.4.21 <0.7.0;

import "@openzeppelin/contracts/access/AccessControl.sol";
import "./ValistRepository.sol";

contract ValistOrganization is AccessControl {
    bytes32 constant ORG_OWNER = keccak256("ORG_OWNER_ROLE");
    bytes32 constant ORG_ADMIN = keccak256("ORG_ADMIN_ROLE");

    string public meta; // org metadata (full name, image, description, etc)

    mapping(string => ValistRepository) public repos;

    constructor(address _owner, string memory _meta) public {
        _setupRole(ORG_OWNER, _owner);
        meta = _meta;
    }

    function getRepositoryAddress(string memory _repoName) public view returns(address) {
        return address(repos[_repoName]);
    }

    function createRepository(string memory _repoName, string memory _meta, address _owner) public returns(address) {
        require(address(repos[_repoName]) == address(0), "Repository already exists!");
        repos[_repoName] = new ValistRepository(_owner, _meta);
        return address(repos[_repoName]);
    }

    function updateOrgMeta(string memory _meta) public returns (string memory) {
        require(hasRole(ORG_OWNER, msg.sender) || hasRole(ORG_ADMIN, msg.sender), "You do not have permission to modify this organization!");
        meta = _meta;
    }

}
