import Layout from '../components/Layout/Layout'
import { CreateOrganizationForm } from '../components/CreateOrganization/CreateOrganizationForm'

// @ts-ignore
export const IndexPage:FunctionComponent<any> = ({pageProps, valist}) => {
  return (
    <Layout title="valist.io">
      <CreateOrganizationForm valist={valist} />
    </Layout>
  );
}

export default IndexPage
