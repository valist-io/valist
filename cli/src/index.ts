#!/usr/bin/env node
import * as yargs from 'yargs';

import { createSignerKey, getSignerAddress } from './utils/crypto';
import { publish } from './utils/publish';

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

yargs.parse();
