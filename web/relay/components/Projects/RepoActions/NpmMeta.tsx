import React, { useEffect, useState, useRef } from 'react';

import copyToCB from '../../../utils/clipboard';

const NpmMeta = (orgName: string, repoName: string) => {
  const registryRef = useRef(null);
  const installRef = useRef(null);
  const installFromRegistryRef = useRef(null);

  const [origin, setOrigin] = useState('https://app.valist.io');
  useEffect(() => {
    setOrigin(window.location.origin);
  });

  return (
            <div>
                <div className="pb-2">
                    <h2 className="text-xl text-gray-900 mt-2">NPM Direct Install From Url</h2>
                </div>
                <div ref={installRef} onClick={() => copyToCB(installRef)}
                  className="border-2 border-solid border-indigo-50 rounded-lg
                  p-2 bg-indigo-50 cursor-pointer break-all mb-8">
                    <svg xmlns="http://www.w3.org/2000/svg" className="h-6 w-6 float-right" fill="none"
                    viewBox="0 0 24 24" stroke="currentColor">
                        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2}
                        d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2
                        2 0 00-2 2v8a2 2 0 002 2z" />
                    </svg>
                    <p>npm i {origin}/api/{orgName}/{repoName}/latest</p>
                </div>
                <div className="pb-2">
                    <h2 className="text-xl text-gray-900 mt-2">Set Package Registry</h2>
                </div>
                <div ref={registryRef} onClick={() => copyToCB(registryRef)}
                  className="border-2 border-solid border-indigo-50 rounded-lg
                  h-auto p-2 bg-indigo-50 cursor-pointer break-all mb-8">
                    <svg xmlns="http://www.w3.org/2000/svg" className="h-6 w-6 float-right" fill="none"
                    viewBox="0 0 24 24" stroke="currentColor">
                        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2}
                        d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2
                        2 0 00-2 2v8a2 2 0 002 2z" />
                    </svg>
                    <p>npm config set registry {origin}/api/npm</p>
                </div>
                <div className="pb-2">
                    <h2 className="text-xl text-gray-900 mt-2" >Install From Registry</h2>
                </div>
                <div ref={installFromRegistryRef} onClick={() => copyToCB(installFromRegistryRef)}
                  className="border-2 border-solid border-indigo-50 rounded-lg
                  h-auto p-2 bg-indigo-50 cursor-pointer break-all">
                    <svg xmlns="http://www.w3.org/2000/svg" className="h-6 w-6 float-right" fill="none"
                    viewBox="0 0 24 24" stroke="currentColor">
                        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2}
                        d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2
                        2 0 00-2 2v8a2 2 0 002 2z" />
                    </svg>
                    <p>npm i @{orgName}/{repoName}</p>
                </div>
            </div>
  );
};

export default NpmMeta;
