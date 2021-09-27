import React, { useEffect, useContext, useState } from 'react';
import { useRouter } from 'next/router';
import { Organization } from 'valist/dist/types';
import ValistContext from '../../components/Valist/ValistContext';
import Layout from '../../components/Layouts/DashboardLayout';
import ProjectList from '../../components/Repositories/ProjectList';
import OrgProfileCard from '../../components/Organizations/OrgProfileCard';
import ManageOrgCard from '../../components/AccessControl/ManageOrgCard';
import ErrorDialog from '../../components/Dialog/ErrorDialog';

export default function Dashboard() {
  const valist = useContext(ValistContext);
  const router = useRouter();
  const orgName = `${router.query.orgName}`;

  const [error, setError] = useState<Error>();
  const [org, setOrg] = useState<Organization>({
    orgID: 'Loading',
    threshold: 0,
    thresholdDate: 0,
    meta: {
      name: 'Loading',
      description: 'Loading',
    },
    metaCID: 'Loading',
    repoNames: [],
  });
  const [orgAdmins, setOrgAdmins] = useState<string[]>([]);
  const [repoNames, setRepoNames] = useState<string[]>([]);

  const fetchAll = () => Promise.all([
    valist.getOrganization(orgName).then(setOrg),
    valist.getRepoNames(orgName).then(setRepoNames),
    valist.getOrgAdmins(orgName).then(setOrgAdmins),
  ]);

  const getData = async () => {
    try {
      await fetchAll();
    } catch (e) {
      setError(e as any);
    } finally {
      console.log('Data');
    }
  };

  useEffect(() => {
    getData();
  }, [valist, orgName]);

  return (
    <Layout>
        <div className="grid grid-cols-1 gap-4 items-start lg:grid-cols-3 lg:gap-8">
        <div className="grid grid-cols-1 gap-4 lg:col-span-2">
          <OrgProfileCard orgName={orgName} orgMeta={org.meta} />
          <section className="rounded-lg bg-white overflow-hidden shadow">
            <ProjectList orgName={orgName} repoNames={repoNames} />
          </section>
        </div>
        <div className="grid grid-cols-1 gap-4">
          <ManageOrgCard orgName={orgName} orgAdmins={orgAdmins} />
        </div>
      </div>
      {error && <ErrorDialog error={error} close={() => setError(undefined)} />}
    </Layout>
  );
}
