import Layout from '../../../components/Layouts/DashboardLayout';
import EditOrgForm from '../../../components/Organizations/EditOrgForm';
import { useRouter } from 'next/router';

export const EditOrgPage = () => {
    const router = useRouter();
    const { orgName } = router.query;

    return (
        <Layout title="Valist | Edit Organization">
            <div className="flex-grow w-full pt-8 max-w-7xl mx-auto xl:px-8 lg:flex">
                <EditOrgForm orgName={`${orgName}`}/>
            </div>
        </Layout>
    );
}

export default EditOrgPage;
