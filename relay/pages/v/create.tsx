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
  const [loadingText, setLoadingText] = useState('Creating Organization...');
  const [error, setError] = useState<Error>();

  const createOrg = async (shortName: string, meta: OrgMeta) => {
    try {
      setLoading(true);
      const { orgID } = await valist.createOrganization(shortName, meta);
      setLoadingText(`Linking orgID ${orgID} to ${shortName}...`);
      await valist.linkNameToID(shortName, orgID);
      router.push(`/v/${shortName}/create`);
    } catch (e) {
      setError(e);
    } finally {
      setLoading(false);
    }
  };

  return (
    <DashboardLayout title="Valist | Create Organization">
      <CreateOrganizationForm createOrg={createOrg}/>
      {loading && <LoadingDialog>{ loadingText }</LoadingDialog>}
      {error && <ErrorDialog error={error} close={() => setError(undefined)} />}
    </DashboardLayout>
  );
}
