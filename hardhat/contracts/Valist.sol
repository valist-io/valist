// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0;
import "hardhat/console.sol";
import "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";
import "./ValistStorage.sol";

contract Valist is ValistStorage {

  constructor(address metaTxForwarder) ValistStorage(metaTxForwarder) {}

  // get paginated list of organization names
  function getOrgNames(uint _page, uint _resultsPerPage) public view returns (string[] memory) {
    uint i = _resultsPerPage * _page - _resultsPerPage;
    uint limit = _page * _resultsPerPage;
    if (limit > getOrgCount()) {
      limit = getOrgCount();
    }
    string[] memory _orgNames = new string[](_resultsPerPage);
    for (i; i < limit; i++) {
      _orgNames[i] = orgNames[i];
    }
    return _orgNames;
  }

  function getOrgCount() public view returns (uint) {
    return orgNames.length;
  }

  function getOrgMeta(string memory _orgName) public view returns (string memory) {
    return orgs[_orgName].metaCID;
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

  function getLatestTag(string memory _orgName, string memory _repoName) public view returns(string memory) {
    string[] storage tags = orgs[_orgName].repos[_repoName].tags;
    return tags[tags.length - 1];
  }

  function getReleaseTags(string memory _orgName, string memory _repoName) public view returns(string[] memory) {
    return orgs[_orgName].repos[_repoName].tags;
  }

  function getLatestRelease(string memory _orgName, string memory _repoName) public view returns(string memory, string memory) {
    string[] storage tags = orgs[_orgName].repos[_repoName].tags;
    return (
      orgs[_orgName].repos[_repoName].releases[tags[tags.length - 1]].releaseCID,
      orgs[_orgName].repos[_repoName].releases[tags[tags.length - 1]].metaCID
    );
  }
}
