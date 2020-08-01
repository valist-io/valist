// SPDX-License-Identifier: MIT
pragma solidity >=0.4.21 <0.7.0;

import "@openzeppelin/contracts/access/AccessControl.sol";
import "./ValistRepository.sol";

contract ValistOrganization is AccessControl {

    bytes32 constant ORG_ADMIN = keccak256("ORG_ADMIN_ROLE");
    bytes32 constant REPO_ADMIN = keccak256("REPO_ADMIN_ROLE");

    address immutable deployer;

    string public orgMeta; // org metadata (full name, image, description, etc)

    event MetaUpdated(string orgMeta);

    event RepositoryCreated(string repoName, string repoMeta);

    event RepositoryDeleted(string repoName);

    mapping(string => ValistRepository) public repos;

    modifier admin() {
        require(hasRole(ORG_ADMIN, msg.sender), "You do not have permission to perform this action!");
        _;
    }

    constructor(address _admin, string memory _orgMeta) public {
        deployer = msg.sender;

        _setupRole(ORG_ADMIN, _admin);
        _setRoleAdmin(ORG_ADMIN, ORG_ADMIN);

        orgMeta = _orgMeta;
    }

    function createRepository(string memory _repoName, string memory _repoMeta) public admin returns(address) {
        require(address(repos[_repoName]) == address(0), "Repository already exists!");

        repos[_repoName] = new ValistRepository(msg.sender, _repoMeta);

        emit RepositoryCreated(_repoName, _repoMeta);

        return address(repos[_repoName]);
    }

    function updateOrgMeta(string memory _orgMeta) public admin returns (string memory) {
        orgMeta = _orgMeta;

        emit MetaUpdated(orgMeta);
    }

    function deleteRepository(string memory _repoName) public {
        require(repos[_repoName].hasRole(REPO_ADMIN, msg.sender), "You do not have permission to perform this action!");

        repos[_repoName]._deleteRepository(msg.sender);

        delete repos[_repoName];

        emit RepositoryDeleted(_repoName);
    }

    function _deleteOrganization(address payable _admin) external {
        require(msg.sender == deployer, "Can only be called from parent contract!");

        selfdestruct(_admin);
    }

}
