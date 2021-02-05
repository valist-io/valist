import Layout from '../../../components/Layout/Layout';
import { CreateRepoForm } from '../../../components/CreateRepoForm/CreateRepoForm';

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
