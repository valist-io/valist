import { Command } from '@oclif/command';
import { initValist } from '../../utils/config';

export default class RepoGrantKey extends Command {
  static description = 'Propose a new developer key for repo';

  static examples = [
    '$ valist repo:key exampleOrg exampleRepo add <key>',
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
      name: 'operation',
      required: true,
    },
    {
      name: 'key',
      required: true,
    },
  ];

  async run() {
    const { args } = this.parse(RepoGrantKey);

    // Create a new valist instance and connect
    const valist = await initValist();

    const { transactionHash } = await valist.voteRepoDev(args.orgName, args.repoName, args.key);

    this.log(`✅ Successfully voted to add Developer key to ${args.orgShortName}/${args.repoName}!`);
    this.log('🔗 Transaction Hash:', transactionHash);

    this.exit(0);
  }
}
