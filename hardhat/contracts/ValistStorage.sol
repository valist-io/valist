// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0;
import "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";
import "hardhat/console.sol";

contract ValistStorage {
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

    // list of pending tags
    string[] pendingReleaseIDs;

    // signer threshold for operations
    uint signerThreshold;

    // mapping of tag => Release
    mapping(string => Release) releases;

    // mapping of keccak256(tag+releaseCID) => pending Release
    mapping(bytes32 => Release) pendingReleases;

    // repository level roles pending approval
    mapping(address => PendingKey) pendingRoles;

    // repository level roles
    mapping(bytes32 => EnumerableSet.AddressSet) roles;
  }

  struct Release {
    // release artifact
    string releaseCID;
    // release metadata
    string metaCID;
    // enumerable set of signers
    EnumerableSet.AddressSet signers;
  }

  struct PendingKey {
    bytes32 role;
    EnumerableSet.AddressSet signers;
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

  event ReleaseEvent(string _orgName, string _repoName, string _tag, string releaseCID, string metaCID);

  function getOrgCount() public view returns (uint) {
    return orgNames.length;
  }

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

  function voteAddRepoKey(string memory _orgName, string memory _repoName, bytes32 _role, address _key) internal repoAdmin(_orgName, _repoName) {
    // stop if key has already signed
    require(!orgs[_orgName].repos[_repoName].pendingRoles[_key].signers.contains(msg.sender), "User already voted to add this key");
    
    // check if key has already been proposed
    if (orgs[_orgName].repos[_repoName].pendingRoles[_key].role.length == 0) {
      orgs[_orgName].repos[_repoName].pendingRoles[_key].role = _role;
    }

    // add user to list of signers approving this key
    orgs[_orgName].repos[_repoName].pendingRoles[_key].signers.add(msg.sender);

    // check if signature threshold has been met
    if (orgs[_orgName].repos[_repoName].pendingRoles[_key].signers.length() >= orgs[_orgName].repos[_repoName].signerThreshold) {
      // add proposed key as repo admin
      orgs[_orgName].repos[_repoName].roles[_role].add(_key);
    }
  }

  function voteRevokeRepoKey(string memory _orgName, string memory _repoName, bytes32 _role, address _key) internal repoAdmin(_orgName, _repoName) {
    // stop if key has already signed
    require(!orgs[_orgName].repos[_repoName].pendingRoles[_key].signers.contains(msg.sender), "User already voted to revoke this key");
    
    // check if key has already been proposed
    if (orgs[_orgName].repos[_repoName].pendingRoles[_key].role.length == 0) {
      orgs[_orgName].repos[_repoName].pendingRoles[_key].role = _role;
    }

    // add user to list of signers approving this key
    orgs[_orgName].repos[_repoName].pendingRoles[_key].signers.add(msg.sender);

    // check if signature threshold has been met
    if (orgs[_orgName].repos[_repoName].pendingRoles[_key].signers.length() >= orgs[_orgName].repos[_repoName].signerThreshold) {
      // add proposed key as repo admin
      orgs[_orgName].repos[_repoName].roles[_role].add(_key);
    }
  }

  function voteAddRepoAdmin(string memory _orgName, string memory _repoName, address _key) public {
    voteAddRepoKey(_orgName, _repoName, REPO_ADMIN, _key);
  }

  function voteAddRepoDev(string memory _orgName, string memory _repoName, address _key) public {
    voteAddRepoKey(_orgName, _repoName, REPO_DEV, _key);
  }

  function voteRevokeRepoAdmin(string memory _orgName, string memory _repoName, address _key) public {
    voteRevokeRepoKey(_orgName, _repoName, REPO_ADMIN, _key);
  }

  function voteRevokeRepoDev(string memory _orgName, string memory _repoName, address _key) public {
    voteRevokeRepoKey(_orgName, _repoName, REPO_DEV, _key);
  }

  function createOrganization(string memory _orgName, string memory _orgMeta) public {
    require(bytes(orgs[_orgName].metaCID).length == 0, "Organization exists");
    require(bytes(_orgName).length > 0, "Must provide orgName");
    require(bytes(_orgMeta).length > 0, "Must provide orgMeta");

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
    require(bytes(orgs[_orgName].metaCID).length > 0, "Organization does not exist");
    require(bytes(orgs[_orgName].repos[_repoName].metaCID).length == 0, "Repository already exists");
    require(bytes(_repoMeta).length > 0, "Must provide repoMeta");

    orgs[_orgName].repoNames.push(_repoName);
    orgs[_orgName].repos[_repoName].index = orgs[_orgName].repoNames.length - 1;
    orgs[_orgName].repos[_repoName].metaCID = _repoMeta;
  }

  function setRepoSignerThreshold(
    string memory _orgName,
    string memory _repoName,
    uint _threshold
  ) public repoAdmin(_orgName, _repoName) {
    require(_threshold <= (
      orgs[_orgName].roles[ORG_OWNER].length() +
      orgs[_orgName].roles[ORG_ADMIN].length() +
      orgs[_orgName].repos[_repoName].roles[REPO_ADMIN].length() +
      orgs[_orgName].repos[_repoName].roles[REPO_DEV].length()),
      "Not enough members in this organization"
    );

    orgs[_orgName].repos[_repoName].signerThreshold = _threshold;
  }

  function voteRelease(
    string memory _orgName,
    string memory _repoName,
    string memory _tag,
    string memory _releaseCID,
    string memory _metaCID
  ) public repoDev(_orgName, _repoName) {
    require(bytes(orgs[_orgName].repos[_repoName].metaCID).length > 0, "Repository does not exist");
    require(bytes(orgs[_orgName].repos[_repoName].releases[_tag].releaseCID).length == 0, "Tag already released");

    bytes32 releaseID = keccak256(abi.encodePacked(_tag, _releaseCID));

    // propose release if tag not used
    if (bytes(orgs[_orgName].repos[_repoName].pendingReleases[releaseID].releaseCID).length == 0) {
      orgs[_orgName].repos[_repoName].pendingReleases[releaseID].releaseCID = _releaseCID;
      orgs[_orgName].repos[_repoName].pendingReleases[releaseID].metaCID = _metaCID;
      orgs[_orgName].repos[_repoName].pendingReleases[releaseID].signers.add(msg.sender);
      emit PendingReleaseEvent(_orgName, _repoName, _tag);
    } else {
      // release already proposed, continue with adding signers
      require(!orgs[_orgName].repos[_repoName].pendingReleases[releaseID].signers.contains(msg.sender), "User already signed this release");
      require(keccak256(bytes(_releaseCID)) == keccak256(bytes(orgs[_orgName].repos[_repoName].pendingReleases[releaseID].releaseCID)), "releaseCID does not match proposed");
      require(keccak256(bytes(_metaCID)) == keccak256(bytes(orgs[_orgName].repos[_repoName].pendingReleases[releaseID].metaCID)), "metaCID does not match proposed");

      // add user to list of signers for this release
      orgs[_orgName].repos[_repoName].pendingReleases[releaseID].signers.add(msg.sender);

      // if signature threshold has been met, finalize release
      if (orgs[_orgName].repos[_repoName].pendingReleases[releaseID].signers.length() == orgs[_orgName].repos[_repoName].signerThreshold) {
        Release storage release = orgs[_orgName].repos[_repoName].pendingReleases[releaseID];
        orgs[_orgName].repos[_repoName].tags.push(_tag);
        orgs[_orgName].repos[_repoName].releases[_tag].releaseCID = release.releaseCID;
        orgs[_orgName].repos[_repoName].releases[_tag].metaCID = release.metaCID;

        // copy signers from pendingRelease to finalized Release
        for (uint i = 0; i < release.signers.length(); i++) {
          orgs[_orgName].repos[_repoName].releases[_tag].signers.add(release.signers.at(i));
        }

        emit ReleaseEvent(_orgName, _repoName, _tag, release.releaseCID, release.metaCID);
      }
    }
  }

}