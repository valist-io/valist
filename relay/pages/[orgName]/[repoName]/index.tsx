import React from 'react';
import IndexLayout from '../../../components/Layout/IndexLayout'
import ProjectMetaBar from '../../../components/ProjectMetaBar/ProjectMetaBar';
import ProfileActionSidebar from '../../../components/ActionSidebar/ProfileActionSidebar';
import ReleaseList from '../../../components/List/ReleaseList';

import { useRouter } from 'next/router';

export const ReposPage = () => {
    const router = useRouter();
    const { orgName, repoName } = router.query;

    return (
        <IndexLayout title={`${orgName} | ${repoName}`}>
            <div className="flex-grow w-full max-w-7xl mx-auto xl:px-8 lg:flex">
                <div className="flex-1 min-w-0 bg-white xl:flex">
                    <ProfileActionSidebar />
                    <ReleaseList orgName={`${orgName}`} repoName={`${repoName}`} />
                    <ProjectMetaBar orgName={`${orgName}`} repoName={`${repoName}`} />
                </div>
            </div>
        </IndexLayout>
    );
}

export default ReposPage
