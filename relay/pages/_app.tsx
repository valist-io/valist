import { AppProps } from 'next/app';
import getConfig from "next/config";
import { useEffect, useState } from 'react';

import Valist from 'valist';
import ValistContext from '../components/ValistContext/ValistContext';
import LoginContext from '../components/LoginContext/LoginContext';

import LoadingDialog from '../components/LoadingDialog/LoadingDialog';
import LoginForm from '../components/LoginForm/LoginForm';
import { Magic } from 'magic-sdk';

import Ceramic from '@ceramicnetwork/http-client';
import { IDX } from '@ceramicstudio/idx';

import '../styles/main.css';

const threeID = require("3id-connect");

type ProviderType = "magic" | "metaMask" | "readOnly";

const { publicRuntimeConfig } = getConfig();

function App({ Component, pageProps }: AppProps) {

  const [valist, setValist] = useState<Valist>();

  const [email, setEmail] = useState("");

  const [showLogin, setShowLogin] = useState(false);

  const [loggedIn, setLoggedIn] = useState(false);

  const [magic, setMagic] = useState<Magic | undefined>();

  const providers = {
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
            // @ts-ignore
            if (window.ethereum) {
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

  const handleLogin = async (providerType: ProviderType) => {
    let provider;

    try {
      provider = await providers[providerType]();
      window.localStorage.setItem("loginType", providerType);
    } catch (e) {
      console.log("Could not set provider, falling back to readOnly", e);
      provider = await providers["readOnly"]();
      window.localStorage.setItem("loginType", "readOnly");
    }

    try {
      const valist = new Valist({
        web3Provider: provider,
        metaTx: provider == publicRuntimeConfig.WEB3_PROVIDER ? false : publicRuntimeConfig.METATX_ENABLED
      });

      await valist.connect();
      setValist(valist);

      if (window.localStorage.getItem('loginType') !== "readOnly") {
        console.log('Valist', valist);
        console.log("Current Account: ", valist.defaultAccount);

        const threeIdConnect = new threeID.ThreeIdConnect();
        const authProvider = new threeID.EthereumAuthProvider(provider, valist.defaultAccount);
        await threeIdConnect.connect(authProvider);

        const didProvider = await threeIdConnect.getDidProvider();
        console.log("DID Provider", didProvider);

        const ceramic: any = new Ceramic('https://gateway-clay.ceramic.network');
        await ceramic.setDIDProvider(didProvider);
        console.log("DID PROVIDER SET");
        const idx = new IDX(ceramic);
        console.log("IDX Instance", idx);
        console.log("Basic Profile", await idx.get('basicProfile'));
      }


      // @ts-ignore keep for dev purposes
      window.valist = valist;
    } catch (e) {
      console.error("Could not initialize Valist object", e);
      try {
        await handleLogin("readOnly");
      } catch (e) {
        console.error("Critical error, could not login with desired method or readOnly", e);
      }
    }
  }

  const logOut = async () => {
    window.localStorage.clear();
    setShowLogin(false);
    setLoggedIn(false);
    if (magic) {
      magic.user.logout();
      setMagic(undefined);
    }
    await handleLogin("readOnly");
  }

  const loginObject = {
    setShowLogin: setShowLogin,
    loggedIn: loggedIn,
    logOut: logOut
  }

  useEffect(() => {
    (async function () {
      // check login type on first load and set provider to previous state
      const loginType = window.localStorage.getItem("loginType");

      if (loginType === "readOnly" || loginType === "magic" || loginType === "metaMask") {
        await handleLogin(loginType);
      } else {
        await handleLogin("readOnly");
        window.localStorage.setItem("loginType", "readOnly");
      }

    })();
  }, []);

  return (
    <LoginContext.Provider value={loginObject}>
      { valist ?
        <ValistContext.Provider value={valist}>
          <Component loggedIn={loggedIn} setShowLogin={setShowLogin} {...pageProps} />
          { showLogin && <LoginForm setShowLogin={setShowLogin} handleLogin={handleLogin} setEmail={setEmail} /> }
        </ValistContext.Provider>
        : <LoadingDialog>Loading...</LoadingDialog>
      }
    </LoginContext.Provider>
  )
}

// Only uncomment this method if you have blocking data requirements for
// every single page in your application. This disables the ability to
// perform automatic static optimization, causing every page in your app to
// be server-side rendered.
//
// MyApp.getInitialProps = async (appContext: AppContext) => {
//   // calls page's `getInitialProps` and fills `appProps.pageProps`
//   const appProps = await App.getInitialProps(appContext);

//   return { ...appProps }
// }

export default App;
