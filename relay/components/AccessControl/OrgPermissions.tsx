import React, { useContext, useState } from 'react';
import AddressIdenticon from '../Identicons/AddressIdenticon';
import ValistContext from '../Valist/ValistContext';
import IsOrgAdmin from './IsOrgAdmin';

const OrganizationPermissions = ({ orgName, orgAdmins }: { orgName: string, orgAdmins: string[] }) => {
  const valist = useContext(ValistContext);
  const [key, setKey] = useState('');

  const grantRole = async () => {
    try {
      if (valist.web3.utils.isAddress(key)) {
        await valist.voteOrgAdmin(orgName, key);
        setKey('');
      } else {
        alert('Please enter a valid Ethereum address');
      }
    } catch (e) {
      console.error('Could not grant role', e);
    }
  };

  const revokeRole = async (address: string) => {
    try {
      await valist.revokeOrgAdmin(orgName, address);
    } catch (e) {
      console.error('Could not revoke role', e);
    }
  };

  return (
    <div className="flex flex-col w-full">
      <IsOrgAdmin orgName={orgName}>
        <div className="col-span-3 sm:col-span-2 pb-8">
          <div className="mt-1 flex shadow-sm">
              <input onChange={(e) => setKey(e.target.value)} type="text" value={key}
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
                {orgAdmins && orgAdmins[0] !== '0x0' && orgAdmins.map((address) => (
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
                      Active
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
    </div>
  );
};

export default OrganizationPermissions;
