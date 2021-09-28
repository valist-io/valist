import React, { useEffect, useState } from 'react';
import { projectTypes, GetActions } from '../../utils/Actions';

interface RepoMetaCardProps {
  repoMeta: any,
  orgName: string,
  repoName: string
}

const RepoMetaCard = (props:RepoMetaCardProps) => {
  const [actions, setActions] = useState({} as Record<string, any>);
  const { repoMeta } = props;

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
              <div className="lg:col-span-8 col-span-12">
                  <pre style={{ overflow: 'scroll' }} className="p-2 bg-indigo-50 rounded-lg">
                  <code>
                    {actions[projectTypes[repoMeta.projectType].default].command}
                  </code>
                  </pre>
              </div>
          </div>}
          {repoMeta.repository
              && <div className="pb-4">
                  <h1 className="text-xl text-gray-900 mb-1">Repository</h1>
                  <a className="text-blue-600" href={repoMeta.repository}>{repoMeta.repository}</a>
              </div>
          }

          {props.repoMeta.homepage
              && <div className="pb-4">
                  <h1 className="text-xl text-gray-900 mb-1">Homepage</h1>
                  <a className="text-blue-600" href={repoMeta.homepage}>{repoMeta.homepage}</a>
              </div>
          }
          {repoMeta.version
              && <div className="pb-4">
              <h1 className="text-base font-medium text-gray-900">Version</h1>
              <div className="text-gray-600">{'0.0.4'}</div>
          </div>}

          {repoMeta.License
              && <div className="pb-4">
              <h1 className="text-base font-medium text-gray-900">License</h1>
              <div className="text-gray-600">{'MPL-2.0'}</div>
          </div>}
        </div>
  );
};

export default RepoMetaCard;
