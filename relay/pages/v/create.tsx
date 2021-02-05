import Layout from '../../components/Layout/Layout';
import { CreateOrganizationForm } from '../../components/CreateOrganization/CreateOrganizationForm';

export const CreateOrgPage = () => {
    return (
        <Layout title="Valist | Create Organization">
            <CreateOrganizationForm />
        </Layout>
    );
}

export default CreateOrgPage;
