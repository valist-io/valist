import * as yaml from 'js-yaml';
import * as fs from 'fs';

// eslint-disable-next-line
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
