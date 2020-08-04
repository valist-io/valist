import Link from 'next/link'
import Layout from '../components/Layout/Layout'
import { CreateOrganizationForm } from '../components/CreateOrganization/CreateOrganizationForm'

// @ts-ignore
export const IndexPage:FunctionComponent<any> = ({pageProps, valist}) => {
  console.log("Incoming Index Page", valist)

  return (
    <Layout title="valist.io">
      <CreateOrganizationForm valist={valist} />
    </Layout>
  );
}

export default IndexPage
