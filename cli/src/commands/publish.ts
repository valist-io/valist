import { Command } from '@oclif/command';
import { publish } from '../utils/publish';

export default class Publish extends Command {
  static description = 'Publish a package to Valist';

  static examples = [
    '$ valist publish',
  ];

  async run() {
    await publish();
    this.exit(0);
  }
}
