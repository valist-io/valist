// SPDX-License-Identifier: MIT
pragma solidity >=0.4.21 <0.7.0;

import "@openzeppelin/contracts/access/AccessControl.sol";

contract ValistRepository is AccessControl {
    bytes32 constant REPO_OWNER = keccak256("REPO_OWNER_ROLE");
    bytes32 constant REPO_ADMIN = keccak256("REPO_ADMIN_ROLE");
    bytes32 constant REPO_DEV = keccak256("REPO_DEV_ROLE");

    string public meta; // ipfs URI for metadata (image, description, etc)
    string public latestRelease; // latest release hash (should equal releases[-1])

    string[] public changelog; // list of IPFS uris for any changelogs (also emitted as an event during update)
    string[] public releases; // list of previous release ipfs hashes

    event Release(string changelog, string release);

    modifier admin() {
      require(hasRole(REPO_OWNER, msg.sender) || hasRole(REPO_ADMIN, msg.sender), "You do not have permission to modify this repository!");
      _;
    }

    modifier developers() {
      require(hasRole(REPO_OWNER, msg.sender) ||
              hasRole(REPO_ADMIN, msg.sender) ||
              hasRole(REPO_DEV, msg.sender),
              "You do not have permission to modify this repository!");
      _;
    }

    constructor(address _owner, string memory _meta) public {
        _setupRole(REPO_OWNER, _owner);
        meta = _meta;
    }

    function updateRepoMeta(string memory _meta) public admin returns (string memory) {
        meta = _meta;
    }

    function publishUpdate(string memory _changelog, string memory _release) public admin developers {
        changelog.push(_changelog);
        releases.push(_release);

        emit Release(_changelog, _release);
    }

}
