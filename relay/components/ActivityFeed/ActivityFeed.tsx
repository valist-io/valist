import React, { useState, useEffect, useContext } from 'react';
import ValistContext from '../ValistContext';

export const ActivityFeed = () => {
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
        <div className="bg-gray-50 pr-4 sm:pr-6 lg:pr-8 lg:flex-shrink-0 lg:border-l lg:border-gray-200 xl:pr-0">
            <div className="pl-6 lg:w-80">
                <div className="pt-6 pb-2">
                    <h2 className="text-sm leading-5 font-semibold">Activity</h2>
                </div>
                <div>
                {orgs.map((orgName: string) => (
                    <ul className="divide-y divide-gray-200" key={orgName}>
                        <li className="py-4">
                        <div className="flex space-x-3">
                            <div className="flex-1 space-y-1">
                                <div className="flex items-center justify-between">
                                    <h3 className="text-sm font-medium leading-5">{orgName}</h3>
                                </div>
                                <p className="text-sm leading-5 text-gray-500">Organization {orgName} created!</p>
                            </div>
                        </div>
                        </li>
                    </ul>
                ))}
            </div>
        </div>
    </div>
    )
}

export default ActivityFeed;
