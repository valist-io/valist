import { NextApiRequest, NextApiResponse } from 'next'

export default async function test(req: NextApiRequest, res: NextApiResponse) {

    const {
      query: { test },
    } = req;

    console.log(test)

    return res.redirect(`https://ipfs.io/ipfs/${test}`);
}
