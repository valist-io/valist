// We require the Hardhat Runtime Environment explicitly here. This is optional
// but useful for running the script in a standalone fashion through `node <script>`.
//
// When running the script with `hardhat run <script>` you'll find the Hardhat
// Runtime Environment's members available in the global scope.
const hre = require("hardhat");

async function main() {
  // Hardhat always runs the compile task when running scripts with its command
  // line interface.
  //
  // If this script is run directly using `node` you may want to call compile
  // manually to make sure everything is compiled
  // await hre.run('compile');

  const trustedForwarders = {
    80001: '0x9399BB24DBB5C4b782C70c2969F58716Ebbd6a3b',
    137: '0x86C80a8aa58e0A4fa09A69624c31Ab2a6CAD56b8'
  }
  const Valist = await hre.ethers.getContractFactory("Valist");
  const valist = await Valist.deploy(trustedForwarders[80001]);

  await valist.deployed();

  const ValistRegistry = await hre.ethers.getContractFactory("ValistRegistry");
  const registry = await ValistRegistry.deploy(trustedForwarders[80001]);

  console.log("Valist deployed to:", valist.address);
  console.log("ValistRegistry deployed to:", registry.address);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main()
  .then(() => process.exit(0))
  .catch(error => {
    console.error(error);
    process.exit(1);
  });
