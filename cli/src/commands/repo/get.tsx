import { Command } from '@oclif/command';
import { initValist } from '../../utils/config';

export default class OrgGet extends Command {
  static description = 'print organization info';

  static examples = [
    '$ valist repo:get exampleOrg exampleRepo',
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
  ];

  async run() {
    const { args } = this.parse(OrgGet);
    const valist = await initValist();
    const repoData = await valist.getRepository(args.orgName, args.repoName);
    this.log();
    this.log(`Org ID: ${repoData.orgID}`);
    this.log(`Repo: ${args.orgName}/${args.repoName}`);
    this.log(`Name: ${repoData.meta.name}`);
    this.log(`Description: ${repoData.meta.description}`);
    this.log(`Signature Threshold: ${repoData.threshold}`);
    this.log();
    this.exit(0);
  }
}
