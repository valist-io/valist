import React, { FunctionComponent, useState, useEffect} from 'react';

import Valist from 'valist';

export const PublishReleaseForm:FunctionComponent<any> = ({ valist, orgName, repoName }: { valist: Valist, orgName: string, repoName: string }) => {

    const [account, setAccount] = useState("");
    const [releaseMeta, setReleaseMeta] = useState("")
    const [projectTag, setProjectTag] = useState("")
    const [releaseData, setReleaseData] = useState<File | null> (null)

    useEffect(() => {
        if (valist) {
            (async function () {
                try {
                    const accounts = await valist.web3.eth.getAccounts();
                    setAccount(accounts[0]);
                    console.log(releaseData);
                } catch (error) {
                    alert(`Failed to load accounts.`);
                    console.log(error);
                }
            })();
        }
    }, [valist]);

    const readUploadedFileAsBuffer = (inputFile: any) => {
        const temporaryFileReader = new FileReader();

        return new Promise((resolve, reject) => {
            temporaryFileReader.onerror = () => {
                temporaryFileReader.abort();
                reject(new DOMException("Problem parsing input file."));
            };

            temporaryFileReader.onload = () => {
                resolve(temporaryFileReader.result);
            };

            // @ts-ignore
            temporaryFileReader.readAsArrayBuffer(inputFile);
        });
    };

    const handleUpload = async (file: any) => {
        console.log("Incoming File:", file)
        try {
            const fileContents = await readUploadedFileAsBuffer(file)
            return await valist.addFileToIPFS(fileContents)
        } catch (e) {
            console.warn(e.message)
        }
    }

    const createRelease = async () => {
        const releaseHash = await handleUpload(releaseData);
        const meta = await valist.addJSONtoIPFS(releaseMeta);
        const release = {
            tag: projectTag,
            hash: releaseHash,
            meta
        };

        await valist.publishRelease(orgName, repoName, release, account);
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
                    Publish a New Release
                </h2>
                </div>
                <div className="mt-12">
                <form className="grid grid-cols-1 gap-y-6 sm:grid-cols-2 sm:gap-x-8">
                    <div className="sm:col-span-2">
                        <label htmlFor="ProjectMeta" className="block text-sm font-medium leading-5 text-gray-700">Metadata</label>
                        <div className="mt-1 relative rounded-md shadow-sm">
                            <input onChange={(e) => setReleaseMeta(e.target.value)} id="ProjectMeta" className="form-input py-3 px-4 block w-full transition ease-in-out duration-150" />
                        </div>
                    </div>
                    <div className="sm:col-span-2">
                        <label htmlFor="ReleaseTag" className="block text-sm font-medium leading-5 text-gray-700">Release Tag</label>
                        <div className="mt-1 relative rounded-md shadow-sm">
                            <input onChange={(e) => setProjectTag(e.target.value)} id="ReleaseTag" className="form-input py-3 px-4 block w-full transition ease-in-out duration-150" />
                        </div>
                    </div>
                    <div className="sm:col-span-2">
                        <label htmlFor="ReleaseData" className="block text-sm font-medium leading-5 text-gray-700">Release Data</label>
                        <div className="mt-1 relative rounded-md shadow-sm">
                            <input type="file" onChange={(e) => setReleaseData(e.target.files && e.target.files[0])} id="ReleaseData" className="form-input py-3 px-4 block w-full transition ease-in-out duration-150" />
                        </div>
                    </div>
                    <div className="sm:col-span-2">
                        <span className="w-full inline-flex rounded-md shadow-sm">
                            <button onClick={createRelease} value="Submit" type="button" className="w-full inline-flex items-center justify-center px-6 py-3 border border-transparent text-base leading-6 font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-500 focus:outline-none focus:border-indigo-700 focus:shadow-outline-indigo active:bg-indigo-700 transition ease-in-out duration-150">
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

export default PublishReleaseForm;

/*

    const reducer = (state: any, action: any) => {
        switch (action.type) {
            case 'SET_DROP_DEPTH':
                return { ...state, dropDepth: action.dropDepth }
            case 'SET_IN_DROP_ZONE':
                return { ...state, inDropZone: action.inDropZone };
            case 'ADD_FILE_TO_LIST':
                return { ...state, fileList: state.fileList.concat(action.files) };
            default:
                return state;
        }
    };

    const [data, dispatch] = React.useReducer(
        reducer, { dropDepth: 0, inDropZone: false, fileList: [] }
    )

                    <DragAndDrop data={data} dispatch={dispatch}/>
                    <ol className="dropped-files">
                        {data.fileList.map((f: any) => {
                            return (
                                <li key={f.name}>{f.name}</li>
                            )
                        })}
                    </ol>
*/
