import { Command } from '@oclif/command';
import { buildRelease } from '../utils/build';
import { parseValistConfig } from '../utils/config';

export default class Build extends Command {
  static description = 'Build the target valist project';

  static examples = [
    '$ valist build',
  ];

  async run(): Promise<void> {
    // Get current config from valist.yml
    const config = parseValistConfig();

    const releaseFiles = await buildRelease(config);
    this.log('🔨 Project built to path:', releaseFiles[releaseFiles.length - 1].path);
    this.exit(0);
  }
}
