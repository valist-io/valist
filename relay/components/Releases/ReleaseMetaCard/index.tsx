import { ProjectType, RepoMeta } from 'valist/dist/types';
import BinaryMeta from './BinaryMeta';
import NpmMeta from './NpmMeta';
import PipMeta from './PipMeta';
import DockerMeta from './DockerMeta';

interface ProjectMetaBarProps {
  orgName: string,
  repoName: string,
  repoMeta: RepoMeta
}

export default function ProjectMetaBar(props: ProjectMetaBarProps): JSX.Element {
  const getProjectType = (projectType: ProjectType) => {
    switch (projectType) {
      case 'node':
        return NpmMeta(props.orgName, props.repoName, props.repoMeta);
      case 'python':
        return PipMeta(props.orgName, props.repoName, props.repoMeta);
      case 'docker':
        return DockerMeta(props.orgName, props.repoName, props.repoMeta);
      default:
        return BinaryMeta(props.orgName, props.repoName, props.repoMeta);
    }
  };

  return (
      <div className="rounded-lg bg-white overflow-hidden shadow p-6">
        { getProjectType(props.repoMeta.projectType) }
      </div>
  );
}
