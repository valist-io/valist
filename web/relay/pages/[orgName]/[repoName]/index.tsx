import React, { useEffect, useState, useContext } from 'react';
import { useRouter } from 'next/router';
import { Repository } from 'valist/dist/types';
import Layout from '../../../components/Layouts/DashboardLayout';
import ReleaseList from '../../../components/Releases/ReleaseList';
import ProjectProfileCard from '../../../components/Projects/ProjectProfileCard';
import ReleaseMetaCard from '../../../components/Releases/ReleaseMetaCard';
import ManageProjectCard from '../../../components/AccessControl/ManageProjectCard';
import ValistContext from '../../../components/Valist/ValistContext';
import LoadingDialog from '../../../components/Dialog/LoadingDialog';
import ErrorDialog from '../../../components/Dialog/ErrorDialog';

export default function Dashboard() {
  const valist = useContext(ValistContext);
  const router = useRouter();
  const orgName = `${router.query.orgName}`;
  const repoName = `${router.query.repoName}`;

  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<Error>();
  const [repo, setRepo] = useState<Repository>();
  const [repoTags, setRepoTags] = useState<string[]>([]);
  const [repoDevs, setRepoDevs] = useState<string[]>([]);
  const [orgAdmins, setOrgAdmins] = useState<string[]>([]);

  const fetchAll = () => Promise.all([
    valist.getRepository(orgName, repoName).then(setRepo),
    valist.getReleaseTags(orgName, repoName).then(setRepoTags),
    valist.getRepoDevs(orgName, repoName).then(setRepoDevs),
    valist.getOrgAdmins(orgName).then(setOrgAdmins),
  ]);

  const getRelease = async (tag: string) => {
    try {
      const { releaseCID } = await valist.getReleaseByTag(orgName, repoName, tag);
      window.location.assign(`https://gateway.valist.io/ipfs/${releaseCID}`);
    } catch (e) {
      setError(e);
    } finally {
      setLoading(false);
    }
  };

  const getData = async () => {
    try {
      setLoading(true);
      await fetchAll();
    } catch (e) {
      setError(e);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    getData();
  }, [valist, orgName, repoName]);

  return (
    <Layout>
        <div className="grid grid-cols-1 gap-4 items-start lg:grid-cols-3 lg:gap-8">
          <div className="grid grid-cols-1 gap-4 lg:col-span-2">
            { repo && <ProjectProfileCard orgName={orgName} repoName={repoName} repoMeta={repo.meta} /> }
            <section className="rounded-lg bg-white overflow-hidden shadow">
              <ReleaseList orgName={orgName} repoName={repoName} repoTags={repoTags} getRelease={getRelease} />
            </section>
          </div>
          <div className="grid grid-cols-1 gap-4">
            { repo && <ReleaseMetaCard orgName={orgName} repoName={repoName} repoMeta={repo.meta} /> }
            <ManageProjectCard orgName={orgName} repoName={repoName} repoDevs={repoDevs} orgAdmins={orgAdmins} />
          </div>
      </div>
      {loading && <LoadingDialog>Loading...</LoadingDialog>}
      {error && <ErrorDialog error={error} close={() => setError(undefined)} />}
    </Layout>
  );
}
