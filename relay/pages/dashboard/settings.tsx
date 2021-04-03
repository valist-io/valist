import Layout from '../../components/Dashboard/Layout';

const settings = () => {
  return (
    <Layout>
      <div className="grid grid-cols-1 gap-4 items-start lg:grid-cols-3 lg:gap-8">
        <div className="grid grid-cols-1 gap-4 lg:col-span-2">
          <section aria-labelledby="profile-overview-title">
            <div style={{minHeight: "500px"}} className="rounded-lg bg-white p-10 overflow-hidden shadow">
              <h2 style={{fontSize: '25px'}}>Settings</h2>
            </div>
          </section>
        </div>
      </div>
    </Layout>
  )
}

export default settings;