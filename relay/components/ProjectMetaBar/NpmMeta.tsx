import React, { useEffect, useState } from 'react';

const NpmMeta = (orgName: string = "organization", repoName: string = "repo", projectMeta: any = {}) => {

    console.log(projectMeta, "in component")
    const [origin, setOrigin] = useState("https://app.valist.io");
    useEffect(() => {
        setOrigin(window.location.origin);
    });

    return (
        <div>
            <div className="pl-6 lg:w-80">
                <div className="pt-6 pb-2">
                    <h1 className="flex-1 text-3xl leading-7">Project Metadata</h1>
                </div>
                <div className="pt-6 pb-2">
                    Homepage {projectMeta['homepage']}
                </div>
                <div className="pt-6 pb-2">
                    Github {projectMeta['github']}
                </div>
                <div className="pt-6 pb-2">
                    <h1 className="flex-1 text-3xl leading-7">Installation</h1>
                </div>
                <div className="pt-6 pb-2">
                    <h1 className="flex-1 text-lg leading-7 font-medium">NPM Direct Install From Url</h1>
                </div>
                <div className="border-2 border-solid border-black-200 rounded-lg p-2">
                    npm install {origin}/api/{orgName}/{repoName}/latest
                </div>
                <div className="pt-6 pb-2">
                    <h1 className="flex-1 text-lg leading-7 font-medium">Set Package Registry</h1>
                </div>
                <div className="border-2 border-solid border-black-200 rounded-lg h-auto break-words p-2">
                    npm config set registry {origin}/api/npm
                </div>
                <div className="pt-6 pb-2">
                    <h1 className="flex-1 text-md leading-7 font-medium" >Install From Registry</h1>
                </div>
                <div className="border-2 border-solid border-black-200 rounded-lg h-auto break-words p-2">
                    npm install @{orgName}/{repoName}
                </div>
            </div>
        </div>
    );
}

export default NpmMeta;
