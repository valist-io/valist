import { Magic } from 'magic-sdk';
import getConfig from "next/config";
const { publicRuntimeConfig } = getConfig();

function getProviders(setMagic: any, setLoggedIn: any, email:any) {
  return {
    magic: async () => {
        try {
            const customNodeOptions = {
                rpcUrl: publicRuntimeConfig.WEB3_PROVIDER
            };

            const magicObj = new Magic(publicRuntimeConfig.MAGIC_PUBKEY, { network: customNodeOptions });
            const magicLoggedIn = await magicObj.user.isLoggedIn();
            setMagic(magicObj);

            if (magicLoggedIn) {
              setLoggedIn(true);
              return magicObj.rpcProvider;
            } else if (email) {
              await magicObj.auth.loginWithMagicLink({ email });
              setLoggedIn(true);
              return magicObj.rpcProvider;
            }

        } catch (e) {
            console.error("Could not set Magic as provider", e);
        }
    },
    metaMask: async () => {
        if ((window as any).ethereum) {
            // @ts-ignore
            await window.ethereum.enable();
            setLoggedIn(true);
            // @ts-ignore
            return window.ethereum;
        }
    },
    readOnly: async () => {
        setLoggedIn(false);
        return publicRuntimeConfig.WEB3_PROVIDER;
    }
  }
}

export default getProviders;