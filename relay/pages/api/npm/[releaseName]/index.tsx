import { NextApiRequest, NextApiResponse } from 'next'
import Valist, { Web3Providers } from 'valist';

export default async function getReleasesFromRepo(req: NextApiRequest, res: NextApiResponse) {

  // set .env.local to your local chain or set in production deployment
    if (process.env.WEB3_PROVIDER) {
        const provider = new Web3Providers.HttpProvider(process.env.WEB3_PROVIDER);

        const valist = new Valist(provider);
        await valist.connect();

        const {
            query: { releaseName },
        } = req;

        const [orgName, repoName] = decodeURIComponent(releaseName.toString().replace("@", "")).split("/");
        const releases = await valist.getReleasesFromRepo(orgName.toString(), repoName.toString());
        // const release = releases[0].returnValues

        /*
        let releaseMetaJson = await got(`https://ipfs.io/ipfs/${release.releaseMeta}`);
        let releaseMeta = JSON.parse(releaseMetaJson)
        */
        const latestTag = await valist.getLatestTagFromRepo(orgName, repoName);
        let version_object = {}

        for(let i = 0; i < releases.length; i++) {
            let currentRelease = releases[i];
            let modified_release = {
                name: repoName,
                version: releases[i].tag,
                repository: "",
                contributors: "",
                bugs: "",
                homepage: "",
                dependencies: {},
                dist: {
                    tarball: `https://ipfs.io/ipfs/${releases[i].releaseCID}`
                }
            }
            // @ts-ignore
            version_object[currentRelease.tag] = modified_release
        }

        return res.status(200).json({
            _id: "",
            name: "",
            "dist-tags": {
                latest: latestTag
            },
            versions: version_object
        });

    } else {
        return res.status(500).json({statusCode: 500, message: "No Web3 Provider!"});
    }
}
