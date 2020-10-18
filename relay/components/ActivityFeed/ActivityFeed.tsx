import React, { useState, useEffect } from 'react';

export const ActivityFeed = ({valist}: {valist: any}) => {

    const [created, setCreated] = useState([{ returnValues: { orgName: "Loading...", orgMeta: "Loading..." }, blockNumber: 0, transactionHash: "0x0" }]);

    useEffect(() => {
        (async function() {
            if (valist) {
                setCreated(await valist.getCreatedOrganizations());
            }
        })()
    }, [valist])

    return (
        <div className="bg-gray-50 pr-4 sm:pr-6 lg:pr-8 lg:flex-shrink-0 lg:border-l lg:border-gray-200 xl:pr-0">
            <div className="pl-6 lg:w-80">
                <div className="pt-6 pb-2">
                    <h2 className="text-sm leading-5 font-semibold">Activity</h2>
                </div>
                <div>
                {created.map((org: any) => (
                    <ul className="divide-y divide-gray-200" key={org.transactionHash}>
                        <li className="py-4">
                        <div className="flex space-x-3">
                            <img className="h-6 w-6 rounded-full" src="https://images.unsplash.com/photo-1517365830460-955ce3ccd263?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=256&h=256&q=80" alt="" />
                            <div className="flex-1 space-y-1">
                                <div className="flex items-center justify-between">
                                    <h3 className="text-sm font-medium leading-5">{org.returnValues.orgName}</h3>
                                </div>
                                <p className="text-sm leading-5 text-gray-500">Organization {org.returnValues.orgName} created!
                                <br/><a href={`https://ropsten.etherscan.io/tx/${org.transactionHash}`} className="font-bold">View Tx</a></p>
                            </div>
                        </div>
                        </li>
                    </ul>
                ))}
                <div className="py-4 text-sm leading-5 border-t border-gray-200">
                    <a href="#" className="text-indigo-600 font-semibold hover:text-indigo-900">View all activity &rarr;</a>
                </div>
            </div>
        </div>
    </div>
    )
}

export default ActivityFeed;
