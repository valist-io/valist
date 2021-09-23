import React, { useEffect, useState, useContext } from 'react';
import { useRouter } from 'next/router';
import { Repository, Release } from 'valist/dist/types';
import Layout from '../../../components/Layouts/DashboardLayout';
import RepoContent from '../../../components/Projects/RepoContent';
import ProjectProfileCard from '../../../components/Projects/ProjectProfileCard';
import RepoMetaCard from '../../../components/Projects/RepoMeta';
import ValistContext from '../../../components/Valist/ValistContext';
import ErrorDialog from '../../../components/Dialog/ErrorDialog';

export default function Dashboard() {
  const valist = useContext(ValistContext);
  const router = useRouter();
  const orgName = `${router.query.orgName}`;
  const repoName = `${router.query.repoName}`;

  const [error, setError] = useState<Error>();
  const [repo, setRepo] = useState<Repository>();
  const [repoDevs, setRepoDevs] = useState<string[]>([]);
  const [orgAdmins, setOrgAdmins] = useState<string[]>([]);
  const [repoReleases, setRepoReleases] = useState<Release[]>([]);
  const [repoView, setRepoView] = useState<string>('meta');

  const fetchAll = () => Promise.all([
    valist.getRepository(orgName, repoName).then(setRepo),
    valist.getReleases(orgName, repoName, 1, 10).then(setRepoReleases),
    valist.getRepoDevs(orgName, repoName).then(setRepoDevs),
    valist.getOrgAdmins(orgName).then(setOrgAdmins),
  ]);

  const getData = async () => {
    try {
      await fetchAll();
    } catch (e) {
      setError(e as any);
    }
  };

  useEffect(() => {
    getData();
  }, [valist, orgName, repoName]);

  return (
    <Layout>
        <div className="grid grid-cols-1 gap-4 items-start lg:grid-cols-3 lg:gap-8">
          <div className="grid grid-cols-1 gap-4 lg:col-span-2">
            { repo && <ProjectProfileCard
                        setRepoView={setRepoView}
                        orgName={orgName}
                        repoName={repoName}
                        repoMeta={repo.meta} />
            }
            <section className="rounded-lg bg-white overflow-hidden shadow">
              {repo && <RepoContent
                repoReleases={repoReleases}
                view={repoView}
                orgName={orgName}
                repoName={repoName}
                repoMeta={repo.meta}
                repoDevs={repoDevs}
                orgAdmins={orgAdmins}/>}
            </section>
          </div>
          <div className="grid grid-cols-1 gap-4">
            { repo && <RepoMetaCard repoMeta={repo.meta} /> }
          </div>
      </div>
      {error && <ErrorDialog error={error} close={() => setError(undefined)} />}
    </Layout>
  );
}
