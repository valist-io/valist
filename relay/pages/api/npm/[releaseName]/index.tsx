import { NextApiRequest, NextApiResponse } from 'next'
import Valist, { Web3Providers } from 'valist';

export default async function getReleasesFromRepo(req: NextApiRequest, res: NextApiResponse) {

  // set .env.local to your local chain or set in production deployment
    if (process.env.WEB3_PROVIDER) {

        const valist = new Valist({ web3Provider: new Web3Providers.HttpProvider(process.env.WEB3_PROVIDER), metaTx: false });
        await valist.connect();

        const {
            query: { releaseName },
        } = req;

        const [orgName, repoName] = decodeURIComponent(releaseName.toString().replace("@", "")).split("/");
        const releases = await valist.getReleasesFromRepo(orgName.toString(), repoName.toString());

        const latestTag = await valist.getLatestTagFromRepo(orgName, repoName);
        let versions: any = {};

        for (let i = 0; i < releases.length; i++) {
            versions[releases[i].tag] = {
                name: repoName,
                version: releases[i].tag,
                repository: "",
                contributors: "",
                bugs: "",
                homepage: "",
                dependencies: {},
                dist: {
                    tarball: `https://ipfs.fleek.co/ipfs/${releases[i].releaseCID}`
                }
            };
        }

        return res.status(200).json({
            _id: "",
            name: "",
            "dist-tags": {
                latest: latestTag
            },
            versions
        });

    } else {
        return res.status(500).json({statusCode: 500, message: "No Web3 Provider!"});
    }
}
