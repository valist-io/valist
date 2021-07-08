import { Command } from '@oclif/command';
import { initValist } from '../../utils/config';

export default class OrgKey extends Command {
  static description = 'Add, remove, or rotate organization key';

  static examples = [
    '$ valist org:key exampleOrg grant <key>',
    '$ valist org:key exampleOrg revoke <key>',
    '$ valist org:key exampleOrg rotate <key>',
  ];

  static args = [
    {
      name: 'orgName',
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
    const { args } = this.parse(OrgKey);
    const valist = await initValist();
    if (!['grant', 'revoke', 'rotate'].includes(args.operation)) {
      this.log('Invalid operation', args.operation);
      this.exit(1);
    }
    const operations: Record<string, any> = {
      grant: valist.voteOrgAdmin,
      revoke: valist.revokeOrgAdmin,
      rotate: valist.rotateOrgAdmin,
    };

    const { transactionHash } = await operations[args.operation](args.orgName, args.key);

    this.log(`✅ Successfully voted to ${args.operation} Admin key on ${args.orgName}!`);
    this.log('🔗 Transaction Hash:', transactionHash);

    this.exit(0);
  }
}
