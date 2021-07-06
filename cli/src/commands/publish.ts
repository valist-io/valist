import { Command } from '@oclif/command';
import { publish } from '../utils/publish';

export default class Publish extends Command {
  static description = 'publish a package to valist'

  static examples = [
    `$ valist publish`,
  ]

  async run() {
    await publish();
  }
}
