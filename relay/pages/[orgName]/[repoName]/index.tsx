import React from 'react';
import Layout from '../../../components/Layouts/DashboardLayout';
import ReleaseList from '../../../components/Releases/ReleaseList';
import ProjectProfileCard from '../../../components/Projects/ProjectProfileCard';
import { useRouter } from 'next/router';
import ReleaseBox from '../../../components/Releases/ReleaseMetaCard';
import ManageProjectCard from '../../../components/AccessControl/ManageProjectCard';

export default function Dashboard(){
  const router = useRouter();
  const { orgName, repoName }: any = router.query;

  return (
    <Layout>
        <div className="grid grid-cols-1 gap-4 items-start lg:grid-cols-3 lg:gap-8">
          <div className="grid grid-cols-1 gap-4 lg:col-span-2">
            <ProjectProfileCard orgName={orgName} projectName={repoName}/>
            <section className="rounded-lg bg-white overflow-hidden shadow">
              <ReleaseList orgName={orgName} repoName={repoName} />
            </section>
          </div>
          <div className="grid grid-cols-1 gap-4">
            <ReleaseBox orgName={orgName} repoName={repoName} />
            <ManageProjectCard orgName={orgName} projectName={repoName} />
          </div>
      </div>
    </Layout>
  )
}
