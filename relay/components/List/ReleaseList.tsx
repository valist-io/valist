import React, { useContext, useEffect, useState } from 'react';
import Link from 'next/link';
import ValistContext from '../ValistContext/ValistContext';
import NavTree from '../Nav/NavTree';

export const ReleaseList= ({ orgName, repoName }: { orgName: string, repoName: string }) => {
    const valist = useContext(ValistContext)

    const [releases, setReleases] = useState(["Loading..."]);

    useEffect(() => {
        (async function() {
            if (valist) {
                try {
                    setReleases(await valist.getReleaseTagsFromRepo(orgName, repoName));
                } catch (e) {}
            }
        })()
    }, [valist]);

    return (
        <div className="bg-white lg:min-w-0 lg:flex-1">
            <div className="pl-4 pr-6 pt-4 pb-4 border-b border-t border-gray-200 sm:pl-6 lg:pl-8 xl:pl-6 xl:pt-6 xl:border-t-0">
                <div className="flex items-center">
                    <NavTree orgName={orgName} repoName={repoName} />
                </div>
            </div>
            <ul className="relative z-0 divide-y divide-gray-200 border-b border-gray-200">
            {[...releases].reverse().map((tag) => {
                return (
                    <Link href={`/api/${orgName}/${repoName}/${tag}`} key={tag}>
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
                                                    {tag}
                                                </a>
                                            </h2>
                                        </span>
                                    </div>
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

export default ReleaseList;
