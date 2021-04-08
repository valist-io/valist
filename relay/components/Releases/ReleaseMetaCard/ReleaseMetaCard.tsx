import React, { useState, useEffect, useContext } from 'react';
import ValistContext from '../../Valist/ValistContext';

import BinaryMeta from './BinaryMeta';
import NpmMeta from './NpmMeta';
import PipMeta from './PipMeta';
import DockerMeta from './DockerMeta';
import { ProjectType } from 'valist';

export const ProjectMetaBar = ({ orgName, repoName }: { orgName: string, repoName: string }) => {
    const valist = useContext(ValistContext);
    const [type, setType] = useState<ProjectType>("binary");
    const [projectMeta, setProjectMeta] = useState();

    const projectTypes = {
        "binary": BinaryMeta(orgName, repoName, projectMeta),
        "npm": NpmMeta(orgName, repoName, projectMeta),
        "pip": PipMeta(orgName, repoName, projectMeta),
        "docker": DockerMeta(orgName, repoName, projectMeta)
    }

    useEffect(() => {
        (async function() {
            if (valist) {
                try {
                    const repoMeta = await valist.getRepoMeta(orgName, repoName);
                    const projectType = repoMeta['projectType'] as ProjectType;
                    setType(projectType);
                    setProjectMeta(repoMeta);
                } catch (e) {}
            }
        })()
    }, [valist]);

    return (
        <div className="rounded-lg bg-white overflow-hidden shadow p-6">
            {projectTypes[type] || projectTypes["binary"]}
        </div>
    )
}

export default ProjectMetaBar;
