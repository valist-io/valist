import { RepoMeta } from 'valist/dist/types';

interface RepoDependenciesProps {
  orgName: string,
  repoName: string,
  repoMeta: RepoMeta,
  releaseMeta: any,
}

export default function RepoDependencies(props: RepoDependenciesProps): JSX.Element {
  const { repoMeta, releaseMeta } = props;
  let currentVersion;
  let dependencies;
  let devDependencies;

  if (repoMeta.projectType === 'npm' && releaseMeta) {
    currentVersion = releaseMeta.versions[Object.keys(releaseMeta.versions)[0]];
    dependencies = currentVersion.dependencies;
    devDependencies = currentVersion.devDependencies;
  }

  return (
      <div className="p-8">
        <h2 className="text-xl">Dependencies</h2>
        <hr/>
        <div className="flex flex-wrap mt-2 mb-4">
          {Object.keys(dependencies).map((dependency) => (
            <div className="text-indigo-500 py-4 pr-4" key={dependency}>
              {dependency}
            </div>
          ))}
        </div>

        <h2 className="text-xl">Dev Dependencies</h2>
        <hr/>
        <div className="flex flex-wrap mt-2">
          {Object.keys(devDependencies).map((dependency) => (
            <div className="text-indigo-500 py-4 pr-4" key={dependency}>
              {dependency}
            </div>
          ))}
        </div>
      </div>
  );
}
