import { Release, RepoMeta } from 'valist/dist/types';
import RepoActions from './RepoActions';
import RepoMemberList from './RepoMemberList';
import ReleaseList from '../Releases/ReleaseList';
import RepoReadme from './RepoReadme';
import PublishReleaseSteps from '../Releases/PublishReleaseSteps';

interface ReleaseListProps {
  repoReleases: Release[]
  repoReadme: string,
  view: string,
  orgName: string,
  repoMeta: RepoMeta
  repoName: string,
  repoDevs: any,
  orgAdmins: any,
}

export default function RepoContent(props: ReleaseListProps): JSX.Element {
  const getRepoView = (view: string) => {
    let currentView = view;
    if (view === 'versions' && props.repoReleases.length === 0) {
      currentView = 'releaseSteps';
    }

    switch (currentView) {
      case 'readme':
        return <RepoReadme repoReadme={props.repoReadme} />;
      case 'install':
        return (<RepoActions
          orgName={props.orgName}
          repoName={props.repoName}
          repoMeta={props.repoMeta}
        />);
      case 'members':
        return (<RepoMemberList
          repoDevs={props.repoDevs}
          orgAdmins={props.orgAdmins} />);
      case 'versions':
        return (<ReleaseList
          repoReleases={props.repoReleases}
          orgName={props.orgName}
          repoName={props.repoName}
          />);
      case 'releaseSteps':
        return <PublishReleaseSteps />;
      default:
        return <RepoReadme repoReadme={props.repoReadme} />;
    }
  };

  return (
    <div className="bg-white lg:min-w-0 lg:flex-1">
        <div className="border-b border-t border-gray-200 xl:border-t-0">
          {getRepoView(props.view)}
        </div>
    </div>
  );
}
