import React, { useContext, useEffect, useState } from 'react';
import LoadingDialog from '../LoadingDialog/LoadingDialog';
import ValistContext from '../Valist/ValistContext';
import IsOrgAdmin from './IsOrgAdmin';
import UserAccessCard from './UserAccessCard';

const OrganizationPermissions = ({ orgName }: { orgName: string }) => {
  const valist = useContext(ValistContext);

  const [orgAdmins, setOrgAdmins] = useState(['0x0']);
  const [orgOwners, setOrgOwners] = useState(['0x0']);
  const [grantee, setGrantee] = useState('');

  const [renderLoading, setRenderLoading] = useState(false);

  const updateData = async () => {
    if (valist) {
      try {
        setOrgAdmins(await valist.getOrgAdmins(orgName) || ['0x0']);
        setOrgOwners(await valist.getOrgOwners(orgName) || ['0x0']);
      } catch (e) {
        console.error('Could not fetch ACL data', e);
      }
    }
  };

  const grantRole = async () => {
    try {
      if (valist.web3.utils.isAddress(grantee)) {
        await valist.grantOrgAdmin(orgName, valist.defaultAccount, grantee);
        await updateData();
        setGrantee('');
      } else {
        alert('Please enter a valid Ethereum address');
      }
    } catch (e) {
      console.error('Could not grant role', e);
    }
  };

  const revokeRole = async (address: string) => {
    try {
      await valist.revokeOrgAdmin(orgName, valist.defaultAccount, address);
      await updateData();
    } catch (e) {
      console.error('Could not revoke role', e);
    }
  };

  useEffect(() => {
    updateData();
  }, [valist]);

  return (
        <div>
          <IsOrgAdmin orgName={orgName}>
            <div className="col-span-3 sm:col-span-2 pb-8">
                <div className="mt-1 flex shadow-sm">
                    <input onChange={(e) => setGrantee(e.target.value)} type="text" value={grantee}
                    className="form-input flex-1 block rounded-l-md w-full rounded-none transition duration-150
                    ease-in-out sm:text-sm sm:leading-5 shadow-sm"
                    placeholder="0x0123456789012345678901234567890123456789" />
                    <div>
                        <select id="role" className="form-select rounded-none block w-full
                        text-base leading-6 border-gray-300 focus:outline-none h-13">
                            <option value="ORG_ADMIN_ROLE">Admin</option>
                        </select>
                    </div>
                    <button value="Submit" type="button" className="inline-flex items-center justify-center
                    px-6 py-3 border border-transparent text-base leading-6 font-medium text-white bg-indigo-600
                    hover:bg-indigo-500 focus:outline-none focus:border-indigo-700 focus:shadow-outline-indigo
                    active:bg-indigo-700 transition ease-in-out duration-150 rounded-r-md"
                    onClick={async () => { setRenderLoading(true); await grantRole(); setRenderLoading(false); }}>
                      Grant Role
                    </button>
                </div>
              </div>
            </IsOrgAdmin>
            <ul className="grid grid-cols-1 gap-6 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4">
                {orgOwners[0] !== '0x0' && orgOwners.map((address) => (
                  <UserAccessCard
                    key={address}
                    address={address}
                    orgName={orgName}
                    setRenderLoading={setRenderLoading}
                    revokeRole={revokeRole}
                    roleType={'orgOwner'}
                  />
                ))}
                {orgAdmins[0] !== '0x0' && orgAdmins.map((address) => (
                  <UserAccessCard
                    key={address}
                    address={address}
                    orgName={orgName}
                    setRenderLoading={setRenderLoading}
                    revokeRole={revokeRole}
                    roleType={'orgAdmin'}
                  />
                ))}
            </ul>
            { renderLoading && <LoadingDialog>Updating Access Control List...</LoadingDialog> }
        </div>
  );
};

export default OrganizationPermissions;
