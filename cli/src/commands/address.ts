import {Command } from '@oclif/command';
import { getSignerAddress } from '../utils/crypto';

export default class Address extends Command {
  static description = 'print current signer address'

  static examples = [
    `$ valist address`,
  ]

  async run() {
    const address = await getSignerAddress();
    this.log(address);
  }
}