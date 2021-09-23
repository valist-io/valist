import React from 'react';

const RepoMetaCard = (repoMeta: any) => {
  const meta = repoMeta.repoMeta;
  return (
            <div className="rounded-lg bg-white overflow-hidden shadow p-6">
                <div className="pb-4">
                    <h1 className="text-xl text-gray-900 mb-2">Install</h1>
                    <div className="lg:col-span-8 col-span-12">
                        <pre style={{ overflow: 'scroll' }} className="p-2 bg-indigo-50 rounded-lg">
                        <code>
                            {'npm i valist/sdk --registry=http://localhost:3000/api/'}
                        </code>
                        </pre>
                    </div>
                </div>
                {repoMeta && <div>
                    {meta.repository
                        && <div className="pb-4">
                            <h1 className="text-xl text-gray-900 mb-1">Repository</h1>
                            <a className="text-blue-600" href={meta.repository}>{meta.repository}</a>
                        </div>
                    }

                    {meta.homepage
                        && <div className="pb-4">
                            <h1 className="text-xl text-gray-900 mb-1">Homepage</h1>
                            <a className="text-blue-600" href={meta.homepage}>{meta.homepage}</a>
                        </div>
                    }
                    {meta.version
                        && <div className="pb-4">
                        <h1 className="text-base font-medium text-gray-900">Version</h1>
                        <div className="text-gray-600">{'0.0.4'}</div>
                    </div>}

                    {meta.License
                        && <div className="pb-4">
                        <h1 className="text-base font-medium text-gray-900">License</h1>
                        <div className="text-gray-600">{'MPL-2.0'}</div>
                    </div>}

                    {meta.published
                        && <div className="pb-4">
                        <h1 className="text-base font-medium text-gray-900">Last Publish</h1>
                        <div className="text-gray-600">{'a month ago'}</div>
                    </div>}
                </div>
                }
            </div>
  );
};

export default RepoMetaCard;
