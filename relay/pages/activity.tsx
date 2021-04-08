import ActivityFeed from "../components/Activity/ActivityList";
import Layout from "../components/Layouts/DashboardLayout";

const ActivityPage = () => {
  return (
    <Layout>
      <div className="grid grid-cols-1 gap-4 items-start">
        <div className="grid grid-cols-1 gap-4 lg:col-span-2">
          <section aria-labelledby="profile-overview-title">
              <div style={{minHeight: "500px"}} className="rounded-lg bg-white p-10 overflow-hidden shadow">
                <h2 style={{fontSize: '25px'}}>Activity</h2>
                <div className="flow-root mt-6 overflow-hidden">
                  <ActivityFeed />
                </div>
              </div>
          </section>
        </div>
      </div>
    </Layout>
  )
}

export default ActivityPage;