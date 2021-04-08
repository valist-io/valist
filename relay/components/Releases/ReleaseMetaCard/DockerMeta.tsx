import React, { useEffect, useState, useRef } from 'react';

import { copyToCB } from '../../../utils/clipboard';

const DockerMeta = (orgName: string = "organization", repoName: string = "repo", projectMeta: any = {}) => {

    const curlRef = useRef(null);

    const [origin, setOrigin] = useState("https://app.valist.io");
    useEffect(() => {
        setOrigin(window.location.origin);
    });

    return (
        <div>
            <div className="lg:w-80">
                {projectMeta &&
                    <div>
                        {projectMeta['homepage'] &&
                            <div className="pb-2">
                                <h1 className="flex-1 text-lg leading-7 font-medium">Homepage</h1>
                                <a className="text-blue-600" href={projectMeta['homepage']}>{projectMeta['homepage']}</a>
                            </div>
                        }

                        {projectMeta['repository'] &&
                            <div className="pb-2">
                                <h1 className="flex-1 text-lg leading-7 font-medium">Repository</h1>
                                <a className="text-blue-600" href={projectMeta['repository']}>{projectMeta['repository']}</a>
                            </div>
                        }
                    </div>
                }
                <div className="pt-6 pb-2">
                    <h1 className="flex-1 text-lg leading-7 font-medium">Load Container from Url</h1>
                </div>
                <div ref={curlRef} onClick={() => copyToCB(curlRef)} className="border-2 border-solid border-black-200 rounded-lg p-2 bg-gray-200 cursor-pointer break-all">
                    curl -L {origin}/api/{orgName}/{repoName}/latest | docker load
                </div>
            </div>
        </div>
    );
}

export default DockerMeta;
