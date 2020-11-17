import { FunctionComponent, useEffect, useState, useContext } from 'react';
import ValistContext from '../ValistContext/ValistContext';

const IsRepoDev:FunctionComponent<any> = (props) => {
    const valist = useContext(ValistContext);
    const [isRepoAdmin, setIsRepoAdmin] = useState(false);
    const [isRepoDev, setIsRepoDev] = useState(false);

    useEffect(() => {
        (async function() {
            if (valist) {
                let accounts = await valist.web3.eth.getAccounts();
                try {
                    setIsRepoAdmin(await valist.isRepoAdmin(props.orgName, props.repoName, accounts[0]));
                    setIsRepoDev(await valist.isRepoDev(props.orgName, props.repoName, accounts[0]));
                } catch (e) {
                    setIsRepoAdmin(false);
                    setIsRepoDev(false);
                }
            }
        })()
    }, [valist]);

    if (isRepoAdmin || isRepoDev) {
        return props.children;
    } else {
        return null;
    }
}

export default IsRepoDev;
