#!/usr/bin/env node
import yargs from 'yargs';

import { createSignerKey, getSignerAddress } from './utils/crypto';
import { publish } from './utils/publish';
import { createOrg, createRepo } from './utils/create';
import { voteOrgAdmin, voteRepoDev } from './utils/voting';

yargs.command('address', 'Get current signer address', () => {}, async () => {
  const address = await getSignerAddress();
  console.log('Your signer address is:', address);
  process.exit(0);
});

yargs.command('keygen', 'Create a new signer key', () => {}, async () => {
  console.log('🛠 Generating new signer key...');
  const address = await createSignerKey();
  console.log('🔒 Successfully stored in keychain/keyring');
  console.log('📇 Your new signer address is:', address);
  process.exit(0);
});

yargs.command('publish', 'Publish package to Valist', () => {}, async () => {
  await publish();
  process.exit(0);
});

yargs.command('createOrg <orgName> <orgMeta>', 'Create a Valist organization', () => {},
  async (argv) => {
    await createOrg((argv.orgName as string), (argv.orgMeta as string));
    process.exit(0);
  });

yargs.command('createRepo <orgName> <repoName> <repoMeta>', 'Create a Valist repository', () => {},
  async (argv) => {
    await createRepo((argv.orgName as string), (argv.repoName as string), (argv.repoMeta as string));
    process.exit(0);
  });

yargs.command('voteOrgAdmin <orgName> <key>', 'Propose a new developer key for repo', () => {},
  async (argv) => {
    await voteOrgAdmin((argv.orgName as string), (argv.key as string));
    process.exit(0);
  }).string('key');

yargs.command('voteRepoDev <orgName> <repoName> <key>', 'Propose a new developer key for repo', () => {},
  async (argv) => {
    await voteRepoDev((argv.orgName as string), (argv.repoName as string), (argv.key as string));
    process.exit(0);
  }).string('key');

yargs.parse();
