import { useRouter } from 'next/router';
import DashboardLayout from '../../../components/Layouts/DashboardLayout';
import OrgPermissions from '../../../components/AccessControl/OrgPermissions';

export const OrgPermissionsPage = () => {
  const router = useRouter();

  const { orgName } = router.query;

  return (
      <DashboardLayout title="Valist | Manage Organization Permissions">
        <div className="grid grid-cols-1 gap-4 items-start">
          <div className="grid grid-cols-1 gap-4 lg:col-span-2">
            <section aria-labelledby="profile-overview-title"></section>
              <div style={{ minHeight: '500px' }} className="rounded-lg bg-white p-10 overflow-hidden shadow">
                <div className="text-center">
                  <h2 className="text-3xl">
                    Manage Access & Permissions:
                    <span className="text-2xl text-violet-400">
                      {orgName}
                    </span>
                  </h2>
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

export default OrgPermissionsPage;
