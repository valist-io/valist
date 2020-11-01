import React from 'react';
import IndexLayout from '../../../components/Layout/IndexLayout'
import ActivityFeed from '../../../components/ActivityFeed/ActivityFeed';
import ProfileSidebar from '../../../components/ProfileSidebar/ProfileSidebar';
import ReleaseList from '../../../components/List/ReleaseList';

import { useRouter } from 'next/router';

export const ReposPage = ({valist}: {valist: any}) => {
    const router = useRouter();
    const { orgName, repoName } = router.query;

    return (
        <IndexLayout title={`${orgName} | ${repoName}`}>
            <div className="flex-grow w-full max-w-7xl mx-auto xl:px-8 lg:flex">
                <div className="flex-1 min-w-0 bg-white xl:flex">
                    <ProfileSidebar valist={valist} />
                    <ReleaseList valist={valist} orgName={`${orgName}`} repoName={`${repoName}`} />
                    <ActivityFeed valist={valist} />
                </div>
            </div>
        </IndexLayout>
    );
}

export default ReposPage
