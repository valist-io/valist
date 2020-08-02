import Link from 'next/link'
import Layout from '../components/Layout/Layout'
import { Web3Container } from '../components/Web3/Web3Container'

const IndexPage = () => (
  <Layout title="valist.io">
    <Web3Container />
  </Layout>
)

export default IndexPage
