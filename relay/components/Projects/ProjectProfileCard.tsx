import React, { useState, useEffect, useContext } from 'react';
import Link from 'next/link';
import ValistContext from '../Valist/ValistContext';
import AddressIdenticon from '../Identicons/AddressIdenticon';
import IsOrgAdmin from '../AccessControl/IsOrgAdmin';
import IsRepoDev from '../AccessControl/IsRepoDev';

const ProjectProfileCard = ({ orgName, projectName }: { orgName: string, projectName: string }) => {
  const valist = useContext(ValistContext);
  const [projectMeta, setProjectMeta] = useState({ name: 'Loading...', description: 'Loading...' });

  useEffect(() => {
    (async () => {
      if (valist && orgName && projectName) {
        try {
          const repo = await valist.getRepository(orgName, projectName);
          setProjectMeta(repo.meta);
        } catch (e) {
          console.log(e);
        }
      }
    })();
  }, [valist]);

  return (
    <section aria-labelledby="profile-overview-title">
      <div className="rounded-lg bg-white overflow-hidden shadow">
        <h2 className="sr-only" id="profile-overview-title">Profile Overview</h2>
        <div className="bg-white p-6">
          <div className="sm:flex sm:items-center sm:justify-between">
            <div className="sm:flex sm:space-x-5">
              <div className="flex-shrink-0">
                <AddressIdenticon address={projectMeta.name} height={20}/>
              </div>
              <div className="mt-4 text-center sm:mt-0 sm:pt-1 sm:text-left max-w-sm">
                <p className="text-sm font-medium text-gray-600"></p>
                <p className="text-xl font-bold text-gray-900 sm:text-2xl">{`${orgName}/${projectName}`}</p>
                <p className="text-sm font-medium text-gray-600">{projectMeta.description}</p>
              </div>
            </div>
            <div className="mt-5 flex justify-center sm:mt-0">
              <IsOrgAdmin orgName={orgName}>
                <Link href={`/v/${orgName}/${projectName}/edit`}>
                  <a className="flex justify-center items-center px-4 py-2 border
                  border-gray-300 shadow-sm text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50">
                    Edit Project
                  </a>
                </Link>
              </IsOrgAdmin>
              <IsRepoDev orgName={orgName} repoName={projectName}>
                <Link href={`/v/${orgName}/${projectName}/publish`}>
                  <a className="ml-2 flex justify-center items-center px-4 py-2 border
                  border-transparent text-sm leading-5 font-medium rounded-md text-white
                  bg-indigo-600 hover:bg-indigo-500 focus:outline-none focus:border-indigo-700
                  focus:shadow-outline-indigo active:bg-indigo-700 transition ease-in-out duration-150">
                    Publish Release
                  </a>
                </Link>
              </IsRepoDev>
            </div>
          </div>
        </div>
      </div>
    </section>
  );
};

export default ProjectProfileCard;
