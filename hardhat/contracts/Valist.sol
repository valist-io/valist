// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0;
import "hardhat/console.sol";
import "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";
import "./ValistStorage.sol";

contract Valist is ValistStorage {
  
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

  function getOrganization(string memory _orgName) public view returns (uint, string memory, string[] memory) {
    Organization storage org = orgs[_orgName];
    return (org.index, org.metaCID, org.repoNames);
  }

  function getLatestRelease(string memory _orgName, string memory _repoName) public view returns(string memory, string memory) {
    string[] storage tags = orgs[_orgName].repos[_repoName].tags;
    return (
      orgs[_orgName].repos[_repoName].releases[tags[tags.length - 1]].releaseCID,
      orgs[_orgName].repos[_repoName].releases[tags[tags.length - 1]].metaCID
    );
  }

  /*
  function getRepository(string memory _orgName, string memory _repoName) public view returns(string memory, string[] memory) {
    return (valist.orgs[_orgName].repos[_repoName].metaCID, valist.orgs[_orgName].repos[_repoName].tags);
  }*/
}