import React, { useContext, useEffect, useState } from 'react';
import LoadingDialog from '../LoadingDialog/LoadingDialog';
import ValistContext from '../Valist/ValistContext';
import AddressIdenticon from '../Identicons/AddressIdenticon';
import IsRepoAdmin from './IsRepoAdmin';

const ProjectPermissions = ({ orgName, repoName }: { orgName: string, repoName: string }): JSX.Element => {
  const valist = useContext(ValistContext);

  const [repoAdmins, setRepoAdmins] = useState(['0x0']);
  const [repoDevs, setRepoDevs] = useState(['0x0']);
  const [grantee, setGrantee] = useState('');
  const [role, setRole] = useState<string | 'REPO_ADMIN_ROLE' | 'REPO_DEV_ROLE'>('REPO_DEV_ROLE');

  const [renderLoading, setRenderLoading] = useState(false);

  const updateData = async () => {
    if (valist) {
      try {
        setRepoAdmins(await valist.getRepoAdmins(orgName, repoName));
        setRepoDevs(await valist.getRepoDevs(orgName, repoName));
      } catch (e) {
        console.error('Could not fetch ACL data', e);
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

  const revokeAdmin = async (address: string) => {
    try {
      await valist.revokeRepoAdmin(orgName, repoName, valist.defaultAccount, address);
      await updateData();
    } catch (e) {
      console.error('Could not revoke admin', e);
    }
  };

  const revokeDev = async (address: string) => {
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
                    <li key={address}
                      className="col-span-1 flex flex-col text-center bg-white rounded-lg border shadow-lg">
                        <div className="flex-1 flex flex-col p-8">
                          <AddressIdenticon address={address} height={32}/>
                          <h3 className="mt-6 text-gray-900 text-sm leading-5 font-medium break-words">{address}</h3>
                          <dl className="mt-1 flex-grow flex flex-col justify-between">
                                  <dd className="mt-3">
                                  <span className="px-2 py-1 text-xs leading-4 font-medium bg-teal-100 rounded-full">
                                    Admin
                                    </span>
                              </dd>
                          </dl>
                        </div>
                        <IsRepoAdmin orgName={orgName} repoName={repoName}>
                          <div className="border-t border-gray-200">
                            <div className="-mt-px flex">
                                <div className="w-0 flex-1 flex border-r border-gray-200">
                                <a href="#" onClick={ async () => {
                                  setRenderLoading(true);
                                  await revokeAdmin(address);
                                  setRenderLoading(false);
                                }}
                                className="relative -mr-px w-0 flex-1 inline-flex items-center justify-center py-4
                                  text-sm leading-5 text-gray-700 font-medium border border-transparent rounded-bl-lg
                                  hover:text-gray-500 focus:outline-none focus:shadow-outline-blue focus:border-blue-300
                                  focus:z-10 transition ease-in-out duration-150">
                                    <svg className="w-5 h-5 text-gray-400"
                                      xmlns="http://www.w3.org/2000/svg" fill="none"
                                        viewBox="0 0 24 24" stroke="currentColor">
                                        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2}
                                          d="M18.364 18.364A9 9 0 005.636 5.636m12.728 12.728A9 9
                                          0 015.636 5.636m12.728 12.728L5.636 5.636" />
                                    </svg>
                                    <span className="ml-3">Revoke Admin Role</span>
                                </a>
                                </div>
                            </div>
                          </div>
                        </IsRepoAdmin>
                    </li>
                ))}

                {repoDevs[0] !== '0x0' && repoDevs.map((address) => (
                    <li key={address} className="col-span-1 flex flex-col text-center border rounded-lg shadow-lg">
                        <div className="flex-1 flex flex-col p-8">
                          <AddressIdenticon address={address} height={32}/>
                          <h3 className="mt-6 text-gray-900 text-sm leading-5 font-medium break-words">{address}</h3>
                          <dl className="mt-1 flex-grow flex flex-col justify-between">
                                  <dd className="mt-3">
                                  <span
                                    className="px-2 py-1 text-xs leading-4 font-medium bg-orange-300 rounded-full">
                                      Developer
                                  </span>
                              </dd>
                          </dl>
                        </div>
                        <IsRepoAdmin orgName={orgName} repoName={repoName}>
                          <div className="border-t border-gray-200">
                              <div className="-mt-px flex">
                                  <div className="w-0 flex-1 flex border-r border-gray-200">
                                  <a href="#" onClick={ async () => {
                                    setRenderLoading(true);
                                    await revokeDev(address);
                                    setRenderLoading(false);
                                  }}
                                    className="relative -mr-px w-0 flex-1 inline-flex items-center justify-center py-4
                                      text-sm leading-5 text-gray-700 font-medium border border-transparent
                                      rounded-bl-lg hover:text-gray-500 focus:outline-none focus:shadow-outline-blue
                                      focus:border-blue-300 focus:z-10 transition ease-in-out duration-150">
                                      <svg className="w-5 h-5 text-gray-400"
                                        xmlns="http://www.w3.org/2000/svg" fill="none"
                                        viewBox="0 0 24 24" stroke="currentColor">
                                          <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2}
                                          d="M18.364 18.364A9 9 0 005.636 5.636m12.728 12.728A9 9
                                            0 015.636 5.636m12.728 12.728L5.636 5.636" />
                                      </svg>
                                      <span className="ml-3">Revoke Developer Role</span>
                                  </a>
                                  </div>
                              </div>
                        </div>
                      </IsRepoAdmin>
                    </li>
                ))}
            </ul>
            { renderLoading && <LoadingDialog>Updating Access Control List...</LoadingDialog> }
        </div>
  );
};

export default ProjectPermissions;
