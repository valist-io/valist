// SPDX-License-Identifier: MIT
pragma solidity >=0.4.21 <0.7.0;

import "@openzeppelin/contracts/access/AccessControl.sol";
import "./ValistRepository.sol";

contract ValistOrganization is AccessControl {

    bytes32 constant ORG_ADMIN = keccak256("ORG_ADMIN_ROLE");

    string public meta; // org metadata (full name, image, description, etc)

    event MetaUpdate(string meta);

    mapping(string => ValistRepository) public repos;

    modifier admin() {
        require(hasRole(ORG_ADMIN, msg.sender), "You do not have permission to perform this action!");
        _;
    }

    constructor(address _owner, string memory _meta) public {
        _setupRole(ORG_ADMIN, _owner);
        _setRoleAdmin(ORG_ADMIN, ORG_ADMIN);

        meta = _meta;
    }

    function createRepository(address _owner, string memory _repoName, string memory _meta) public admin returns(address) {
        require(address(repos[_repoName]) == address(0), "Repository already exists!");

        repos[_repoName] = new ValistRepository(_owner, _meta);

        return address(repos[_repoName]);
    }

    function updateOrgMeta(string memory _meta) public admin returns (string memory) {
        meta = _meta;

        emit MetaUpdate(meta);
    }

}
