import { Magic } from 'magic-sdk';
import WalletConnectProvider from '@walletconnect/web3-provider';
import getConfig from 'next/config';

const { publicRuntimeConfig } = getConfig();

function getProviders(setMagic: any, setLoggedIn: any, email: string) {
  return {
    magic: async () => {
      const customNodeOptions = {
        rpcUrl: publicRuntimeConfig.WEB3_PROVIDER,
      };

      const magicObj = new Magic(publicRuntimeConfig.MAGIC_PUBKEY, { network: customNodeOptions });
      const magicLoggedIn = await magicObj.user.isLoggedIn();
      setMagic(magicObj);

      if (magicLoggedIn) {
        setLoggedIn(true);
        return magicObj.rpcProvider;
      }

      await magicObj.auth.loginWithMagicLink({ email });
      setLoggedIn(true);
      return magicObj.rpcProvider;
    },
    walletConnect: async () => {
      const provider = new WalletConnectProvider({
        rpc: {
          80001: publicRuntimeConfig.WEB3_PROVIDER,
        },
      });

      //  Enable session (triggers QR Code modal)
      await provider.enable();
      setLoggedIn(true);
      return provider;
    },
    metaMask: async () => {
      await (window as any).ethereum.send('eth_requestAccounts');
      setLoggedIn(true);
      return (window as any).ethereum;
    },
    readOnly: async () => {
      setLoggedIn(false);
      return publicRuntimeConfig.WEB3_PROVIDER;
    },
  };
}

export default getProviders;
