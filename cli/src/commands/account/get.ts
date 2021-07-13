import { Command } from '@oclif/command';
import { getSignerAddress } from '../../utils/crypto';

export default class AccountGet extends Command {
  static description = 'print account info';

  static examples = [
    '$ valist account:get',
  ];

  async run(): Promise<void> {
    const address = await getSignerAddress();
    this.log(`📇 Your signer address is: ${address}`);
    this.exit(0);
  }
}
