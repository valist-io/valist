import AddressIdenticon from '../Identicons/AddressIdenticon';

interface OrgAccessCardListItemProps {
  address: string
}

export default function RepoMemberListItem(props: OrgAccessCardListItemProps): JSX.Element {
  return (
    <tr key={props.address}>
      <td className="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
        <AddressIdenticon address={props.address} height={32} />
      </td>
      <td className="px-6 py-4 text-left whitespace-nowrap text-sm text-gray-500">{props.address}</td>
      <td className="px-6 py-4 text-left whitespace-nowrap text-sm text-gray-500">{'Org Admin'}</td>
    </tr>
  );
}
