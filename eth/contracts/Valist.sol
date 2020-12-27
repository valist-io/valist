// SPDX-License-Identifier: MIT
pragma solidity >=0.7.0;
pragma experimental ABIEncoderV2;

import "./ValistOrganization.sol";

contract Valist {

    bytes32 constant ORG_ADMIN = keccak256("ORG_ADMIN_ROLE");

    struct Organization {
        uint index;
        ValistOrganization org;
    }

    // map an organization/username handle to an Organization contract
    // this enables the following schema:
    // [valist.io/other relays]/[organization]/[repository]
    mapping(string => Organization) public orgs;

    string[] public orgNames;

    function getOrgNames() public view returns (string[] memory) {
        return orgNames;
    }

    // register organization/username to the global valist namespace
    function createOrganization(string memory _orgName, string memory _orgMeta) public returns(address) {
        require(address(orgs[_orgName].org) == address(0), "Organization exists");

        orgNames.push(_orgName);

        orgs[_orgName].index = orgNames.length - 1;

        orgs[_orgName].org = new ValistOrganization(msg.sender, _orgMeta);

        return address(orgs[_orgName].org);
    }

    function deleteOrganization(string memory _orgName) public {
        require(orgs[_orgName].org.hasRole(ORG_ADMIN, msg.sender), "Access Denied");

        ValistOrganization org = orgs[_orgName].org;

        delete orgNames[orgs[_orgName].index];
        delete orgs[_orgName];

        org._deleteOrganization(msg.sender);
    }

}
