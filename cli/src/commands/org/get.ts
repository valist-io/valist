import { Command } from '@oclif/command';
import { initValist } from '../../utils/config';

export default class OrgGet extends Command {
  static description = 'print organization info';

  static examples = [
    '$ valist org:get valist',
  ];

  static args = [
    {
      name: 'orgName',
      required: true,
    },
  ];

  async run() {
    const { args } = this.parse(OrgGet);
    const valist = await initValist();
    const orgData = await valist.getOrganization(args.orgName);
    this.log(`Name: ${orgData.meta.name}, Description: ${orgData.meta.description}`);
    this.exit(0);
  }
}
