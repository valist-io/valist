import { NextApiRequest, NextApiResponse } from 'next'
import Valist, { Web3Providers } from 'valist';

export default async function getLatestReleaseTag(req: NextApiRequest, res: NextApiResponse) {

  // set .env.local to your local chain or set in production deployment
  if (process.env.WEB3_PROVIDER) {
    const provider = new Web3Providers.HttpProvider(process.env.WEB3_PROVIDER);

    const valist = new Valist(provider);
    await valist.connect();

    const {
      query: { orgName, repoName },
    } = req;

    const latestTag = await valist.getLatestTagFromRepo(orgName.toString(), repoName.toString());

    return res.status(200).json({latestTag});

  } else {
    return res.status(500).json({statusCode: 500, message: "No Web3 Provider!"});
  }
}
