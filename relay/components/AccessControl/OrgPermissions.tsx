import React, { useContext, useEffect, useState } from 'react';
import LoadingDialog from '../LoadingDialog/LoadingDialog';
import AddressIdenticon from '../Identicons/AddressIdenticon';
import ValistContext from '../Valist/ValistContext';
import IsOrgAdmin from './IsOrgAdmin';

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
    setRenderLoading(true);
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
    setRenderLoading(false);
  };

  const revokeRole = async (address: string) => {
    setRenderLoading(true);
    try {
      await valist.revokeOrgAdmin(orgName, valist.defaultAccount, address);
      await updateData();
    } catch (e) {
      console.error('Could not revoke role', e);
    }
    setRenderLoading(false);
  };

  useEffect(() => {
    updateData();
  }, [valist, orgName]);

  return (
    <div className="flex flex-col w-full">
      <IsOrgAdmin orgName={orgName}>
        <div className="col-span-3 sm:col-span-2 pb-8">
          <div className="mt-1 flex shadow-sm">
              <input onChange={(e) => setGrantee(e.target.value)} type="text" value={grantee}
              className="form-input flex-1 block rounded-l-md w-full rounded-none transition duration-150
              ease-in-out sm:text-sm sm:leading-5 shadow-sm"
              placeholder="0x0123456789012345678901234567890123456789" />
              <button value="Submit" type="button" className="inline-flex items-center justify-center
              px-6 py-3 border border-transparent text-base leading-6 font-medium text-white bg-indigo-600
              hover:bg-indigo-500 focus:outline-none focus:border-indigo-700 focus:shadow-outline-indigo
              active:bg-indigo-700 transition ease-in-out duration-150 rounded-r-md"
              onClick={async () => { await grantRole(); }}>
                Add Key
              </button>
          </div>
        </div>
      </IsOrgAdmin>
      <div className="-my-2 overflow-x-auto sm:-mx-6 lg:-mx-8">
        <div className="py-2 align-middle inline-block min-w-full sm:px-6 lg:px-8">
          <div className="shadow overflow-hidden border-b border-gray-200 sm:rounded-lg">
            <table className="min-w-full divide-y divide-gray-200">
              <thead className="bg-gray-50">
                <tr>
                  <th scope="col"
                  className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                    Address
                  </th>
                  <th scope="col"
                  className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                    Role
                  </th>
                  <th scope="col"
                  className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                    Status
                  </th>
                  <th scope="col" className="relative px-6 py-3">
                    <span className="sr-only">Action</span>
                  </th>
                </tr>
              </thead>
              <tbody>
                {orgOwners[0] !== '0x0' && orgOwners.map((address) => (
                  <tr className="bg-white" key={address}>
                    <td className="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                      <div className="flex items-center">
                        <div>
                          <AddressIdenticon address={address} height={8}/>
                        </div>
                        <div className="ml-8 font-mono">
                          {address}
                        </div>
                      </div>
                    </td>
                    <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                      Org Owner
                    </td>
                    <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                      Pending
                    </td>
                    <td className="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
                      <a href="#" className="text-indigo-600 hover:text-indigo-900"
                      onClick={async () => revokeRole(address)}>Revoke Role</a>
                    </td>
                  </tr>
                ))}
                {orgAdmins[0] !== '0x0' && orgAdmins.map((address) => (
                  <tr className="bg-white" key={address}>
                    <td className="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                      <div className="flex items-center">
                        <div>
                          <AddressIdenticon address={address} height={8}/>
                        </div>
                        <div className="ml-8 font-mono">
                          {address}
                        </div>
                      </div>
                    </td>
                    <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                      Org Admin
                    </td>
                    <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                      Pending
                    </td>
                    <td className="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
                      <a href="#" className="text-indigo-600 hover:text-indigo-900"
                      onClick={async () => revokeRole(address)}>Revoke Role</a>
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        </div>
      </div>
      { renderLoading && <LoadingDialog>Updating Access Control List...</LoadingDialog> }
    </div>
  );
};

export default OrganizationPermissions;
