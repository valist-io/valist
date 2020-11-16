import Layout from '../../../components/Layout/Layout';
import { CreateRepoForm } from '../../../components/CreateRepoForm/CreateRepoForm';

import { useRouter } from 'next/router';

export const CreateRepoPage = ({valist}: any) => {
    const router = useRouter();
    const { orgName } = router.query;

    return (
        <Layout title={`Valist | ${orgName}`}>
            <CreateRepoForm valist={valist} orgName={orgName} />
        </Layout>
    );
}

export default CreateRepoPage;
