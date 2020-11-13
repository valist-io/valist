import React, { useContext, useEffect, useState } from 'react';
import Link from 'next/link';
import ValistContext from '../ValistContext/ValistContext';

export const OrgList= () => {

    const [orgs, setOrgs] = useState([{ returnValues: { orgName: "Loading...", orgMeta: "Loading..." }, blockNumber: 0, transactionHash: "0x0" }]);
    const valist = useContext(ValistContext)

    useEffect(() => {
        (async function() {
            if (valist) {
                // @ts-ignore
                setOrgs(await valist.getCreatedOrganizations());
            }
        })()
    }, [valist]);

    return (
        <div className="bg-white lg:min-w-0 lg:flex-1">
            <div className="pl-4 pr-6 pt-4 pb-4 border-b border-t border-gray-200 sm:pl-6 lg:pl-8 xl:pl-6 xl:pt-6 xl:border-t-0">
            <div className="flex items-center">
                <h1 className="flex-1 text-lg leading-7 font-medium">Organizations</h1>
                <div className="relative">
                <span className="rounded-md shadow-sm">
                    <Link href="/create">
                        <button type="button" className="w-full inline-flex items-center justify-center px-4 py-2 border border-transparent text-sm leading-5 font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-500 focus:outline-none focus:border-indigo-700 focus:shadow-outline-indigo active:bg-indigo-700 transition ease-in-out duration-150">
                            Create Organization
                        </button>
                    </Link>
                </span>
                <div className={"origin-top-right z-10 absolute right-0 mt-2 w-56 rounded-md shadow-lg hidden"}>
                    <div className="rounded-md bg-white shadow-xs">
                        <div className="py-1" role="menu" aria-orientation="vertical" aria-labelledby="sort-menu">
                            <a href="#" className="block px-4 py-2 text-sm leading-5 text-gray-700 hover:bg-gray-100 hover:text-gray-900 focus:outline-none focus:bg-gray-100 focus:text-gray-900" role="menuitem">Name</a>
                            <a href="#" className="block px-4 py-2 text-sm leading-5 text-gray-700 hover:bg-gray-100 hover:text-gray-900 focus:outline-none focus:bg-gray-100 focus:text-gray-900" role="menuitem">Date modified</a>
                            <a href="#" className="block px-4 py-2 text-sm leading-5 text-gray-700 hover:bg-gray-100 hover:text-gray-900 focus:outline-none focus:bg-gray-100 focus:text-gray-900" role="menuitem">Date created</a>
                        </div>
                    </div>
                </div>
                </div>
            </div>
            </div>
            <ul className="relative z-0 divide-y divide-gray-200 border-b border-gray-200">
            {orgs.map((org: any) => {
                return (
                    <Link href={`${org.returnValues.orgName}`} key={org.transactionHash}>
                        <li className="relative pl-4 pr-6 py-5 hover:bg-gray-50 sm:py-6 sm:pl-6 lg:pl-8 xl:pl-6">
                            <div className="flex items-center justify-between space-x-4">
                            <div className="min-w-0 space-y-3">
                                <div className="flex items-center space-x-3">
                                <span aria-label="Running" className="h-4 w-4 bg-green-100 rounded-full flex items-center justify-center">
                                    <span className="h-2 w-2 bg-green-400 rounded-full"></span>
                                </span>

                                <span className="block">
                                    <h2 className="text-sm font-medium leading-5">
                                    <a href="#">
                                        <span className="absolute inset-0"></span>
                                        {org.returnValues.orgName}
                                    </a>
                                    </h2>
                                </span>
                                </div>
                                <a href="#" className="relative group flex items-center space-x-2.5">
                                <div className="text-sm leading-5 text-gray-500 group-hover:text-gray-900 font-medium truncate">
                                    {org.returnValues.orgMeta}
                                </div>
                                </a>
                            </div>
                            <div className="sm:hidden">
                                <svg className="h-5 w-5 text-gray-400" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
                                <path fillRule="evenodd" d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z" clipRule="evenodd" />
                                </svg>
                            </div>

                            <div className="hidden sm:flex flex-col flex-shrink-0 items-end space-y-3">
                                <p className="flex text-gray-500 text-sm leading-5 space-x-2">
                                <span>Block: {org.blockNumber}</span>
                                </p>
                            </div>
                            </div>
                        </li>
                    </Link>
                )
            })}
            </ul>
        </div>
    );
}

export default OrgList;
