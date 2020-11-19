import React, { useState, useEffect, useContext} from 'react';
import ValistContext from '../ValistContext/ValistContext';

const EditRepoForm = ({ orgName, repoName }: { orgName: string, repoName: string}) => {
    const valist = useContext(ValistContext);
    const [projectHomepage, setProjectHomepage] = useState("")
    const [projectRepository, setProjectRepository] = useState("")
    const [projectDescription, setProjectDescription] = useState("")

    const getCurrentMeta = async () =>{
        if (valist) {
            try {
                const repoMeta = await valist.getRepoMeta(orgName, repoName);
                setProjectHomepage(repoMeta['homepage']);
                setProjectRepository(repoMeta['repository']);
                setProjectDescription(repoMeta['description']);
            } catch (e) {}
        }
    }

    const updateRepoMeta = async () =>{
        const meta = {
            homepage: projectHomepage,
            repository: projectRepository,
            description: projectDescription
        };

        try {
            await valist.setRepoMeta(orgName, repoName, meta, valist.defaultAccount);
        } catch(e) {}

        await getCurrentMeta();
    }

    useEffect(() => {
        getCurrentMeta();
    }, [valist]);

    return (
    <div>
        <div className="bg-white shadow px-4 py-5 sm:rounded-lg sm:p-6">
            <div className="md:grid md:grid-cols-3 md:gap-6">
                <div className="md:col-span-1">
                    <h3 className="text-lg font-medium leading-6 text-gray-900">{orgName}/{repoName} Metadata</h3>
                    <p className="mt-1 text-sm leading-5 text-gray-500">
                        This information will be displayed publicly so be careful what you share.
                    </p>
                </div>
                <div className="mt-5 md:mt-0 md:col-span-2">
                    <form className="grid grid-cols-1 gap-y-6 sm:grid-cols-2 sm:gap-x-8">
                        <div className="sm:col-span-2">
                            <label htmlFor="ProjectHomepage" className="block text-sm font-medium leading-5 text-gray-700">Homepage</label>
                            <div className="mt-1 relative rounded-md shadow-sm">
                                <input value={projectHomepage} onChange={(e) => setProjectHomepage(e.target.value)} id="ProjectHomepage" className="form-input block w-full sm:text-sm sm:leading-5 transition ease-in-out duration-150" />
                            </div>
                        </div>
                        <div className="sm:col-span-2">
                            <label htmlFor="ProjectRepository" className="block text-sm font-medium leading-5 text-gray-700">Repository</label>
                            <div className="mt-1 relative rounded-md shadow-sm">
                                <input value={projectRepository} onChange={(e) => setProjectRepository(e.target.value)} id="ProjectRepository" className="form-input block w-full sm:text-sm sm:leading-5 transition ease-in-out duration-150" />
                            </div>
                        </div>
                        <div className="sm:col-span-2">
                            <label htmlFor="ProjectDescription" className="block text-sm font-medium leading-5 text-gray-700">Description</label>
                            <div className="mt-1 relative rounded-md shadow-sm">
                                <textarea value={projectDescription} onChange={(e) => setProjectDescription(e.target.value)} id="ProjectDescription" className="h-20 form-input block w-full sm:text-sm sm:leading-5 transition ease-in-out duration-150" />
                            </div>
                        </div>
                        <div className="sm:col-span-2">
                        <span className="w-full inline-flex rounded-md shadow-sm">
                            <button onClick={updateRepoMeta} value="Submit" type="button" className="w-full inline-flex items-center justify-center px-6 py-3 border border-transparent text-base leading-6 font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-500 focus:outline-none focus:border-indigo-700 focus:shadow-outline-indigo active:bg-indigo-700 transition ease-in-out duration-150">
                                Update Project Meta
                            </button>
                        </span>
                        </div>
                    </form>
                </div>
            </div>
        </div>
    </div>
    );
}

export default EditRepoForm;
