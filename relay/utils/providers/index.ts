import { Magic } from 'magic-sdk';
import getConfig from 'next/config';

const { publicRuntimeConfig } = getConfig();

function getProviders(setMagic: any, setLoggedIn: any, email:string) {
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
    metaMask: async () => {
      await (window as any).ethereum.enable();
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
