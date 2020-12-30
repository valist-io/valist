// SPDX-License-Identifier: MIT
pragma solidity >=0.7.0;
pragma experimental ABIEncoderV2;

import "@openzeppelin/contracts/access/AccessControl.sol";

contract ValistRepository is AccessControl {

    address immutable deployer;

    bytes32 constant REPO_ADMIN = keccak256("REPO_ADMIN_ROLE");
    bytes32 constant REPO_DEV = keccak256("REPO_DEV_ROLE");

    string public repoMeta; // ipfs URI for metadata (image, description, etc)

    struct Release {
        string tag;
        string releaseCID;
        string metaCID;
    }

    mapping(string => Release) public releases;

    string[] public tags;

    Release public latestRelease; // latest release

    event NewRelease(string tag);

    modifier admin() {
        require(hasRole(REPO_ADMIN, msg.sender), "Access Denied");
        _;
    }

    modifier developer() {
        require(hasRole(REPO_ADMIN, msg.sender) || hasRole(REPO_DEV, msg.sender), "Access Denied");
        _;
    }

    constructor(address _admin, string memory _repoMeta) {
        deployer = msg.sender;

        _setupRole(REPO_ADMIN, _admin);
        _setRoleAdmin(REPO_ADMIN, REPO_ADMIN);
        _setRoleAdmin(REPO_DEV, REPO_ADMIN);

        repoMeta = _repoMeta;
    }

    function getTags() public view returns(string[] memory) {
        return tags;
    }

    function updateRepoMeta(string memory _repoMeta) public admin {
        repoMeta = _repoMeta;
    }

    function publishRelease(string memory _tag, string memory _latestRelease, string memory _releaseMeta) public developer {
        require(bytes(releases[_tag].releaseCID).length == 0, "Tag already used");

        tags.push(_tag);

        releases[_tag] = Release(_tag, _latestRelease, _releaseMeta);

        latestRelease = releases[_tag];

        emit NewRelease(_tag);
    }

    function _deleteRepository(address payable _admin) external {
        require(msg.sender == deployer, "Can only be called from parent contract");

        selfdestruct(_admin);
    }

}
