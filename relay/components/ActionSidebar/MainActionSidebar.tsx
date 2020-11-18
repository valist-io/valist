import React, { FunctionComponent, useState, useEffect, useContext } from 'react';
import Link from 'next/link';
import ValistContext from '../ValistContext/ValistContext';

const ProfileActionSidebar:FunctionComponent<any> = () => {
    const valist = useContext(ValistContext)
    const [profile, setProfile] = useState({ address: "0x0" });

    useEffect(() => {
        (async function() {
            if (valist) {
                let accounts = await valist.web3.eth.getAccounts() || ["0x0"];
                setProfile({ address: accounts[0] || "0x0" });
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
                                        <img className="h-12 w-12 rounded-full" src={`https://identicon-api.herokuapp.com/${profile.address}/32?format=png`} alt="" />
                                    </div>
                                    <div className="space-y-1">
                                        <div className="text-sm leading-5 font-medium text-gray-900">{profile.address.substring(0, 12)}...</div>
                                    </div>
                                </div>
                                <div className="flex flex-col space-y-3 sm:space-y-0 sm:space-x-3 sm:flex-row xl:flex-col xl:space-x-0 xl:space-y-3">
                                <span className="rounded-md shadow-sm">
                                    <Link href="/v/create">
                                        <button type="button" className="w-full inline-flex items-center justify-center px-4 py-2 border border-transparent text-sm leading-5 font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-500 focus:outline-none focus:border-indigo-700 focus:shadow-outline-indigo active:bg-indigo-700 transition ease-in-out duration-150">
                                            Create Organization
                                        </button>
                                    </Link>
                                </span>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
    )
}

export default ProfileActionSidebar;
