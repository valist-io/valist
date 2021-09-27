import RepoMemberListItem from './RepoMemberListItem';

interface RepoMemberListItemProps {
  repoDevs: any,
  orgAdmins: any,
}

export default function ManageProjectAccessCard(props: RepoMemberListItemProps): JSX.Element {
  return (
    <div className="flex flex-col">
      <div className="-my-2 overflow-x-auto sm:-mx-6 lg:-mx-8">
        <div className="py-2 align-middle inline-block min-w-full sm:px-6 lg:px-8">
          <div className="shadow overflow-hidden border-b border-gray-200 sm:rounded-lg">
            <table className="min-w-full divide-y divide-gray-200">
            <thead className="bg-gray-50">
                <tr>
                  <th
                    scope="col"
                    className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
                  >
                  </th>
                  <th
                    scope="col"
                    className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider truncate"
                  >
                    Address
                  </th>
                  <th
                    scope="col"
                    className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
                  >
                    Role
                  </th>
                </tr>
              </thead>
              <tbody className="bg-white divide-y divide-gray-200">
                { props.repoDevs.map((address: string) => <RepoMemberListItem key={address} address={address} />)}
                { props.orgAdmins.map((address: string) => <RepoMemberListItem key={address} address={address} />)}
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </div>
  );
}
