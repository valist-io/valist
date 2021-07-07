// SPDX-License-Identifier: MPL-2.0
pragma solidity >=0.8.4;
import "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";
import "./BaseRelayRecipient.sol";
// import "hardhat/console.sol";

contract Valist is BaseRelayRecipient {

  constructor(address metaTxForwarder) {
    trustedForwarder = metaTxForwarder;
  }

  string public override versionRecipient = "2.2.0";

  using EnumerableSet for EnumerableSet.AddressSet;

  // organization level role
  bytes32 constant ORG_ADMIN = keccak256("ORG_ADMIN_ROLE");
  // repository level role
  bytes32 constant REPO_DEV = keccak256("REPO_DEV_ROLE");
  // key operations
  bytes32 constant ADD_KEY = keccak256("ADD_KEY_OPERATION");
  bytes32 constant REVOKE_KEY = keccak256("REVOKE_KEY_OPERATION");
  bytes32 constant ROTATE_KEY = keccak256("ROTATE_KEY_OPERATION");

  struct Organization {
    // multi-party threshold
    uint threshold;
    // date threshold was set
    uint thresholdDate;
    // metadata CID
    string metaCID;
    // list of repo names
    string[] repoNames;
  }

  struct Repository {
    // multi-party threshold
    uint threshold;
    // date threshold was set
    uint thresholdDate;
    // metadata CID
    string metaCID;
    // list of release tags
    string[] tags;
  }

  struct Release {
    // release artifact
    string releaseCID;
    // release metadata
    string metaCID;
    // release signers
    address[] signers;
  }

  struct PendingRelease {
    // release tag (can be SemVer, CalVer, etc)
    string tag;
    // release artifact
    string releaseCID;
    // release metadata
    string metaCID;
  }

  struct PendingVote {
    // time of first proposal + 7 days
    uint expiration;
    // role signers
    EnumerableSet.AddressSet signers;
  }

  // incrementing orgNumber used for assigning unique IDs to organizations
  // this + chainID also prevents hash collision attacks on mapping selectors across orgs/repos
  uint public orgCount;

  // list of unique orgIDs
  // keccak256(abi.encodePacked(++orgCount, block.chainid))[]
  bytes32[] public orgIDs;

  // list of unique orgNames
  string[] public orgNames;

  // orgName => orgID (can be governed by a DAO in the future)
  mapping(string => bytes32) public orgIDByName;

  // orgID => Organization
  mapping(bytes32 => Organization) public orgs;

  // keccak256(abi.encodePacked(orgID, repoName)) => Repository
  mapping(bytes32 => Repository) public repos;

  // keccak256(abi.encode(orgID, repoName, tag)) => Release
  // using abi.encode prevents hash collisions since there are multiple dynamic types here
  mapping(bytes32 => Release) public releases;

  // keccak256(abi.encodePacked(orgID, ORG_ADMIN)) => orgAdminSet
  // keccak256(abi.encodePacked(orgID, repoName, REPO_DEV)) => repoDevSet
  mapping(bytes32 => EnumerableSet.AddressSet) roles;

  // keccak256(abi.encodePacked(orgID, pendingOrgAdminAddress)) => orgAdminModifiedTimestamp
  // keccak256(abi.encodePacked(orgID, repoName, pendingRepoDevAddress)) => repoDevModifiedTimestamp
  // this is primarily used to auto-expire any pending operations on the same key once a vote is cast for said key
  mapping(bytes32 => uint) public roleModifiedTimestamps;

  // keccak256(abi.encodePacked(orgID, repoName)) => PendingRelease[]
  mapping(bytes32 => PendingRelease[]) public pendingReleaseRequests;

  // orgID => pendingOrgAdminSet
  // keccak256(abi.encodePacked(orgID, repoName)) => pendingRepoDevSet
  mapping(bytes32 => address[]) public pendingRoleRequests;

  // orgID => pendingOrgThreshold[]
  // keccak256(abi.encodePacked(orgID, repoName)) => pendingRepoThreshold[]
  mapping(bytes32 => uint[]) public pendingThresholdRequests;

  // pendingOrgAdminVotes: keccak256(abi.encodePacked(orgID, ORG_ADMIN, OPERATION, pendingOrgAdminAddress)) => signers
  // pendingOrgThresholdVotes: keccak256(abi.encodePacked(orgID, pendingOrgThreshold)) => signers
  // pendingRepoDevVotes: keccak256(abi.encodePacked(orgID, repoName, REPO_DEV, OPERATION, pendingRepoDevAddress)) => signers
  // pendingRepoThresholdVotes: keccak256(abi.encodePacked(orgID, repoName, pendingRepoThreshold)) => signers
  // pendingReleaseVotes: keccak256(abi.encode(orgID, repoName, tag, releaseCID, metaCID)) => PendingVote
  // using abi.encode prevents hash collisions since there are multiple dynamic types here
  mapping(bytes32 => PendingVote) pendingVotes;

  modifier orgAdmin(bytes32 _orgID) {
    require(isOrgAdmin(_orgID, _msgSender()), "Denied");
    _;
  }

  modifier repoDev(bytes32 _orgID, string memory _repoName) {
    require(isRepoDev(_orgID, _repoName, _msgSender()), "Denied");
    _;
  }

  event MetaUpdate(bytes32 _orgID, string _repoName, address _signer, string _metaCID);

  event VoteThresholdEvent(bytes32 _orgID, string _repoName, address _signer, uint _pendingThreshold, uint _sigCount, uint _threshold);

  event VoteKeyEvent(bytes32 _orgID, string _repoName, address _signer, bytes32 _operation, address _key, uint _sigCount, uint _threshold);

  event VoteReleaseEvent(bytes32 _orgID, string _repoName, string _tag, string _releaseCID, string _metaCID, address _signer, uint _sigCount, uint _threshold);

  // check if user has orgAdmin role
  function isOrgAdmin(bytes32 _orgID, address _address) public view returns (bool) {
    bytes32 selector = keccak256(abi.encodePacked(_orgID, ORG_ADMIN));
    return roles[selector].contains(_address);
  }

  // check if user has at least repoDev role
  function isRepoDev(bytes32 _orgID, string memory _repoName, address _address) public view returns (bool) {
    bytes32 selector = keccak256(abi.encodePacked(_orgID, _repoName, REPO_DEV));
    return roles[selector].contains(_address) || isOrgAdmin(_orgID, _address);
  }

  // create an organization and claim an orgName and orgID
  function createOrganization(string memory _orgName, string memory _orgMeta) public {
    require(orgIDByName[_orgName] == 0, "Org exists");
    require(bytes(_orgName).length > 0, "No orgName");
    require(bytes(_orgMeta).length > 0, "No orgMeta");
    // generate new orgID by incrementing and hashing orgCount
    bytes32 orgID = keccak256(abi.encodePacked(++orgCount, block.chainid));
    // map orgName to orgID
    orgIDByName[_orgName] = orgID;
    // set Organization ID and metadata
    orgs[orgID].metaCID = _orgMeta;
    // add to list of orgIDs
    orgIDs.push(orgID);
    // add to list of orgNames
    orgNames.push(_orgName);
    // add creator of org to orgAdmin role
    roles[keccak256(abi.encodePacked(orgID, ORG_ADMIN))].add(_msgSender());
  }

  function createRepository(bytes32 _orgID, string memory _repoName, string memory _repoMeta) public orgAdmin(_orgID) {
    bytes32 repoSelector = keccak256(abi.encodePacked(_orgID, _repoName));
    require(bytes(repos[repoSelector].metaCID).length == 0 && repos[repoSelector].tags.length == 0, "Repo exists");
    require(bytes(_repoName).length > 0, "No repoName");
    require(bytes(_repoMeta).length > 0, "No repoMeta");
    // add repoName to org
    orgs[_orgID].repoNames.push(_repoName);
    // set metadata for repo
    repos[repoSelector].metaCID = _repoMeta;
  }

  function voteRelease(
    bytes32 _orgID,
    string memory _repoName,
    string memory _tag,
    string memory _releaseCID,
    string memory _metaCID
  ) public repoDev(_orgID, _repoName) {
    require(bytes(_tag).length > 0, "No tag");
    require(bytes(_releaseCID).length > 0, "No releaseCID");
    require(bytes(_metaCID).length > 0, "No metaCID");
    bytes32 releaseSelector = keccak256(abi.encode(_orgID, _repoName, _tag));
    require(bytes(releases[releaseSelector].releaseCID).length == 0, "Tag used");
    bytes32 pendingReleaseSelector = keccak256(abi.encode(_orgID, _repoName, _tag, _releaseCID, _metaCID));
    bytes32 repoSelector = keccak256(abi.encodePacked(_orgID, _repoName));
    // if threshold is 1 or 0, skip straight to publishing
    if (repos[repoSelector].threshold <= 1) {
      // create release
      releases[releaseSelector].releaseCID = _releaseCID;
      releases[releaseSelector].metaCID = _metaCID;
      // add user to release signers
      releases[releaseSelector].signers.push(_msgSender());
      // push tag to repo
      repos[repoSelector].tags.push(_tag);
      // fire ReleaseEvent to notify live clients
      emit VoteReleaseEvent(_orgID, _repoName, _tag, _releaseCID, _metaCID, _msgSender(), 1, repos[repoSelector].threshold);
    } else {
      // propose release if this combination of tag, releaseCID, and metaCID has not been used
      if (pendingVotes[pendingReleaseSelector].expiration == 0) {
        // set expiration to now + 7 days
        pendingVotes[pendingReleaseSelector].expiration = block.timestamp + 7 days;
        // add release proposer to signers
        pendingVotes[pendingReleaseSelector].signers.add(_msgSender());
        // push PendingRelease to pendingReleaseRequests for clients to see
        pendingReleaseRequests[repoSelector].push(PendingRelease(_tag, _releaseCID, _metaCID));
      } else {
        require(block.timestamp <= pendingVotes[pendingReleaseSelector].expiration, "Expired");
        require(!pendingVotes[pendingReleaseSelector].signers.contains(_msgSender()), "User voted");
        // add user to list of signers
        pendingVotes[pendingReleaseSelector].signers.add(_msgSender());
        // clean up any revoked keys
        for (uint i = 0; i < pendingVotes[pendingReleaseSelector].signers.length(); ++i) {
          if (!isRepoDev(_orgID, _repoName, pendingVotes[pendingReleaseSelector].signers.at(i))) {
            pendingVotes[pendingReleaseSelector].signers.remove(pendingVotes[pendingReleaseSelector].signers.at(i));
            i = 0; // order not guaranteed in EnumerableSet, need to restart loop :c
          }
        }
        // if threshold has been met, move PendingRelease to Release
        if (pendingVotes[pendingReleaseSelector].signers.length() >= repos[repoSelector].threshold) {
          // add Release to releases for this repo + tag
          releases[releaseSelector].releaseCID = _releaseCID;
          releases[releaseSelector].metaCID = _metaCID;
          // add all signers from PendingVote to finalized array of signers
          for (uint i = 0; i < pendingVotes[pendingReleaseSelector].signers.length(); ++i) {
            releases[releaseSelector].signers.push(pendingVotes[pendingReleaseSelector].signers.at(i));
          }
          // push tag to repo
          repos[repoSelector].tags.push(_tag);
        }
      }
      // fire VoteReleaseEvent to notify live clients
      emit VoteReleaseEvent(
        _orgID,
        _repoName,
        _tag,
        _releaseCID,
        _metaCID,
        _msgSender(),
        pendingVotes[pendingReleaseSelector].signers.length(),
        repos[repoSelector].threshold
      );
    }
  }

  function setOrgMeta(bytes32 _orgID, string memory _metaCID) public orgAdmin(_orgID) {
    orgs[_orgID].metaCID = _metaCID;
    emit MetaUpdate(_orgID, "", _msgSender(), _metaCID);
  }

  function setRepoMeta(bytes32 _orgID, string memory _repoName, string memory _metaCID) public orgAdmin(_orgID) {
    repos[keccak256(abi.encodePacked(_orgID, _repoName))].metaCID = _metaCID;
    emit MetaUpdate(_orgID, _repoName, _msgSender(), _metaCID);
  }

  function voteKey(bytes32 _orgID, string memory _repoName, bytes32 _operation, address _key) public repoDev(_orgID, _repoName) {
    require(_operation == ADD_KEY || _operation == REVOKE_KEY || _operation == ROTATE_KEY, "Invalid op");
    bool isRepoOperation = bytes(_repoName).length > 0;
    bytes32 voteSelector;
    bytes32 roleSelector;
    bytes32 repoSelector;
    bytes32 timestampSelector = keccak256(abi.encodePacked(_orgID, _repoName, _key));
    uint currentThreshold;
    if (isRepoOperation) {
      voteSelector = keccak256(abi.encodePacked(_orgID, _repoName, REPO_DEV, _operation, _key));
      roleSelector = keccak256(abi.encodePacked(_orgID, _repoName, REPO_DEV));
      repoSelector = keccak256(abi.encodePacked(_orgID, _repoName));
      currentThreshold = repos[repoSelector].threshold;
    } else {
      voteSelector = keccak256(abi.encodePacked(_orgID, ORG_ADMIN, _operation, _key));
      roleSelector = keccak256(abi.encodePacked(_orgID, ORG_ADMIN));
      currentThreshold = orgs[_orgID].threshold;
    }
    // when revoking key, ensure key is in role
    if (_operation == REVOKE_KEY) {
      require(roles[roleSelector].contains(_key), "No Key");
    } else {
      // otherwise, if adding or rotating, ensure key is not in role
      require(!roles[roleSelector].contains(_key), "Key exists");
    }
    // allow self-serve key rotation
    if (_operation == ROTATE_KEY) {
      // double check role -- if _msgSender() is an orgAdmin, just relying on the repoDev modifier would allow
      // this function to bypass the threshold and would add the new key and keep the old key
      require(roles[roleSelector].contains(_msgSender()), "Denied");
      roles[roleSelector].remove(_msgSender());
      roles[roleSelector].add(_key);
    } else if (currentThreshold <= 1) {
      // if threshold is 1 or 0, skip straight to role modification
      if (_operation == ADD_KEY) {
        roles[roleSelector].add(_key);
      } else {
        roles[roleSelector].remove(_key);
      }
    } else {
      // propose key if vote not started
      if (pendingVotes[voteSelector].expiration == 0) {
        pendingVotes[voteSelector].expiration = block.timestamp + 7 days;
        if (isRepoOperation) {
          pendingRoleRequests[repoSelector].push(_key);
        } else {
          pendingRoleRequests[_orgID].push(_key);
        }
      } else {
        require(
          block.timestamp <= pendingVotes[voteSelector].expiration ||
          roleModifiedTimestamps[timestampSelector] <= pendingVotes[voteSelector].expiration - 7 days,
          "Expired"
        );
      }
      // add current user to signers
      pendingVotes[voteSelector].signers.add(_msgSender());
      // clean up any revoked keys
      for (uint i = 0; i < pendingVotes[voteSelector].signers.length(); ++i) {
        if (
          isRepoOperation && !isRepoDev(_orgID, _repoName, pendingVotes[voteSelector].signers.at(i)) ||
          !isRepoOperation && !isOrgAdmin(_orgID, pendingVotes[voteSelector].signers.at(i))
        ) {
          pendingVotes[voteSelector].signers.remove(pendingVotes[voteSelector].signers.at(i));
          i = 0; // order not guaranteed in EnumerableSet, need to restart loop :c
        }
      }
      // if threshold met, finalize vote
      if (pendingVotes[voteSelector].signers.length() >= currentThreshold) {
        if (_operation == ADD_KEY) {
          roles[roleSelector].add(_key);
        } else {
          roles[roleSelector].remove(_key);
          uint totalOrgAdmins = roles[keccak256(abi.encodePacked(_orgID, ORG_ADMIN))].length();
          if (
            isRepoOperation
            && (totalOrgAdmins + roles[roleSelector].length() - 1)
            <= currentThreshold
          ) {
            // ensure that threshold does not lock existing members
            repos[repoSelector].threshold--;
          } else if (
            !isRepoOperation && roles[roleSelector].length() - 1 <= currentThreshold
          ) {
            orgs[_orgID].threshold--;
          }
        }
        roleModifiedTimestamps[timestampSelector] = block.timestamp;
        // client needs to now call clearPendingRepoKey
      }
    }
    emit VoteKeyEvent(_orgID, _repoName, _msgSender(), _operation, _key, pendingVotes[voteSelector].signers.length(), currentThreshold);
  }

  function voteThreshold(bytes32 _orgID, string memory _repoName, uint _threshold) public repoDev(_orgID, _repoName) {
    bool isRepoOperation = bytes(_repoName).length > 0;
    bytes32 repoSelector = keccak256(abi.encodePacked(_orgID, _repoName));
    bytes32 requestSelector;
    uint currentThreshold;
    if (isRepoOperation) {
      currentThreshold = repos[repoSelector].threshold;
      requestSelector = repoSelector;
    } else {
      currentThreshold = orgs[_orgID].threshold;
      requestSelector = _orgID;
    }
    require(_threshold != currentThreshold, "Threshold set");
    bytes32 voteSelector = keccak256(abi.encodePacked(_orgID, _repoName, _threshold));
    require(!pendingVotes[voteSelector].signers.contains(_msgSender()), "User voted");
    require(
      isRepoOperation && (
        _threshold <=
        roles[keccak256(abi.encodePacked(_orgID, _repoName, REPO_DEV))].length() +
        roles[keccak256(abi.encodePacked(_orgID, ORG_ADMIN))].length() - 1
      ) ||
      !isRepoOperation && (_threshold <=roles[keccak256(abi.encodePacked(_orgID, ORG_ADMIN))].length() - 1),
      "Not enough members"
    );
    // if threshold not requested yet, set expiration and add to pendingThresholdRequests[]
    if (pendingVotes[voteSelector].expiration == 0) {
      pendingVotes[voteSelector].expiration = block.timestamp + 7 days;
      pendingThresholdRequests[requestSelector].push(_threshold);
    } else {
      require(
        block.timestamp <= pendingVotes[voteSelector].expiration ||
        isRepoOperation && repos[repoSelector].thresholdDate <= pendingVotes[voteSelector].expiration - 7 days ||
        !isRepoOperation && orgs[_orgID].thresholdDate <= pendingVotes[voteSelector].expiration - 7 days,
        "Expired"
      );
    }
    // add user to signers
    pendingVotes[voteSelector].signers.add(_msgSender());
    // clean up any revoked keys
    for (uint i = 0; i < pendingVotes[voteSelector].signers.length(); ++i) {
      if (
        isRepoOperation && !isRepoDev(_orgID, _repoName, pendingVotes[voteSelector].signers.at(i)) ||
        !isRepoOperation && !isOrgAdmin(_orgID, pendingVotes[voteSelector].signers.at(i))
      ) {
        pendingVotes[voteSelector].signers.remove(pendingVotes[voteSelector].signers.at(i));
        i = 0; // order not guaranteed in EnumerableSet, need to restart loop :c
      }
    }
    // if threshold met, finalize vote
    if (
      pendingVotes[voteSelector].signers.length() >= currentThreshold &&
      pendingVotes[voteSelector].signers.length() >= _threshold
    ) {
      if (isRepoOperation) {
        repos[repoSelector].threshold = _threshold;
        repos[repoSelector].thresholdDate = block.timestamp;
      } else {
        orgs[_orgID].threshold = _threshold;
        orgs[_orgID].thresholdDate = block.timestamp;
      }
      // client needs to now call clearPendingRepoThreshold
    }
    emit VoteThresholdEvent(
      _orgID,
      _repoName,
      _msgSender(),
      _threshold,
      pendingVotes[voteSelector].signers.length(),
      currentThreshold
    );
  }

  function clearPendingRelease(
    bytes32 _orgID,
    string memory _repoName,
    string memory _tag,
    string memory _releaseCID,
    string memory _metaCID,
    uint _requestIndex
  ) public repoDev(_orgID, _repoName) {
    bytes32 releaseSelector = keccak256(abi.encode(_orgID, _repoName, _tag));
    bytes32 pendingReleaseSelector = keccak256(abi.encode(_orgID, _repoName, _tag, _releaseCID, _metaCID));
    require(block.timestamp >= pendingVotes[pendingReleaseSelector].expiration || bytes(releases[releaseSelector].releaseCID).length > 0, "Not expired");
    bytes32 repoSelector = keccak256(abi.encodePacked(_orgID, _repoName));
    require(keccak256(abi.encodePacked(pendingReleaseRequests[repoSelector][_requestIndex].tag)) == keccak256(abi.encodePacked(_tag)), "Wrong tag");
    // copy last element to index
    pendingReleaseRequests[repoSelector][_requestIndex] = pendingReleaseRequests[repoSelector][pendingReleaseRequests[repoSelector].length - 1];
    // delete last element
    pendingReleaseRequests[repoSelector].pop();
  }

  function clearPendingKey(bytes32 _orgID, string memory _repoName, bytes32 _operation, address _key, uint _requestIndex) public repoDev(_orgID, _repoName) {
    require(_operation == ADD_KEY || _operation == REVOKE_KEY, "Invalid op");
    bytes32 voteSelector;
    bytes32 requestSelector;
    bytes32 timestampSelector = keccak256(abi.encodePacked(_orgID, _repoName, _key));
    if (bytes(_repoName).length > 0) {
      voteSelector = keccak256(abi.encodePacked(_orgID, _repoName, REPO_DEV, _operation, _key));
      requestSelector = keccak256(abi.encodePacked(_orgID, _repoName));
    } else {
      voteSelector = keccak256(abi.encodePacked(_orgID, ORG_ADMIN, _operation, _key));
      requestSelector = _orgID;
    }
    require(
      block.timestamp >= pendingVotes[voteSelector].expiration ||
      roleModifiedTimestamps[timestampSelector] >= pendingVotes[voteSelector].expiration - 7 days,
      "Not expired"
    );
    require(pendingRoleRequests[requestSelector][_requestIndex] == _key, "Wrong key");
    // remove all pending votes
    while(pendingVotes[voteSelector].signers.length() > 0) {
      pendingVotes[voteSelector].signers.remove(pendingVotes[voteSelector].signers.at(0));
    }
    pendingVotes[voteSelector].expiration = 0;
    // copy last element to index
    pendingRoleRequests[requestSelector][_requestIndex] = pendingRoleRequests[requestSelector][pendingRoleRequests[requestSelector].length - 1];
    // delete last element
    pendingRoleRequests[requestSelector].pop();
  }

  function clearPendingThreshold(bytes32 _orgID, string memory _repoName, uint _threshold, uint _requestIndex) public repoDev(_orgID, _repoName) {
    bool isRepoOperation = bytes(_repoName).length > 0;
    bytes32 voteSelector = keccak256(abi.encodePacked(_orgID, _repoName, _threshold));
    bytes32 requestSelector;
    if (isRepoOperation) {
      requestSelector = keccak256(abi.encodePacked(_orgID, _repoName));
    } else {
      requestSelector = _orgID;
    }
    require(
      block.timestamp >= pendingVotes[voteSelector].expiration ||
      isRepoOperation && repos[requestSelector].thresholdDate >= pendingVotes[voteSelector].expiration - 7 days ||
      !isRepoOperation && orgs[_orgID].thresholdDate >= pendingVotes[voteSelector].expiration - 7 days,
      "Not expired"
    );
    // remove all pending votes
    while (pendingVotes[voteSelector].signers.length() > 0) {
      pendingVotes[voteSelector].signers.remove(pendingVotes[voteSelector].signers.at(0));
    }
    pendingVotes[voteSelector].expiration = 0;
    // copy last element to index
    pendingThresholdRequests[requestSelector][_requestIndex] = pendingThresholdRequests[requestSelector][pendingThresholdRequests[requestSelector].length - 1];
    // delete last element
    pendingThresholdRequests[requestSelector].pop();
  }

  function getLatestRelease(bytes32 _orgID, string memory _repoName) public view returns (string memory, string memory, string memory, address[] memory) {
    Repository storage repo = repos[keccak256(abi.encodePacked(_orgID, _repoName))];
    Release storage release = releases[keccak256(abi.encode(_orgID, _repoName, repo.tags[repo.tags.length - 1]))];
    return (repo.tags[repo.tags.length - 1], release.releaseCID, release.metaCID, release.signers);
  }

  // get paginated list of organization names
  function getOrgNames(uint _page, uint _resultsPerPage) public view returns (string[] memory) {
    uint i = _resultsPerPage * _page - _resultsPerPage;
    uint limit = _page * _resultsPerPage;
    if (limit > orgNames.length) {
      limit = orgNames.length;
    }
    string[] memory _orgNames = new string[](_resultsPerPage);
    for (i; i < limit; ++i) {
      _orgNames[i] = orgNames[i];
    }
    return _orgNames;
  }

  // get paginated list of repo names
  function getRepoNames(bytes32 _orgID, uint _page, uint _resultsPerPage) public view returns (string[] memory) {
    uint i = _resultsPerPage * _page - _resultsPerPage;
    uint limit = _page * _resultsPerPage;
    if (limit > orgs[_orgID].repoNames.length) {
      limit = orgs[_orgID].repoNames.length;
    }
    string[] memory _repoNames = new string[](_resultsPerPage);
    for (i; i < limit; ++i) {
      _repoNames[i] = orgs[_orgID].repoNames[i];
    }
    return _repoNames;
  }

  // get paginated list of release tags from repo
  function getReleaseTags(bytes32 _selector, uint _page, uint _resultsPerPage) public view returns (string[] memory) {
    uint i = _resultsPerPage * _page - _resultsPerPage;
    uint limit = _page * _resultsPerPage;
    if (limit > repos[_selector].tags.length) {
      limit = repos[_selector].tags.length;
    }
    string[] memory _tags = new string[](_resultsPerPage);
    for (i; i < limit; ++i) {
      _tags[i] = repos[_selector].tags[i];
    }
    return _tags;
  }

  function getPendingVotes(bytes32 _selector) public view returns (uint, address[] memory) {
    address[] memory signers = new address[](pendingVotes[_selector].signers.length());
    for (uint i = 0; i < pendingVotes[_selector].signers.length(); ++i) {
      signers[i] = pendingVotes[_selector].signers.at(i);
    }
    return (pendingVotes[_selector].expiration, signers);
  }

  function getRoleMembers(bytes32 _selector) public view returns (address[] memory) {
    address[] memory members = new address[](roles[_selector].length());
    for (uint i = 0; i < roles[_selector].length(); ++i) {
      members[i] = roles[_selector].at(i);
    }
    return members;
  }

  function getPendingReleaseCount(bytes32 _selector) public view returns (uint) {
    return pendingReleaseRequests[_selector].length;
  }

  function getRoleRequestCount(bytes32 _selector) public view returns (uint) {
    return pendingRoleRequests[_selector].length;
  }

  function getThresholdRequestCount(bytes32 _selector) public view returns (uint) {
    return pendingThresholdRequests[_selector].length;
  }

}
