import React, { FunctionComponent, useState, useEffect, useRef } from 'react';

export const PublishReleaseForm:FunctionComponent<any> = ({valist}) => {

    const [account, setAccount] = useState("");
    const [orgName, setOrgName] = useState("")
    const [projectName, setProjectName] = useState("")
    const [projectMeta, setProjectMeta] = useState("")
    const refInput = useRef();

    useEffect(() => {
        if (valist) {
            (async function () {
                try {
                    const accounts = await valist.web3.eth.getAccounts();
                    setAccount(accounts[0]);
                    setOrgName("")
                    setProjectName("")
                    setProjectMeta("")
                } catch (error) {
                    alert(`Failed to load accounts.`);
                    console.log(error);
                }
            })();
        }
    }, [valist]);

    const createRepo = async () => {
        const releaseBody = {
            name: projectName,
            description: projectMeta
        };

        await valist.publishRelease(orgName, projectName, releaseBody, account)
    }


    const handleDragEnter = (e:any) => {
        e.preventDefault();
        e.stopPropagation();
    };
    const handleDragLeave = (e: any) => {
        e.preventDefault();
        e.stopPropagation();
    };
    const handleDragOver = (e: any) => {
        e.preventDefault();
        e.stopPropagation();
    };
    const handleDrop = (e: any) => {
        e.preventDefault();
        e.stopPropagation();
    };


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
                    Publish a New Release
                </h2>
                </div>
                <div className="mt-12">
                <form className="grid grid-cols-1 gap-y-6 sm:grid-cols-2 sm:gap-x-8">
                    <div className="sm:col-span-2">
                        <label htmlFor="OrgName" className="block text-sm font-medium leading-5 text-gray-700">Organization Name</label>
                        <div className="mt-1 relative rounded-md shadow-sm">
                            <input onChange={(e) => setOrgName(e.target.value)} id="OrgName" className="form-input py-3 px-4 block w-full transition ease-in-out duration-150" />
                        </div>
                    </div>
                    <div className="sm:col-span-2">
                        <label htmlFor="RepoName" className="block text-sm font-medium leading-5 text-gray-700">Project Name</label>
                        <div className="mt-1 relative rounded-md shadow-sm">
                            <input onChange={(e) => setProjectName(e.target.value)} id="RepoName" className="form-input py-3 px-4 block w-full transition ease-in-out duration-150" />
                        </div>
                    </div>
                    <div className="sm:col-span-2">
                        <label htmlFor="ProjectMeta" className="block text-sm font-medium leading-5 text-gray-700">Project Meta</label>
                        <div className="mt-1 relative rounded-md shadow-sm">
                            <input onChange={(e) => setProjectMeta(e.target.value)} id="ProjectMeta" className="form-input py-3 px-4 block w-full transition ease-in-out duration-150" />
                        </div>
                    </div>
                    <div className="sm:col-span-2">
                    <div className="border-4 border-dashed border-gray-200 rounded-lg h-96 relative h-32 ">
                        <div onDrop={e => handleDrop(e)}
                            onDragOver={e => handleDragOver(e)}
                            onDragEnter={e => handleDragEnter(e)}
                            onDragLeave={e => handleDragLeave(e)}
                            className="absolute inset-0 flex items-center justify-center">
                            Drop Files Here
                        </div> 
                    </div>
                    </div>
                    <div className="sm:col-span-2">
                        <span className="w-full inline-flex rounded-md shadow-sm">
                            <button onClick={createRepo} value="Submit" type="button" className="w-full inline-flex items-center justify-center px-6 py-3 border border-transparent text-base leading-6 font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-500 focus:outline-none focus:border-indigo-700 focus:shadow-outline-indigo active:bg-indigo-700 transition ease-in-out duration-150">
                                Publish Release
                            </button>
                        </span>
                    </div>
                </form>
                </div>
            </div>
        </div>
    );
}
