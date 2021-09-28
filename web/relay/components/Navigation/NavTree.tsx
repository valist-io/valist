import React from 'react';
import Link from 'next/link';

const NavTree = ({ orgName = '', repoName = '' }: { orgName?: string, repoName?: string }) => (
    <nav className="flex divide-y" aria-label="Breadcrumb">
        <ol className="flex items-center space-x-4">
            <li>
                <div>
                    <Link href="/">
                        <a className="text-gray-400 hover:text-gray-500">
                            <svg className="flex-shrink-0 h-5 w-5 transition duration-150 ease-in-out"
                                xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
                                <path d="M10.707 2.293a1 1 0 00-1.414 0l-7 7a1 1 0 001.414
                                1.414L4 10.414V17a1 1 0 001 1h2a1 1 0 001-1v-2a1 1 0 011-1h2a1 1 0 011
                                1v2a1 1 0 001 1h2a1 1 0 001-1v-6.586l.293.293a1 1 0 001.414-1.414l-7-7z" />
                            </svg>
                            <span className="sr-only">Home</span>
                        </a>
                    </Link>
                </div>
            </li>

            {orgName && <li>
                <div className="flex items-center space-x-4">
                    <svg className="flex-shrink-0 h-5 w-5 text-gray-300"
                    xmlns="http://www.w3.org/2000/svg" fill="currentColor" viewBox="0 0 20 20">
                    <path d="M5.555 17.776l8-16 .894.448-8 16-.894-.448z" />
                    </svg>
                    <Link href={`/${orgName}`}>
                        <a className="w-24 min-w-full text-sm leading-5 font-medium text-gray-500 hover:text-gray-700
                        transition duration-150 ease-in-out">{orgName}</a>
                    </Link>
                </div>
            </li>}

            {orgName && repoName && <li>
                <div className="flex items-center space-x-4">
                    <svg className="flex-shrink-0 h-5 w-5 text-gray-300"
                    xmlns="http://www.w3.org/2000/svg" fill="currentColor" viewBox="0 0 20 20">
                    <path d="M5.555 17.776l8-16 .894.448-8 16-.894-.448z" />
                    </svg>
                    <Link href={`/${orgName}/${repoName}`}>
                        <a className="text-sm leading-5 font-medium text-gray-500
                        hover:text-gray-700 transition duration-150 ease-in-out">
                            {repoName}
                        </a>
                    </Link>
                </div>
            </li>}

        </ol>
        <div className="border-b-2 border-fuchsia-600" style={{ width: '100%' }}></div>
    </nav>
);

export default NavTree;
