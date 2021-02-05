import Layout from '../../../components/Layout/Layout';
import OrgPermissions from '../../../components/AccessControl/OrgPermissions';
import { useRouter } from 'next/router';

export const OrgPermissionsPage = () => {
    const router = useRouter();

    const { orgName } = router.query;

    return (
        <Layout title="Valist | Manage Organization Permissions">
            <div className="flex-grow w-full pt-8 max-w-7xl mx-auto xl:px-8 lg:flex">
                <OrgPermissions orgName={`${orgName}`} />
            </div>
        </Layout>
    );
}

export default OrgPermissionsPage;
