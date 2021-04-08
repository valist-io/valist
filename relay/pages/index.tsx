import React from 'react';
import Layout from '../components/Layouts/DashboardLayout';
import OrgList from '../components/Organizations/OrgList';
import ProfileBox from '../components/Users/UserProfileCard';
import ActivityBox from '../components/Activity/ActivityBox';

export default function Dashboard(){
  return(
    <Layout>
      <div className="grid grid-cols-1 gap-4 items-start lg:grid-cols-3 lg:gap-8">
        <div className="grid grid-cols-1 gap-4 lg:col-span-2">
          <ProfileBox />
          <section className="rounded-lg bg-white overflow-hidden shadow">
            <OrgList />
          </section>
        </div>
        <div className="grid grid-cols-1 gap-4">
          <ActivityBox />
        </div>
      </div>
    </Layout>
  )
}
