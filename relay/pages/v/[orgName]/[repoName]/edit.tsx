import { useState, useEffect, useContext } from 'react';
import { useRouter } from 'next/router';
import {
  Repository, RepoMeta, VoteKeyEvent, VoteThresholdEvent, VoteReleaseEvent, Release,
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
  const [keyVotes, setKeyVotes] = useState<VoteKeyEvent[]>([]);
  const [thresholdVotes, setThresholdVotes] = useState<VoteThresholdEvent[]>([]);
  const [releaseVotes, setReleaseVotes] = useState<VoteReleaseEvent[]>([]);

  const fetchData = async () => Promise.all([
    valist.getRepository(orgName, repoName).then(setRepo),
    valist.getRepoDevs(orgName, repoName).then(setRepoDevs),
    valist.getVoteKeyEvents(orgName, repoName).then(setKeyVotes),
    valist.getVoteThresholdEvents(orgName, repoName).then(setThresholdVotes),
    valist.getVoteReleaseEvents(orgName, repoName).then(setReleaseVotes),
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
                  <div className="text-center">
                    <h2 className="text-3xl">Manage Project</h2>
                  </div>
                  <div className="flex-grow w-full pt-8 max-w-7xl mx-auto xl:px-8 lg:flex">
                      {repo && <EditProjectMetaForm meta={repo.meta} repoName={repoName}
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
                      <Votes keyVotes={keyVotes} thresholdVotes={thresholdVotes}
                      grantKey={grantKey} revokeKey={revokeKey} voteThreshold={voteThreshold}
                      clearPendingKey={clearPendingKey} clearPendingThreshold={clearPendingThreshold}
                      releaseVotes={releaseVotes} voteRelease={voteRelease}
                      clearPendingRelease={clearPendingRelease} />
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
