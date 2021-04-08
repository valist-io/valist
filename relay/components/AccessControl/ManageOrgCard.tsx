import Link from 'next/link';
import AddressIdenticon from '../Indenticons/AddressIdenticon';

const ManageAccessCard = ({orgName, projectName}: {orgName: string, projectName?: string}) => {
  return (
    <section aria-labelledby="recent-hires-title">
      <div className="rounded-lg bg-white overflow-hidden shadow">
        <div className="p-6">
          <h2 className="text-base font-medium text-gray-900" id="recent-hires-title">Manage Permissions</h2>
          <div className="flow-root mt-6">
            <ul className="-my-5 divide-y divide-gray-200">
              <li className="py-4">
                <div className="flex items-center space-x-4">
                  <div className="flex-shrink-0">
                    <AddressIdenticon address={"0xD978bb2AD7d67290a1780098BBBd8b3293a315E6"} height={8} />
                  </div>
                  <div className="flex-1 min-w-0">
                    <p className="text-sm font-medium text-gray-900 truncate">
                      0xD978bb2AD7d67290a1780098BBBd8b3293a315E6
                    </p>
                  </div>
                  <div>
                    <a href="#" className="inline-flex items-center shadow-sm px-2.5 py-0.5 border border-gray-300 text-sm leading-5 font-medium rounded-full text-gray-700 bg-white hover:bg-gray-50">
                      View
                    </a>
                  </div>
                </div>
              </li>
            </ul>
          </div>
          <div className="mt-6">
            <Link href={"/v/" + orgName + '/permissions'}>
              <a className="w-full flex justify-center items-center px-4 py-2 border border-gray-300 shadow-sm text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50">
                View all
              </a>
            </Link>
          </div>
        </div>
      </div>
    </section>
  );
}

export default ManageAccessCard;