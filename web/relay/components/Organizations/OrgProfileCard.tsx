import { OrgMeta } from 'valist/dist/types';
import AddressIdenticon from '../Identicons/AddressIdenticon';

interface OrgProfileCardProps {
  orgName: string,
  orgMeta: OrgMeta
}

export default function OrgProfileCard(props: OrgProfileCardProps): JSX.Element {
  return (
    <section aria-labelledby="profile-overview-title">
      <div className="rounded-lg bg-white overflow-hidden shadow">
        <h2 className="sr-only" id="profile-overview-title">Profile Overview</h2>
        <div className="bg-white p-6">
          <div className="sm:flex sm:items-center sm:justify-between">
            <div className="sm:flex sm:space-x-5">
              <div className="flex-shrink-0">
                <AddressIdenticon address={props.orgMeta.name} height={20}/>
              </div>
              <div>
                <p className="text-sm font-medium text-gray-600"></p>
                <p className="text-xl font-bold text-gray-900 sm:text-2xl">{props.orgMeta.name}</p>
                <p className="text-sm font-medium text-gray-600">{props.orgMeta.description}</p>
                { props.orgMeta.homepage && <a href={props.orgMeta.homepage}
                  className="text-sm font-medium text-gray-600">{props.orgMeta.homepage}</a>}
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
  );
}
