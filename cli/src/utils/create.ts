import * as fs from 'fs';
import * as path from 'path';
import { initValist } from './config';

export const createOrg = async (orgName:string, orgMeta:string) => {
  // Create a new valist instance and connect
  const valist = await initValist();

  // Look for path to meta file from current working directory
  const metaData = JSON.parse(fs.readFileSync(path.join(process.cwd(), orgMeta), 'utf8'));
  const { transactionHash } = await valist.createOrganization(orgName, metaData, valist.defaultAccount);

  console.log(`✅ Successfully Created ${orgName}!`);
  console.log('🔗 Transaction Hash:', transactionHash);
};

export const createRepo = async (orgName:string, repoName:string, repoMeta:string) => {
  // Create a new valist instance and connect
  const valist = await initValist();

  // Look for path to meta file from current working directory
  const metaData = JSON.parse(fs.readFileSync(path.join(process.cwd(), repoMeta), 'utf8'));
  const { transactionHash } = await valist.createRepository(orgName, repoName, metaData);

  console.log(`✅ Successfully Created ${orgName}/${repoName}!`);
  console.log('🔗 Transaction Hash:', transactionHash);
};
