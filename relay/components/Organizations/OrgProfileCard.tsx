import Link from 'next/link';
import { OrgMeta } from 'valist/dist/types';
import AddressIdenticon from '../Identicons/AddressIdenticon';
import IsOrgAdmin from '../AccessControl/IsOrgAdmin';

interface OrgProfileCardProps {
  orgName: string,
  orgMeta: OrgMeta
}

export default function OrgProfileCard(props: OrgProfileCardProps): JSX.Element {
  return (
    <section aria-labelledby="profile-overview-title">
      <div className="rounded-lg bg-white overflow-hidden shadow">
        <h2 className="sr-only" id="profile-overview-title">Profile Overview</h2>
        <div className="bg-white p-6">
          <div className="sm:flex sm:items-center sm:justify-between">
            <div className="sm:flex sm:space-x-5">
              <div className="flex-shrink-0">
                <AddressIdenticon address={props.orgMeta.name} height={20}/>
              </div>
              <div className="mt-4 text-center sm:mt-0 sm:pt-1 sm:text-left">
                <p className="text-sm font-medium text-gray-600"></p>
                <p className="text-xl font-bold text-gray-900 sm:text-2xl">{props.orgMeta.name}</p>
                <p className="text-sm font-medium text-gray-600">{props.orgMeta.description}</p>
              </div>
            </div>
            <div className="mt-5 flex justify-center sm:mt-0">
              <IsOrgAdmin orgName={props.orgName}>
                <Link href={`/v/${props.orgName}/edit/`}>
                  <a className="flex justify-center items-center px-4 py-2 border
                  border-gray-300 shadow-sm text-sm font-medium rounded-md text-gray-700
                  bg-white hover:bg-gray-50">
                    Manage Organization
                  </a>
                </Link>

                <Link href={`/v/${props.orgName}/create/`}>
                  <a className="ml-2 flex justify-center items-center px-4 py-2 border
                  border-transparent text-sm leading-5 font-medium rounded-md text-white
                  bg-indigo-600 hover:bg-indigo-500 focus:outline-none focus:border-indigo-700
                  focus:shadow-outline-indigo active:bg-indigo-700 transition ease-in-out duration-150">
                    New Project
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
