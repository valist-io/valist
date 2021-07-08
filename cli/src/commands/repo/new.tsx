import * as fs from 'fs';
import * as path from 'path';
import { Command } from '@oclif/command';
import { initValist } from '../../utils/config';

export default class RepoNew extends Command {
  static description = 'Create a Valist repository'

  static examples = [
    `$ valist repo:new exampleOrg exampleRepo meta/repoMeta.json`,
  ]

  static args = [
    {
      name: 'orgName',
      required: true
    },
    {
      name: 'repoName',
      required: true
    },
    {
      name: 'repoMeta',
      required: true
    }
  ]

  async run() {
    const { args } = this.parse(RepoNew);

    // Create a new valist instance and connect
    const valist = await initValist();

    // Look for path to meta file from current working directory
    const metaData = JSON.parse(fs.readFileSync(path.join(process.cwd(), args.repoMeta), 'utf8'));
    const { transactionHash } = await valist.createRepository(args.orgName, args.repoName, metaData, valist.defaultAccount);

    this.log(`✅ Successfully Created ${args.orgName}/${args.repoName}!`);
    this.log('🔗 Transaction Hash:', transactionHash);

    this.exit(0);
  }
}
