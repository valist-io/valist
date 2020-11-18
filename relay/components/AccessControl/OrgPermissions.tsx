import React from 'react';

const OrganizationPermissions = () => {
    return (
        <ul className="grid grid-cols-1 gap-6 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4">
            <li className="col-span-1 flex flex-col text-center bg-white rounded-lg shadow">
                <div className="flex-1 flex flex-col p-8">
                <img className="w-32 h-32 flex-shrink-0 mx-auto bg-black rounded-full" src="https://images.unsplash.com/photo-1494790108377-be9c29b29330?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=4&w=256&h=256&q=60" alt="" />
                <h3 className="mt-6 text-gray-900 text-sm leading-5 font-medium">Jane Cooper</h3>
                <dl className="mt-1 flex-grow flex flex-col justify-between">
                    <dt className="sr-only">Title</dt>
                    <dd className="text-gray-500 text-sm leading-5">Paradigm Representative</dd>
                    <dt className="sr-only">Role</dt>
                    <dd className="mt-3">
                    <span className="px-2 py-1 text-teal-800 text-xs leading-4 font-medium bg-teal-100 rounded-full">Admin</span>
                    </dd>
                </dl>
                </div>
            </li>
            </ul>
    )
}

export default OrganizationPermissions;
