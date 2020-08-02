import ValistOrganizationABI from './abis/ValistOrganization.json';

const getValistContract = async (web3: any, deployedAddress: string) => {
    // get network ID and the deployed address
    const networkId = await web3.eth.net.getId()

    // create the instance
    const instance = new web3.eth.Contract(
        ValistOrganizationABI.abi,
        deployedAddress
    )
    return instance
    }

export default getValistContract

