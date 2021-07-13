import { Command } from '@oclif/command';
import { initValist } from '../../utils/config';
import { ADD_KEY, REVOKE_KEY, ROTATE_KEY } from '@valist/sdk/dist/constants';

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
      options: ['grant', 'revoke', 'rotate'],
    },
    {
      name: 'key',
      required: true,
    },
  ];

  async run() {
    const { args } = this.parse(RepoKey);
    const valist = await initValist();

    let opFunc: (orgName: string, repoName: string, key: string) => Promise<any>;
    let opName: string;

    switch(args.operation) {
      case 'grant':
        opFunc = async (orgName: string, repoName: string, key: string) => valist.voteRepoDev(orgName, repoName, key);
        opName = ADD_KEY;
        break;
      case 'revoke':
        opFunc = async (orgName: string, repoName: string, key: string) => valist.revokeRepoDev(orgName, repoName, key);
        opName = REVOKE_KEY;
        break;
      case 'rotate':
        opFunc = async (orgName: string, repoName: string, key: string) => valist.rotateRepoDev(orgName, repoName, key);
        opName = ROTATE_KEY;
        break;
      default:
        this.log('invalid operation');
        this.exit(0);
    }

    const { threshold } = await valist.getRepository(args.orgName, args.repoName);
    const { transactionHash } = await opFunc(args.orgName, args.repoName, args.key);
    const { signers } = await valist.getPendingRepoDevVotes(args.orgName, args.repoName, opName, args.key);

    if (signers.length < threshold) {
      this.log(`🗳 Voted to ${args.operation} ${args.key} on ${args.orgName}/${args.repoName}: ${signers.length}/${threshold}`);
    } else {
      this.log(`✅ Approved ${args.operation} ${args.key} on ${args.orgName}/${args.repoName}!`);
    }

    this.log('🔗 Transaction Hash:', transactionHash);
    this.exit(0);
  }
}
