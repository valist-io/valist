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

    const projectTypes = {
        "binary": BinaryMeta(orgName, repoName),
        "npm": NpmMeta(orgName, repoName),
        "pip": PipMeta(orgName, repoName),
        "docker": DockerMeta(orgName, repoName)
    }

    useEffect(() => {
        (async function() {
            if (valist) {
                const rawMeta = await valist.getRepoMeta(orgName, repoName);
                const metaResponse = await valist.fetchJSONfromIPFS(rawMeta);
                const projectType = metaResponse.projectType;
                setType(projectType);
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
