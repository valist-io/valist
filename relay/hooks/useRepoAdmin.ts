import { useEffect, useState, useContext } from 'react';
import ValistContext from '../components/Valist/ValistContext';

/**
 * Check if the current user has REPO_ADMIN_ROLE.
 * @param orgName Organization short name.
 * @param repoName Repository short name.
 * @returns True if the current user has REPO_ADMIN_ROLE, else false.
 */
export default function useRepoAdmin(orgName: string, repoName: string) {
  const valist = useContext(ValistContext);
  const [isRepoAdmin, setIsRepoAdmin] = useState(false);

  useEffect(() => {
    (async () => {
      if (valist && valist.defaultAccount !== '0x0') {
        try {
          setIsRepoAdmin(await valist.isRepoAdmin(orgName, repoName, valist.defaultAccount));
        } catch (e) {
          setIsRepoAdmin(false);
        }
      }
    })();
  }, [valist, orgName, repoName]);

  return isRepoAdmin;
}
