import { AppProps } from 'next/app';
import getConfig from "next/config";
import React, { useEffect, useState } from 'react';

import Valist from 'valist';
import ValistContext from '../components/ValistContext/ValistContext';
import { Magic } from 'magic-sdk';

import LoginForm from '../components/LoginForm/LoginForm';

import '../styles/main.css';

const { publicRuntimeConfig } = getConfig();

function App({ Component, pageProps }: AppProps) {

  const [valist, setValist] = useState<Valist>();

  const [magic, setMagic] = useState<Magic>();

  const [loggedIn, setLoggedIn] = useState(false);

  // const [loginMethod, setLoginMethod] = useState<"magic" | "metamask" | "github">("magic");

  // initialize web3 and valist object on document load (this effect is only triggered once)
  useEffect(() => {
    (async function () {
        try {

          const customNodeOptions = {
            rpcUrl: publicRuntimeConfig.WEB3_PROVIDER
          };

          const magicObj = new Magic(publicRuntimeConfig.MAGIC_PUBKEY, { network: customNodeOptions });
          setMagic(magicObj);

        } catch (e) {
          console.log(e);
        }
    })();
  }, []);

  useEffect(() => {
    (async function () {
      if (magic && loggedIn) {
        try {
          console.log(publicRuntimeConfig)
          // @ts-ignore Magic's RPCProviderModule doesn't fit the web3.js provider types perfectly yet
          const valist = new Valist({ web3Provider: magic.rpcProvider, metaTx: publicRuntimeConfig.METATX_ENABLED });
          await valist.connect();

          setValist(valist);

          console.log("Current Account: ", valist.defaultAccount);
          console.log("Current Account Balance: ", valist.web3.utils.fromWei(await valist.web3.eth.getBalance(valist.defaultAccount), 'ether'));

          // @ts-ignore keep for dev purposes
          window.valist = valist;
        } catch (e) {
          console.error("Could not initialize Valist object", e);
        }
      }
    })();
  }, [magic, loggedIn]);

  return loggedIn ? (
    // @ts-ignore
    <ValistContext.Provider value={valist}>
      <Component {...pageProps} />
    </ValistContext.Provider>
  ) : <LoginForm magic={magic} setLoggedIn={setLoggedIn} {...pageProps} />
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

export default App
