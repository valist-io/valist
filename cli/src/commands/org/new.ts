import * as fs from 'fs';
import * as path from 'path';
import { Command } from '@oclif/command';
import { initValist } from '../../utils/config';

export default class OrgNew extends Command {
  static description = 'Create a Valist organization'

  static examples = [
    `$ valist org:new valist`,
  ]

  static args = [
    {
      name: 'orgName',
      required: true
    },
    {
      name: 'orgMeta',
      required: true
    }
  ]

  async run() {
    const { args } = this.parse(OrgNew);

    // Create a new valist instance and connect
    const valist = await initValist();

    // Look for path to meta file from current working directory
    const metaData = JSON.parse(fs.readFileSync(path.join(process.cwd(), args.orgMeta), 'utf8'));
    const { transactionHash } = await valist.createOrganization(args.orgName, metaData, valist.defaultAccount);

    this.log(`✅ Successfully Created ${args.orgName}!`);
    this.log('🔗 Transaction Hash:', transactionHash);

    this.exit(0);
  }
}
