import { FunctionComponent, useEffect, useState, useContext } from 'react';
import ValistContext from '../ValistContext/ValistContext';

const IsRepoAdmin:FunctionComponent<any> = (props) => {
    const valist = useContext(ValistContext);
    const [isRepoAdmin, setIsRepoAdmin] = useState(false);

    useEffect(() => {
        (async function() {
            if (valist) {
                let accounts = await valist.web3.eth.getAccounts();
                setIsRepoAdmin(await valist.isRepoAdmin(props.orgName, props.repoName, accounts[0]));
            }
        })()
    }, [valist]);

    if (isRepoAdmin) {
        return props.children;
    } else {
        return null;
    }
}

export default IsRepoAdmin;
