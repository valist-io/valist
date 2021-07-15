import { useState, useContext } from 'react';
import { useRouter } from 'next/router';
import { OrgMeta } from 'valist/dist/types';
import ValistContext from '../../components/Valist/ValistContext';
import DashboardLayout from '../../components/Layouts/DashboardLayout';
import CreateOrganizationForm from '../../components/Organizations/CreateOrganizationForm';
import LoadingDialog from '../../components/Dialog/LoadingDialog';
import ErrorDialog from '../../components/Dialog/ErrorDialog';

export default function CreateOrgPage(): JSX.Element {
  const valist = useContext(ValistContext);
  const router = useRouter();

  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<Error>();

  const createOrg = async (shortName: string, meta: OrgMeta) => {
    try {
      setLoading(true);
      await valist.createOrganization(shortName, meta, valist.defaultAccount);
      router.push(`/v/${shortName}/create`);
    } catch (e) {
      setError(e);
    } finally {
      setLoading(false);
    }
  };

  return (
    <DashboardLayout title="Valist | Create Organization">
      <CreateOrganizationForm createOrg={createOrg} />
      {loading && <LoadingDialog>Loading...</LoadingDialog>}
      {error && <ErrorDialog error={error} close={() => setError(undefined)} />}
    </DashboardLayout>
  );
}
