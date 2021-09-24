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
  const [repo, setRepo] = useState<Repository>({
    orgID: 'Loading',
    threshold: 0,
    thresholdDate: 0,
    meta: {
      projectType: 'binary',
      description: 'Loading',
      name: 'Loading',
    },
    metaCID: 'Loading',
    tags: [],
  });
  const [repoDevs, setRepoDevs] = useState<string[]>([]);
  const [orgAdmins, setOrgAdmins] = useState<string[]>([]);
  const [repoReleases, setRepoReleases] = useState<Release[]>([]);
  const [repoView, setRepoView] = useState<string>('meta');
  const [repoReadme, setRepoReadme] = useState<string>('');

  const fetchReadme = async (releases: Release[]): Promise<string> => {
    setRepoReleases(releases.reverse());
    const release = releases[0];
    let markdown = '';
    if (release && release.metaCID !== '') {
      const req = await fetch(`https://gateway.valist.io/ipfs/${release.metaCID.replace('/ipfs/', '')}`);
      markdown = await req.text();
    }
    setRepoReadme(markdown);
    return markdown;
  };

  const fetchAll = () => Promise.all([
    valist.getRepository(orgName, repoName).then(setRepo),
    valist.getReleases(orgName, repoName, 1, 10).then(fetchReadme),
    valist.getRepoDevs(orgName, repoName).then(setRepoDevs),
    valist.getOrgAdmins(orgName).then(setOrgAdmins),
  ]);

  const getData = async () => {
    try {
      if (orgName && repoName) {
        await fetchAll();
      }
    } catch (e) {
      setError(e as any);
    }
  };

  useEffect(() => {
    getData();
  }, [valist, orgName, repoName]);

  return (
    <Layout>
        <div className="grid grid-cols-1 gap-4 items-start lg:grid-cols-6 lg:gap-8">
          <div className="grid grid-cols-1 gap-4 lg:col-span-4">
            <ProjectProfileCard
              setRepoView={setRepoView}
              orgName={orgName}
              repoName={repoName}
              repoMeta={repo.meta} />
            <section className="rounded-lg bg-white overflow-hidden shadow">
              {repo && <RepoContent
                repoReleases={repoReleases}
                repoReadme={repoReadme}
                view={repoView}
                orgName={orgName}
                repoName={repoName}
                repoMeta={repo.meta}
                repoDevs={repoDevs}
                orgAdmins={orgAdmins}/>}
            </section>
          </div>
          <div className="grid grid-cols-1 gap-4 lg:col-span-2">
            <RepoMetaCard
            repoMeta={repo.meta}
            orgName={orgName}
            repoName={repoName} />
          </div>
      </div>
      {error && <ErrorDialog error={error} close={() => setError(undefined)} />}
    </Layout>
  );
}
