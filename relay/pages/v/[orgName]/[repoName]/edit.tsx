import { useState, useEffect, useContext } from 'react';
import { useRouter } from 'next/router';
import Link from 'next/link';
import {
  Repository, RepoMeta, Release, PendingRelease,
} from 'valist/dist/types';
import ValistContext from '../../../../components/Valist/ValistContext';
import DashboardLayout from '../../../../components/Layouts/DashboardLayout';
import EditProjectMetaForm from '../../../../components/Projects/EditProjectMetaForm';
import EditProjectThresholdForm from '../../../../components/Projects/EditProjectThresholdForm';
import LoadingDialog from '../../../../components/Dialog/LoadingDialog';
import ErrorDialog from '../../../../components/Dialog/ErrorDialog';
import Votes from '../../../../components/AccessControl/Votes';
import Permissions from '../../../../components/AccessControl/Permissions';

export const EditProjectPage: React.FC = (): JSX.Element => {
  const valist = useContext(ValistContext);
  const router = useRouter();
  const orgName = `${router.query.orgName}`;
  const repoName = `${router.query.repoName}`;

  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<Error>();
  const [repo, setRepo] = useState<Repository>();
  const [repoDevs, setRepoDevs] = useState<string[]>([]);
  const [repoEvents, setRepoEvents] = useState<any[]>([]);
  const [pendingKeys, setPendingKeys] = useState<string[]>([]);
  const [pendingThresholds, setPendingThresholds] = useState<string[]>([]);
  const [pendingReleases, setPendingReleases] = useState<PendingRelease[]>([]);

  const fetchData = async () => Promise.all([
    valist.getRepository(orgName, repoName).then(setRepo),
    valist.getRepoDevs(orgName, repoName).then(setRepoDevs),
    valist.getRepoEvents(orgName, repoName).then(setRepoEvents),
    valist.getPendingRepoDevs(orgName, repoName).then(setPendingKeys),
    valist.getPendingRepoThresholds(orgName, repoName).then(setPendingThresholds),
    valist.getPendingReleases(orgName, repoName).then(setPendingReleases),
  ]);

  const getRepoData = async () => {
    try {
      setLoading(true);
      await fetchData();
    } catch (e) {
      setError(e);
    } finally {
      setLoading(false);
    }
  };

  const updateMeta = async (meta: RepoMeta) => {
    try {
      setLoading(true);
      await valist.setRepoMeta(orgName, repoName, meta, valist.defaultAccount);
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
      await valist.voteRepoThreshold(orgName, repoName, parseInt(threshold, 10));
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
      await valist.voteRepoDev(orgName, repoName, key);
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
      await valist.revokeRepoDev(orgName, repoName, key);
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
      await valist.rotateRepoDev(orgName, repoName, key);
      await fetchData();
    } catch (e) {
      setError(e);
    } finally {
      setLoading(false);
    }
  };

  const voteRelease = async (release: Release) => {
    try {
      setLoading(true);
      await valist.publishRelease(orgName, repoName, release, valist.defaultAccount);
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
      await valist.clearPendingRepoKey(orgName, repoName, operation, key, index);
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
      await valist.clearPendingRepoThreshold(orgName, repoName, parseInt(threshold, 10), index);
      await fetchData();
    } catch (e) {
      setError(e);
    } finally {
      setLoading(false);
    }
  };

  const clearPendingRelease = async (release: Release, index: number) => {
    try {
      setLoading(true);
      await valist.clearPendingRelease(orgName, repoName, release, index);
      await fetchData();
    } catch (e) {
      setError(e);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    getRepoData();
  }, [valist, orgName, repoName]);

  return (
        <DashboardLayout title="Valist | Manage Project">
          <div className="grid grid-cols-1 gap-4 items-start">
            <div className="grid grid-cols-1 gap-4 lg:col-span-2">
              <section aria-labelledby="profile-overview-title"></section>
              <div style={{ minHeight: '500px' }} className="rounded-lg bg-white p-10 overflow-hidden shadow">
                  <div className="text-center text-3xl">
                    <Link href={`/${orgName}`}>{orgName}</Link>
                    <span className="mx-8">/</span>
                    <Link href={`/${orgName}/${repoName}`}>{repoName}</Link>
                  </div>
                  <div className="flex-grow w-full pt-8 max-w-7xl mx-auto xl:px-8 lg:flex">
                      {repo && <EditProjectMetaForm meta={repo.meta} orgName={orgName} repoName={repoName}
                      setRepoMeta={updateMeta} /> }
                  </div>
                  <div className="flex-grow w-full pt-8 max-w-7xl mx-auto xl:px-8 lg:flex">
                      { repo && <EditProjectThresholdForm threshold={repo.threshold}
                      voteThreshold={voteThreshold} /> }
                  </div>
                  <div className="text-center pt-8">
                    <h2 className="text-3xl">Multi-factor Votes</h2>
                  </div>
                  <div className="flex-grow w-full pt-8 max-w-7xl mx-auto xl:px-8 lg:flex">
                    <Votes pendingKeys={pendingKeys} pendingThresholds={pendingThresholds}
                      pendingReleases={pendingReleases} votes={repoEvents} grantKey={grantKey}
                      revokeKey={revokeKey} voteThreshold={voteThreshold} clearPendingKey={clearPendingKey}
                      clearPendingThreshold={clearPendingThreshold} clearPendingRelease={clearPendingRelease}
                      voteRelease={voteRelease} />
                  </div>
                  <div className="text-center pt-8">
                    <h2 className="text-3xl">Manage Access & Permissions</h2>
                  </div>
                  <div className="flex-grow w-full pt-8 max-w-7xl mx-auto xl:px-8 lg:flex">
                      <Permissions keys={repoDevs} rotateKey={rotateKey}
                      grantKey={grantKey} revokeKey={revokeKey} />
                  </div>
              </div>
            </div>
          </div>
          {loading && <LoadingDialog>Loading project...</LoadingDialog>}
          {error && <ErrorDialog error={error} close={() => setError(undefined)} />}
        </DashboardLayout>
  );
};

export default EditProjectPage;
