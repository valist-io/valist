import { Command } from '@oclif/command';
import { initValist } from '../../utils/config';

export default class RepoKey extends Command {
  static description = 'Add, remove, or rotate repository key';

  static examples = [
    '$ valist repo:key exampleOrg exampleRepo grant <key>',
    '$ valist repo:key exampleOrg exampleRepo revoke <key>',
    '$ valist repo:key exampleOrg exampleRepo rotate <key>',
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
    const { args } = this.parse(RepoKey);
    const valist = await initValist();
    if (!['grant', 'revoke', 'rotate'].includes(args.operation)) {
      this.log('Invalid operation', args.operation);
      this.exit(1);
    }
    const operations: Record<string, any> = {
      grant: (orgName: string, repoName:string, key: string) => valist.voteRepoDev(orgName, repoName, key),
      revoke: (orgName: string, repoName:string, key: string) => valist.revokeRepoDev(orgName, repoName, key),
      rotate: (orgName: string, repoName:string, key: string) => valist.rotateRepoDev(orgName, repoName, key),
    };

    const { transactionHash } = await operations[args.operation](args.orgName, args.repoName, args.key);

    this.log(`✅ Successfully voted to ${args.operation} Developer key on ${args.orgName}/${args.repoName}!`);
    this.log('🔗 Transaction Hash:', transactionHash);

    this.exit(0);
  }
}
