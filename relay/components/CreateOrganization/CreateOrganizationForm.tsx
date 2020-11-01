import React, { FunctionComponent, useState, useEffect, useContext } from 'react';
import ValistContext from '../ValistContext/ValistContext';

export const CreateOrganizationForm:FunctionComponent<any> = () => {
    const valist = useContext(ValistContext)
    const [account, setAccount] = useState("");

    const [orgShortName, setOrgShortName] = useState("")
    const [orgFullName, setOrgFullName] = useState("")
    const [orgDescription, setOrgDescription] = useState("")

    useEffect(() => {
        if (valist) {
            (async function () {
                try {
                    const accounts = await valist.web3.eth.getAccounts();
                    setAccount(accounts[0]);
                    setOrgShortName("")
                    setOrgDescription("")
                    setOrgDescription("")
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

        await valist.createOrganization(orgShortName, meta, account);
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
                    Create a new <b>organization</b> to begin creating new projects.
                </p>
                </div>
                <div className="mt-12">
                <form className="grid grid-cols-1 gap-y-6 sm:grid-cols-2 sm:gap-x-8">
                    <div className="sm:col-span-2">
                        <label htmlFor="company" className="block text-sm font-medium leading-5 text-gray-700">Organization Short Name</label>
                        <div className="mt-1 relative rounded-md shadow-sm">
                            <input onChange={(e) => setOrgShortName(e.target.value)} id="company" className="form-input py-3 px-4 block w-full transition ease-in-out duration-150" />
                        </div>
                    </div>
                    <div className="sm:col-span-2">
                        <label htmlFor="OrgFullName" className="block text-sm font-medium leading-5 text-gray-700">Organization Full Name</label>
                        <div className="mt-1 relative rounded-md shadow-sm">
                            <input onChange={(e) => setOrgFullName(e.target.value)} id="OrgFullName" className="form-input py-3 px-4 block w-full transition ease-in-out duration-150" />
                        </div>
                    </div>
                    <div className="sm:col-span-2">
                        <label htmlFor="OrgDescription" className="block text-sm font-medium leading-5 text-gray-700">Organization Description</label>
                        <div className="mt-1 relative rounded-md shadow-sm">
                            <input onChange={(e) => setOrgDescription(e.target.value)} id="OrgDescription" className="form-input py-3 px-4 block w-full transition ease-in-out duration-150" />
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
