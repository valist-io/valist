import Link from 'next/link';
import ReleaseList from '../Releases/ReleaseList';

interface ReleaseListProps {
  orgName: string,
  repoName: string,
  repoTags: string[],
  getRelease: (tag: string) => Promise<void>
  // eslint-disable-next-line @typescript-eslint/ban-types
  setRepoView: Function,
}

export default function ReleaseCard(props: ReleaseListProps): JSX.Element {
  return (
    <div className="rounded-lg bg-white overflow-hidden shadow p-6">
        <h2 className="text-base font-medium text-gray-900" id="recent-hires-title">Versions</h2>
        <ReleaseList
          orgName={props.orgName}
          repoName={props.repoName}
          repoTags={props.repoTags}
          getRelease={props.getRelease}/>
      <div className="mt-6">
        <Link href={`/${props.orgName}/${props.repoName}/versions`}>
          <div className="w-full flex justify-center items-center px-4 py-2 border border-gray-300 shadow-sm
                        text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50">
            View all
          </div>
        </Link>
      </div>
    </div>
  );
}
