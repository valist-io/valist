import React from 'react';
import Layout from '../../../components/Dashboard/Layout';
import ProjectList from '../../../components/List/ProjectList';
import ProfileBox from '../../../components/Dashboard/ProfileBox';
import ActivityBox from '../../../components/Dashboard/Activity/ActivityBox';
import { useRouter } from 'next/router';

export default function Dashboard(){
  const router = useRouter();
  const { orgName }: any = router.query;

  return(
    <Layout>
        <div className="grid grid-cols-1 gap-4 items-start lg:grid-cols-3 lg:gap-8">
        <div className="grid grid-cols-1 gap-4 lg:col-span-2">
          <ProfileBox />
          <ProjectList orgName={orgName} />
        </div>
        <div className="grid grid-cols-1 gap-4">
          <ActivityBox />
        </div>
      </div>
    </Layout>
  )
}
