import React, { useEffect, useState, useRef } from 'react';

import copyToCB from '../../../utils/clipboard';

const NpmMeta = (orgName: string, repoName: string, projectMeta: any = {}) => {
  const registryRef = useRef(null);
  const installRef = useRef(null);
  const installFromRegistryRef = useRef(null);

  const [origin, setOrigin] = useState('https://app.valist.io');
  useEffect(() => {
    setOrigin(window.location.origin);
  });

  return (
        <div>
            <div className="lg:w-80">
                {projectMeta
                    && <div>
                        {projectMeta.homepage
                            && <div className="pb-2">
                                <h1 className="flex-1 text-lg leading-7 font-medium">Homepage</h1>
                                <a className="text-blue-600" href={projectMeta.homepage}>{projectMeta.homepage}</a>
                            </div>
                        }

                        {projectMeta.repository
                            && <div className="pt-6 pb-2">
                                <h1 className="flex-1 text-lg leading-7 font-medium">Github</h1>
                                <a className="text-blue-600" href={projectMeta.repository}>{projectMeta.repository}</a>
                            </div>
                        }
                    </div>
                }
                <div className="pb-2">
                    <h1 className="flex-1 text-lg leading-7 font-medium">NPM Direct Install From Url</h1>
                </div>
                <div ref={installRef} onClick={() => copyToCB(installRef)}
                  className="border-2 border-solid border-black-200 rounded-lg
                  p-2 bg-gray-200 cursor-pointer break-all mb-8">
                    npm install {origin}/api/{orgName}/{repoName}/latest
                </div>
                <div className="pb-2">
                    <h1 className="flex-1 text-lg leading-7 font-medium">Set Package Registry</h1>
                </div>
                <div ref={registryRef} onClick={() => copyToCB(registryRef)}
                  className="border-2 border-solid border-black-200 rounded-lg
                  h-auto p-2 bg-gray-200 cursor-pointer break-all mb-8">
                    npm config set registry {origin}/api/npm
                </div>
                <div className="pb-2">
                    <h1 className="flex-1 text-lg leading-7 font-medium" >Install From Registry</h1>
                </div>
                <div ref={installFromRegistryRef} onClick={() => copyToCB(installFromRegistryRef)}
                  className="border-2 border-solid border-black-200 rounded-lg
                  h-auto p-2 bg-gray-200 cursor-pointer break-all">
                    npm install @{orgName}/{repoName}
                </div>
            </div>
        </div>
  );
};

export default NpmMeta;
