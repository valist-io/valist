import Link from 'next/link';
import { RepoMeta } from 'valist/dist/types';
import AddressIdenticon from '../Identicons/AddressIdenticon';
import IsOrgAdmin from '../AccessControl/IsOrgAdmin';

interface ProjectProfileCardProps {
  orgName: string,
  repoName: string,
  repoMeta: RepoMeta
}

export default function ProjectProfileCard(props: ProjectProfileCardProps): JSX.Element {
  return (
    <section aria-labelledby="profile-overview-title">
      <div className="rounded-lg bg-white overflow-hidden shadow">
        <h2 className="sr-only" id="profile-overview-title">Profile Overview</h2>
        <div className="bg-white p-6">
          <div className="sm:flex sm:items-center sm:justify-between">
            <div className="sm:flex sm:space-x-5">
              <div className="flex-shrink-0">
                <AddressIdenticon address={props.repoMeta.name} height={20}/>
              </div>
              <div>
                <p className="text-sm font-medium text-gray-600"></p>
                <p className="text-xl font-bold text-gray-900 sm:text-2xl">{`${props.orgName}/${props.repoName}`}</p>
                <p className="text-sm font-medium text-gray-600">{props.repoMeta.description}</p>
              </div>
            </div>
            <div className="mt-5 flex justify-center sm:mt-0">
              <IsOrgAdmin orgName={props.orgName}>
                <Link href={`/v/${props.orgName}/${props.repoName}/edit`}>
                  <a className="flex justify-center items-center px-4 py-2 border
                  border-gray-300 shadow-sm text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50">
                    Manage Project
                  </a>
                </Link>
              </IsOrgAdmin>
            </div>
          </div>
        </div>
      </div>
    </section>
  );
}
