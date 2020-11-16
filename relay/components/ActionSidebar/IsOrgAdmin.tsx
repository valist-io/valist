import React, { FunctionComponent, useEffect, useState, useContext } from 'react';
import ValistContext from '../ValistContext/ValistContext';

const IsOrgAdmin:FunctionComponent<any> = (props) => {
    const valist = useContext(ValistContext)
    const [isOrgAdmin, setIsOrgAdmin] = useState(false);

    useEffect(() => {
        (async function() {
            if (valist) {
                let accounts = await valist.web3.eth.getAccounts();
                setIsOrgAdmin(await valist.isOrgAdmin(props.orgName, accounts[0]))
            }
        })()
    }, [valist]);

    if(isOrgAdmin){
        return props.children
    }else{
        return null
    }
}

export default IsOrgAdmin;

