import { NextApiRequest, NextApiResponse } from 'next'
import Valist, { Web3Providers } from 'valist';

export default async function orgMeta(req: NextApiRequest, res: NextApiResponse) {

  // @TODO toggle between local and infura based on NODE_ENV, load from .env
  const provider = new Web3Providers.HttpProvider('http://127.0.0.1:9545');

  const valist = new Valist(provider, false);
  await valist.connect();

  const {
    query: { orgName },
  } = req

  const orgMeta = await valist.getOrganizationMeta(orgName.toString());

    return res.status(200).json({orgMeta});

}
