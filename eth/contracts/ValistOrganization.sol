// SPDX-License-Identifier: MIT
pragma solidity >=0.7.0;
pragma experimental ABIEncoderV2;

import "@openzeppelin/contracts/access/AccessControl.sol";
import "./ValistRepository.sol";

contract ValistOrganization is AccessControl {

    address immutable deployer;

    bytes32 constant ORG_ADMIN = keccak256("ORG_ADMIN_ROLE");
    bytes32 constant REPO_ADMIN = keccak256("REPO_ADMIN_ROLE");

    string public orgMeta; // org metadata (full name, image, description, etc)

    struct Repository {
        uint index;
        ValistRepository repo;
    }

    mapping(string => Repository) public repos;

    string[] public repoNames;

    modifier admin() {
        require(hasRole(ORG_ADMIN, msg.sender), "Access Denied");
        _;
    }

    constructor(address _admin, string memory _orgMeta) {
        deployer = msg.sender;

        _setupRole(ORG_ADMIN, _admin);
        _setRoleAdmin(ORG_ADMIN, ORG_ADMIN);

        orgMeta = _orgMeta;
    }

    function getRepoNames() public view returns(string[] memory) {
        return repoNames;
    }

    function createRepository(string memory _repoName, string memory _repoMeta) public admin returns(address) {
        require(address(repos[_repoName].repo) == address(0), "Repo exists");

        repoNames.push(_repoName);

        repos[_repoName].index = repoNames.length - 1;

        repos[_repoName].repo = new ValistRepository(msg.sender, _repoMeta);

        return address(repos[_repoName].repo);
    }

    function updateOrgMeta(string memory _orgMeta) public admin {
        orgMeta = _orgMeta;
    }

    function deleteRepository(string memory _repoName) public {
        require(repos[_repoName].repo.hasRole(REPO_ADMIN, msg.sender), "Access Denied");

        ValistRepository repo = repos[_repoName].repo;

        delete repoNames[repos[_repoName].index];
        delete repos[_repoName];

        repo._deleteRepository(msg.sender);
    }

    function _deleteOrganization(address payable _admin) external {
        require(msg.sender == deployer, "Can only be called from parent contract");
        require(repoNames.length == 0, "You must delete all repos before deleting the organization");

        selfdestruct(_admin);
    }

}
