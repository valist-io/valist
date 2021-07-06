import { useEffect, useState, useContext } from 'react';
import ValistContext from '../components/Valist/ValistContext';

/**
 * Check if the current user has REPO_DEV_ROLE.
 * @param orgName Organization short name.
 * @param repoName Repository short name.
 * @returns True if the current user has REPO_DEV_ROLE, else false.
 */
export default function useRepoDev(orgName: string, repoName: string) {
  const valist = useContext(ValistContext);
  const [isRepoDev, setIsRepoDev] = useState(false);

  useEffect(() => {
    (async () => {
      if (valist && valist.defaultAccount !== '0x0' && orgName && repoName) {
        try {
          setIsRepoDev(await valist.isRepoDev(orgName, repoName, valist.defaultAccount));
        } catch (e) {
          setIsRepoDev(false);
        }
      }
    })();
  }, [valist, orgName, repoName]);

  return isRepoDev;
}
