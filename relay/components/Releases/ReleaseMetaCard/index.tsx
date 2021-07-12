import React, { useState, useEffect, useContext } from 'react';
import { ProjectType, RepoMeta } from 'valist/dist/types';
import ValistContext from '../../Valist/ValistContext';

import BinaryMeta from './BinaryMeta';
import NpmMeta from './NpmMeta';
import PipMeta from './PipMeta';
import DockerMeta from './DockerMeta';

export const ProjectMetaBar = ({ orgName, repoName }: { orgName: string, repoName: string }) => {
  const valist = useContext(ValistContext);
  const [type, setType] = useState<ProjectType>('binary');
  const [projectMeta, setProjectMeta] = useState<RepoMeta | undefined>();

  const projectTypes = {
    binary: BinaryMeta(orgName, repoName, projectMeta),
    node: NpmMeta(orgName, repoName, projectMeta),
    python: PipMeta(orgName, repoName, projectMeta),
    docker: DockerMeta(orgName, repoName, projectMeta),
    go: BinaryMeta(orgName, repoName, projectMeta),
    rust: BinaryMeta(orgName, repoName, projectMeta),
    'c++': BinaryMeta(orgName, repoName, projectMeta),
  };

  useEffect(() => {
    (async () => {
      if (valist) {
        try {
          const repo = await valist.getRepository(orgName, repoName);
          const { projectType } = repo.meta;
          setType(projectType);
          setProjectMeta(repo.meta);
        } catch (e) { console.log(e); }
      }
    })();
  }, [valist]);

  return (
        <div className="rounded-lg bg-white overflow-hidden shadow p-6">
            {projectTypes[type] || projectTypes.binary}
        </div>
  );
};

export default ProjectMetaBar;
