import React, { useContext, useEffect, useState } from 'react';
import ValistContext from '../ValistContext/ValistContext';

const ProjectPermissions = ({ orgName, repoName }: { orgName: string, repoName: string }) => {
    const valist = useContext(ValistContext);

    const [repoAdmins, setRepoAdmins] = useState(["0x0"]);
    const [repoDevs, setRepoDevs] = useState(["0x0"]);

    const updateData = async () => {
        if (valist) {
            try {
                setRepoAdmins(await valist.getRepoAdmins(orgName, repoName));
                setRepoDevs(await valist.getRepoDevs(orgName, repoName));
            } catch (e) {}
        }
    }

    useEffect(() => {
        updateData();
    }, [valist]);

    return (
        <ul className="grid grid-cols-1 gap-6 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4">
            {repoAdmins[0] !== "0x0" && repoAdmins.map((address) => (
                <li key={address} className="col-span-1 flex flex-col text-center bg-white rounded-lg shadow">
                    <div className="flex-1 flex flex-col p-8">
                        <img className="w-32 h-32 flex-shrink-0 mx-auto bg-black rounded-full" src={`https://identicon-api.herokuapp.com/${address}/128?format=png`} alt="" />
                        <h3 className="mt-6 text-gray-900 text-sm leading-5 font-medium">{address}</h3>
                        <dl className="mt-1 flex-grow flex flex-col justify-between">
                                <dd className="mt-3">
                                <span className="px-2 py-1 text-xs leading-4 font-medium bg-teal-100 rounded-full">Admin</span>
                            </dd>
                        </dl>
                    </div>
                    <div className="border-t border-gray-200">
                        <div className="-mt-px flex">
                            <div className="w-0 flex-1 flex border-r border-gray-200">
                            <a href="#" onClick={async () => {await valist.revokeRepoAdmin(orgName, repoName, valist.defaultAccount, address); await updateData()}} className="relative -mr-px w-0 flex-1 inline-flex items-center justify-center py-4 text-sm leading-5 text-gray-700 font-medium border border-transparent rounded-bl-lg hover:text-gray-500 focus:outline-none focus:shadow-outline-blue focus:border-blue-300 focus:z-10 transition ease-in-out duration-150">
                                <svg className="w-5 h-5 text-gray-400" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M18.364 18.364A9 9 0 005.636 5.636m12.728 12.728A9 9 0 015.636 5.636m12.728 12.728L5.636 5.636" />
                                </svg>
                                <span className="ml-3">Revoke Admin Role</span>
                            </a>
                            </div>
                        </div>
                    </div>
                </li>
            ))}

            {repoDevs[0] !== "0x0" && repoDevs.map((address) => (
                <li key={address} className="col-span-1 flex flex-col text-center bg-white rounded-lg shadow">
                    <div className="flex-1 flex flex-col p-8">
                        <img className="w-32 h-32 flex-shrink-0 mx-auto bg-black rounded-full" src={`https://identicon-api.herokuapp.com/${address}/128?format=png`} alt="" />
                        <h3 className="mt-6 text-gray-900 text-sm leading-5 font-medium">{address}</h3>
                        <dl className="mt-1 flex-grow flex flex-col justify-between">
                                <dd className="mt-3">
                                <span className="px-2 py-1 text-xs leading-4 font-medium bg-orange-300 rounded-full">Developer</span>
                            </dd>
                        </dl>
                    </div>
                    <div className="border-t border-gray-200">
                        <div className="-mt-px flex">
                            <div className="w-0 flex-1 flex border-r border-gray-200">
                            <a href="#" onClick={async () => {await valist.revokeRepoDev(orgName, repoName, valist.defaultAccount, address); await updateData()}} className="relative -mr-px w-0 flex-1 inline-flex items-center justify-center py-4 text-sm leading-5 text-gray-700 font-medium border border-transparent rounded-bl-lg hover:text-gray-500 focus:outline-none focus:shadow-outline-blue focus:border-blue-300 focus:z-10 transition ease-in-out duration-150">
                                <svg className="w-5 h-5 text-gray-400" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M18.364 18.364A9 9 0 005.636 5.636m12.728 12.728A9 9 0 015.636 5.636m12.728 12.728L5.636 5.636" />
                                </svg>
                                <span className="ml-3">Revoke Developer Role</span>
                            </a>
                            </div>
                        </div>
                    </div>
                </li>
            ))}
        </ul>
    )
}

export default ProjectPermissions;
