interface ActivityProps {
  orgNames: string[]
}

export default function Activity(props: ActivityProps): JSX.Element {
  return (
    <ul className="-my-5 divide-y divide-gray-200">
      {[...props.orgNames].reverse().map((orgName: string) => (
        <li className="py-4" key={orgName}>
          <div className="flex space-x-3">
              <div className="flex-1 space-y-1">
                  <div className="flex items-center justify-between">
                      <h3 className="text-sm font-medium leading-5">{orgName}</h3>
                  </div>
                  <p className="text-sm leading-5 text-gray-500">Organization {orgName} created!</p>
              </div>
          </div>
        </li>
      ))}
    </ul>
  );
}
