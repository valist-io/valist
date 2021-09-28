import React, { useEffect, useState, useRef } from 'react';

import copyToCB from '../../../utils/clipboard';

const PipMeta = (orgName = 'organization', repoName = 'repo') => {
  const pipRef = useRef(null);

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
            <div className="pb-2">
                <h2 className="flex-1 text-xl">Pip Install From Url</h2>
            </div>
            <div ref={pipRef}
            className="border-2 border-solid border-indigo-50 rounded-lg
            p-2 bg-indigo-50 break-all">
                <svg
                onClick={() => copyToCB(pipRef)}
                xmlns="http://www.w3.org/2000/svg"
                className="h-6 w-6 float-right  cursor-pointer"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2}
                    d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2
                    2 0 00-2 2v8a2 2 0 002 2z" />
                </svg>
                <p>pip install {origin}/api/{orgName}/{repoName}/latest</p>
            </div>
        </div>
  );
};

export default PipMeta;
