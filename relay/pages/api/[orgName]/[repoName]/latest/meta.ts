import { NextApiRequest, NextApiResponse } from 'next'
import Valist, { Web3Providers } from 'valist';

export default async function getLatestReleaseMeta(req: NextApiRequest, res: NextApiResponse) {

  // set .env.local to your local chain or set in production deployment
  if (process.env.WEB3_PROVIDER) {

    const valist = new Valist({ web3Provider: new Web3Providers.HttpProvider(process.env.WEB3_PROVIDER) });
    await valist.connect();

    const {
      query: { orgName, repoName },
    } = req;

    const latestRelease = await valist.getLatestReleaseFromRepo(orgName.toString(), repoName.toString());

    if (latestRelease) {
      //return res.status(200).json({releaseMeta});
      return res.redirect(`https://cloudflare-ipfs.com/ipfs/${latestRelease.metaCID}`);
    } else {
      return res.status(404).json({statusCode: 404, message: "No release found!"});
    }

  } else {
    return res.status(500).json({statusCode: 500, message: "No Web3 Provider!"});
  }
}
