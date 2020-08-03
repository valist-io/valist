import Link from 'next/link'
import Layout from '../components/Layout/Layout'
import { CreateOrganizationForm } from '../components/CreateOrganization/CreateOrganizationForm'

// @ts-ignore
export const IndexPage:FunctionComponent<any> = ({pageProps, contract}) => {
  console.log("Incomeing Index Page", contract)
  
  return (
    <Layout title="valist.io">
      <CreateOrganizationForm contract={contract} />
    </Layout>
  );
}

export default IndexPage
