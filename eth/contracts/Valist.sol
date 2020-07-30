pragma solidity >=0.4.21 <0.7.0;

import "@openzeppelin/contracts/access/AccessControl.sol";

contract Valist is AccessControl {

    bytes32 public constant APP_OWNER = keccak256("APP_OWNER_ROLE");
    bytes32 public constant ORG_OWNER = keccak256("ORG_OWNER_ROLE");

    bytes32 public constant APP_ADMIN = keccak256("APP_ADMIN_ROLE");
    bytes32 public constant ORG_ADMIN = keccak256("ORG_ADMIN_ROLE");

    bytes32 public constant APP_DEV = keccak256("APP_DEV_ROLE");

    struct Repository {
        mapping(bytes32 => address[]) usersByRole;
        string name;
        bytes32 id;
    }

    struct Organization {
        mapping(bytes32 => address[]) usersByRole;
        string name;
        bytes32 id;
    }

    mapping(string => Repository[]) repos;

    mapping(string => Organization) organization;

}
