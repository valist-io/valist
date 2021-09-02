import AddressIdenticon from '../Identicons/AddressIdenticon';
import IsOrgAdmin from './IsOrgAdmin';

const UserAccessCard = ({
  address, orgName, setRenderLoading, revokeRole, roleType,
} : { address: string, orgName: string, setRenderLoading: any, revokeRole: any, roleType: string }) => {
  const roleStyle: { [key: string]: any } = {
    orgOwner: { label: 'Org Owner', color: 'red' },
    orgAdmin: { label: 'Org Admin', color: 'red' },
    repoAdmin: { label: 'Repo Admin', color: 'orange' },
    repoDev: { label: 'Repo Developer', color: 'teal' },
  };
  return (
    <li key={address} className="col-span-1 flex flex-col text-center bg-white rounded-lg shadow-md">
        <div className="flex-1 flex flex-col p-8">
            <AddressIdenticon address={address} height={32}/>
            <h3 className="mt-6 text-gray-900 text-sm leading-5 font-medium break-words">{address}</h3>
            <dl className="mt-1 flex-grow flex flex-col justify-between">
                    <dd className="mt-3">
                    <span className={`px-2 py-1 text-xs leading-4 
                    font-medium bg-${roleStyle[roleType].color}-100 rounded-full`}>
                      { roleStyle[roleType].label }
                    </span>
                </dd>
            </dl>
        </div>
        {roleType !== 'orgOwner' && <IsOrgAdmin orgName={orgName}>
          <div className="border-t border-gray-200">
              <div className="-mt-px flex">
                  <div className="w-0 flex-1 flex border-r border-gray-200">
                  <a href="#" onClick={async () => {
                    setRenderLoading(true);
                    await revokeRole(address);
                    setRenderLoading(false);
                  }}
                  className="relative -mr-px w-0 flex-1 inline-flex items-center justify-center py-4
                  text-sm leading-5 text-gray-700 font-medium border border-transparent rounded-bl-lg
                  hover:text-gray-500 focus:outline-none focus:shadow-outline-blue focus:border-blue-300
                  focus:z-10 transition ease-in-out duration-150">
                      <svg className="w-5 h-5 text-gray-400" xmlns="http://www.w3.org/2000/svg"
                      fill="none" viewBox="0 0 24 24" stroke="currentColor">
                          <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2}
                          d="M18.364 18.364A9 9 0 005.636
                          5.636m12.728 12.728A9 9 0 015.636 5.636m12.728 12.728L5.636 5.636" />
                      </svg>
                      <span className="ml-3">Revoke {roleStyle[roleType].label} Role</span>
                  </a>
                  </div>
              </div>
          </div>
        </IsOrgAdmin>}
    </li>
  );
};

export default UserAccessCard;
