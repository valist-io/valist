import Layout from '../../components/Layout/Layout'
import { CreateRepoForm } from '../../components/CreateRepoForm/CreateRepoForm'

export const CreateRepoPage = ({valist}: any) => {
    return (
        <Layout title="valist.io">
            <CreateRepoForm valist={valist} />
        </Layout>
    );
}

export default CreateRepoPage
