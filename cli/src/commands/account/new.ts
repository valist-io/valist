import cli from 'cli-ux';
import { Command } from '@oclif/command';
import { createSignerKey, getSignerKey } from '../../utils/crypto';
import AccountGet from './get';

export default class AccountNew extends Command {
  static description = 'Create a new account';

  static examples = [
    '$ valist account:new',
  ];

  async run(): Promise<void> {
    let overwrite = true;
    try {
      if (await getSignerKey()) {
        overwrite = await cli.confirm('⚠️  existing key found. overwrite? y/n');
      }
    } catch (e) {
      // noop
    }
    if (overwrite) {
      this.log('🛠  Generating new signer key...');
      await createSignerKey();
      this.log('🔒 Successfully stored in keychain/keyring');
      await AccountGet.run();
    }
    this.exit(0);
  }
}
