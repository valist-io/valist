import { Command } from '@oclif/command';
import { createSignerKey } from '../utils/crypto';

export default class Keygen extends Command {
  static description = 'create a new signer key'

  static examples = [
    `$ valist keygen`,
  ]

  async run() {
    this.log('🛠 Generating new signer key...');
    const address = await createSignerKey();
    this.log('🔒 Successfully stored in keychain/keyring');
    this.log('📇 Your new signer address is:', address);
  }
}