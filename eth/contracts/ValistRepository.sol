// SPDX-License-Identifier: MIT
pragma solidity >=0.4.21 <0.7.0;

import "@openzeppelin/contracts/access/AccessControl.sol";

contract ValistRepository is AccessControl {

    bytes32 constant REPO_ADMIN = keccak256("REPO_ADMIN_ROLE");
    bytes32 constant REPO_DEV = keccak256("REPO_DEV_ROLE");

    string public meta; // ipfs URI for metadata (image, description, etc)
    string public latestRelease; // latest release hash (should equal releases[-1])

    string[] public releases; // list of previous release ipfs hashes

    event Release(string release, string changelog);

    event MetaUpdate(string meta);

    modifier admin() {
      require(hasRole(REPO_ADMIN, msg.sender), "You do not have permission to perform this action!");
      _;
    }

    modifier developer() {
      require(hasRole(REPO_ADMIN, msg.sender) || hasRole(REPO_DEV, msg.sender), "You do not have permission to perform this action!");
      _;
    }

    constructor(address _owner, string memory _meta) public {
        _setupRole(REPO_ADMIN, _owner);
        _setRoleAdmin(REPO_ADMIN, REPO_ADMIN);
        _setRoleAdmin(REPO_DEV, REPO_ADMIN);

        meta = _meta;
    }

    function updateRepoMeta(string memory _meta) public admin returns (string memory) {
        meta = _meta;

        emit MetaUpdate(_meta);
    }

    function publishRelease(string memory _changelog, string memory _release) public developer {
        releases.push(_release);

        emit Release(_release, _changelog);
    }

}
