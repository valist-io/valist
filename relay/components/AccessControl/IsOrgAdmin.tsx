import {
  FunctionComponent, useEffect, useState, useContext,
} from 'react';
import ValistContext from '../Valist/ValistContext';

const IsOrgAdmin:FunctionComponent<any> = (props) => {
  const valist = useContext(ValistContext);
  const [isOrgAdmin, setIsOrgAdmin] = useState(false);

  useEffect(() => {
    (async () => {
      if (valist && valist.defaultAccount !== '0x0') {
        try {
          setIsOrgAdmin(await valist.isOrgAdmin(props.orgName, valist.defaultAccount));
        } catch (e) {
          setIsOrgAdmin(false);
        }
      }
    })();
  }, [valist]);

  if (isOrgAdmin) {
    return props.children;
  }
  return null;
};

export default IsOrgAdmin;
