import { NextApiRequest, NextApiResponse } from 'next';
import Valist from 'valist';
import { Web3Providers } from 'valist/dist/utils';

export default async function getLatestReleaseMeta(req: NextApiRequest, res: NextApiResponse) {
  // set .env.local to your local chain or set in production deployment
  if (process.env.WEB3_PROVIDER) {
    const valist = new Valist({
      web3Provider: new Web3Providers.HttpProvider(process.env.WEB3_PROVIDER),
      metaTx: false,
    });
    await valist.connect();

    const {
      query: { orgName, repoName },
    } = req;

    const latestRelease = await valist.getLatestReleaseFromRepo(orgName.toString(), repoName.toString());

    if (latestRelease) {
      // return res.status(200).json({releaseMeta});
      return res.redirect(`https://gateway.valist.io/ipfs/${latestRelease.metaCID}`);
    }
    return res.status(404).json({ statusCode: 404, message: 'No release found!' });
  }
  return res.status(500).json({ statusCode: 500, message: 'No Web3 Provider!' });
}
