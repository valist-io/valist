import { FunctionComponent, useEffect, useState, useContext } from 'react';
import ValistContext from '../ValistContext/ValistContext';

const IsRepoAdmin:FunctionComponent<any> = (props) => {
    const valist = useContext(ValistContext);
    const [isRepoAdmin, setIsRepoAdmin] = useState(false);

    useEffect(() => {
        (async function() {
            if (valist) {
                try {
                    setIsRepoAdmin(await valist.isRepoAdmin(props.orgName, props.repoName, valist.defaultAccount));
                } catch (e) {
                    setIsRepoAdmin(false);
                }
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
