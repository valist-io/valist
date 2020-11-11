import { NextApiRequest, NextApiResponse } from 'next'
import Valist, { Web3Providers } from 'valist';

export default async function getReleaseByTag(req: NextApiRequest, res: NextApiResponse) {

    // set .env.local to your local chain or set in production deployment
    if (process.env.WEB3_PROVIDER) {
        const provider = new Web3Providers.HttpProvider(process.env.WEB3_PROVIDER);

        const valist = new Valist(provider);
        await valist.connect();

        const {
            query: { releaseName, tag },
        } = req;

        const [orgName, repoName] = decodeURIComponent(releaseName.toString().replace("@", "")).split("/");
        const release = await valist.getReleaseByTag(orgName, repoName, tag.toString());

        if (release) {
            return res.status(200).json({
                name: repoName,
                version: release.tag,
                repository: "",
                contributors: "",
                bugs: "",
                homepage: "",
                dependencies: {},
                dist: {
                    tarball: `https://ipfs.io/ipfs/${release.release}`
                }
            });
            //return res.redirect(`https://ipfs.io/ipfs/${release.release}`);
        } else {
            return res.status(404).json({statusCode: 404, message: "No release found!"});
        }

        } else {
            return res.status(500).json({statusCode: 500, message: "No Web3 Provider!"});
        }
    }
