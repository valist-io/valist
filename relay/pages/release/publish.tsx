import Layout from '../../components/Layout/Layout'
import { PublishReleaseForm } from '../../components/PublishReleaseForm/PublishReleaseForm'

export const CreateRepoPage = ({valist}: any) => {
    return (
        <Layout title="valist.io">
            <PublishReleaseForm valist={valist} />
        </Layout>
    );
}

export default CreateRepoPage
