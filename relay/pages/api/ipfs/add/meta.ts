import { NextApiRequest, NextApiResponse } from 'next'
import Valist, { Web3Providers } from 'valist';

export default async function addOrgMetaIPFS(req: NextApiRequest, res: NextApiResponse) {
    if (process.env.WEB3_PROVIDER && req.method === 'POST') {

      const valist = new Valist({ web3Provider: new Web3Providers.HttpProvider(process.env.WEB3_PROVIDER) });
      await valist.connect();

      const {
        body: { metaJSON },
      } = req;

      try {
        const ipfsResponse = await valist.addJSONtoIPFS(metaJSON);

        res.setHeader('Content-Type', 'application/json');
        res.status(200).json({ipfsResponse});
      } catch (err) {
        // If error
        res.setHeader('Content-Type', 'application/json');
        res.status(500).json({ statusCode: 500, message: err.message });
      }
    } else {
      res.status(500).json({ statusCode: 500, message: "This endpoint only supports POST" });
    }
}
