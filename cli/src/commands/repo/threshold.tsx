import { Command } from '@oclif/command';
import { initValist } from '../../utils/config';

export default class RepoThreshold extends Command {
  static description = 'Set or update repository threshold';

  static examples = [
    '$ valist repo:threshold exampleOrg exampleRepo <thresholdNumber>',
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
      name: 'thresholdNumber',
      required: true,
    },
  ];

  async run(): Promise<void> {
    const { args } = this.parse(RepoThreshold);
    const valist = await initValist();

    const { transactionHash } = await valist.voteRepoThreshold(args.orgName, args.repoName, args.thresholdNumber);

    this.log(`✅ Successfully voted to set threshold for ${args.orgName}/${args.repoName}!`);
    this.log('🔗 Transaction Hash:', transactionHash);

    this.exit(0);
  }
}
