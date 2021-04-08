import React, { useState, useEffect, useContext } from 'react';
import ValistContext from '../Valist/ValistContext';

export const Activity = () => {
    const valist = useContext(ValistContext)
    const [orgs, setOrgs] = useState(["Loading..."]);

    useEffect(() => {
        (async function() {
            if (valist) {
                try {
                    setOrgs(await valist.getOrganizationNames());
                } catch (e) {}
            }
        })()
    }, [valist]);

    return (
      <ul className="-my-5 divide-y divide-gray-200">
        {orgs.map((orgName: string) => (
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

export default Activity;
