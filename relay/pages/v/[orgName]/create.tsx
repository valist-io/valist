import Layout from '../../../components/Layouts/DashboardLayout';
import CreateRepoForm from '../../../components/Projects/CreateProjectForm';

import { useRouter } from 'next/router';

export const CreateRepoPage = () => {
    const router = useRouter();
    const { orgName } = router.query;

    return (
        <Layout title={`Valist | ${orgName}`}>
            <CreateRepoForm orgName={orgName} />
        </Layout>
    );
}

export default CreateRepoPage;
