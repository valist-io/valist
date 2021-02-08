import { NextApiRequest, NextApiResponse } from 'next'
import Valist, { Web3Providers } from 'valist';

export default async function getReleaseByTag(req: NextApiRequest, res: NextApiResponse) {

  // set .env.local to your local chain or set in production deployment
  if (process.env.WEB3_PROVIDER) {

    const valist = new Valist({ web3Provider: new Web3Providers.HttpProvider(process.env.WEB3_PROVIDER), metaTx: false });
    await valist.connect();

    const {
      query: { orgName, repoName, tag },
    } = req;

    const release = await valist.getReleaseByTag(orgName.toString(), repoName.toString(), tag.toString());

    if (release) {
      //return res.status(200).json({release});
      return res.redirect(`https://ipfs.fleek.co/ipfs/${release.releaseCID}`);
    } else {
      return res.status(404).json({statusCode: 404, message: "No release found!"});
    }

  } else {
    return res.status(500).json({statusCode: 500, message: "No Web3 Provider!"});
  }
}
