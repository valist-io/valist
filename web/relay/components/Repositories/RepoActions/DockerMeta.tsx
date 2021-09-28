import React, { useEffect, useState, useRef } from 'react';

import copyToCB from '../../../utils/clipboard';

const DockerMeta = (orgName = 'organization', repoName = 'repo', projectMeta: any = {}) => {
  const curlRef = useRef(null);

  const [origin, setOrigin] = useState('https://app.valist.io');
  useEffect(() => {
    // @TODO respect local config here
    if (window.location.origin === 'http://localhost:3000') {
      setOrigin('http://localhost:9000');
    } else {
      setOrigin(window.location.origin);
    }
  });

  return (
        <div>
                {projectMeta
                    && <div>
                        {projectMeta.homepage
                            && <div className="pb-2">
                                <h1 className="flex-1 text-lg leading-7 font-medium">Homepage</h1>
                                <a className="text-blue-600" href={projectMeta.homepage}>{projectMeta.homepage}</a>
                            </div>
                        }

                        {projectMeta.repository
                            && <div className="pb-2">
                                <h1 className="flex-1 text-lg leading-7 font-medium">Repository</h1>
                                <a className="text-blue-600" href={projectMeta.repository}>{projectMeta.repository}</a>
                            </div>
                        }
                    </div>
                }
                <div className="pb-2 mb-2">
                    <h1 className="flex-1 text-lg leading-7 font-medium">Load Container from Url</h1>
                </div>
                <div ref={curlRef}
                  className="border-2 border-solid border-indigo-50 rounded-lg p-2
                  bg-indigo-50 break-all">
                    <svg
                      onClick={() => copyToCB(curlRef)}
                      xmlns="http://www.w3.org/2000/svg"
                      className="h-6 w-6 float-right cursor-pointer"
                      fill="none"
                      viewBox="0 0 24 24"
                      stroke="currentColor">
                        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2}
                        d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2
                        2 0 00-2 2v8a2 2 0 002 2z" />
                    </svg>
                    <p>curl -L {origin}/api/{orgName}/{repoName}/latest | docker load</p>
                </div>
        </div>
  );
};

export default DockerMeta;
