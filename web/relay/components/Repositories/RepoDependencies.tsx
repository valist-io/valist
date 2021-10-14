import { RepoMeta } from 'valist/dist/types';

interface RepoDependenciesProps {
  orgName: string,
  repoName: string,
  repoMeta: RepoMeta,
  releaseMeta: any,
}

export default function RepoDependencies(props: RepoDependenciesProps): JSX.Element {
  const { releaseMeta } = props;

  return (
      <div className="p-8">
        <h2 className="text-xl">Dependencies</h2>
        <hr/>
        <div className="flex flex-wrap mt-2 mb-4">
          {releaseMeta.dependencies && releaseMeta.dependencies.map((dependency: string) => (
            <div className="text-indigo-500 py-4 pr-4" key={dependency}>
              {dependency}
            </div>
          ))}
        </div>

        {releaseMeta.devDependencies && <div>
          <h2 className="text-xl">Dev Dependencies</h2>
          <hr/>
          <div className="flex flex-wrap mt-2">
            {releaseMeta.devDependencies.map((dependency: string) => (
              <div className="text-indigo-500 py-4 pr-4" key={dependency}>
                {dependency}
              </div>
            ))}
          </div>
        </div>}
      </div>
  );
}
