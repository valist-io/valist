import { useRouter } from 'next/router';
import DashboardLayout from '../../../components/Layouts/DashboardLayout';
import EditOrgMetadataForm from '../../../components/Organizations/EditOrgMetadataForm';
import EditOrgMultifactorForm from '../../../components/Organizations/EditOrgMultifactorForm';
import OrgPermissions from '../../../components/AccessControl/OrgPermissions';
import OrgVotes from '../../../components/AccessControl/OrgVotes';

export const EditOrgPage = () => {
  const router = useRouter();
  const { orgName } = router.query;

  return (
      <DashboardLayout title="Valist | Manage Organization">
          <div className="grid grid-cols-1 gap-4 items-start">
              <div className="grid grid-cols-1 gap-4 lg:col-span-2">
                <section aria-labelledby="profile-overview-title"></section>
                <div style={{ minHeight: '500px' }} className="rounded-lg bg-white p-10 overflow-hidden shadow">
                    <div className="text-center">
                      <h2 className="text-3xl">Manage Organization</h2>
                    </div>
                    <div className="flex-grow w-full pt-8 max-w-7xl mx-auto xl:px-8 lg:flex">
                        <EditOrgMetadataForm orgName={`${orgName}`}/>
                    </div>
                    <div className="flex-grow w-full pt-8 max-w-7xl mx-auto xl:px-8 lg:flex">
                        <EditOrgMultifactorForm orgName={`${orgName}`}/>
                    </div>
                    <div className="text-center pt-8">
                      <h2 className="text-3xl">Multi-factor Votes</h2>
                    </div>
                    <div className="flex-grow w-full pt-8 max-w-7xl mx-auto xl:px-8 lg:flex">
                        <OrgVotes orgName={`${orgName}`} />
                    </div>
                    <div className="text-center pt-8">
                      <h2 className="text-3xl">Manage Access & Permissions</h2>
                    </div>
                    <div className="flex-grow w-full pt-8 max-w-7xl mx-auto xl:px-8 lg:flex">
                        <OrgPermissions orgName={`${orgName}`} />
                    </div>
                </div>
              </div>
          </div>
      </DashboardLayout>
  );
};

export default EditOrgPage;
