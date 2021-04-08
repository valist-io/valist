import DashboardLayout from '../../components/Layouts/DashboardLayout';
import { CreateOrganizationForm } from '../../components/Organizations/CreateOrganizationForm';

const CreateOrgPage = () => {
    return (
        <DashboardLayout title="Valist | Create Organization">
            <CreateOrganizationForm />
        </DashboardLayout>
    );
}

export default CreateOrgPage;
