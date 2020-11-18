import Layout from '../../../../components/Layout/Layout';
import ProjectPermissions from '../../../../components/AccessControl/ProjectPermissions';
import { useRouter } from 'next/router';

export const ProjectPermissionsPage = () => {
    const router = useRouter();

    const { orgName, repoName } = router.query;

    return (
        <Layout title="Valist | Manage Project Permissions">
            <ProjectPermissions orgName={`${orgName}`} repoName={`${repoName}`} />
        </Layout>
    );
}

export default ProjectPermissionsPage;
