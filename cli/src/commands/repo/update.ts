import * as fs from 'fs';
import * as path from 'path';
import { Command } from '@oclif/command';
import { initValist } from '../../utils/config';

export default class RepoUpdate extends Command {
  static description = 'Update repository metadata';

  static examples = [
    '$ valist repo:update exampleOrg exampleRepo meta/repoMeta.json',
  ];

  static args = [
    {
      name: 'orgName',
      required: true,
    },
    {
      name: 'repoName',
      required: true,
    },
    {
      name: 'repoMeta',
      required: true,
    },
  ];

  async run(): Promise<void> {
    const { args } = this.parse(RepoUpdate);
    const valist = await initValist();

    // Look for path to meta file from current working directory
    const metaData = JSON.parse(fs.readFileSync(path.join(process.cwd(), args.repoMeta), 'utf8'));
    const { transactionHash } = await valist.setRepoMeta(args.orgName, args.repoName, metaData);

    this.log(`✅ Successfully updated ${args.orgName}!`);
    this.log('🔗 Transaction Hash:', transactionHash);

    this.exit(0);
  }
}
