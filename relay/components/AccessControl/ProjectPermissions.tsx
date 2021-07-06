import React, { useContext, useEffect, useState } from 'react';
import LoadingDialog from '../LoadingDialog/LoadingDialog';
import ValistContext from '../Valist/ValistContext';
import UserAccessCard from './UserAccessCard';

const ProjectPermissions = ({ orgName, repoName }: { orgName: string, repoName: string }): JSX.Element => {
  const valist = useContext(ValistContext);

  const [repoDevs, setRepoDevs] = useState(['0x0']);
  const [orgAdmins, setOrgAdmins] = useState(['0x0']);
  const [grantee, setGrantee] = useState('');

  const [renderLoading, setRenderLoading] = useState(false);

  const updateData = async () => {
    if (valist) {
      try {
        setRepoDevs(await valist.getRepoDevs(orgName, repoName));
      } catch (e) {
        console.error('Could not fetch ACL data for repo devs', e);
      }

      try {
        setOrgAdmins(await valist.getOrgAdmins(orgName));
      } catch (e) {
        console.error('Could not fetch ACL data org owners', e);
      }
    }
  };

  const addKey = async () => {
    try {
      if (valist.web3.utils.isAddress(grantee)) {
        await valist.voteRepoDev(orgName, repoName, grantee);
        await updateData();
        setGrantee('');
      } else {
        alert('Please enter a valid Ethereum address');
      }
    } catch (e) {
      console.error('Could not grant role', e);
    }
  };

  const revokeRepoDev = async (address: string) => {
    try {
      if (valist.web3.utils.isAddress(grantee)) {
        await valist.revokeRepoDev(orgName, repoName, address);
        await updateData();
      } else {
        alert('Please enter a valid Ethereum address');
      }
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
                <button value="Submit" type="button"
                  className="inline-flex items-center justify-center px-6 py-3 border border-transparent text-base
                            leading-6 font-medium text-white bg-indigo-600 hover:bg-indigo-500 focus:outline-none
                            focus:border-indigo-700 focus:shadow-outline-indigo active:bg-indigo-700
                            transition ease-in-out duration-150 rounded-r-md"
                  onClick={ async () => { setRenderLoading(true); await addKey(); setRenderLoading(false); }}>
                  Add Key
                </button>
            </div>
        </div>
        <ul className="grid grid-cols-1 gap-6 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4">
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

            {orgAdmins[0] !== '0x0' && orgAdmins.map((address) => (
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
