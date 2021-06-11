// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0;
import "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";
import "hardhat/console.sol";

contract Valist {
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

    // mapping of tag => PendingRelease
    mapping(string => PendingRelease) pendingReleases;

    // mapping of tag => Release
    mapping(string => Release) releases;

    // repository level roles
    mapping(bytes32 => EnumerableSet.AddressSet) roles;
  }

  struct PendingRelease {
    string releaseCID;
    string metaCID;
    // address[] signers;
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
    require(isOrgOwner(_orgName, msg.sender), "Access Denied");
    _;
  }

  modifier orgAdmin(string memory _orgName) {
    require(isOrgAdmin(_orgName, msg.sender), "Access Denied");
    _;
  }

  modifier repoAdmin(string memory _orgName, string memory _repoName) {
    require(isRepoAdmin(_orgName, _repoName, msg.sender), "Access Denied");
    _;
  }

  modifier repoDev(string memory _orgName, string memory _repoName) {
    require(isRepoDev(_orgName, _repoName, msg.sender), "Access Denied");
    _;
  }


  event PendingReleaseEvent(string _orgName, string _repoName, string _tag);
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

  function createOrganization(string memory _orgName, string memory _orgMeta) public {
      require(bytes(orgs[_orgName].metaCID).length == 0, "Organization exists");
      orgs[_orgName].roles[ORG_OWNER].add(msg.sender);
      orgNames.push(_orgName);
      orgs[_orgName].index = orgNames.length - 1;
      orgs[_orgName].metaCID = _orgMeta;
  }

  function createRepository(
    string memory _orgName,
    string memory _repoName,
    string memory _repoMeta
  ) public orgAdmin(_orgName) {
      orgs[_orgName].repoNames.push(_repoName);
      orgs[_orgName].repos[_repoName].index = orgs[_orgName].repoNames.length - 1;
      orgs[_orgName].repos[_repoName].metaCID = _repoMeta;
  }

  function proposeRelease(
    string memory _orgName,
    string memory _repoName,
    string memory _tag,
    string memory _releaseCID,
    string memory _metaCID
  ) public {
    // Check if release tag is already in use
    require(bytes(orgs[_orgName].repos[_repoName].releases[_tag].releaseCID).length == 0, "Tag used in the past");
    orgs[_orgName].repos[_repoName].pendingReleases[_tag] = PendingRelease(_releaseCID, _metaCID);
		emit PendingReleaseEvent(_orgName, _repoName, _tag);
  }

  function signRelease(
    string memory _orgName, 
    string memory _repoName, 
    string memory _tag
  ) public repoDev(_orgName, _repoName){
		// require(orgs[_orgName].repos[_repoName].pendingReleases[_tag].signers[msg.sender]);
		// orgs[_orgName].repos[_repoName].pendingReleases[_tag].signers.push(msg.sender);
  }

  function finalizeRelease(
      string memory _orgName,
      string memory _repoName,
      string memory _tag
  ) public repoDev(_orgName, _repoName) {
      emit ReleaseEvent(_orgName, _repoName, _tag);
  }

  function getLatestRelease(string memory _orgName, string memory _repoName) public view returns(Release memory) {
      string[] storage tags = orgs[_orgName].repos[_repoName].tags;
      return orgs[_orgName].repos[_repoName].releases[tags[tags.length - 1]];
  }
}