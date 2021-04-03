import React from 'react';
import Layout from '../../components/Dashboard/Layout';
import OrgList from '../../components/Dashboard/OrgList';
import ProfileBox from '../../components/Dashboard/ProfileBox';
import ActivityBox from '../../components/Dashboard/Activity/ActivityBox';

export default function Dashboard(){
  return(
    <Layout>
      <div className="grid grid-cols-1 gap-4 items-start lg:grid-cols-3 lg:gap-8">
        <div className="grid grid-cols-1 gap-4 lg:col-span-2">
          <ProfileBox />
          <OrgList />
        </div>
        <div className="grid grid-cols-1 gap-4">
          <ActivityBox />
        </div>
      </div>
    </Layout>
  )
}
