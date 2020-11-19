import React, { useState, useEffect, useContext } from 'react';
import ValistContext from '../ValistContext/ValistContext';

import BinaryMeta from './BinaryMeta';
import NpmMeta from './NpmMeta';
import PipMeta from './PipMeta';
import DockerMeta from './DockerMeta';

type ProjectType = "binary" | "npm" | "pip" | "docker";

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
                    const projectType = repoMeta['projectType'];
                    setType(projectType);
                    setProjectMeta(repoMeta);
                } catch (e) {}
            }
        })()
    }, [valist]);

    return (
        <div className="bg-gray-50 pr-4 sm:pr-6 lg:pr-8 lg:flex-shrink-0 lg:border-l lg:border-gray-200 xl:pr-0">
            {projectTypes[type] || projectTypes["binary"]}
        </div>
    )
}

export default ProjectMetaBar;
