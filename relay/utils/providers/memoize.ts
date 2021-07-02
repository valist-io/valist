import getConfig from 'next/config';
import Valist from 'valist';
import { Web3Providers } from 'valist/dist/utils';

let valist: Valist;

// eslint-disable-next-line import/prefer-default-export
export async function getMemoizedValist() {
  if (!valist) {
    const { publicRuntimeConfig } = getConfig();

    // set .env.local to your local chain or set in production deployment
    if (publicRuntimeConfig.WEB3_PROVIDER) {
      valist = new Valist({
        web3Provider: new Web3Providers.HttpProvider(publicRuntimeConfig.WEB3_PROVIDER),
        metaTx: false,
      });
      await valist.connect();
    } else {
      throw new Error('Missing WEB3_PROVIDER');
    }
  }
  return valist;
}
