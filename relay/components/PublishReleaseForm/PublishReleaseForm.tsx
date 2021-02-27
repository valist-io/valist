import React, { FunctionComponent, useState, useContext } from 'react';
import { useRouter } from 'next/router';
import axios from 'axios';

import ValistContext from '../ValistContext/ValistContext';
import LoadingDialog from '../LoadingDialog/LoadingDialog';

export const PublishReleaseForm:FunctionComponent<any> = ({ orgName, repoName }: { orgName: string, repoName: string }) => {
    const valist = useContext(ValistContext);
    const router = useRouter();

    const [releaseMeta, setReleaseMeta] = useState("");
    const [projectTag, setProjectTag] = useState("");
    const [file, setFile] = useState<File | null>(null);

    const [loadingMessage, setLoadingMessage] = useState("");

    const [renderLoading, setRenderLoading] = useState(false);

    const handleUpload = async (file: File) => {
        try {
            const filename = encodeURIComponent(file.name);

            // get presigned url for uploading to bucket
            const res = await axios.post(`/api/ipfs/add/file?name=${filename}`);

            const { url, fields } = await res.data;
            const formData = new FormData();

            Object.entries({ ...fields, file }).forEach(([key, value]: [key: string, value: any]) => {
                formData.append(key, value);
            });

            const upload = await axios.post(url, formData, {
                headers: { "Content-Type": "multipart/form-data" },
                onUploadProgress: (progressEvent) => setLoadingMessage(`Uploading file to IPFS: ${Math.round((progressEvent.loaded * 100) / progressEvent.total)}%`),
            });

            const responseHash = upload.headers["etag"].replaceAll(`"`, "");
            return responseHash;

            // generate only IPFS hash of file
            // const expectedHash = await valist.addFileToIPFS(file, true);

            // if (responseHash === expectedHash) {
            //    return expectedHash;
            // } else {
            //     const msg = `Integrity check failed, response hash "${responseHash}" does not equal expected hash "${expectedHash}"!`;
            //     console.error(msg);
            //     throw new Error(msg);
            // }

        } catch (e) {
            console.error("Could not upload file", e);
            throw e;
        }
    }

    const createRelease = async () => {
        try {
            if (!file) {
                console.error("No file selected");
                return;
            }

            setLoadingMessage("Uploading file to IPFS...");

            const releaseCID = await handleUpload(file);

            setLoadingMessage("Uploading metadata JSON to IPFS...");

            const metaCID = await valist.addJSONtoIPFS(releaseMeta);

            const release = {
                tag: projectTag,
                releaseCID,
                metaCID
            };

            setLoadingMessage("Publishing Release...");

            await valist.publishRelease(orgName, repoName, release);
            router.push(`/${orgName}/${repoName}`);

        } catch (e) {
            console.error("Could not publish release", e);
            return;
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
                            <input type="file" onChange={(e) => setFile(e.target.files && e.target.files[0])} id="ReleaseData" className="form-input py-3 px-4 block w-full transition ease-in-out duration-150" />
                        </div>
                    </div>
                    <div className="sm:col-span-2">
                        <span className="w-full inline-flex rounded-md shadow-sm">
                            <button onClick={async () => { setRenderLoading(true); await createRelease(); setRenderLoading(false); }} value="Submit" type="button" className="w-full inline-flex items-center justify-center px-6 py-3 border border-transparent text-base leading-6 font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-500 focus:outline-none focus:border-indigo-700 focus:shadow-outline-indigo active:bg-indigo-700 transition ease-in-out duration-150">
                                Publish Release
                            </button>
                        </span>
                    </div>
                </form>
                </div>
            </div>
            { renderLoading && <LoadingDialog>{loadingMessage}</LoadingDialog> }
        </div>
    );
}

export default PublishReleaseForm;