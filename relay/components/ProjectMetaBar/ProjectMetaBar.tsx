import React, { useState, useEffect, useContext } from 'react';
import ValistContext from '../ValistContext/ValistContext';

export const ProjectMetaBar = ({ orgName, repoName }: { orgName: string, repoName: string }) => {
    const valist = useContext(ValistContext);
    const [meta, setMeta] = useState<JSX.Element>();

    const PipMeta = () => {
        return (
            <div>
                <div className="pl-6 lg:w-80">
                    <div className="pt-6 pb-2">
                        <h1 className="flex-1 text-lg leading-7 font-medium">Pip Install From Url</h1>
                    </div>
                    <div className="border-2 border-solid border-black-200 rounded-lg p-2">
                        pip install {window.location.origin}/api/{orgName}/{repoName}/latest
                    </div>
                </div>
            </div>
        );
    }

    const NpmMeta = () => {
        return (
            <div>
                <div className="pl-6 lg:w-80">
                    <div className="pt-6 pb-2">
                        <h1 className="flex-1 text-lg leading-7 font-medium">NPM Direct Install From Url</h1>
                    </div>
                    <div className="border-2 border-solid border-black-200 rounded-lg p-2">
                        npm install {window.location.origin}/api/{orgName}/{repoName}/latest
                    </div>
                    <div className="pt-6 pb-2">
                        <h1 className="flex-1 text-lg leading-7 font-medium">Set Package Registry</h1>
                    </div>
                    <div className="border-2 border-solid border-black-200 rounded-lg h-auto break-words p-2">
                        npm config set registry {window.location.origin}/api/npm
                    </div>
                    <div className="pt-6 pb-2">
                        <h1 className="flex-1 text-md leading-7 font-medium" >Install From Registry</h1>
                    </div>
                    <div className="border-2 border-solid border-black-200 rounded-lg h-auto break-words p-2">
                        npm install ${repoName}
                    </div>
                </div>
            </div>
        );
    }

    const BinaryMeta = () => {
        return (
            <div>
                <div className="pl-6 lg:w-80">
                    <div className="pt-6 pb-2">
                        <h1 className="flex-1 text-lg leading-7 font-medium">Download (GET) from Url</h1>
                    </div>
                    <div className="border-2 border-solid border-black-200 rounded-lg p-2">
                        curl -L -o {repoName}-latest.tar.gz {window.location.origin}/api/{orgName}/{repoName}/latest
                    </div>
                </div>
            </div>
        );
    }

    const DockerMeta = () => {
        return (
            <div>
                <div className="pl-6 lg:w-80">
                    <div className="pt-6 pb-2">
                        <h1 className="flex-1 text-lg leading-7 font-medium">Load Container from Url</h1>
                    </div>
                    <div className="border-2 border-solid border-black-200 rounded-lg p-2">
                        curl -L {window.location.origin}/api/{orgName}/{repoName}/latest | docker load
                    </div>
                </div>
            </div>
        );
    }

    const package_type_table = {
        "binary": BinaryMeta,
        "npm": NpmMeta,
        "pip": PipMeta,
        "docker": DockerMeta
    }

    useEffect(() => {
        (async function() {
            if (valist) {
                const rawMeta = await valist.getRepoMeta(orgName, repoName);
                const metaResponse: { projectType: "binary" | "npm" | "pip" | "docker" | undefined }= await valist.fetchJSONfromIPFS(rawMeta);
                const repoMeta = package_type_table[metaResponse.projectType || "binary"];
                setMeta(repoMeta);
            }
        })()
    }, [valist]);

    return (
        <div className="bg-gray-50 pr-4 sm:pr-6 lg:pr-8 lg:flex-shrink-0 lg:border-l lg:border-gray-200 xl:pr-0">
            {meta}
        </div>
    )
}

export default ProjectMetaBar;
