import Layout from '../../../components/Layout/Layout';
import EditOrgForm from '../../../components/EditOrgForm/EditOrgForm';

export const EditOrgPage = () => {
    return (
        <Layout title="Valist | Edit Organization">
            <div className="flex-grow w-full pt-8 max-w-7xl mx-auto xl:px-8 lg:flex">
                <EditOrgForm />
            </div>
        </Layout>
    );
}

export default EditOrgPage;
