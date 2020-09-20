// SPDX-License-Identifier: MIT
pragma solidity >=0.4.21 <0.7.0;

import "@openzeppelin/contracts/access/AccessControl.sol";

contract ValistRepository is AccessControl {

    bytes32 constant REPO_ADMIN = keccak256("REPO_ADMIN_ROLE");
    bytes32 constant REPO_DEV = keccak256("REPO_DEV_ROLE");

    address immutable deployer;

    string public repoMeta; // ipfs URI for metadata (image, description, etc)\
    string public releaseMeta; // version/build number, changelog, other release specific metadata (valist json schema)
    string public latestRelease; // latest release hash
    string public tag;

    event Release(string indexed tag, string release, string releaseMeta);

    event MetaUpdated(string repoMeta);

    modifier admin() {
        require(hasRole(REPO_ADMIN, msg.sender), "Access Denied");
        _;
    }

    modifier developer() {
        require(hasRole(REPO_ADMIN, msg.sender) || hasRole(REPO_DEV, msg.sender), "Access Denied");
        _;
    }

    constructor(address _admin, string memory _repoMeta) public {
        deployer = msg.sender;

        _setupRole(REPO_ADMIN, _admin);
        _setRoleAdmin(REPO_ADMIN, REPO_ADMIN);
        _setRoleAdmin(REPO_DEV, REPO_ADMIN);

        repoMeta = _repoMeta;
    }

    function updateRepoMeta(string memory _repoMeta) public admin returns (string memory) {
        repoMeta = _repoMeta;

        emit MetaUpdated(repoMeta);
    }

    function publishRelease(string memory _tag, string memory _latestRelease, string memory _releaseMeta) public developer {
        tag = _tag;
        latestRelease = _latestRelease;
        releaseMeta = _releaseMeta;

        emit Release(tag, latestRelease, releaseMeta);
    }

    function _deleteRepository(address payable _admin) external {
        require(msg.sender == deployer, "Can only be called from parent contract");

        selfdestruct(_admin);
    }

}
