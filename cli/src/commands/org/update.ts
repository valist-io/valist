import * as fs from 'fs';
import * as path from 'path';
import { Command } from '@oclif/command';
import { initValist } from '../../utils/config';

export default class OrgUpdate extends Command {
  static description = 'Update organization metadata';

  static examples = [
    '$ valist org:update exampleOrg meta/orgMeta.json',
  ];

  static args = [
    {
      name: 'orgName',
      required: true,
    },
    {
      name: 'orgMeta',
      required: true,
    },
  ];

  async run() {
    const { args } = this.parse(OrgUpdate);
    const valist = await initValist();

    // Look for path to meta file from current working directory
    const metaData = JSON.parse(fs.readFileSync(path.join(process.cwd(), args.orgMeta), 'utf8'));
    const { transactionHash } = await valist.setOrgMeta(args.orgName, metaData);

    this.log(`✅ Successfully updated ${args.orgName}!`);
    this.log('🔗 Transaction Hash:', transactionHash);

    this.exit(0);
  }
}
