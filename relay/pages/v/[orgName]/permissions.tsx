import Layout from '../../../components/Layout/Layout';
import OrgPermissions from '../../../components/AccessControl/OrgPermissions';
import { useRouter } from 'next/router';

export const OrgPermissionsPage = () => {
    const router = useRouter();

    const { orgName } = router.query;

    return (
        <Layout title="Valist | Manage Organization Permissions">
            <OrgPermissions orgName={`${orgName}`} />
        </Layout>
    );
}

export default OrgPermissionsPage;
