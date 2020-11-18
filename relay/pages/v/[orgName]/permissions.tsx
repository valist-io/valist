import Layout from '../../../components/Layout/Layout';
import OrgPermissions from '../../../components/AccessControl/OrgPermissions';

export const OrgPermissionsPage = () => {
    return (
        <Layout title="Valist | Manage Organization Permissions">
            <OrgPermissions />
        </Layout>
    );
}

export default OrgPermissionsPage;
