import DashboardLayout from '../../../../components/Layouts/DashboardLayout';
import EditRepoForm from '../../../../components/Projects/EditProjectForm';
import { useRouter } from 'next/router';

export const EditProjectPage = () => {
    const router = useRouter();
    const { orgName, repoName } = router.query;

    return (
        <DashboardLayout title="Valist | Edit Project">
            <div className="flex-grow w-full pt-8 max-w-7xl mx-auto xl:px-8 lg:flex">
                <EditRepoForm orgName={`${orgName}`} repoName={`${repoName}`}/>
            </div>
        </DashboardLayout>
    );
}

export default EditProjectPage;
