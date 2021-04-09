import { useRouter } from 'next/router';
import Layout from '../../../../components/Layouts/DashboardLayout';
import PublishReleaseForm from '../../../../components/Releases/PublishReleaseForm';

export const PublishReleasePage = () => {
  const router = useRouter();
  const { orgName, repoName } = router.query;

  return (
        <Layout title={`Publish Release | ${repoName}`}>
            <PublishReleaseForm orgName={orgName} repoName={repoName} />
        </Layout>
  );
};

export default PublishReleasePage;
