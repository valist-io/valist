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

        const [orgName, repoName] = decodeURIComponent(releaseName.toString()).split("/");
        const releases = await valist.getReleasesFromRepo(orgName.toString(), repoName.toString());
        const release = releases[0]

        return res.status(200).json({
            _id: "",
            name: "",
            versions: [
                {
                    name: release,
                    version: "0.0.1",
                    repository: "",
                    contributors: "",
                    bugs: "",
                    homepage: "",
                    dependencies: {},
                    dist: {
                        tarball: `https://ipfs.io/ipfs/${release}`
                    }
                }
            ]
        });

    } else {
        return res.status(500).json({statusCode: 500, message: "No Web3 Provider!"});
    }
}
