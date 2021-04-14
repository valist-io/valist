import React, { useContext, useEffect, useState } from 'react';
import LoadingDialog from '../LoadingDialog/LoadingDialog';
import ValistContext from '../Valist/ValistContext';
import UserAccessCard from './UserAccessCard';

const ProjectPermissions = ({ orgName, repoName }: { orgName: string, repoName: string }): JSX.Element => {
  const valist = useContext(ValistContext);

  const [repoAdmins, setRepoAdmins] = useState(['0x0']);
  const [repoDevs, setRepoDevs] = useState(['0x0']);
  const [orgOwners, setOrgOwners] = useState(['0x0']);
  const [grantee, setGrantee] = useState('');
  const [role, setRole] = useState<string | 'REPO_ADMIN_ROLE' | 'REPO_DEV_ROLE'>('REPO_DEV_ROLE');

  const [renderLoading, setRenderLoading] = useState(false);

  const updateData = async () => {
    if (valist) {
      try {
        setRepoAdmins(await valist.getRepoAdmins(orgName, repoName));
      } catch (e) {
        console.error('Could not fetch ACL data for repo admins', e);
      }

      try {
        setRepoDevs(await valist.getRepoDevs(orgName, repoName));
      } catch (e) {
        console.error('Could not fetch ACL data for repo devs', e);
      }

      try {
        setOrgOwners(await valist.getOrgOwners(orgName));
      } catch (e) {
        console.error('Could not fetch ACL data org owners', e);
      }
    }
  };

  const grantRole = async () => {
    try {
      if (valist.web3.utils.isAddress(grantee)) {
        if (role === 'REPO_ADMIN_ROLE') {
          await valist.grantRepoAdmin(orgName, repoName, valist.defaultAccount, grantee);
        } else {
          await valist.grantRepoDev(orgName, repoName, valist.defaultAccount, grantee);
        }
        await updateData();
        setGrantee('');
      } else {
        alert('Please enter a valid Ethereum address');
      }
    } catch (e) {
      console.error('Could not grant role', e);
    }
  };

  const revokeRepoAdmin = async (address: string) => {
    try {
      await valist.revokeRepoAdmin(orgName, repoName, valist.defaultAccount, address);
      await updateData();
    } catch (e) {
      console.error('Could not revoke admin', e);
    }
  };

  const revokeRepoDev = async (address: string) => {
    try {
      await valist.revokeRepoDev(orgName, repoName, valist.defaultAccount, address);
      await updateData();
    } catch (e) {
      console.error('Could not revoke dev', e);
    }
  };

  useEffect(() => {
    updateData();
  }, [valist]);

  return (
    <div>
        <div className="col-span-3 sm:col-span-2 pb-8">
            <div className="mt-1 flex border shadow-md">
                <input onChange={(e) => setGrantee(e.target.value)} value={grantee} type="text"
                  className="form-input flex-1 block rounded-l-md w-full rounded-none
                            transition duration-150 ease-in-out sm:text-sm sm:leading-5"
                  placeholder="0x0123456789012345678901234567890123456789" />
                <div>
                    <select onChange={(e) => setRole(e.target.value)} id="role"
                      className={`form-select rounded-none block w-full
                                text-base leading-6 border-gray-300 focus:outline-none h-13`}>
                        <option value="REPO_DEV_ROLE">Developer</option>
                        <option value="REPO_ADMIN_ROLE">Admin</option>
                    </select>
                </div>
                <button value="Submit" type="button"
                  className="inline-flex items-center justify-center px-6 py-3 border border-transparent text-base
                            leading-6 font-medium text-white bg-indigo-600 hover:bg-indigo-500 focus:outline-none
                            focus:border-indigo-700 focus:shadow-outline-indigo active:bg-indigo-700
                            transition ease-in-out duration-150 rounded-r-md"
                  onClick={ async () => { setRenderLoading(true); await grantRole(); setRenderLoading(false); }}>
                  Grant Role
                </button>
            </div>
        </div>
        <ul className="grid grid-cols-1 gap-6 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4">
            {repoAdmins[0] !== '0x0' && repoAdmins.map((address) => (
              <UserAccessCard
                key={address}
                address={address}
                orgName={orgName}
                setRenderLoading={setRenderLoading}
                revokeRole={revokeRepoAdmin}
                roleType={'repoAdmin'}
              />
            ))}

            {repoDevs[0] !== '0x0' && repoDevs.map((address) => (
              <UserAccessCard
                key={address}
                address={address}
                orgName={orgName}
                setRenderLoading={setRenderLoading}
                revokeRole={revokeRepoDev}
                roleType={'repoDev'}
              />
            ))}

            {orgOwners[0] !== '0x0' && orgOwners.map((address) => (
              <UserAccessCard
                key={address}
                address={address}
                orgName={orgName}
                setRenderLoading={setRenderLoading}
                revokeRole={revokeRepoDev}
                roleType={'orgOwner'}
              />
            ))}
        </ul>
        { renderLoading && <LoadingDialog>Updating Access Control List...</LoadingDialog> }
    </div>
  );
};

export default ProjectPermissions;
