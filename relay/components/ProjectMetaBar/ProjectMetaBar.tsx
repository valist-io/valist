import React, { useState, useEffect, useContext } from 'react';
import ValistContext from '../ValistContext/ValistContext';

export const ProjectMetaBar = ({ orgName, repoName }: { orgName: string, repoName: string }) => {
    const valist = useContext(ValistContext)
    const [meta, setMeta] = useState({'install': 'Loading', 'registry': 'Loading'})

    const ProcessRepoMeta = (meta: any) => {
        const packageType = "npm"
        //@ts-ignore
        const metaObject = {}
    
        if(packageType == "npm"){
            console.log(meta)
            //@ts-ignore
            metaObject.install = `npm install https://app.valist.io/api/${orgName}/${repoName}/latest`
            //@ts-ignore
            metaObject.registry = `npm config set registry https://app.valist.io/api/npm`
        }
        return metaObject
    }
    
    useEffect(() => {
        (async function() {
            if (valist) {
                const rawMeta = await valist.getReleasesFromRepo(orgName, repoName)
                const repoMeta = ProcessRepoMeta(rawMeta)
                // @ts-ignore
                setMeta(repoMeta);
            }
        })()
    }, [valist]);

    return (
        <div className="bg-gray-50 pr-4 sm:pr-6 lg:pr-8 lg:flex-shrink-0 lg:border-l lg:border-gray-200 xl:pr-0">
            <div className="pl-6 lg:w-80">
                <div className="pt-6 pb-2">
                    <h1 className="flex-1 text-lg leading-7 font-medium">NPM Direct Install From Url</h1>
                </div>
                <div className="border-2 border-solid border-black-200 rounded-lg p-2">
                    {meta.install}
                </div>

                <div className="pt-6 pb-2">
                    <h1 className="flex-1 text-lg leading-7 font-medium" >Set Package Registry</h1>
                </div>
                <div className="border-2 border-solid border-black-200 rounded-lg h-auto break-words p-2">
                    {meta.registry}
                </div>
                <div className="pt-6 pb-2">
                    <h1 className="flex-1 text-md leading-7 font-medium" >Install From Registry</h1>
                </div>
                <div className="border-2 border-solid border-black-200 rounded-lg h-auto break-words p-2">
                    {`npm install ${repoName}`}
                </div>
            </div>
        </div>
    )
}

export default ProjectMetaBar;
