import { NextApiRequest, NextApiResponse } from 'next'

export default async function test(req: NextApiRequest, res: NextApiResponse) {

    const {
      query: { test },
    } = req;

    return res.redirect(`https://ipfs.io/ipfs/${test}`);
}
