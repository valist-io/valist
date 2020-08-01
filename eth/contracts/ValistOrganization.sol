// SPDX-License-Identifier: MIT
pragma solidity >=0.4.21 <0.7.0;

import "@openzeppelin/contracts/access/AccessControl.sol";
import "./ValistRepository.sol";

contract ValistOrganization is AccessControl {

    bytes32 constant ORG_ADMIN = keccak256("ORG_ADMIN_ROLE");
    bytes32 constant REPO_ADMIN = keccak256("REPO_ADMIN_ROLE");

    address immutable deployer;

    string public orgMeta; // org metadata (full name, image, description, etc)

    uint16 internal repoCount;

    event MetaUpdated(string orgMeta);

    event RepositoryCreated(string repoName, string repoMeta);

    event RepositoryDeleted(string repoName);

    mapping(string => ValistRepository) public repos;

    modifier admin() {
        require(hasRole(ORG_ADMIN, msg.sender), "Access Denied");
        _;
    }

    constructor(address _admin, string memory _orgMeta) public {
        deployer = msg.sender;

        _setupRole(ORG_ADMIN, _admin);
        _setRoleAdmin(ORG_ADMIN, ORG_ADMIN);

        orgMeta = _orgMeta;
    }

    function createRepository(string memory _repoName, string memory _repoMeta) public admin returns(address) {
        require(address(repos[_repoName]) == address(0), "Repo exists");

        repos[_repoName] = new ValistRepository(msg.sender, _repoMeta);

        repoCount++;

        emit RepositoryCreated(_repoName, _repoMeta);

        return address(repos[_repoName]);
    }

    function updateOrgMeta(string memory _orgMeta) public admin returns (string memory) {
        orgMeta = _orgMeta;

        emit MetaUpdated(orgMeta);
    }

    function deleteRepository(string memory _repoName) public {
        require(repos[_repoName].hasRole(REPO_ADMIN, msg.sender), "Access Denied");

        repos[_repoName]._deleteRepository(msg.sender); // will fail if does not exist, no need for safemath on repoCount

        delete repos[_repoName];

        repoCount--;

        emit RepositoryDeleted(_repoName);
    }

    function _deleteOrganization(address payable _admin) external {
        require(msg.sender == deployer, "Can only be called from parent contract");
        require(repoCount == 0, "You must delete all repos before deleting the organization");

        selfdestruct(_admin);
    }

}
