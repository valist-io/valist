import { Command } from '@oclif/command';
import { createSignerKey } from '../../utils/crypto';
import AccountGet from './get';

export default class AccountNew extends Command {
  static description = 'create a new account'

  static examples = [
    `$ valist account:new`,
  ]

  async run() {
    this.log('🛠 Generating new signer key...');
    await createSignerKey();
    this.log('🔒 Successfully stored in keychain/keyring');
    await AccountGet.run();
    this.exit(0);
  }
}