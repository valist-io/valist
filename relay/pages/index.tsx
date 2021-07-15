import React, { useState, useEffect, useContext } from 'react';
import Layout from '../components/Layouts/DashboardLayout';
import OrgList from '../components/Organizations/OrgList';
import UserProfileBox from '../components/Users/UserProfileCard';
import ActivityBox from '../components/Activity/ActivityBox';
import ValistContext from '../components/Valist/ValistContext';
import LoadingDialog from '../components/Dialog/LoadingDialog';
import ErrorDialog from '../components/Dialog/ErrorDialog';

export default function Dashboard() {
  const valist = useContext(ValistContext);

  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<Error>();
  const [orgNames, setOrgNames] = useState<string[]>([]);

  const getData = async () => {
    try {
      setLoading(true);
      setOrgNames(await valist.getOrganizationNames());
    } catch (e) {
      setError(e);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    getData();
  }, [valist]);

  return (
    <Layout>
      <div className="grid grid-cols-1 gap-4 items-start lg:grid-cols-3 lg:gap-8">
        <div className="grid grid-cols-1 gap-4 lg:col-span-2">
          <UserProfileBox />
          <section className="rounded-lg bg-white overflow-hidden shadow">
            <OrgList orgNames={orgNames} />
          </section>
        </div>
        <div className="grid grid-cols-1 gap-4">
          <ActivityBox orgNames={orgNames} />
        </div>
      </div>
      {loading && <LoadingDialog>Loading...</LoadingDialog>}
      {error && <ErrorDialog error={error} close={() => setError(undefined)} />}
    </Layout>
  );
}
