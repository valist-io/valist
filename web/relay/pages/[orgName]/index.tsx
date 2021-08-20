import React, { useEffect, useContext, useState } from 'react';
import { useRouter } from 'next/router';
import { Organization } from 'valist/dist/types';
import ValistContext from '../../components/Valist/ValistContext';
import Layout from '../../components/Layouts/DashboardLayout';
import ProjectList from '../../components/Projects/ProjectList';
import OrgProfileCard from '../../components/Organizations/OrgProfileCard';
import ActivityBox from '../../components/Activity/ActivityBox';
import ManageOrgCard from '../../components/AccessControl/ManageOrgCard';
import LoadingDialog from '../../components/Dialog/LoadingDialog';
import ErrorDialog from '../../components/Dialog/ErrorDialog';

export default function Dashboard() {
  const valist = useContext(ValistContext);
  const router = useRouter();
  const orgName = `${router.query.orgName}`;

  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<Error>();
  const [org, setOrg] = useState<Organization>();
  const [orgAdmins, setOrgAdmins] = useState<string[]>([]);
  const [orgNames, setOrgNames] = useState<string[]>([]);
  const [repoNames, setRepoNames] = useState<string[]>([]);

  const fetchAll = () => Promise.all([
    valist.getOrganizationNames().then(setOrgNames),
    valist.getOrganization(orgName).then(setOrg),
    valist.getRepoNames(orgName).then(setRepoNames),
    valist.getOrgAdmins(orgName).then(setOrgAdmins),
  ]);

  const getData = async () => {
    try {
      setLoading(true);
      await fetchAll();
    } catch (e) {
      setError(e);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    getData();
  }, [valist, orgName]);

  return (
    <Layout>
        <div className="grid grid-cols-1 gap-4 items-start lg:grid-cols-3 lg:gap-8">
        <div className="grid grid-cols-1 gap-4 lg:col-span-2">
          { org && <OrgProfileCard orgName={orgName} orgMeta={org.meta} /> }
          <section className="rounded-lg bg-white overflow-hidden shadow">
            <ProjectList orgName={orgName} repoNames={repoNames} />
          </section>
        </div>
        <div className="grid grid-cols-1 gap-4">
          <ManageOrgCard orgName={orgName} orgAdmins={orgAdmins} />
          <ActivityBox orgNames={orgNames} />
        </div>
      </div>
      {loading && <LoadingDialog>Loading...</LoadingDialog>}
      {error && <ErrorDialog error={error} close={() => setError(undefined)} />}
    </Layout>
  );
}
