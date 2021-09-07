import { useEffect, useState, useContext } from 'react';
import ValistContext from '../components/Valist/ValistContext';
import ActivityFeed from '../components/Activity/ActivityList';
import Layout from '../components/Layouts/DashboardLayout';
import LoadingDialog from '../components/Dialog/LoadingDialog';
import ErrorDialog from '../components/Dialog/ErrorDialog';

const ActivityPage = () => {
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
      <div className="grid grid-cols-1 gap-4 items-start">
        <div className="grid grid-cols-1 gap-4 lg:col-span-2">
          <section aria-labelledby="profile-overview-title">
              <div style={{ minHeight: '500px' }} className="rounded-lg bg-white p-10 overflow-hidden shadow">
                <h2 style={{ fontSize: '25px' }}>Activity</h2>
                <div className="flow-root mt-6 overflow-hidden">
                  <ActivityFeed orgNames={orgNames} />
                </div>
              </div>
          </section>
        </div>
      </div>
      {loading && <LoadingDialog>Loading...</LoadingDialog>}
      {error && <ErrorDialog error={error} close={() => setError(undefined)} />}
    </Layout>
  );
};

export default ActivityPage;
