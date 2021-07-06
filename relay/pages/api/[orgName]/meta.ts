import { NextApiRequest, NextApiResponse } from 'next';
import getConfig from 'next/config';
import Valist from 'valist';
import { Web3Providers } from 'valist/dist/utils';

export default async function getOrganizationMeta(req: NextApiRequest, res: NextApiResponse) {
  const { publicRuntimeConfig } = getConfig();

  // set .env.local to your local chain or set in production deployment
  if (publicRuntimeConfig.WEB3_PROVIDER) {
    const valist = new Valist({
      web3Provider: new Web3Providers.HttpProvider(publicRuntimeConfig.WEB3_PROVIDER),
      metaTx: false,
    });
    await valist.connect();

    const {
      query: { orgName },
    } = req;

    const org = await valist.getOrganization(orgName.toString());

    if (org.meta) {
      return res.status(200).json({ orgMeta: org.meta });
    }
    return res.status(404).json({ statusCode: 404, message: 'No organization found!' });
  }
  return res.status(500).json({ statusCode: 500, message: 'No Web3 Provider!' });
}
