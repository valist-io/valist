import { useEffect, useState, useContext } from 'react';
import ValistContext from '../components/Valist/ValistContext';

/**
 * Check if the current user has ORG_ADMIN_ROLE.
 * @param orgName Organization short name.
 * @returns True if the current user has ORG_ADMIN_ROLE, else false.
 */
export default function useOrgAdmin(orgName: string) {
  const valist = useContext(ValistContext);
  const [isOrgAdmin, setIsOrgAdmin] = useState(false);

  useEffect(() => {
    (async () => {
      if (valist && valist.defaultAccount !== '0x0' && orgName) {
        try {
          setIsOrgAdmin(await valist.isOrgAdmin(orgName, valist.defaultAccount));
        } catch (e) {
          setIsOrgAdmin(false);
        }
      }
    })();
  }, [valist, orgName]);

  return isOrgAdmin;
}
