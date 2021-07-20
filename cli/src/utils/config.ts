import * as yaml from 'js-yaml';
import * as fs from 'fs';
import * as path from 'path';
import Valist from '@valist/sdk';
import type { ValistConfig, ProjectType } from '@valist/sdk/dist/types';
import { getWeb3Provider, getSignerKey } from './crypto';

export const initValist = async (): Promise<Valist> => {
  console.log('📡 Connecting to Valist...');
  try {
    let signer = await getSignerKey();

    const provider = await getWeb3Provider(signer);
    const valist = new Valist({ web3Provider: provider });

    valist.signer = signer;
    signer = '';

    await valist.connect();

    console.log('⚡️ Connected!');
    console.log('📇 Account:', valist.defaultAccount);

    return valist;
  } catch (e) {
    const msg = '😢 Could not connect to Valist';
    console.error(msg, e);
    throw e;
  }
};

export const supportedTypes: ProjectType[] = ['binary', 'go', 'node'];

// will need to tweak these over time
export const defaultImages: Record<ProjectType, string> = {
  binary: 'gcc:bullseye',
  node: 'node:buster',
  go: 'golang:buster',
  rust: 'rust:buster',
  python: 'python:buster',
  docker: 'scratch',
  'c++': 'gcc:bullseye',
  static: '',
};

export const defaultInstalls: Record<ProjectType, string> = {
  binary: 'make install',
  node: 'npm install',
  go: 'go get',
  rust: 'cargo install',
  python: 'pip install -r requirements.txt',
  docker: '',
  'c++': 'make install',
  static: '',
};

export const defaultBuilds: Record<ProjectType, string> = {
  binary: 'make build',
  node: 'npm run build',
  go: 'go build',
  rust: 'cargo build',
  python: 'python3 -m build',
  docker: '',
  'c++': 'make build',
  static: '',
};

export type PackageJson = {
  name: string,
  version: string
};

export const parsePackageJson = ():PackageJson => {
  const { name, version } = JSON.parse(fs.readFileSync(path.join(process.cwd(), 'package.json'), 'utf-8'));
  const packageJSON: PackageJson = {
    name,
    version,
  };

  return packageJSON;
};

export const parseValistConfig = (): ValistConfig => {
  try {
    const configFile: any = yaml.load(fs.readFileSync('./valist.yml', 'utf8'));

    if (!supportedTypes.includes(configFile.type)) {
      console.error('🚧 Project type not supported!');
      process.exit(1);
    }

    const config: ValistConfig = {
      type: configFile.type,
      org: configFile.org,
      repo: configFile.repo,
      tag: configFile.tag,
      meta: configFile.type === 'node' ? 'package.json' : configFile.meta,
      image: configFile.image || defaultImages[configFile.type as ProjectType],
      build: configFile.build,
      install: configFile.install,
      out: configFile.out,
    };

    if (!config.meta) {
      console.error('Metadata file required for this project type');
      process.exit(1);
    }

    // @TODO enforce all parameters are not null by this point

    return config;
  } catch (e) {
    const msg = 'Could not load valist.yml';
    console.error(msg, e);
    throw e;
  }
};
