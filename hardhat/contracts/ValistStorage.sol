// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0;
import "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";
import "./BaseRelayRecipient.sol";

contract ValistStorage is BaseRelayRecipient {

  constructor(address metaTxForwarder) {
    trustedForwarder = metaTxForwarder;
  }

  string public override versionRecipient = "2.2.0";

  using EnumerableSet for EnumerableSet.AddressSet;

  // organization level role
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

    // signer threshold for operations
    uint signerThreshold;

    // list of proposed thresholds (cleared when threshold finalized)
    uint[] pendingThresholds;

    // mapping of threshold => signers
    mapping(uint => EnumerableSet.AddressSet) pendingThreshold;
  }

  struct Repository {
    // keep track of shortname in repoNames
    uint index;

    // repo metadata
    string metaCID;

    // list of release tags
    string[] tags;

    // list of pending tags
    bytes32[] pendingReleaseIDs;

    // mapping of tag => Release
    mapping(string => Release) releases;

    // mapping of keccak256(tag+releaseCID) => pending Release
    mapping(bytes32 => PendingRelease) pendingReleases;

    // repository level roles pending approval
    mapping(address => PendingKey) pendingRoles;

    // repository level roles
    mapping(bytes32 => EnumerableSet.AddressSet) roles;

    // signer threshold for operations
    uint signerThreshold;

    // list of proposed thresholds (cleared when threshold finalized)
    uint[] pendingThresholds;

    // mapping of threshold => signers
    mapping(uint => EnumerableSet.AddressSet) pendingThreshold;
  }

  struct Release {
    // release artifact
    string releaseCID;
    // release metadata
    string metaCID;
    // enumerable set of signers
    EnumerableSet.AddressSet signers;
  }

  struct PendingRelease {
    uint index;
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

  modifier orgAdmin(string memory _orgName) {
    require(isOrgAdmin(_orgName, _msgSender()), "Denied");
    _;
  }

  modifier repoAdmin(string memory _orgName, string memory _repoName) {
    require(isRepoAdmin(_orgName, _repoName, _msgSender()), "Denied");
    _;
  }

  modifier repoDev(string memory _orgName, string memory _repoName) {
    require(isRepoDev(_orgName, _repoName, _msgSender()), "Denied");
    _;
  }

  event PendingReleaseEvent(string _orgName, string _repoName, bytes32 _releaseID);

  event ReleaseEvent(string _orgName, string _repoName, string _tag, string releaseCID, string metaCID);

  function isOrgAdmin(string memory _orgName, address _address) public view returns (bool) {
    return orgs[_orgName].roles[ORG_ADMIN].contains(_address);
  }

  function isRepoAdmin(string memory _orgName, string memory _repoName, address _address) public view returns (bool) {
    return orgs[_orgName].repos[_repoName].roles[REPO_ADMIN].contains(_address) || isOrgAdmin(_orgName, _address);
  }

  function isRepoDev(string memory _orgName, string memory _repoName, address _address) public view returns (bool) {
    return orgs[_orgName].repos[_repoName].roles[REPO_DEV].contains(_address) || isRepoAdmin(_orgName, _repoName, _address);
  }

  function getOrganization(string memory _orgName) public view returns(uint, string memory, string[] memory) {
    Organization storage org = orgs[_orgName];
    return (org.index, org.metaCID, org.repoNames);
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

  function getTotalMembers(string memory _orgName, string memory _repoName) public view returns(uint) {
    return (
      orgs[_orgName].roles[ORG_ADMIN].length() +
      orgs[_orgName].repos[_repoName].roles[REPO_ADMIN].length() +
      orgs[_orgName].repos[_repoName].roles[REPO_DEV].length()
    );
  }

  function setRepoMeta(string memory _orgName, string memory _repoName, string memory _repoMeta) public repoAdmin(_orgName, _repoName) {
    orgs[_orgName].repos[_repoName].metaCID = _repoMeta;
  }

  function voteOrgThreshold(string memory _orgName, uint _threshold) public orgAdmin(_orgName) {
    require(_threshold <= orgs[_orgName].roles[ORG_ADMIN].length() - 1, "Not enough members in this organization to support a threshold this high");
    require(!orgs[_orgName].pendingThreshold[_threshold].contains(_msgSender()), "User already voted on threshold");
    // add user to list approving threshold
    orgs[_orgName].pendingThreshold[_threshold].add(_msgSender());
    // check if signature threshold has been met
    if (orgs[_orgName].pendingThreshold[_threshold].length() >= orgs[_orgName].signerThreshold) {
      // set target threshold for repo
      orgs[_orgName].signerThreshold = _threshold;
      for (uint i = 0; i < orgs[_orgName].pendingThreshold[_threshold].length(); i++) {
        orgs[_orgName].pendingThreshold[_threshold].remove(orgs[_orgName].pendingThreshold[_threshold].at(0));
      }
      // clear pending threshold array
      delete orgs[_orgName].pendingThresholds;
    }
  }

  function voteRepoThreshold(string memory _orgName, string memory _repoName, uint _threshold) public repoDev(_orgName, _repoName) {
    uint memberCount = getTotalMembers(_orgName, _repoName);
    require(_threshold <= memberCount - 1, "Not enough members in this organization to support a threshold this high");
    require(!orgs[_orgName].repos[_repoName].pendingThreshold[_threshold].contains(_msgSender()), "User already voted on threshold");
    // add user to list approving threshold
    orgs[_orgName].repos[_repoName].pendingThreshold[_threshold].add(_msgSender());
    // check if signature threshold has been met
    if (orgs[_orgName].repos[_repoName].pendingThreshold[_threshold].length() >= orgs[_orgName].repos[_repoName].signerThreshold) {
      // set target threshold for repo
      orgs[_orgName].repos[_repoName].signerThreshold = _threshold;
      for (uint i = 0; i < orgs[_orgName].repos[_repoName].pendingThreshold[_threshold].length(); i++) {
        orgs[_orgName].pendingThreshold[_threshold].remove(orgs[_orgName].pendingThreshold[_threshold].at(0));
      }
      // clear pending threshold array
      delete orgs[_orgName].repos[_repoName].pendingThresholds;
    }
  }

  function voteAddRepoKey(string memory _orgName, string memory _repoName, bytes32 _role, address _key) internal repoAdmin(_orgName, _repoName) {
    require(!(
      orgs[_orgName].repos[_repoName].roles[REPO_DEV].contains(_key) ||
      orgs[_orgName].repos[_repoName].roles[REPO_ADMIN].contains(_key)
    ), "User already repoDev or repoAdmin");
    require(!orgs[_orgName].repos[_repoName].pendingRoles[_key].signers.contains(_msgSender()), "User already voted");
    // check if key has already been proposed
    if (orgs[_orgName].repos[_repoName].pendingRoles[_key].role.length == 0) {
      orgs[_orgName].repos[_repoName].pendingRoles[_key].role = _role;
    }
    // add user to list of signers approving this key
    orgs[_orgName].repos[_repoName].pendingRoles[_key].signers.add(_msgSender());
    // check if signature threshold has been met
    if (orgs[_orgName].repos[_repoName].pendingRoles[_key].signers.length() >= orgs[_orgName].repos[_repoName].signerThreshold) {
      // add proposed key as repo admin
      orgs[_orgName].repos[_repoName].roles[_role].add(_key);
    }
  }

  function voteRevokeRepoKey(string memory _orgName, string memory _repoName, bytes32 _role, address _key) internal repoAdmin(_orgName, _repoName) {
    require((
      orgs[_orgName].repos[_repoName].roles[REPO_DEV].contains(_key) ||
      orgs[_orgName].repos[_repoName].roles[REPO_ADMIN].contains(_key)
    ), "User not at least repoDev");
    require(!orgs[_orgName].repos[_repoName].pendingRoles[_key].signers.contains(_msgSender()), "User already voted");
    // check if key has already been proposed
    if (orgs[_orgName].repos[_repoName].pendingRoles[_key].role.length == 0) {
      orgs[_orgName].repos[_repoName].pendingRoles[_key].role = _role;
    }
    // add user to list of signers approving this key
    orgs[_orgName].repos[_repoName].pendingRoles[_key].signers.add(_msgSender());
    // check if signature threshold has been met
    if (orgs[_orgName].repos[_repoName].pendingRoles[_key].signers.length() >= orgs[_orgName].repos[_repoName].signerThreshold) {
      // add proposed key as repo admin
      orgs[_orgName].repos[_repoName].roles[_role].add(_key);
    }
  }

  function voteAddRepoAdmin(string memory _orgName, string memory _repoName, address _key) public {
    voteAddRepoKey(_orgName, _repoName, REPO_ADMIN, _key);
  }

  function voteRevokeRepoAdmin(string memory _orgName, string memory _repoName, address _key) public {
    voteRevokeRepoKey(_orgName, _repoName, REPO_ADMIN, _key);
  }

  function voteAddRepoDev(string memory _orgName, string memory _repoName, address _key) public {
    voteAddRepoKey(_orgName, _repoName, REPO_DEV, _key);
  }

  function voteRevokeRepoDev(string memory _orgName, string memory _repoName, address _key) public {
    voteRevokeRepoKey(_orgName, _repoName, REPO_DEV, _key);
  }

  function createOrganization(string memory _orgName, string memory _orgMeta) public {
    require(bytes(orgs[_orgName].metaCID).length == 0, "Organization exists");
    require(bytes(_orgName).length > 0, "Must provide orgName");
    require(bytes(_orgMeta).length > 0, "Must provide orgMeta");

    orgs[_orgName].roles[ORG_ADMIN].add(_msgSender());
    orgNames.push(_orgName);
    orgs[_orgName].index = orgNames.length - 1;
    orgs[_orgName].metaCID = _orgMeta;
  }

  function setOrgMeta(string memory _orgName, string memory _orgMeta) public orgAdmin(_orgName) {
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

  function getRelease(string memory _orgName, string memory _repoName, string memory _tag) public view returns(string memory, string memory, address[] memory) {
    Release storage release = orgs[_orgName].repos[_repoName].releases[_tag];
    EnumerableSet.AddressSet storage _signers = orgs[_orgName].repos[_repoName].releases[_tag].signers;
    address[] memory signers = new address[](_signers.length());
    for (uint i = 0; i < _signers.length(); i++) {
      signers[i] = _signers.at(i);
    }
    return (release.releaseCID, release.metaCID, signers);
  }

  function voteRelease(
    string memory _orgName,
    string memory _repoName,
    string memory _tag,
    string memory _releaseCID,
    string memory _metaCID
  ) public repoDev(_orgName, _repoName) {
    require(bytes(orgs[_orgName].repos[_repoName].metaCID).length > 0, "Repo does not exist");
    require(bytes(orgs[_orgName].repos[_repoName].releases[_tag].releaseCID).length == 0, "Tag already released");

    bytes32 releaseID = keccak256(abi.encodePacked(_tag, _releaseCID));

    // propose release if tag not used
    if (bytes(orgs[_orgName].repos[_repoName].pendingReleases[releaseID].releaseCID).length == 0) {
      orgs[_orgName].repos[_repoName].pendingReleaseIDs.push(releaseID);
      PendingRelease storage release = orgs[_orgName].repos[_repoName].pendingReleases[releaseID];
      release.index = orgs[_orgName].repos[_repoName].pendingReleaseIDs.length - 1;
      release.releaseCID = _releaseCID;
      release.metaCID = _metaCID;
      release.signers.add(_msgSender());
      emit PendingReleaseEvent(_orgName, _repoName, releaseID);
    } else {
      // release already proposed, continue with adding signers
      require(!orgs[_orgName].repos[_repoName].pendingReleases[releaseID].signers.contains(_msgSender()), "User already signed");
      require(keccak256(bytes(_releaseCID)) == keccak256(bytes(orgs[_orgName].repos[_repoName].pendingReleases[releaseID].releaseCID)), "releaseCID does not match");
      require(keccak256(bytes(_metaCID)) == keccak256(bytes(orgs[_orgName].repos[_repoName].pendingReleases[releaseID].metaCID)), "metaCID does not match");

      // add user to list of signers for this release
      orgs[_orgName].repos[_repoName].pendingReleases[releaseID].signers.add(_msgSender());

      // if signature threshold has been met, finalize release
      if (orgs[_orgName].repos[_repoName].pendingReleases[releaseID].signers.length() == orgs[_orgName].repos[_repoName].signerThreshold) {
        PendingRelease storage release = orgs[_orgName].repos[_repoName].pendingReleases[releaseID];
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

  function clearPendingRelease(string memory _orgName,
    string memory _repoName,
    string memory _tag,
    string memory _releaseCID
  ) public repoDev(_orgName, _repoName) {
    require(bytes(orgs[_orgName].repos[_repoName].releases[_tag].releaseCID).length > 0, "Tag not released");
    bytes32 releaseID = keccak256(abi.encodePacked(_tag, _releaseCID));
    PendingRelease storage release = orgs[_orgName].repos[_repoName].pendingReleases[releaseID];
    _deleteBytesFromArray(orgs[_orgName].repos[_repoName].pendingReleaseIDs, release.index);
    for (uint i = 0; i < release.signers.length(); i++) {
      orgs[_orgName].repos[_repoName].pendingReleases[releaseID].signers.remove(
        orgs[_orgName].repos[_repoName].pendingReleases[releaseID].signers.at(0)
      );
    }
    delete orgs[_orgName].repos[_repoName].pendingReleases[releaseID];
  }

  // this function does not preserve array order, but is more gas efficient
  function _deleteBytesFromArray(bytes32[] storage _array, uint _index) internal {
    // copy last element to index
    _array[_array.length - 1] = _array[_index];
    // delete last element
    _array.pop();
  }
}
