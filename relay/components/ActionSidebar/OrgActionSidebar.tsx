import React, { FunctionComponent, useState, useEffect, useContext } from 'react';
import Link from 'next/link';
import ValistContext from '../ValistContext/ValistContext';
import IsOrgAdmin from '../AccessControl/IsOrgAdmin';

export const OrgActionSidebar:FunctionComponent<any> = ({orgName}: { orgName: string }) => {
    const valist = useContext(ValistContext)
    const [profile, setProfile] = useState({ address: "0x0",  });

    useEffect(() => {
        (async function() {
            if (valist) {
                setProfile({ address: valist.defaultAccount });
            }
        })()
    }, [valist]);

    return (
            <div className="xl:flex-shrink-0 xl:w-64 xl:border-r xl:border-gray-200 bg-white">
                <div className="pl-4 pr-6 py-6 sm:pl-6 lg:pl-8 xl:pl-0">
                    <div className="flex items-center justify-between">
                        <div className="flex-1 space-y-8">
                            <div className="space-y-8 sm:space-y-0 sm:flex sm:justify-between sm:items-center xl:block xl:space-y-8">
                                <div className="flex items-center space-x-3">
                                    <div className="flex-shrink-0 h-12 w-12">
                                        <img className="h-12 w-12 rounded-full bg-black" src={`https://identicon-api.herokuapp.com/${profile.address}/32?format=png`} alt="" />
                                    </div>
                                    <div className="space-y-1">
                                        <div className="text-sm leading-5 font-medium text-gray-900">{profile.address.substring(0, 12)}...</div>
                                    </div>
                                </div>
                                <div className="flex flex-col space-y-3 sm:space-y-0 sm:space-x-3 sm:flex-row xl:flex-col xl:space-x-0 xl:space-y-3">
                                <IsOrgAdmin orgName={orgName}>
                                    <span className="rounded-md shadow-sm">
                                        <Link href={`/v/${orgName}/create`}>
                                            <button type="button" className="w-full inline-flex items-center justify-center px-4 py-2 border border-transparent text-sm leading-5 font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-500 focus:outline-none focus:border-indigo-700 focus:shadow-outline-indigo active:bg-indigo-700 transition ease-in-out duration-150">
                                                New Project
                                            </button>
                                        </Link>
                                    </span>
                                    <span className="inline-flex rounded-md shadow-sm">
                                        <Link href={`/v/${orgName}/permissions`}>
                                            <button type="button" className="w-full inline-flex items-center justify-center px-4 py-2 border border-gray-300 text-sm leading-5 font-medium rounded-md text-gray-700 bg-white hover:text-gray-500 focus:outline-none focus:border-blue-300 focus:shadow-outline-blue active:text-gray-800 active:bg-gray-50 transition ease-in-out duration-150">
                                                Manage Permissions
                                            </button>
                                        </Link>
                                    </span>
                                    <span className="inline-flex rounded-md shadow-sm">
                                        <Link href={`/v/${orgName}/edit`}>
                                            <button type="button" className="w-full inline-flex items-center justify-center px-4 py-2 border border-gray-300 text-sm leading-5 font-medium rounded-md text-gray-700 bg-white hover:text-gray-500 focus:outline-none focus:border-blue-300 focus:shadow-outline-blue active:text-gray-800 active:bg-gray-50 transition ease-in-out duration-150">
                                                Edit Organization
                                            </button>
                                        </Link>
                                    </span>
                                </IsOrgAdmin>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
    )
}

export default OrgActionSidebar;
