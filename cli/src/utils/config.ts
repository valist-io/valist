import * as yaml from 'js-yaml';
import * as fs from 'fs';
import Valist from 'valist';
import { getWeb3Provider, getSignerKey } from './crypto';

export const initValist = async () => {
  try {
    let signer: string | null = await getSignerKey();

    const valist = new Valist({ web3Provider: await getWeb3Provider(signer) });

    valist.signer = signer;
    signer = null;

    const waitForMetaTx: boolean = true;

    await valist.connect(waitForMetaTx);

    console.log('📇 Account:', valist.defaultAccount);

    return valist;
  } catch (e) {
    const msg = '😢 Could not connect to Valist';
    console.error(msg, e);
    throw e;
  }
};

export const parseValistConfig = () => {
  try {
    const config: any = yaml.load(fs.readFileSync('./valist.yml', 'utf8'));
    return config || {};
  } catch (e) {
    const msg = 'Could not load valist.yml';
    console.error(msg, e);
    throw e;
  }
};
