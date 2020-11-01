import React from 'react';
import IndexLayout from '../../components/Layout/IndexLayout'
import ProjectList from '../../components/List/ProjectList'
import ActivityFeed from '../../components/ActivityFeed/ActivityFeed';
import ProfileSidebar from '../../components/ProfileSidebar/ProfileSidebar';

import { useRouter } from 'next/router';

export const ProjectsPage = () => {
    const router = useRouter();
    const { orgName } = router.query;

    return (
        <IndexLayout title="valist.io">
            <div className="flex-grow w-full max-w-7xl mx-auto xl:px-8 lg:flex">
                <div className="flex-1 min-w-0 bg-white xl:flex">
                    <ProfileSidebar />
                    <ProjectList orgName={`${orgName}`} />
                    <ActivityFeed />
                </div>
            </div>
        </IndexLayout>
    );
}

export default ProjectsPage;
