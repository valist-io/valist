import React, { FunctionComponent, useState, useEffect, useContext } from 'react';
import Link from 'next/link';
import ValistContext from '../ValistContext/ValistContext';

const ProfileActionSidebar:FunctionComponent<any> = () => {
    const valist = useContext(ValistContext)
    const [profile, setProfile] = useState({ address: "0x0",  });

    useEffect(() => {
        (async function() {
            if (valist) {
                let accounts = await valist.web3.eth.getAccounts() || ["0x0"];
                setProfile({ address: accounts[0] });
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
                                        {/*
                                        <a href="#" className="group flex items-center space-x-2.5">

                                        <svg className="w-5 h-5 text-gray-400 group-hover:text-gray-500" viewBox="0 0 18 18" fill="none" xmlns="http://www.w3.org/2000/svg">
                                            <path fillRule="evenodd" clipRule="evenodd" d="M8.99917 0C4.02996 0 0 4.02545 0 8.99143C0 12.9639 2.57853 16.3336 6.15489 17.5225C6.60518 17.6053 6.76927 17.3277 6.76927 17.0892C6.76927 16.8762 6.76153 16.3104 6.75711 15.5603C4.25372 16.1034 3.72553 14.3548 3.72553 14.3548C3.31612 13.316 2.72605 13.0395 2.72605 13.0395C1.9089 12.482 2.78793 12.4931 2.78793 12.4931C3.69127 12.5565 4.16643 13.4198 4.16643 13.4198C4.96921 14.7936 6.27312 14.3968 6.78584 14.1666C6.86761 13.5859 7.10022 13.1896 7.35713 12.965C5.35873 12.7381 3.25756 11.9665 3.25756 8.52116C3.25756 7.53978 3.6084 6.73667 4.18411 6.10854C4.09129 5.88114 3.78244 4.96654 4.27251 3.72904C4.27251 3.72904 5.02778 3.48728 6.74717 4.65082C7.46487 4.45101 8.23506 4.35165 9.00028 4.34779C9.76494 4.35165 10.5346 4.45101 11.2534 4.65082C12.9717 3.48728 13.7258 3.72904 13.7258 3.72904C14.217 4.96654 13.9082 5.88114 13.8159 6.10854C14.3927 6.73667 14.7408 7.53978 14.7408 8.52116C14.7408 11.9753 12.6363 12.7354 10.6318 12.9578C10.9545 13.2355 11.2423 13.7841 11.2423 14.6231C11.2423 15.8247 11.2313 16.7945 11.2313 17.0892C11.2313 17.3299 11.3937 17.6097 11.8501 17.522C15.4237 16.3303 18 12.9628 18 8.99143C18 4.02545 13.97 0 8.99917 0Z" fill="currentcolor" />
                                        </svg>
                                        <div className="text-sm leading-5 text-gray-500 group-hover:text-gray-900 font-medium">awantoch</div>
                                        </a>
                                        */}
                                    </div>
                                </div>
                                <div className="flex flex-col space-y-3 sm:space-y-0 sm:space-x-3 sm:flex-row xl:flex-col xl:space-x-0 xl:space-y-3">
                                <span className="inline-flex rounded-md shadow-sm">
                                    <Link href="/settings">
                                        <button type="button" className="w-full inline-flex items-center justify-center px-4 py-2 border border-transparent text-sm leading-5 font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-500 focus:outline-none focus:border-indigo-700 focus:shadow-outline-indigo active:bg-indigo-700 transition ease-in-out duration-150">
                                            Account Settings
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
