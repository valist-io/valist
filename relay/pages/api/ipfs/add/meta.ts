import { NextApiRequest, NextApiResponse } from 'next';
import getConfig from 'next/config';
import Valist from 'valist';
import { Web3Providers } from 'valist/dist/utils';
import { withSentry } from '@sentry/nextjs';

const addOrgMetaIPFS = async (req: NextApiRequest, res: NextApiResponse) => {
  const { publicRuntimeConfig } = getConfig();

  if (publicRuntimeConfig.WEB3_PROVIDER && req.method === 'POST') {
    const valist = new Valist({
      web3Provider: new Web3Providers.HttpProvider(publicRuntimeConfig.WEB3_PROVIDER),
      metaTx: false,
    });
    await valist.connect();

    const {
      body: { metaJSON },
    } = req;

    try {
      const ipfsResponse = await valist.addJSONtoIPFS(metaJSON);

      res.setHeader('Content-Type', 'application/json');
      res.status(200).json({ ipfsResponse });
    } catch (err) {
      // If error
      res.setHeader('Content-Type', 'application/json');
      res.status(500).json({ statusCode: 500, message: err.message });
    }
  } else {
    res.status(500).json({ statusCode: 500, message: 'This endpoint only supports POST' });
  }
}

export default withSentry(addOrgMetaIPFS);