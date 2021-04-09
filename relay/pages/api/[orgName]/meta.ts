import { NextApiRequest, NextApiResponse } from 'next';
import Valist, { Web3Providers } from 'valist';

export default async function getOrganizationMeta(req: NextApiRequest, res: NextApiResponse) {
  // set .env.local to your local chain or set in production deployment
  if (process.env.WEB3_PROVIDER) {
    const valist = new Valist({
      web3Provider: new Web3Providers.HttpProvider(process.env.WEB3_PROVIDER),
      metaTx: false,
    });
    await valist.connect();

    const {
      query: { orgName },
    } = req;

    const orgMeta = await valist.getOrganizationMeta(orgName.toString());

    if (orgMeta) {
      return res.status(200).json({ orgMeta });
    }
    return res.status(404).json({ statusCode: 404, message: 'No organization found!' });
  }
  return res.status(500).json({ statusCode: 500, message: 'No Web3 Provider!' });
}
