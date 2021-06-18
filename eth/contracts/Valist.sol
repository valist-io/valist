// SPDX-License-Identifier: MIT
pragma solidity >=0.7.0;
pragma experimental ABIEncoderV2;

import "./EIP712MetaTransaction.sol";
import "@openzeppelin/contracts/utils/EnumerableSet.sol";

contract Valist is EIP712MetaTransaction("Valist","0") {

    using EnumerableSet for EnumerableSet.AddressSet;

    // organization level roles
    bytes32 constant ORG_OWNER = keccak256("ORG_OWNER_ROLE");
    bytes32 constant ORG_ADMIN = keccak256("ORG_ADMIN_ROLE");

    // repository level roles
    bytes32 constant REPO_ADMIN = keccak256("REPO_ADMIN_ROLE");
    bytes32 constant REPO_DEV = keccak256("REPO_DEV_ROLE");

    struct Organization {
        // keep track of shortname in orgNames
        uint index;

        // organization metadata
        string metaCID;

        // list of repo names
        string[] repoNames;

        // mapping of repoName => Repository
        mapping(string => Repository) repos;

        // organization level roles
        mapping(bytes32 => EnumerableSet.AddressSet) roles;
    }

    struct Repository {
        // keep track of shortname in repoNames
        uint index;

        // repo metadata
        string metaCID;

        // list of release tags
        string[] tags;

        // mapping of tag => Release
        mapping(string => Release) releases;

        // repository level roles
        mapping(bytes32 => EnumerableSet.AddressSet) roles;
    }

    struct Release {
        // release artifact
        string releaseCID;

        // release metadata
        string metaCID;

    }

    // list of shortnames
    string[] public orgNames;

    // orgName => Organization
    mapping(string => Organization) public orgs;

    modifier orgOwner(string memory _orgName) {
        require(isOrgOwner(_orgName, msgSender()), "Access Denied");
        _;
    }

    modifier orgAdmin(string memory _orgName) {
        require(isOrgAdmin(_orgName, msgSender()), "Access Denied");
        _;
    }

    modifier repoAdmin(string memory _orgName, string memory _repoName) {
        require(isRepoAdmin(_orgName, _repoName, msgSender()), "Access Denied");
        _;
    }

    modifier repoDev(string memory _orgName, string memory _repoName) {
        require(isRepoDev(_orgName, _repoName, msgSender()), "Access Denied");
        _;
    }

    event ReleaseEvent(string _orgName, string _repoName, string _tag);

    function isOrgOwner(string memory _orgName, address _address) public view returns (bool) {
        return orgs[_orgName].roles[ORG_OWNER].contains(_address);
    }

    function isOrgAdmin(string memory _orgName, address _address) public view returns (bool) {
        return orgs[_orgName].roles[ORG_ADMIN].contains(_address) || isOrgOwner(_orgName, _address);
    }

    function isRepoAdmin(string memory _orgName, string memory _repoName, address _address) public view returns (bool) {
        return orgs[_orgName].repos[_repoName].roles[REPO_ADMIN].contains(_address) || isOrgAdmin(_orgName, _address);
    }

    function isRepoDev(string memory _orgName, string memory _repoName, address _address) public view returns (bool) {
        return orgs[_orgName].repos[_repoName].roles[REPO_DEV].contains(_address) || isRepoAdmin(_orgName, _repoName, _address);
    }

    function grantOrgOwner(string memory _orgName, address _address) public orgOwner(_orgName) {
        orgs[_orgName].roles[ORG_OWNER].add(_address);
    }

    function revokeOrgOwner(string memory _orgName, address _address) public orgOwner(_orgName) {
        orgs[_orgName].roles[ORG_OWNER].remove(_address);
    }

    function grantOrgAdmin(string memory _orgName, address _address) public orgAdmin(_orgName) {
        orgs[_orgName].roles[ORG_ADMIN].add(_address);
    }

    function revokeOrgAdmin(string memory _orgName, address _address) public orgAdmin(_orgName) {
        orgs[_orgName].roles[ORG_ADMIN].remove(_address);
    }

    function grantRepoAdmin(string memory _orgName, string memory _repoName, address _address) public orgAdmin(_orgName) {
        orgs[_orgName].repos[_repoName].roles[REPO_ADMIN].add(_address);
    }

    function revokeRepoAdmin(string memory _orgName, string memory _repoName, address _address) public orgAdmin(_orgName) {
        orgs[_orgName].repos[_repoName].roles[REPO_ADMIN].remove(_address);
    }

    function grantRepoDev(string memory _orgName, string memory _repoName, address _address) public repoAdmin(_orgName, _repoName) {
        orgs[_orgName].repos[_repoName].roles[REPO_DEV].add(_address);
    }

    function revokeRepoDev(string memory _orgName, string memory _repoName, address _address) public repoAdmin(_orgName, _repoName) {
        orgs[_orgName].repos[_repoName].roles[REPO_DEV].remove(_address);
    }

    function getOrgOwners(string memory _orgName) public view returns(address[] memory) {
        EnumerableSet.AddressSet storage ownerSet = orgs[_orgName].roles[ORG_OWNER];

        address[] memory owners = new address[](ownerSet.length());

        for (uint i = 0; i < ownerSet.length(); i++) {
            owners[i] = ownerSet.at(i);
        }

        return owners;
    }

    function getOrgAdmins(string memory _orgName) public view returns(address[] memory) {
        EnumerableSet.AddressSet storage adminSet = orgs[_orgName].roles[ORG_ADMIN];

        address[] memory admins = new address[](adminSet.length());

        for (uint i = 0; i < adminSet.length(); i++) {
            admins[i] = adminSet.at(i);
        }

        return admins;
    }

    function getRepoAdmins(string memory _orgName, string memory _repoName) public view returns(address[] memory) {
        EnumerableSet.AddressSet storage adminSet = orgs[_orgName].repos[_repoName].roles[REPO_ADMIN];

        address[] memory admins = new address[](adminSet.length());

        for (uint i = 0; i < adminSet.length(); i++) {
            admins[i] = adminSet.at(i);
        }

        return admins;
    }

    function getRepoDevs(string memory _orgName, string memory _repoName) public view returns(address[] memory) {
        EnumerableSet.AddressSet storage devSet = orgs[_orgName].repos[_repoName].roles[REPO_DEV];

        address[] memory devs = new address[](devSet.length());

        for (uint i = 0; i < devSet.length(); i++) {
            devs[i] = devSet.at(i);
        }

        return devs;
    }

    function getOrgCount() public view returns(uint) {
        return orgNames.length;
    }

    function getOrgNames() public view returns(string[] memory) {
        return orgNames;
    }

    function getOrganization(string memory _orgName) public view returns (string memory, string[] memory) {
        return (orgs[_orgName].metaCID, orgs[_orgName].repoNames);
    }

    function getOrgMeta(string memory _orgName) public view returns (string memory) {
        return orgs[_orgName].metaCID;
    }

    function setOrgMeta(string memory _orgName, string memory _orgMeta) public orgAdmin(_orgName) {
        orgs[_orgName].metaCID = _orgMeta;
    }

    function createOrganization(string memory _orgName, string memory _orgMeta) public {
        require(bytes(orgs[_orgName].metaCID).length == 0, "Organization exists");

        orgs[_orgName].roles[ORG_OWNER].add(msgSender());

        orgNames.push(_orgName);

        orgs[_orgName].index = orgNames.length - 1;
        orgs[_orgName].metaCID = _orgMeta;
    }

    function getRepoNames(string memory _orgName) public view returns (string[] memory) {
        return orgs[_orgName].repoNames;
    }

    function getRepository(string memory _orgName, string memory _repoName) public view returns(string memory, string[] memory) {
        return (orgs[_orgName].repos[_repoName].metaCID, orgs[_orgName].repos[_repoName].tags);
    }

    function getRepoMeta(string memory _orgName, string memory _repoName) public view returns (string memory) {
        return orgs[_orgName].repos[_repoName].metaCID;
    }

    function setRepoMeta(string memory _orgName, string memory _repoName, string memory _repoMeta) public repoAdmin(_orgName, _repoName) {
        orgs[_orgName].repos[_repoName].metaCID = _repoMeta;
    }

    function createRepository(string memory _orgName, string memory _repoName, string memory _repoMeta) public orgAdmin(_orgName) {
        orgs[_orgName].repoNames.push(_repoName);

        orgs[_orgName].repos[_repoName].index = orgs[_orgName].repoNames.length - 1;
        orgs[_orgName].repos[_repoName].metaCID = _repoMeta;
    }

    function getLatestTag(string memory _orgName, string memory _repoName) public view returns(string memory) {
        string[] storage tags = orgs[_orgName].repos[_repoName].tags;

        return tags[tags.length - 1];
    }

    function getReleaseTags(string memory _orgName, string memory _repoName) public view returns(string[] memory) {
        return orgs[_orgName].repos[_repoName].tags;
    }

    function getRelease(string memory _orgName, string memory _repoName, string memory _tag) public view returns(Release memory) {
        return orgs[_orgName].repos[_repoName].releases[_tag];
    }

    function getLatestRelease(string memory _orgName, string memory _repoName) public view returns(Release memory) {
        string[] storage tags = orgs[_orgName].repos[_repoName].tags;

        return orgs[_orgName].repos[_repoName].releases[tags[tags.length - 1]];
    }

    function publishRelease(
        string memory _orgName,
        string memory _repoName,
        string memory _tag,
        string memory _releaseCID,
        string memory _metaCID
    ) public repoDev(_orgName, _repoName) {
        require(bytes(orgs[_orgName].repos[_repoName].releases[_tag].releaseCID).length == 0, "Tag used in the past");

        orgs[_orgName].repos[_repoName].tags.push(_tag);

        orgs[_orgName].repos[_repoName].releases[_tag] = Release(_releaseCID, _metaCID);

        emit ReleaseEvent(_orgName, _repoName, _tag);
    }

    /*
    function markReleaseCompromised() public{ 
    }
    */

    function deleteRepository(string memory _orgName, string memory _repoName) public orgAdmin(_orgName) {
        _deleteStringFromArray(orgs[_orgName].repoNames, orgs[_orgName].repos[_repoName].index);
        delete orgs[_orgName].repos[_repoName].roles[REPO_ADMIN];
        delete orgs[_orgName].repos[_repoName].roles[REPO_DEV];
        delete orgs[_orgName].repos[_repoName];
    }

    function deleteOrganization(string memory _orgName) public orgOwner(_orgName) {
        _deleteStringFromArray(orgNames, orgs[_orgName].index);
        delete orgs[_orgName].roles[ORG_OWNER];
        delete orgs[_orgName].roles[ORG_ADMIN];
        delete orgs[_orgName];
    }

    // this function does not preserve array order, but is more gas efficient
    function _deleteStringFromArray(string[] storage _array, uint _index) internal {
        // copy last element to index
        _array[_array.length - 1] = _array[_index];
        // delete last element
        _array.pop();
    }

}
