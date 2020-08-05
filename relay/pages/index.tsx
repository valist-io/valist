import Layout from '../components/Layout/Layout'
import { CreateOrganizationForm } from '../components/CreateOrganization/CreateOrganizationForm'

export const IndexPage = ({valist}: any) => {
  return (
    <Layout title="valist.io">
      <CreateOrganizationForm valist={valist} />
    </Layout>
  );
}

export default IndexPage
