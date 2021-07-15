import { useState, useContext } from 'react';
import { useRouter } from 'next/router';
import { RepoMeta } from 'valist/dist/types';
import ValistContext from '../../../components/Valist/ValistContext';
import Layout from '../../../components/Layouts/DashboardLayout';
import CreateRepoForm from '../../../components/Projects/CreateProjectForm';
import LoadingDialog from '../../../components/Dialog/LoadingDialog';
import ErrorDialog from '../../../components/Dialog/ErrorDialog';

export default function CreateRepoPage(): JSX.Element {
  const valist = useContext(ValistContext);
  const router = useRouter();
  const orgName = `${router.query.orgName}`;

  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<Error>();

  const createRepo = async (name: string, meta: RepoMeta) => {
    try {
      setLoading(true);
      await valist.createRepository(orgName, name, meta, valist.defaultAccount);
      router.push(`/${orgName}/${name}`);
    } catch (e) {
      setError(e);
    } finally {
      setLoading(false);
    }
  };

  return (
      <Layout title={`Valist | ${orgName}`}>
        <CreateRepoForm orgName={orgName} createRepo={createRepo} />
        {loading && <LoadingDialog>Loading...</LoadingDialog>}
        {error && <ErrorDialog error={error} close={() => setError(undefined)} />}
      </Layout>
  );
}
