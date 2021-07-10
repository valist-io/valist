import fs from 'fs';
import yaml from 'js-yaml';
import cli from 'cli-ux';
import { Command } from '@oclif/command';
import { ProjectType } from '@valist/sdk/dist/types';
import {
  supportedTypes,
  defaultBuilds,
  defaultInstalls,
  defaultImages,
} from '../utils/config';

export default class ValistInit extends Command {
  static description = 'Generate a new valist project';

  static examples = [
    '$ valist init',
  ];

  async run(): Promise<void> {
    const orgName = await cli.prompt('organization name');
    const repoName = await cli.prompt('repository name');

    let projectType: ProjectType = await cli.prompt('repository type (binary)');
    while (!supportedTypes.includes(projectType)) {
      this.log('Unsupported project type! Try one of the following:', supportedTypes);
      // eslint-disable-next-line no-await-in-loop
      projectType = await cli.prompt('repository type (binary)');
    }

    const install = await cli.prompt(`install command (${defaultInstalls[projectType]})`,
      { required: false }) || defaultInstalls[projectType];

    const build = await cli.prompt(`build command (${defaultBuilds[projectType]})`,
      { required: false }) || defaultBuilds[projectType];

    const tag = await cli.prompt('release tag (0.0.1)', { required: false }) || '0.0.1';

    let meta: string | undefined;
    if (projectType !== 'node') {
      meta = await cli.prompt('release meta file (README.md)', { required: false }) || 'README.md';
    }

    const image: string | undefined = await cli.prompt(`docker image (${defaultImages[projectType]})`,
      { required: false });

    const out = await cli.prompt('output file/directory (dist)');

    const configToWrite: any = {};
    configToWrite.type = projectType;
    configToWrite.org = orgName;
    configToWrite.repo = repoName;
    configToWrite.tag = tag;
    if (meta) configToWrite.meta = meta;
    if (image) configToWrite.image = image;
    configToWrite.install = install;
    configToWrite.build = build;
    configToWrite.out = out;

    const valistFile = yaml.dump(configToWrite);

    fs.writeFileSync('valist.yml', valistFile, 'utf8');
    this.exit(0);
  }
}
