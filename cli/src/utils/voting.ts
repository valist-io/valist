import { initValist } from './config';

export const voteRepoDev = async (orgShortName: string, repoName:string, key:string): Promise<void> => {
  const valist = await initValist();
  const { transactionHash } = await valist.voteRepoDev(orgShortName, repoName, key);

  console.log(`✅ Successfully voted to add Developer key to ${orgShortName}/${repoName}!`);
  console.log('🔗 Transaction Hash:', transactionHash);
};

export const voteOrgAdmin = async (orgShortName: string, key:string): Promise<void> => {
  const valist = await initValist();
  const { transactionHash } = await valist.voteOrgAdmin(orgShortName, key);

  console.log(`✅ Successfully voted to add Admin key to ${orgShortName}!`);
  console.log('🔗 Transaction Hash:', transactionHash);
};
