import ValistABI from './abis/Valist.json';

const getValistContract = async (web3: any) => {
    // get network ID and the deployed address
    const networkId = await web3.eth.net.getId()
    // @ts-ignore
    const deployedAddress: any = ValistABI.networks[networkId].address

    // create the instance
    const instance = new web3.eth.Contract(
        ValistABI.abi,
        deployedAddress
    )
    return instance
    }

export default getValistContract