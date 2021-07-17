import NavTree from '../Navigation/NavTree';

interface ReleaseListProps {
  orgName: string,
  repoName: string,
  repoTags: string[],
  getRelease: (tag: string) => Promise<void>
}

export default function ReleaseList(props: ReleaseListProps): JSX.Element {
  const example = `# Install the Valist CLI
npm i -g @valist/cli --registry=https://valist.io/api/npm

# Initialize your project
valist init

# Publish a release
valist publish`;

  return (
    <div className="bg-white lg:min-w-0 lg:flex-1">
        <div className="pl-4 pr-6 pt-4 pb-4 border-b border-t border-gray-200
        sm:pl-6 lg:pl-8 xl:pl-6 xl:pt-6 xl:border-t-0">
            <div className="flex items-center">
                <NavTree orgName={props.orgName} repoName={props.repoName} />
            </div>
        </div>
        { props.repoTags.length === 0
          && <div className="p-8">
            <a className="float-right text-indigo-500" href="https://docs.valist.io">docs</a>
            <h2 className="text-xl">Publish a new release</h2>
            <pre className="p-4 my-4 bg-gray-200 rounded-lg overflow-x-scroll"><code>{ example }</code></pre>
          </div>
        }
        <ul className="relative z-0 divide-y divide-gray-200 border-b border-gray-200">
        {[...props.repoTags].reverse().map((tag) => (
          <li key={tag} onClick={() => props.getRelease(tag)}
          className="relative pl-4 pr-6 py-5 hover:bg-gray-50 sm:py-6 sm:pl-6 lg:pl-8 xl:pl-6">
              <div className="flex items-center justify-between space-x-4">
                  <div className="min-w-0 space-y-3">
                      <div className="flex items-center space-x-3">
                          <span aria-label="Running" className="h-4 w-4 bg-green-100 rounded-full
                          flex items-center justify-center">
                              <span className="h-2 w-2 bg-green-400 rounded-full"></span>
                          </span>
                          <span className="block">
                              <h2 className="text-sm font-medium leading-5">
                                  <a href="#">
                                      <span className="absolute inset-0"></span>
                                      {tag}
                                  </a>
                              </h2>
                          </span>
                      </div>
                  </div>
              </div>
          </li>
        ))}
        </ul>
    </div>
  );
}
