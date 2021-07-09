import { useState, useEffect, useContext } from 'react';
import { useRouter } from 'next/router';
import {
  Organization, OrgMeta, VoteKeyEvent, VoteThresholdEvent,
} from 'valist/dist/types';
import ValistContext from '../../../components/Valist/ValistContext';
import DashboardLayout from '../../../components/Layouts/DashboardLayout';
import EditOrgMetadataForm from '../../../components/Organizations/EditOrgMetadataForm';
import EditOrgThresholdForm from '../../../components/Organizations/EditOrgThresholdForm';
import Permissions from '../../../components/AccessControl/Permissions';
import Votes from '../../../components/AccessControl/Votes';
import LoadingDialog from '../../../components/Dialog/LoadingDialog';
import ErrorDialog from '../../../components/Dialog/ErrorDialog';

export const EditOrgPage: React.FC = (): JSX.Element => {
  const valist = useContext(ValistContext);
  const router = useRouter();
  const orgName = `${router.query.orgName}`;

  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<Error>();
  const [org, setOrg] = useState<Organization>();
  const [orgAdmins, setOrgAdmins] = useState<string[]>([]);
  const [orgKeyVotes, setOrgKeyVotes] = useState<VoteKeyEvent[]>([]);
  const [orgThresholdVotes, setOrgThresholdVotes] = useState<VoteThresholdEvent[]>([]);

  const fetchData = async () => Promise.all([
    valist.getOrganization(orgName).then(setOrg),
    valist.getOrgAdmins(orgName).then(setOrgAdmins),
    valist.getVoteKeyEvents(orgName).then(setOrgKeyVotes),
    valist.getVoteThresholdEvents(orgName).then(setOrgThresholdVotes),
  ]);

  const getOrgData = async () => {
    try {
      setLoading(true);
      await fetchData();
    } catch (e) {
      setError(e);
    } finally {
      setLoading(false);
    }
  };

  const updateMeta = async (meta: OrgMeta) => {
    try {
      setLoading(true);
      await valist.setOrgMeta(orgName, meta, valist.defaultAccount);
      await fetchData();
    } catch (e) {
      setError(e);
    } finally {
      setLoading(false);
    }
  };

  const voteThreshold = async (threshold: string) => {
    try {
      setLoading(true);
      await valist.voteOrgThreshold(orgName, parseInt(threshold, 10));
      await fetchData();
    } catch (e) {
      setError(e);
    } finally {
      setLoading(false);
    }
  };

  const grantKey = async (key: string) => {
    try {
      setLoading(true);
      await valist.voteOrgAdmin(orgName, key);
      await fetchData();
    } catch (e) {
      setError(e);
    } finally {
      setLoading(false);
    }
  };

  const revokeKey = async (key: string) => {
    try {
      setLoading(true);
      await valist.revokeOrgAdmin(orgName, key);
      await fetchData();
    } catch (e) {
      setError(e);
    } finally {
      setLoading(false);
    }
  };

  const rotateKey = async (key: string) => {
    try {
      setLoading(true);
      await valist.rotateOrgAdmin(orgName, key);
      await fetchData();
    } catch (e) {
      setError(e);
    } finally {
      setLoading(false);
    }
  };

  const clearPendingKey = async (key: string, operation: string, index: number) => {
    try {
      setLoading(true);
      await valist.clearPendingOrgKey(orgName, operation, key, index);
      await fetchData();
    } catch (e) {
      setError(e);
    } finally {
      setLoading(false);
    }
  };

  const clearPendingThreshold = async (threshold: string, index: number) => {
    try {
      setLoading(true);
      await valist.clearPendingOrgThreshold(orgName, parseInt(threshold, 10), index);
      await fetchData();
    } catch (e) {
      setError(e);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    getOrgData();
  }, [valist, orgName]);

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
                      <Votes keyVotes={orgKeyVotes} thresholdVotes={orgThresholdVotes}
                      grantKey={grantKey} revokeKey={revokeKey} voteThreshold={voteThreshold}
                      clearPendingKey={clearPendingKey} clearPendingThreshold={clearPendingThreshold} />
                  </div>
                  <div className="text-center pt-8">
                    <h2 className="text-3xl">Manage Access & Permissions</h2>
                  </div>
                  <div className="flex-grow w-full pt-8 max-w-7xl mx-auto xl:px-8 lg:flex">
                      <Permissions keys={orgAdmins} rotateKey={rotateKey}
                      grantKey={grantKey} revokeKey={revokeKey} />
                  </div>
              </div>
            </div>
        </div>
        {loading && <LoadingDialog>Loading organization...</LoadingDialog>}
        {error && <ErrorDialog error={error} close={() => setError(undefined)} />}
    </DashboardLayout>
  );
};

export default EditOrgPage;
