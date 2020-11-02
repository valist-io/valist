import { NextApiRequest, NextApiResponse } from 'next'
import Valist, { Web3Providers } from 'valist';

export default async function getRepoMeta(req: NextApiRequest, res: NextApiResponse) {

  // set .env.local to your local chain or set in production deployment
  if (process.env.WEB3_PROVIDER) {
    const provider = new Web3Providers.HttpProvider(process.env.WEB3_PROVIDER);

    const valist = new Valist(provider);
    await valist.connect();

    const {
      query: { orgName, repoName },
    } = req;

    const repoMeta = await valist.getRepoMeta(orgName.toString(), repoName.toString());

    if (repoMeta) {
      // return res.status(200).json({repoMeta});
      return res.redirect(`https://ipfs.io/ipfs/${repoMeta}`);
    } else {
      return res.status(404).json({statusCode: 404, message: "No repository found!"});
    }

  } else {
    return res.status(500).json({statusCode: 500, message: "No Web3 Provider!"});
  }
}
