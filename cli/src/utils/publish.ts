import * as fs from 'fs';
import * as path from 'path';
import Valist from '@valist/sdk';
import { buildRelease } from './build';
import { initValist, parseValistConfig } from './config';

const releaseExists = async (valist: Valist, org: string, project: string, tag: string) => {
  const { releaseCID } = await valist.getReleaseByTag(org, project, tag);
  if (releaseCID) {
    return true;
  }
  return false;
};

export const publish = async (): Promise<void> => {
  // Placeholder for release file path
  // let releaseFile;

  // Create a new valist instance and connect
  const valist = await initValist();

  // Get current config from valist.yml
  const config = parseValistConfig();

  // Get org, project, tag, artifact, meta from config
  const {
    org, repo, tag, meta,
  } = config;

  // Check if release exists
  if (await releaseExists(valist, org, repo, tag)) {
    console.log('✅ Release already exists, skipping publish');
    process.exit(0);
  }

  // // Check if environment is CI/CD and artifact exists
  // if (process.env.CI && out) {
  //   // Read artifact and metadata from disk
  //   releaseFile = fs.createReadStream(path.join(process.cwd(), out));
  // } else {
  // Call buildRelease with project type (npm, binary, etc) to return artifact path
  const releaseFiles = await buildRelease(config);
  console.log('files', releaseFiles.length);
  // }

  const metaFile = fs.createReadStream(path.join(process.cwd(), meta as string));

  console.log('🪐 Preparing release on IPFS...');
  const releaseObject = await valist.prepareRelease(config, releaseFiles, metaFile);
  console.log('📦 Release Object:', releaseObject);

  // cleanup generated tarball/build artifact
  if (config.type === 'node') {
    fs.unlinkSync(releaseFiles[0].path);
  }

  // try {
  //   console.log('⚡️ Publishing Release to Valist...');
  //   const { threshold } = await valist.getRepository(org, repo);
  //   const { transactionHash } = await valist.publishRelease(org, repo, releaseObject);
  //   const { signers } = await valist.getPendingReleaseVotes(org, repo, releaseObject);

  //   if (signers.length < threshold) {
  //     console.log(`🗳 Voted to publish release ${org}/${repo}/${tag}: ${signers.length}/${threshold}`);
  //   } else {
  //     console.log(`✅ Approved release ${org}/${repo}/${tag}!`);
  //   }

  //   console.log('📖 IPFS address of release:', `ipfs://${releaseObject.releaseCID}`);
  //   console.log('🔗 Transaction Hash:', transactionHash);
  // } catch (e) {
  //   // noop, error already handled/logged within, move on to cleanup
  // }
};
