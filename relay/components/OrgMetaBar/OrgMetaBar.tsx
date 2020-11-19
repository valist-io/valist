import React, { useState, useEffect, useContext } from 'react';
import ValistContext from '../ValistContext/ValistContext';

export const OrgMetaBar = ({ orgName }: { orgName: string }) => {
    const valist = useContext(ValistContext);
    const [orgMeta, setOrgMeta] = useState({ name: "Loading...", description: "Loading..." });

    useEffect(() => {
        (async function() {
            if (valist) {
                try {
                    const orgMeta = await valist.getOrganizationMeta(orgName);
                    setOrgMeta(orgMeta);
                } catch (e) {}
            }
        })()
    }, [valist]);

    return (
        <div className="bg-gray-50 pr-4 sm:pr-6 lg:pr-8 lg:flex-shrink-0 lg:border-l lg:border-gray-200 xl:pr-0">
            <div className="pl-6 lg:w-80">
                {orgMeta['name'] &&
                    <div className="pt-6 pb-2">
                        <h1 className="flex-1 text-lg leading-7 font-medium">Organization</h1>
                        {orgMeta['name']}
                    </div>
                }

                {orgMeta['description'] &&
                    <div className="pt-6 pb-2">
                        <h1 className="flex-1 text-lg leading-7 font-medium">Description</h1>
                        {orgMeta['description']}
                    </div>
                }
            </div>
        </div>
    );
}

export default OrgMetaBar;
