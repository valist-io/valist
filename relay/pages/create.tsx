import Layout from '../components/Layout/Layout';
import { CreateOrganizationForm } from '../components/CreateOrganization/CreateOrganizationForm';

export const CreateOrgPage = ({valist}: any) => {
    return (
        <Layout title="Valist | Create Organization">
            <CreateOrganizationForm valist={valist} />
        </Layout>
    );
}

export default CreateOrgPage;
