import { FunctionComponent, useEffect, useState, useContext } from 'react';
import ValistContext from '../ValistContext/ValistContext';

const IsOrgAdmin:FunctionComponent<any> = (props) => {
    const valist = useContext(ValistContext);
    const [isOrgAdmin, setIsOrgAdmin] = useState(false);

    useEffect(() => {
        (async function() {
            if (valist && valist.defaultAccount !== "0x0") {
                try {
                    setIsOrgAdmin(await valist.isOrgAdmin(props.orgName, valist.defaultAccount));
                } catch (e) {
                    setIsOrgAdmin(false);
                }
            }
        })()
    }, [valist]);

    if (isOrgAdmin) {
        return props.children;
    } else {
        return null;
    }
}

export default IsOrgAdmin;
