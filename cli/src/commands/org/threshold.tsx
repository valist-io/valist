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
    const { signers } = await valist.getPendingOrgThresholdVotes(args.orgName, args.thresholdNumber);

    if (signers.length < args.thresholdNumber) {
      this.log(
        `🗳  Voted to set threshold for ${args.orgName}: ${signers.length + 1}/${args.thresholdNumber}`,
      );
    } else {
      this.log(`✅ Approved threshold of ${args.thresholdNumber} for ${args.orgName}}!`);
    }
    this.log('🔗 Transaction Hash:', transactionHash);
    this.exit(0);
  }
}
