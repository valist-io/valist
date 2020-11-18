import Layout from '../../../../components/Layout/Layout';
import ProjectPermissions from '../../../../components/AccessControl/ProjectPermissions';

export const ProjectPermissionsPage = () => {
    return (
        <Layout title="Valist | Manage Project Permissions">
            <ProjectPermissions />
        </Layout>
    );
}

export default ProjectPermissionsPage;
