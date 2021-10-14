import Link from 'next/link';
import { RepoMeta } from 'valist/dist/types';
import AddressIdenticon from '../Identicons/AddressIdenticon';
import setLoading from '../../utils/loading';

interface ProjectProfileCardProps {
  // eslint-disable-next-line @typescript-eslint/ban-types
  setRepoView: Function
  orgName: string,
  repoName: string,
  repoMeta: RepoMeta
}

export default function ProjectProfileCard(props: ProjectProfileCardProps): JSX.Element {
  const { repoMeta } = props;
  return (
    <section aria-labelledby="profile-overview-title">
      <div className="rounded-lg bg-white overflow-hidden shadow pt-8 pr-6 pl-6">
        <h2 className="sr-only" id="profile-overview-title">Profile Overview</h2>
          <div className="sm:flex sm:items-center">
            <div className="sm:flex sm:space-x-5">
              <div className="flex-shrink-0 lg:ml-4">
                <AddressIdenticon address={props.repoMeta.name} height={56}/>
              </div>
              <div>
                <p className={`text-2xl font-bold text-gray-900 sm:text-2xl ${setLoading(props.repoName)}`}>
                  {`${props.repoName}`}
                </p>
                <p className="text-sm font-medium text-gray-600">
                  <Link href={`/${props.orgName}`}>
                    {`Published by: ${props.orgName}`}
                  </Link>
                </p>
              </div>
            </div>
          </div>
        <div className="flex flex-col sm:flex-row mt-4 cursor-pointer">
          <div onClick={() => { props.setRepoView('readme'); }}
            className="text-gray-600 text-center py-4 px-6 block hover:text-blue-500 focus:outline-none">
            Readme
          </div>
          <div onClick={() => { props.setRepoView('install'); }}
            className="text-gray-600 text-center py-4 px-6 block hover:text-blue-500 focus:outline-none">
            Install
          </div>
          <div onClick={() => { props.setRepoView('versions'); }}
            className="text-gray-600 text-center py-4 px-6 block hover:text-blue-500 focus:outline-none">
            Versions
          </div>
          <div onClick={() => { props.setRepoView('members'); }}
            className="text-gray-600 text-center py-4 px-6 block hover:text-blue-500 focus:outline-none">
            Maintainers
          </div>
          {(repoMeta.projectType === 'npm' || repoMeta.projectType === 'go')
            && <div onClick={() => { props.setRepoView('dependencies'); }}
                className="text-gray-600 text-center py-4 px-6 block hover:text-blue-500 focus:outline-none">
                Dependencies
              </div>
          }
        </div>
      </div>
    </section>
  );
}
