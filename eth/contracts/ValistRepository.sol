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

    constructor(address _owner, string memory _meta) public {
        _setupRole(REPO_OWNER, _owner);
        meta = _meta;
    }

    function updateRepoMeta(string memory _meta) public returns (string memory) {
        require(hasRole(REPO_OWNER, msg.sender) || hasRole(REPO_ADMIN, msg.sender), "You do not have permission to modify this repository!");
        meta = _meta;
    }

}
