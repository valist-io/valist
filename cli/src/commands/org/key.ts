import { Command } from '@oclif/command';
import { initValist } from '../../utils/config';

export default class OrgGrantKey extends Command {
  static description = 'Propose a new admin key for organization';

  static examples = [
    '$ valist org:key exampleOrg add <key>',
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
    const { args } = this.parse(OrgGrantKey);
    const valist = await initValist();

    const { transactionHash } = await valist.voteOrgAdmin(args.orgName, args.key);

    this.log(`✅ Successfully voted to add Admin key to ${args.orgName}!`);
    this.log('🔗 Transaction Hash:', transactionHash);

    this.exit(0);
  }
}
