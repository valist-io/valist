import Layout from '../../../../components/Layout/Layout';
import ProjectPermissions from '../../../../components/AccessControl/ProjectPermissions';
import { useRouter } from 'next/router';

export const ProjectPermissionsPage = () => {
    const router = useRouter();

    const { orgName, repoName } = router.query;

    return (
        <Layout title="Valist | Manage Project Permissions">
            <div className="flex-grow w-full pt-8 max-w-7xl mx-auto xl:px-8 lg:flex">
                <ProjectPermissions orgName={`${orgName}`} repoName={`${repoName}`} />
            </div>
        </Layout>
    );
}

export default ProjectPermissionsPage;
