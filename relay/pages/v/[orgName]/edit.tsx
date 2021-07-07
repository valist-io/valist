import { useRouter } from 'next/router';
import { useState, useEffect, useContext } from 'react';
import ValistContext from '../../../components/Valist/ValistContext';
import DashboardLayout from '../../../components/Layouts/DashboardLayout';
import EditOrgMetadataForm from '../../../components/Organizations/EditOrgMetadataForm';
import EditOrgThresholdForm from '../../../components/Organizations/EditOrgThresholdForm';
import OrgPermissions from '../../../components/AccessControl/OrgPermissions';
import OrgVotes from '../../../components/AccessControl/OrgVotes';
import LoadingDialog from '../../../components/LoadingDialog/LoadingDialog';
import { Organization } from '../../../../lib/dist/types';

export const EditOrgPage = () => {
  const valist = useContext(ValistContext);
  const router = useRouter();
  const { orgName } = router.query;

  const [loading, setLoading] = useState(false);
  const [org, setOrg] = useState<Organization>();
  const [orgAdmins, setOrgAdmins] = useState(['0x0']);
  const [pendingOrgAdmins, setPendingOrgAdmins] = useState(['0x0']);

  const fetchData = async () => {
    try {
      setOrg(await valist.getOrganization(`${orgName}`));
      setOrgAdmins(await valist.getOrgAdmins(`${orgName}`) || ['0x0']);
      setPendingOrgAdmins(await valist.getPendingOrgAdmins(`${orgName}`) || ['0x0']);
    } catch(e) {
      console.log(e);
    }
  }

  useEffect(() => {
    setLoading(true);
    fetchData();
    setLoading(false);
  }, [valist, orgName]);

  if (loading) return (<LoadingDialog>Loading organization...</LoadingDialog>);

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
                        { org && org.meta && <EditOrgMetadataForm orgName={`${orgName}`} orgMeta={org.meta} /> }
                    </div>
                    <div className="flex-grow w-full pt-8 max-w-7xl mx-auto xl:px-8 lg:flex">
                        { org && <EditOrgThresholdForm orgName={`${orgName}`} orgThreshold={org.threshold} /> }
                    </div>
                    <div className="text-center pt-8">
                      <h2 className="text-3xl">Multi-factor Votes</h2>
                    </div>
                    <div className="flex-grow w-full pt-8 max-w-7xl mx-auto xl:px-8 lg:flex">
                        <OrgVotes orgName={`${orgName}`} pendingOrgAdmins={pendingOrgAdmins} />
                    </div>
                    <div className="text-center pt-8">
                      <h2 className="text-3xl">Manage Access & Permissions</h2>
                    </div>
                    <div className="flex-grow w-full pt-8 max-w-7xl mx-auto xl:px-8 lg:flex">
                        <OrgPermissions orgName={`${orgName}`} orgAdmins={orgAdmins} />
                    </div>
                </div>
              </div>
          </div>
      </DashboardLayout>
  );
};

export default EditOrgPage;
