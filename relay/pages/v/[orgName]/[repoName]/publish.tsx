import Layout from '../../../../components/Layout/Layout'
import PublishReleaseForm from '../../../../components/PublishReleaseForm/PublishReleaseForm'

import { useRouter } from 'next/router';

export const PublishReleasePage = () => {
    const router = useRouter();
    const { orgName, repoName } = router.query;

    return (
        <Layout title={`Publish Release | ${repoName}`}>
            <PublishReleaseForm orgName={orgName} repoName={repoName} />
        </Layout>
    );
}

export default PublishReleasePage;
