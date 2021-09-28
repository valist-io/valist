import { useEffect, useState } from 'react';
import Layout from '../components/Layouts/DashboardLayout';
import { GetActions, integrations, links } from '../utils/Homepage';

export default function Dashboard() {
  const [actions, setActions] = useState({} as Record<string, any>);

  useEffect(() => {
    let { origin } = window.location;
    if (origin === 'http://localhost:3000') {
      origin = 'http://localhost:9000';
    }
    setActions(GetActions(origin));
  }, []);

  return (
    <Layout>
      <div className="grid grid-cols-1 gap-4 items-start lg:grid-cols-3 lg:gap-8">
        <div className="grid grid-cols-1 gap-4 lg:col-span-2">
          <section aria-labelledby="profile-overview-title">
            <div className="rounded-lg bg-white overflow-hidden shadow">
              <div className="bg-white p-6">
                <div className="sm:flex sm:items-center sm:justify-between">
                  <div className="sm:flex sm:space-x-5">
                    <div className="flex-shrink-0">
                      <img className="h-15" src="/images/ValistLogo128.png" />
                    </div>
                    <div className="mt-4 text-center sm:mt-0 sm:pt-1 sm:text-left">
                      <p className="text-3xl font-bold text-gray-900 sm:text-2xl">
                        Welcome to Valist
                      </p>
                      <p className="text-md font-medium text-gray-600">
                        Universal package repository, code-signing, and CDN system.<br/>
                        Build, sign, and distribute any software/firmware globally in just a few steps.
                        Powered by Ethereum, IPFS, and Filecoin.
                      </p>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </section>
        </div>

        <div className="grid grid-cols-1 gap-4">
          <section aria-labelledby="references-title">
            <div className="rounded-lg bg-white overflow-hidden shadow h-44">
              <div className="p-6">
                <h2 className="text-base font-medium text-gray-900" id="references-title">Quick Links</h2>
                <div className="flow-root mt-2">
                  <ul className="list-disc ml-4">
                    {links.map((link) => <li key={link.href}>
                      <span>{link.name} - <a className="text-blue-500" href={link.href}>
                        {link.href}</a>
                      </span></li>)}
                  </ul>
                </div>
              </div>
            </div>
          </section>
        </div>
      </div>

      {integrations.map((integration) => <div key={integration.name}
      className="mt-4 grid grid-cols-1 gap-4 items-start lg:grid-cols-1 lg:gap-8">
          <section aria-labelledby="npm-title">
            <div className="rounded-lg bg-white overflow-hidden shadow">
              <div className="p-6">
                <img style={{ display: 'inline-block' }}
                className="max-h-32 mr-4 mb-4" width="50px" src={integration.image} />
                <h2 style={{ display: 'inline-block' }}
                className="mb-4 text-base text-2xl text-gray-900" id="npm-title">{integration.name}</h2>
                {(Object.keys(actions).length !== 0) && integration.actions.map((action:string) => <div key={action}
                  className="mb-4 grid grid-cols-12 gap-4 lg:gap-8">
                    <div className="lg:col-span-4 col-span-12">
                      {actions[action].description}
                    </div>
                    <div className="lg:col-span-8 col-span-12">
                      <pre style={{ whiteSpace: 'pre-wrap' }} className="p-4 bg-indigo-50 rounded-lg">
                        <code>
                          {actions[action].code}
                        </code>
                      </pre>
                    </div>
                  </div>)}
                <div className="flex">
                  <a href={integration.docs} className="underline p-2 text-blue-500">Documentation</a>
                  {(integration.code !== '')
                    && <a href={integration.code} className="underline p-2 text-blue-500">Example Code</a>
                  }
                </div>
              </div>
            </div>
          </section>
        </div>)}
    </Layout>
  );
}
