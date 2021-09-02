import { NextApiRequest, NextApiResponse } from 'next';
import getConfig from 'next/config';
import Valist from 'valist';
import { Web3Providers } from 'valist/dist/utils';
import { withSentry } from '@sentry/nextjs';

const getOrganizationMeta = async (req: NextApiRequest, res: NextApiResponse) => {
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
      res.status(200).json({ orgMeta: org.meta });
    } else {
      res.status(404).json({ statusCode: 404, message: 'No organization found!' });
    }
  } else {
    res.status(500).json({ statusCode: 500, message: 'No Web3 Provider!' });
  }
};

export default withSentry(getOrganizationMeta);
