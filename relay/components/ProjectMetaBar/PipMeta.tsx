import React, { useEffect, useState, useRef } from 'react';

import { copyToCB } from '../../utils/clipboard';

const PipMeta = (orgName: string = "organization", repoName: string = "repo", projectMeta: any = {}) => {

    const pipRef = useRef(null);

    const [origin, setOrigin] = useState("https://app.valist.io");
    useEffect(() => {
        setOrigin(window.location.origin);
    });

    return (
        <div>
            <div className="pl-6 lg:w-80">
                {projectMeta && <div>
                    <div className="pt-6 pb-2">
                        <h1 className="flex-1 text-2xl leading-7">Project Metadata</h1>
                    </div>
                    <div className="pt-6 pb-2">
                        <h1 className="flex-1 text-lg leading-7 font-medium">Homepage</h1>
                        <a className="text-blue-600" href={projectMeta['homepage']}>{projectMeta['homepage']}</a>
                    </div>
                    <div className="pt-6 pb-2">
                        <h1 className="flex-1 text-lg leading-7 font-medium">Repository</h1>
                        <a className="text-blue-600" href={projectMeta['repository']}>{projectMeta['repository']}</a>
                    </div>
                    <div className="pt-6 pb-2">
                        <h1 className="flex-1 text-lg leading-7 font-medium">Description</h1>
                        {projectMeta['description']}
                    </div>
                </div>}
                <div className="pt-6 pb-2">
                    <h1 className="flex-1 text-lg leading-7 font-medium">Pip Install From Url</h1>
                </div>
                <div ref={pipRef} onClick={() => copyToCB(pipRef)} className="border-2 border-solid border-black-200 rounded-lg p-2 bg-gray-200 cursor-pointer">
                    pip install {origin}/api/{orgName}/{repoName}/latest
                </div>
            </div>
        </div>
    );
}

export default PipMeta;
