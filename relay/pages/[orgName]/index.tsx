import React from 'react';
import IndexLayout from '../../components/Layout/IndexLayout'
import ProjectList from '../../components/List/ProjectList'
import OrgActionSidebar from '../../components/ActionSidebar/OrgActionSidebar';
import OrgMetaBar from '../../components/OrgMetaBar/OrgMetaBar';
import { useRouter } from 'next/router';

export const ProjectsPage = () => {
    const router = useRouter();
    const { orgName } = router.query;

    return (
        <IndexLayout title="valist.io">
            <div className="flex-grow w-full max-w-7xl mx-auto xl:px-8 lg:flex">
                <div className="flex-1 min-w-0 bg-white xl:flex">
                    <OrgActionSidebar orgName={`${orgName}`}/>
                    <ProjectList orgName={`${orgName}`} />
                    <OrgMetaBar orgName={`${orgName}`}/>
                </div>
            </div>
        </IndexLayout>
    );
}

export default ProjectsPage;
