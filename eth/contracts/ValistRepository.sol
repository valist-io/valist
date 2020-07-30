// SPDX-License-Identifier: MIT
pragma solidity >=0.4.21 <0.7.0;

import "@openzeppelin/contracts/access/AccessControl.sol";

contract ValistRepository is AccessControl {
  bytes32 constant REPO_OWNER = keccak256("REPO_OWNER_ROLE");
  bytes32 constant REPO_ADMIN = keccak256("REPO_ADMIN_ROLE");
  bytes32 constant REPO_DEV = keccak256("REPO_DEV_ROLE");

  struct Repository {
        string[] changelog; // list of IPFS uris for any changelogs (also emitted as an event during update)
        string[] releases; // list of previous release ipfs hashes
        string meta; // ipfs URI for metadata (image, description, etc)
        string latest; // latest release hash (should equal releases[-1])
        bool active;
  }

  Repository public repo;

  constructor(string memory meta, address owner) public {
      _setupRole(REPO_OWNER, owner);
      repo.meta = meta;
      repo.active = true;
  }

  function isActive() public view returns(bool) {
    return repo.active;
  }

  function getMeta() public view returns(string memory) {
    return repo.meta;
  }

  function getLatestRelease() public view returns(string memory) {
    return repo.latest;
  }

}
