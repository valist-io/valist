import cli from 'cli-ux';
import { Command } from '@oclif/command';
import { initValist } from '../../utils/config';

export default class OrgNew extends Command {
  static description = 'Create a Valist organization';

  static examples = [
    '$ valist org:new exampleOrg',
  ];

  static args = [
    {
      name: 'orgName',
      required: true,
    },
  ];

  async run(): Promise<void> {
    const { args } = this.parse(OrgNew);

    // Create a new valist instance and connect
    const valist = await initValist();

    // org metadata
    const name = await cli.prompt('organization full name');
    const description = await cli.prompt('description');

    const orgMeta = { name, description };

    this.log();
    this.log('⚙️  Creating organization...');

    const { transactionHash, orgID } = await valist.createOrganization(args.orgName, orgMeta);

    const { transactionHash: registryTxHash } = await valist.linkNameToID(args.orgName, orgID);

    this.log(`✅ Successfully Created orgID ${orgID}!`);
    this.log('🔗 Transaction Hash:', transactionHash);
    this.log();
    this.log(`⚙️  Linking orgID to ${args.orgName}...`);
    this.log(`✅ Successfully Created ${args.orgName}!`);
    this.log('🔗 Transaction Hash:', registryTxHash);
    this.log();
    this.log(`ℹ️  To create a repo within this org, run \`valist repo:new ${args.orgName} exampleRepo\``);
    this.log();

    this.exit(0);
  }
}
