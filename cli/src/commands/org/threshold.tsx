import { Command } from '@oclif/command';
import { initValist } from '../../utils/config';

export default class OrgThreshold extends Command {
  static description = 'Set or update organization threshold';

  static examples = [
    '$ valist repo:threshold exampleOrg <thresholdNumber>',
  ];

  static args = [
    {
      name: 'orgName',
      required: true,
    },
    {
      name: 'thresholdNumber',
      required: true,
    },
  ];

  async run(): Promise<void> {
    const { args } = this.parse(OrgThreshold);
    const valist = await initValist();

    const { transactionHash } = await valist.voteOrgThreshold(args.orgName, args.thresholdNumber);

    this.log(`✅ Successfully voted to set threshold for ${args.orgName}!`);
    this.log('🔗 Transaction Hash:', transactionHash);

    this.exit(0);
  }
}
