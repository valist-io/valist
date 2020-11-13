import React, { FunctionComponent, useState, useEffect } from 'react';
import Valist from 'valist';

export const CreateRepoForm:FunctionComponent<any> = ({valist, orgName}: {valist: Valist, orgName: string}) => {

    const [account, setAccount] = useState("");
    const [repoName, setRepoName] = useState("")
    const [projectType, setProjectType] = useState("")
    const [projectHomepage, setProjectHomepage] = useState("")
    const [projectGithub, setProjectGithub] = useState("")
    const [repoDescription, setRepoDescription] = useState("")

    useEffect(() => {
        if (valist) {
            (async function () {
                try {
                    const accounts = await valist.web3.eth.getAccounts();
                    setAccount(accounts[0]);
                    setRepoName("")
                    setRepoDescription("")
                } catch (error) {
                    alert(`Failed to load accounts.`);
                    console.log(error);
                }
            })();
        }
    }, [valist]);

    const createRepo = async () => {
        const repoMeta = {
            name: repoName,
            description: repoDescription,
            projectType: projectType,
            homepage: projectHomepage,
            github: projectGithub
        };

        await valist.createRepository(orgName, repoName, repoMeta, account)
    }

    const renderPackageMeta = (package_type: any) =>{
        if(package_type == "npm"){
            return (
                <div className="sm:col-span-2">
                    <div className="sm:col-span-2">
                        <label htmlFor="ProjectHomepage" className="block text-sm font-medium leading-5 text-gray-700">Project Homepage</label>
                        <div className="mt-1 relative rounded-md shadow-sm">
                            <input onChange={(e) => setProjectHomepage(e.target.value)} id="ProjectHomepage" className="form-input py-3 px-4 block w-full transition ease-in-out duration-150" />
                        </div>
                    </div>
                    <div className="sm:col-span-2">
                        <label htmlFor="ProjectGithub" className="block text-sm font-medium leading-5 text-gray-700">Project Github Repo</label>
                        <div className="mt-1 relative rounded-md shadow-sm">
                            <input onChange={(e) => setProjectGithub(e.target.value)} id="ProjectGithub" className="form-input py-3 px-4 block w-full transition ease-in-out duration-150" />
                        </div>
                    </div>
                </div>
            )
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
                    Create a New Project
                </h2>
                <p className="mt-4 text-lg leading-6 text-gray-500">
                    Create a new <b>project</b> to begin publishing new releases.
                </p>
                </div>
                <div className="mt-12">
                <form className="grid grid-cols-1 gap-y-6 sm:grid-cols-2 sm:gap-x-8">
                    <div className="sm:col-span-2">
                        <label htmlFor="RepoName" className="block text-sm font-medium leading-5 text-gray-700">Project Name</label>
                        <div className="mt-1 relative rounded-md shadow-sm">
                            <input onChange={(e) => setRepoName(e.target.value)} id="RepoName" className="form-input py-3 px-4 block w-full transition ease-in-out duration-150" />
                        </div>
                    </div>
                    <div className="sm:col-span-2">
                        <label htmlFor="RepoDescription" className="block text-sm font-medium leading-5 text-gray-700">Repo Description</label>
                        <div className="mt-1 relative rounded-md shadow-sm">
                            <input onChange={(e) => setRepoDescription(e.target.value)} id="RepoDescription" className="form-input py-3 px-4 block w-full transition ease-in-out duration-150" />
                        </div>
                    </div>
                    <div>
                        <label htmlFor="projectType" className="block text-sm leading-5 font-medium text-gray-700">Package Type</label>
                        <select onChange={(e) => setProjectType(e.target.value)} id="projectType" className="mt-1 form-select block w-full pl-3 pr-10 py-2 text-base leading-6 border-gray-300 focus:outline-none focus:shadow-outline-blue focus:border-blue-300 sm:text-sm sm:leading-5">
                            <option selected>binary</option>
                            <option >npm</option>
                            <option>pip</option>
                        </select>
                    </div>
                    {renderPackageMeta(projectType)}
                    <div className="sm:col-span-2">
                    <span className="w-full inline-flex rounded-md shadow-sm">
                        <button onClick={createRepo} value="Submit" type="button" className="w-full inline-flex items-center justify-center px-6 py-3 border border-transparent text-base leading-6 font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-500 focus:outline-none focus:border-indigo-700 focus:shadow-outline-indigo active:bg-indigo-700 transition ease-in-out duration-150">
                            Create Project
                        </button>
                    </span>
                    </div>
                </form>
                </div>
            </div>
        </div>
    );
}
