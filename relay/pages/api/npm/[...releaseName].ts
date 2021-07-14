import { NextApiRequest, NextApiResponse } from 'next';
import { getMemoizedValist } from '../../../utils/providers/memoize';
import { withSentry } from '@sentry/nextjs';

const getReleasesFromRepo = async (req: NextApiRequest, res: NextApiResponse) => {
  const valist = await getMemoizedValist();

  const {
    query: { releaseName },
  } = req;

  let orgName: string;
  let repoName: string;
  const cleanReleaseName = releaseName.toString().replace(/,/g, '/');

  if (Array.isArray(releaseName) && releaseName.length > 1) {
    orgName = releaseName[0].toString().replace('@', '');
    repoName = releaseName[1].toString();
  } else {
    [orgName, repoName] = decodeURIComponent(releaseName.toString().replace('@', '')).split('/');
  }

  console.log('Parsed', orgName, repoName, 'from', releaseName);

  if (orgName && repoName) {
    try {
      const releases = await valist.getReleases(orgName.toString(), repoName.toString());
      if (Array.isArray(releases) && releases.length >= 1) {
        const latestTag = releases[releases.length - 1].tag;
        const versions: any = {};

        // eslint-disable-next-line no-plusplus
        for (let i = 0; i < releases.length; i++) {
          const { tag, metaCID, releaseCID } = releases[i];
          versions[tag] = {};
          try {
            // eslint-disable-next-line no-await-in-loop
            const packageJSON = await valist.fetchJSONfromIPFS(metaCID);
            if (typeof packageJSON !== 'string') {
              versions[tag] = packageJSON;
            }
          } catch (e) {
            // noop
          }
          // eslint-disable-next-line no-underscore-dangle
          versions[tag]._id = `${cleanReleaseName}@${tag}`;
          versions[tag].name = cleanReleaseName;
          versions[tag].version = tag;
          versions[tag].dist = {
            tarball: `https://gateway.valist.io/ipfs/${releaseCID}`,
          };
        }

        return res.status(200).json({
          _id: cleanReleaseName,
          name: cleanReleaseName,
          'dist-tags': {
            latest: latestTag,
          },
          versions,
        });
      }
    } catch (e) {
      console.error('Could not find', cleanReleaseName, 'on Valist', e);
    }
  }
  console.log(`Fetching Package ${cleanReleaseName} from https://registry.npmjs.org`);
  res.redirect(`https://registry.npmjs.org/${cleanReleaseName}`);
}

export default withSentry(getReleasesFromRepo);