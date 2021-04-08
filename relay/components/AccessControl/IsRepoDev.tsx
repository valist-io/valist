import { FunctionComponent, useEffect, useState, useContext } from 'react';
import ValistContext from '../Valist';

const IsRepoDev:FunctionComponent<any> = (props) => {
    const valist = useContext(ValistContext);
    const [isRepoAdmin, setIsRepoAdmin] = useState(false);
    const [isRepoDev, setIsRepoDev] = useState(false);

    useEffect(() => {
        (async function() {
            if (valist && valist.defaultAccount !== "0x0") {
                try {
                    setIsRepoAdmin(await valist.isRepoAdmin(props.orgName, props.repoName, valist.defaultAccount));
                    setIsRepoDev(await valist.isRepoDev(props.orgName, props.repoName, valist.defaultAccount));
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
