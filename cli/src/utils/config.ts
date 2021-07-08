import * as yaml from 'js-yaml';
import * as fs from 'fs';
import * as Valist from '@valist/sdk';
import { getWeb3Provider, getSignerKey } from './crypto';
import { MissingKeyError } from './errors';

export type ValistConfig = {
  project: string,
  org: string,
  tag: string,
  meta: string,
  type: 'binary' | 'npm',
  artifact?: string,
};

export const initValist = async (): Promise<Valist> => {
  console.log('📡 Connecting to Valist...');
  try {
    let signer = await getSignerKey();
    if (!signer) throw new MissingKeyError();

    const provider = await getWeb3Provider(signer);
    const valist = new Valist({ web3Provider: provider });

    valist.signer = signer;
    signer = null;

    const waitForMetaTx = true;

    await valist.connect(waitForMetaTx);

    console.log('⚡️ Connected!');
    console.log('📇 Account:', valist.defaultAccount);

    return valist;
  } catch (e) {
    const msg = '😢 Could not connect to Valist';
    console.error(msg, e);
    throw e;
  }
};

export const parseValistConfig = (): ValistConfig => {
  try {
    const config: any = yaml.load(fs.readFileSync('./valist.yml', 'utf8'));

    if (!['binary', 'npm'].includes(config.type)) {
      console.error('🚧 Project type not supported!');
      process.exit(1);
    }

    return config;
  } catch (e) {
    const msg = 'Could not load valist.yml';
    console.error(msg, e);
    throw e;
  }
};
