import { NextApiRequest, NextApiResponse } from 'next';
import Valist from 'valist';
import { Web3Providers } from 'valist/dist/utils';

export default async function getReleasesFromRepo(req: NextApiRequest, res: NextApiResponse) {
  console.log('Pulling package list');

  // set .env.local to your local chain or set in production deployment
  if (process.env.WEB3_PROVIDER) {
    const valist = new Valist({
      web3Provider: new Web3Providers.HttpProvider(process.env.WEB3_PROVIDER),
      metaTx: false,
    });
    await valist.connect();

    const {
      query: { releaseName },
    } = req;

    let orgName: string;
    let repoName: string;

    if (Array.isArray(releaseName) && releaseName.length > 1) {
      orgName = releaseName[0].toString().replace('@', '');
      repoName = releaseName[1].toString();
    } else {
      [orgName, repoName] = decodeURIComponent(releaseName.toString().replace('@', '')).split('/');
    }

    console.log('Parsed', orgName, repoName, 'from', releaseName);

    if (orgName && repoName) {
      try {
        const releases = await valist.getReleasesFromRepo(orgName.toString(), repoName.toString());
        if (releases) {
          const latestTag = await valist.getLatestTagFromRepo(orgName, repoName);
          const versions: any = {};

          // eslint-disable-next-line no-plusplus
          for (let i = 0; i < releases.length; i++) {
            const { tag, metaCID, releaseCID } = releases[i];
            // eslint-disable-next-line no-await-in-loop
            const meta = await valist.fetchJSONfromIPFS(metaCID);
            // overwrite name, version, dist with valist info
            versions[tag] = {
              ...meta,
              name: repoName,
              version: tag,
              dist: {
                tarball: `https://gateway.valist.io/ipfs/${releaseCID}`,
              },
            };
          }

          return res.status(200).json({
            _id: '',
            name: '',
            'dist-tags': {
              latest: latestTag,
            },
            versions,
          });
        }
      } catch (e) {
        console.log('Could not find package in Valist');
      }
    }
    console.log('Package not Registered on Valist');
    console.log(`Fetching Package ${releaseName} from https://registry.npmjs.org`);
    return res.redirect(`https://registry.npmjs.org/${releaseName.toString().replace(',', '/')}`);
  }
  return res.status(404).json({ statusCode: 404, message: 'Package Not Found' });
}
