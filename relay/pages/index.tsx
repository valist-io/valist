import React from 'react';
import IndexLayout from '../components/Layout/IndexLayout'
import OrgList from '../components/List/OrgList'
import ActivityFeed from '../components/ActivityFeed/ActivityFeed';
import ProfileSidebar from '../components/ProfileSidebar/ProfileSidebar';

export const IndexPage = () => {
  return (
    <IndexLayout title="Valist | Securely Distribute Software">
      <div className="flex-grow w-full max-w-7xl mx-auto xl:px-8 lg:flex">
        <div className="flex-1 min-w-0 bg-white xl:flex">
          <ProfileSidebar />
          <OrgList />
          <ActivityFeed />
        </div>
      </div>
    </IndexLayout>
  );
}

export default IndexPage
