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

    const { threshold } = await valist.getRepository(args.orgName, args.repo);
    const { transactionHash } = await valist.voteRepoThreshold(args.orgName, args.repoName, args.thresholdNumber);
    const { signers } = await valist.getPendingRepoThresholdVotes(args.orgName, args.repoName, args.thresholdNumber);

    if (threshold > 1 && signers.length < threshold) {
      this.log(
        `🗳  Voted to set threshold for ${args.orgName}/${args.repoName}: ${signers.length}/${args.thresholdNumber}`,
      );
    } else {
      this.log(`✅ Approved threshold of ${args.thresholdNumber} for ${args.orgName}/${args.repoName}!`);
    }
    this.log('🔗 Transaction Hash:', transactionHash);
    this.exit(0);
  }
}
