import { useRouter } from 'next/router';
import { useState, useEffect, useContext } from 'react';
import {
  Organization, OrgMeta, VoteKeyEvent, VoteThresholdEvent,
} from 'valist/dist/types';
import ValistContext from '../../../components/Valist/ValistContext';
import DashboardLayout from '../../../components/Layouts/DashboardLayout';
import EditOrgMetadataForm from '../../../components/Organizations/EditOrgMetadataForm';
import EditOrgThresholdForm from '../../../components/Organizations/EditOrgThresholdForm';
import OrgPermissions from '../../../components/AccessControl/OrgPermissions';
import OrgVotes from '../../../components/AccessControl/OrgVotes';
import LoadingDialog from '../../../components/LoadingDialog/LoadingDialog';

export const EditOrgPage = () => {
  const valist = useContext(ValistContext);
  const router = useRouter();
  const orgName = `${router.query.orgName}`;

  const [loading, setLoading] = useState(false);
  const [org, setOrg] = useState<Organization>();
  const [orgAdmins, setOrgAdmins] = useState<string[]>([]);
  const [orgKeyVotes, setOrgKeyVotes] = useState<VoteKeyEvent[]>([]);
  const [orgThresholdVotes, setOrgThresholdVotes] = useState<VoteThresholdEvent[]>([]);

  const fetchData = async () => Promise.all([
    valist.getOrganization(orgName).then(setOrg),
    valist.getOrgAdmins(orgName).then(setOrgAdmins),
    valist.getOrgVoteKeyEvents(orgName).then(setOrgKeyVotes),
    valist.getOrgVoteThresholdEvent(orgName).then(setOrgThresholdVotes),
  ]);

  const updateMeta = async (meta: OrgMeta) => {
    try {
      setLoading(true);
      await valist.setOrgMeta(orgName, meta, valist.defaultAccount);
    } catch (e) {
      console.log(e);
    } finally {
      setLoading(false);
    }
  };

  const voteThreshold = async (threshold: number) => {
    try {
      setLoading(true);
      await valist.voteOrgThreshold(orgName, threshold);
    } catch (e) {
      console.log(e);
    } finally {
      setLoading(false);
    }
  };

  const voteAdmin = async (key: string) => {
    try {
      setLoading(true);
      await valist.voteOrgAdmin(orgName, key);
    } catch (e) {
      console.log(e);
    } finally {
      setLoading(false);
    }
  };

  const revokeAdmin = async (key: string) => {
    try {
      setLoading(true);
      await valist.revokeOrgAdmin(orgName, key);
    } catch (e) {
      console.log(e);
    } finally {
      setLoading(false);
    }
  };

  // const rotateAdmin = async (key: string) => {
  //   try {
  //     setLoading(true);
  //     await valist.rotateOrgAdmin(orgName, key);
  //   } catch (e) {
  //     console.log(e);
  //   } finally {
  //     setLoading(false);
  //   }
  // };

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
                        { org && <EditOrgMetadataForm orgMeta={org.meta} updateOrgMeta={updateMeta} /> }
                    </div>
                    <div className="flex-grow w-full pt-8 max-w-7xl mx-auto xl:px-8 lg:flex">
                        { org && <EditOrgThresholdForm orgThreshold={org.threshold}
                        voteThreshold={voteThreshold} /> }
                    </div>
                    <div className="text-center pt-8">
                      <h2 className="text-3xl">Multi-factor Votes</h2>
                    </div>
                    <div className="flex-grow w-full pt-8 max-w-7xl mx-auto xl:px-8 lg:flex">
                        <OrgVotes orgKeyVotes={orgKeyVotes} orgThresholdVotes={orgThresholdVotes}
                        voteAdmin={voteAdmin} revokeAdmin={revokeAdmin} voteThreshold={voteThreshold} />
                    </div>
                    <div className="text-center pt-8">
                      <h2 className="text-3xl">Manage Access & Permissions</h2>
                    </div>
                    <div className="flex-grow w-full pt-8 max-w-7xl mx-auto xl:px-8 lg:flex">
                        <OrgPermissions orgName={orgName} orgAdmins={orgAdmins}
                        voteAdmin={voteAdmin} revokeAdmin={revokeAdmin} />
                    </div>
                </div>
              </div>
          </div>
      </DashboardLayout>
  );
};

export default EditOrgPage;
