import Link from 'next/link';
import NavTree from '../Navigation/NavTree';

interface ProjectsListProps {
  orgName: string,
  repoNames: string[]
}

export default function ProjectsList(props: ProjectsListProps): JSX.Element {
  return (
    <div className="bg-white lg:min-w-0 lg:flex-1">
        <div className="pl-4 pr-6 pt-4 pb-4 border-b border-t border-gray-200
        sm:pl-6 lg:pl-8 xl:pl-6 xl:pt-6 xl:border-t-0">
            <div className="flex items-center">
                <NavTree orgName={props.orgName} />
            </div>
        </div>
        <ul className="relative z-0 divide-y divide-gray-200 border-b border-gray-200">
        {props.repoNames.map((repoName: string) => (
            <Link href={`${props.orgName}/${repoName}`} key={repoName}>
                <li className="relative pl-4 pr-6 py-5 hover:bg-gray-50 sm:py-6 sm:pl-6 lg:pl-8 xl:pl-6">
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
                                    {repoName}
                                </a>
                                </h2>
                            </span>
                            </div>
                        </div>
                        <div className="sm:hidden">
                            <svg className="h-5 w-5 text-gray-400" xmlns="http://www.w3.org/2000/svg"
                            viewBox="0 0 20 20" fill="currentColor">
                            <path fillRule="evenodd" d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293
                            6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z"
                            clipRule="evenodd" />
                            </svg>
                        </div>
                    </div>
                </li>
            </Link>
        ))}
        </ul>
    </div>
  );
}
