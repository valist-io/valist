import fs from 'fs';
import yaml from 'js-yaml';
import { Command } from '@oclif/command';
import cli from 'cli-ux';

export default class ValistInit extends Command {
  static description = 'Generate a new valist project';

  static examples = [
    '$ valist init',
  ];

  async run() {
    const orgName = await cli.prompt('organization name');
    const repoName = await cli.prompt('repository name');
    const projectType = await cli.prompt('repository type (binary)', { required: false }) || 'binary';
    const build = await cli.prompt('build command');
    const tag = await cli.prompt('release tag (0.0.1)', { required: false }) || '0.0.1';
    const meta = await cli.prompt('release meta file (README.md)', { required: false }) || 'README.md';

    const defaultImages: { [string: string]: string } = {
      npm: 'node:buster',
      binary: 'golang:buster',
    };

    const image = await cli.prompt('docker image',
      { required: !(projectType === 'npm' || 'binary') }) || defaultImages[projectType];

    const valistFile = yaml.dump({
      type: projectType,
      org: orgName,
      project: repoName,
      image,
      build,
      tag,
      meta,
    });

    fs.writeFileSync('valist.yml', valistFile, 'utf8');
    this.exit(0);
  }
}
