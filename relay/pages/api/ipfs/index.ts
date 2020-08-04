import Web3 from 'web3';
import { NextApiRequest, NextApiResponse } from 'next'
import Valist from 'valist';

const handler = (_req: NextApiRequest, res: NextApiResponse) => {
  try {
    (async function () {
      const provider = new Web3(new Web3.providers.HttpProvider("https://mainnet.infura.io/v3/3e4419022dae4bcabf20093dde1c42cb"))
      const valist = new Valist(provider, true)
      await valist.connect()

      if (_req.method === 'POST') {
        console.log("post")
      } else {
        // Handle any other HTTP method
        const ipfsResponse = await valist.getFileIpfs("bafkreicayyzlyz2jr2wrjojby524opyrx65nxfdtzfq7vv5h4ioarob6ru")
        res.statusCode = 200
        res.setHeader('Content-Type', 'application/json')
        res.end(JSON.stringify({ipfsResponse}))
      }
    })();
    
  } catch (err) {
    // If error
    res.statusCode = 200
    res.setHeader('Content-Type', 'application/json')
    res.status(500).json({ statusCode: 500, message: err.message })
  }
}

export default handler
