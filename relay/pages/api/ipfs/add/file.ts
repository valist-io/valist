import { NextApiRequest, NextApiResponse } from 'next';
import { IncomingForm } from 'formidable';
import axios from 'axios';
import FormData from 'form-data';
import fs from 'fs';

const pinFileToIPFS = async (pinataPublicKey: string, pinataPrivateKey: string, filePath: string) => {
    const url = `https://api.pinata.cloud/pinning/pinFileToIPFS`;

    // @TODO use local file for now, but will use direct streaming in the future
    let data = new FormData();
    data.append('file', fs.createReadStream(filePath));

    const metadata = JSON.stringify({
        name: 'test',
    });

    // pinataOptions are optional
    const pinataOptions = JSON.stringify({
        cidVersion: 1,
    });

    data.append('pinataMetadata', metadata);
    data.append('pinataOptions', pinataOptions);

    return await axios.post(url, data, {
        maxContentLength: 2000, // 250MB limit
        maxBodyLength: 2000,
        headers: {
            // @ts-ignore _boundary
            'Content-Type': `multipart/form-data; boundary=${data._boundary}`,
            pinata_api_key: pinataPublicKey,
            pinata_secret_api_key: pinataPrivateKey
        }
    });
};

export const config = {
    api: {
        bodyParser: false,
    },
};

export default async function addFiletoIPFS(req: NextApiRequest, res: NextApiResponse) {

    if (!process.env.PINATA_PUBLIC || !process.env.PINATA_PRIVATE) return res.status(500).json({ statusCode: 500, message: "Missing Pinata API Keys" });

    if (req.method === 'POST') {
        try {

            const parseForm = new Promise(function (resolve, reject) {

                const form = new IncomingForm();

                form.parse(req, (e, fields, files) => {
                    if (e) {
                        console.log("Could not parse form data", e);
                        reject(e);
                    }
                    console.log({ fields, files });
                    resolve(files);
                });

            });

            const formData: any = await parseForm;

            const resp = await pinFileToIPFS(process.env.PINATA_PUBLIC, process.env.PINATA_PRIVATE, formData.file.path);

            res.setHeader("Content-Type", "application/json");
            res.status(200).json({ hash: resp.data.IpfsHash });

        } catch (e) {
            console.error(e);
            res.setHeader("Content-Type", "application/json");
            res.status(500).json({ statusCode: 500, message: "Could not pin file to IPFS" });
        }

    } else {
        res.status(500).json({ statusCode: 500, message: "This endpoint only supports POST" });
    }
}
