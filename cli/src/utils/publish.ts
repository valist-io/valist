import * as fs from 'fs';
import * as path from 'path';
import Valist from '@valist/sdk';
import { buildRelease } from './builds';
import { initValist, parseValistConfig } from './config';

// Store status of CI env-variable
const isCI = process.env.CI;

const releaseExists = async (valist: Valist, org: string, project: string, tag: string) => {
  const { releaseCID } = await valist.getReleaseByTag(org, project, tag);
  if (releaseCID && releaseCID.length > 0) {
    return true;
  }
  return false;
};

export const publish = async () => {
  // Placeholder for release file path
  let releaseFile;

  // Create a new valist instance and connect
  const valist = await initValist();

  // Get current config from valist.yml
  const config = parseValistConfig();

  // Get org, project, tag, artifact, meta from config
  const {
    org, project, tag, out, meta,
  } = config;

  // Check if release exists
  if (await releaseExists(valist, org, project, tag)) {
    console.log('✅ Release already exists, skipping publish');
    process.exit(0);
  }

  // Check if environment is CI/CD and artifact exists
  if (isCI && out) {
    // Read artifact and metadata from disk
    releaseFile = fs.createReadStream(path.join(process.cwd(), out));
  } else {
    // Call buildRelease with project type (npm, binary, etc) to return artifact path
    releaseFile = await buildRelease(config);
  }

  const metaFile = fs.createReadStream(path.join(process.cwd(), meta));

  console.log('🪐 Preparing release on IPFS...');
  const releaseObject = await valist.prepareRelease(tag, releaseFile, metaFile);
  console.log('📦 Release Object:', releaseObject);

  try {
    console.log('⚡️ Publishing Release to Valist...');
    const { transactionHash } = await valist.publishRelease(org, project, releaseObject);

    console.log(`✅ Successfully Released ${project} ${tag}!`);
    console.log('📖 IPFS address of release:', `ipfs://${releaseObject.releaseCID}`);
    console.log('🔗 Transaction Hash:', transactionHash);
  } catch (e) {
    // noop, error already handled/logged within, move on to cleanup
  }

  // cleanup generated tarball/build artifact
  if (config.type === 'npm') {
    fs.unlinkSync(releaseFile.path);
  }
};
