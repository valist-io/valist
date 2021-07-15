// SPDX-License-Identifier: MPL-2.0
pragma solidity >=0.8.4;
import "@openzeppelin/contracts/access/AccessControl.sol";
import "@openzeppelin/contracts/metatx/ERC2771Context.sol";

contract ValistRegistry is ERC2771Context, AccessControl {

  // can't run _setupRole in constructor due to AccessControl calling _msgSender
  // since _msgSender is overridden by constructor, the call fails
  // use these to create init function that can only be run once
  address immutable deployer;
  bool initialized = false;

  constructor(address metaTxForwarder) ERC2771Context(metaTxForwarder) {
    deployer = msg.sender;
  }

  function _msgSender() internal view override(Context, ERC2771Context)
    returns (address sender) {
    sender = ERC2771Context._msgSender();
  }

  function _msgData() internal view override(Context, ERC2771Context)
    returns (bytes calldata) {
    return ERC2771Context._msgData();
  }

  string public versionRecipient = "2.2.0";

  bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");

  modifier admin {
    require(hasRole(ADMIN_ROLE, _msgSender()), "Denied");
    _;
  }

  // list of unique names
  string[] public names;

  // orgName or repoName => orgID (can be governed by a DAO in the future)
  mapping(string => bytes32) public nameToID;

  // log new mapping links
  event MappingEvent(bytes32 indexed _orgID, string indexed _nameHash, string _name);

  function getNameCount() public view returns (uint) {
    return names.length;
  }

  // get paginated list of names
  function getNames(uint _page, uint _resultsPerPage) public view returns (string[] memory) {
    uint i = _resultsPerPage * _page - _resultsPerPage;
    uint limit = _page * _resultsPerPage;
    if (limit > names.length) {
      limit = names.length;
    }
    string[] memory _names = new string[](_resultsPerPage);
    for (i; i < limit; ++i) {
      _names[i] = names[i];
    }
    return _names;
  }

  function linkNameToID(bytes32 _orgID, string memory _name) public {
    require(nameToID[_name] == "", "Mapping exists");
    nameToID[_name] = _orgID;
    names.push(_name);
    emit MappingEvent(_orgID, _name, _name);
  }

  function overrideNameToID(bytes32 _orgID, string memory _name) public admin {
    require(nameToID[_name] != "", "No mapping");
    nameToID[_name] = _orgID;
    emit MappingEvent(_orgID, _name, _name);
  }

  function init() public {
    require(!initialized, "Initialized");
    require(_msgSender() == deployer, "Not deployer");
    _setupRole(ADMIN_ROLE, deployer);
    initialized = true;
  }
}
