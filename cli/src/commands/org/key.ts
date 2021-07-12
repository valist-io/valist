import { Command } from '@oclif/command';
import { initValist } from '../../utils/config';
import { ADD_KEY, REVOKE_KEY, ROTATE_KEY } from '@valist/sdk/dist/constants';

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
      options: ['grant', 'revoke', 'rotate'],
    },
    {
      name: 'key',
      required: true,
    },
  ];

  async run() {
    const { args } = this.parse(OrgKey);
    const valist = await initValist();

    let opFunc: (orgName: string, key: string) => Promise<any>;
    let opName: string;

    switch(args.operation) {
      case 'grant':
        opFunc = (orgName: string, key: string) => valist.voteOrgAdmin(orgName, key);
        opName = ADD_KEY;
        break;
      case 'revoke':
        opFunc = (orgName: string, key: string) => valist.revokeOrgAdmin(orgName, key);
        opName = REVOKE_KEY;
        break;
      case 'rotate':
        opFunc = (orgName: string, key: string) => valist.rotateOrgAdmin(orgName, key);
        opName = ROTATE_KEY;
        break;
      default:
        this.log('invalid operation');
        this.exit(0);
    }

    const { threshold } = await valist.getOrganization(args.orgName);
    const { transactionHash } = await opFunc(args.OrgName, args.key);
    const { signers } = await valist.getPendingOrgAdminVotes(args.orgName, opName, args.key);

    if (signers.length < threshold) {
      this.log(`🗳 Voted to ${args.operation} ${args.key} on ${args.orgName}: ${signers.length}/${threshold}`);
    } else {
      this.log(`✅ Approved ${args.operation} ${args.key} on ${args.orgName}!`);
    }
    
    this.log('🔗 Transaction Hash:', transactionHash);
    this.exit(0);
  }
}
