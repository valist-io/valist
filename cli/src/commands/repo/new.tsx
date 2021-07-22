import cli from 'cli-ux';
import { Command } from '@oclif/command';
import { ProjectType, RepoMeta } from '@valist/sdk/dist/types';
import { initValist, supportedTypes } from '../../utils/config';

export default class RepoNew extends Command {
  static description = 'Create a Valist repository';

  static examples = [
    '$ valist repo:new exampleOrg exampleRepo',
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

  async run(): Promise<void> {
    const { args } = this.parse(RepoNew);

    // Create a new valist instance and connect
    const valist = await initValist();

    let projectType: ProjectType = await cli.prompt('repository type (binary)');
    while (!supportedTypes.includes(projectType)) {
      this.log('Unsupported project type! Try one of the following:', supportedTypes);
      // eslint-disable-next-line no-await-in-loop
      projectType = await cli.prompt('repository type (binary)');
    }

    // repo metadata
    const name = await cli.prompt('repo full name');
    const description = await cli.prompt('description');
    const homepage = await cli.prompt('homepage');
    const repository = await cli.prompt('source code repository');

    const repoMeta: RepoMeta = {
      name,
      description,
      homepage,
      repository,
      projectType,
    };

    this.log();
    this.log('⚙️  Creating repo...');

    const { transactionHash } = await valist.createRepository(args.orgName,
      args.repoName, repoMeta, valist.defaultAccount);

    this.log(`✅ Successfully Created ${args.orgName}/${args.repoName}!`);
    this.log('🔗 Transaction Hash:', transactionHash);
    this.log();
    this.log('ℹ️  To get started publishing, run `valist init` to create build & deploy config.');
    this.log();

    this.exit(0);
  }
}
