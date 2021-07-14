import { NextApiRequest, NextApiResponse } from 'next';
import getConfig from 'next/config';
import Valist from 'valist';
import { Web3Providers } from 'valist/dist/utils';
import { withSentry } from '@sentry/nextjs';

const getRepoMeta = async (req: NextApiRequest, res: NextApiResponse) => {
  const { publicRuntimeConfig } = getConfig();

  // set .env.local to your local chain or set in production deployment
  if (publicRuntimeConfig.WEB3_PROVIDER) {
    const valist = new Valist({
      web3Provider: new Web3Providers.HttpProvider(publicRuntimeConfig.WEB3_PROVIDER),
      metaTx: false,
    });
    await valist.connect();

    const {
      query: { orgName, repoName },
    } = req;

    const repo = await valist.getRepository(orgName.toString(), repoName.toString());

    if (repo) {
      res.status(200).json({ repoMeta: repo.meta });
    } else {
      res.status(404).json({ statusCode: 404, message: 'No repository found!' });
    }
  } else {
    res.status(500).json({ statusCode: 500, message: 'No Web3 Provider!' });
  }
}

export default withSentry(getRepoMeta);