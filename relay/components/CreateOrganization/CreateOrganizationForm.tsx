import React, { FunctionComponent, useState, useEffect, useContext } from 'react';
import ValistContext from '../ValistContext/ValistContext';
import { useRouter } from 'next/router';

export const CreateOrganizationForm:FunctionComponent<any> = () => {
    const valist = useContext(ValistContext);
    const router = useRouter();


    const [orgShortName, setOrgShortName] = useState("");
    const [orgFullName, setOrgFullName] = useState("");
    const [orgDescription, setOrgDescription] = useState("");

    useEffect(() => {
        if (valist) {
            (async function () {
                try {
                    setOrgShortName("");
                    setOrgDescription("");
                    setOrgDescription("");
                } catch (error) {
                    alert(`Failed to load accounts.`);
                    console.log(error);
                }
            })();
        }
    }, [valist]);

    const createOrganization = async () => {
        const meta = {
            name: orgFullName,
            description: orgDescription
        };

        if (orgShortName && orgFullName && orgDescription && valist.defaultAccount) {
            await valist.createOrganization(orgShortName, meta, valist.defaultAccount);
            router.push(`/v/${orgShortName}/create`);
        } else {
            alert(`Please complete the required fields`);
        }
    }

    return (
        <div className="bg-white py-16 px-4 overflow-hidden sm:px-6 lg:px-8 lg:py-24">
            <div className="relative max-w-xl mx-auto">
                <svg className="absolute left-full transform translate-x-1/2" width="404" height="404" fill="none" viewBox="0 0 404 404">
                <defs>
                    <pattern id="85737c0e-0916-41d7-917f-596dc7edfa27" x="0" y="0" width="20" height="20" patternUnits="userSpaceOnUse">
                    <rect x="0" y="0" width="4" height="4" className="text-gray-200" fill="currentColor" />
                    </pattern>
                </defs>
                <rect width="404" height="404" fill="url(#85737c0e-0916-41d7-917f-596dc7edfa27)" />
                </svg>
                <svg className="absolute right-full bottom-0 transform -translate-x-1/2" width="404" height="404" fill="none" viewBox="0 0 404 404">
                <defs>
                    <pattern id="85737c0e-0916-41d7-917f-596dc7edfa27" x="0" y="0" width="20" height="20" patternUnits="userSpaceOnUse">
                    <rect x="0" y="0" width="4" height="4" className="text-gray-200" fill="currentColor" />
                    </pattern>
                </defs>
                <rect width="404" height="404" fill="url(#85737c0e-0916-41d7-917f-596dc7edfa27)" />
                </svg>
                <div className="text-center">
                <h2 className="text-3xl leading-9 font-extrabold tracking-tight text-gray-900 sm:text-4xl sm:leading-10">
                    Create a New Organization
                </h2>
                <p className="mt-4 text-lg leading-6 text-gray-500">
                    Create a new <b>organization</b> or <b>username</b> to begin publishing projects.
                </p>
                </div>
                <div className="mt-12">
                <form className="grid grid-cols-1 gap-y-6 sm:grid-cols-2 sm:gap-x-8">
                    <div className="col-span-3 sm:col-span-2">
                        <label htmlFor="OrgShortName" className="block text-sm font-medium leading-5 text-gray-700">
                        Shortname
                        </label>
                        <div className="mt-1 flex rounded-md shadow-sm">
                        <span className="inline-flex items-center px-3 rounded-l-md border border-r-0 border-gray-300 bg-gray-50 text-gray-500 text-sm">
                            https://app.valist.io/
                        </span>
                        <input onChange={(e) => setOrgShortName(e.target.value)} required id="OrgShortName" className="form-input flex-1 block w-full rounded-none rounded-r-md transition duration-150 ease-in-out sm:text-sm sm:leading-5" placeholder="my-organization" />
                        </div>
                        <p className="mt-2 text-sm text-gray-500">You or your organization's username</p>
                    </div>

                    <div className="sm:col-span-2">
                        <label htmlFor="OrgFullName" className="block text-sm font-medium leading-5 text-gray-700">Full Name</label>
                        <div className="mt-1 relative rounded-md shadow-sm">
                            <input onChange={(e) => setOrgFullName(e.target.value)} required id="OrgFullName" className="form-input block w-full sm:text-sm sm:leading-5 transition ease-in-out duration-150" />
                        </div>
                        <p className="mt-2 text-sm text-gray-500">Your name or your organization's name</p>
                    </div>
                    <div className="sm:col-span-2">
                        <label htmlFor="OrgDescription" className="block text-sm font-medium leading-5 text-gray-700">Description</label>
                        <div className="mt-1 relative rounded-md shadow-sm">
                            <textarea onChange={(e) => setOrgDescription(e.target.value)} required id="OrgDescription" className="h-20 form-input block w-full sm:text-sm sm:leading-5 transition ease-in-out duration-150" />
                        </div>
                    </div>
                    <div className="sm:col-span-2">
                    <span className="w-full inline-flex rounded-md shadow-sm">
                        <button onClick={createOrganization} value="Submit" type="button" className="w-full inline-flex items-center justify-center px-6 py-3 border border-transparent text-base leading-6 font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-500 focus:outline-none focus:border-indigo-700 focus:shadow-outline-indigo active:bg-indigo-700 transition ease-in-out duration-150">
                            Create Organization
                        </button>
                    </span>
                    </div>
                </form>
                </div>
            </div>
        </div>
    );
}
