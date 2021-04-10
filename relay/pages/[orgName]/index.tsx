import React from 'react';
import { useRouter } from 'next/router';
import Layout from '../../components/Layouts/DashboardLayout';
import ProjectList from '../../components/Projects/ProjectList';
import OrgProfileCard from '../../components/Organizations/OrgProfileCard';
import ActivityBox from '../../components/Activity/ActivityBox';
import ManageOrgCard from '../../components/AccessControl/ManageOrgCard';

export default function Dashboard() {
  const router = useRouter();
  const { orgName }: any = router.query;

  return (
    <Layout>
        <div className="grid grid-cols-1 gap-4 items-start lg:grid-cols-3 lg:gap-8">
        <div className="grid grid-cols-1 gap-4 lg:col-span-2">
          <OrgProfileCard orgName={orgName}/>
          <section className="rounded-lg bg-white overflow-hidden shadow">
            <ProjectList orgName={orgName} />
          </section>
        </div>
        <div className="grid grid-cols-1 gap-4">
          <ManageOrgCard orgName={orgName} />
          <ActivityBox />
        </div>
      </div>
    </Layout>
  );
}
