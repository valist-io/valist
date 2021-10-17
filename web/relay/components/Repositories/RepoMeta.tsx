import React, { useEffect, useState, useRef } from 'react';
import { projectTypes, GetActions } from '../../utils/Actions';
import copyToCB from '../../utils/clipboard';

interface RepoMetaCardProps {
  releaseMeta: any,
  repoMeta: any,
  orgName: string,
  repoName: string
}

const RepoMetaCard = (props: RepoMetaCardProps) => {
  const [actions, setActions] = useState({} as Record<string, any>);
  const { repoMeta, releaseMeta } = props;
  const installRef = useRef(null);

  useEffect(() => {
    let { origin } = window.location;
    if (origin === 'http://localhost:3000') {
      origin = 'http://localhost:9000';
    }
    setActions(GetActions(origin, props.orgName, props.repoName));
  }, []);

  return (
    <div className="rounded-lg bg-white overflow-hidden shadow p-6">
      {(repoMeta.name !== 'Loading') && projectTypes[repoMeta.projectType].default && Object.keys(actions).length
        && <div className="pb-4">
          <h1 className="text-xl text-gray-900 mb-2">
            {actions[projectTypes[repoMeta.projectType].default].description}
          </h1>
          <div ref={installRef} className="lg:col-span-8 col-span-12 flex bg-indigo-50 rounded-lg">
            <pre style={{ overflow: 'scroll' }} className="p-2 hide-scroll">
              <code>
                {actions[projectTypes[repoMeta.projectType].default].command}
              </code>
            </pre>
            <div className="m-2" style={{ minHeight: '25px', minWidth: '25px' }}>
              <svg onClick={() => copyToCB(installRef)}
                xmlns="http://www.w3.org/2000/svg"
                className="h-6 w-6 float-right cursor-pointer"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor">
                <path strokeLinecap="round"
                  strokeLinejoin="round"
                  strokeWidth={2}
                  d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2
                          2 0 00-2 2v8a2 2 0 002 2z" />
              </svg>
            </div>
          </div>
        </div>}
      {repoMeta.repository
        && <div className="pb-4">
          <h1 className="text-xl text-gray-900 mb-1">Repository</h1>
          <a className="text-blue-600" href={repoMeta.repository}>{repoMeta.repository}</a>
        </div>
      }

      {repoMeta.homepage
        && <div className="pb-4">
          <h1 className="text-xl text-gray-900 mb-1">Homepage</h1>
          <a className="text-blue-600" href={repoMeta.homepage}>{repoMeta.homepage}</a>
        </div>
      }
      {releaseMeta.version
        && <div className="pb-4">
          <h1 className="text-xl text-gray-900 mb-1">Version</h1>
          <div className="text-gray-600">{releaseMeta.version}</div>
        </div>}

      {releaseMeta.license
        && <div className="pb-4">
          <h1 className="text-xl text-gray-900 mb-1">License</h1>
          <div className="text-gray-600">{releaseMeta.license}</div>
        </div>}
    </div>
  );
};

export default RepoMetaCard;
