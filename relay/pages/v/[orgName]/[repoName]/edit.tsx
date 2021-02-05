import Layout from '../../../../components/Layout/Layout';
import EditRepoForm from '../../../../components/EditRepoForm/EditRepoForm';
import { useRouter } from 'next/router';

export const EditProjectPage = () => {
    const router = useRouter();
    const { orgName, repoName } = router.query;

    return (
        <Layout title="Valist | Edit Project">
            <div className="flex-grow w-full pt-8 max-w-7xl mx-auto xl:px-8 lg:flex">
                <EditRepoForm orgName={`${orgName}`} repoName={`${repoName}`}/>
            </div>
        </Layout>
    );
}

export default EditProjectPage;
