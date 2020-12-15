import { AppProps } from 'next/app';
import React, { useEffect, useState } from 'react';

import Valist from 'valist';
import ValistContext from '../components/ValistContext/ValistContext';
import { Magic } from 'magic-sdk';

import LoginForm from '../components/LoginForm/LoginForm';

import '../styles/main.css';

function App({ Component, pageProps }: AppProps) {

  const [valist, setValist] = useState<Valist>();

  const [magic, setMagic] = useState<Magic>();

  const [loggedIn, setLoggedIn] = useState(false);

  // initialize web3 and valist object on document load (this effect is only triggered once)
  useEffect(() => {
    (async function () {
        try {
            // Start Magic Provider Code
            const customNodeOptions = {
              rpcUrl: 'http://127.0.0.1:8545', // Your own node URL
              chainId: 5777 // Your own node's chainId
            }

            const magicObj = new Magic('pk_test_69A0114AF6E0F54E', { network: customNodeOptions });
            setMagic(magicObj);

        } catch (error) {
            console.log(error);
        }
    })();
  }, []);

  useEffect(() => {
    (async function() {
      if (magic) {
        // @ts-ignore Magic's RPCProviderModule doesn't fit the web3.js provider types yet
        const valist = new Valist(magic.rpcProvider, true);

        await valist.connect();

        // @ts-ignore
        window.valist = valist; // keep for testing purposes
        setValist(valist);
      }
    })();
  }, [magic]);

  return loggedIn ? (
    // @ts-ignore
    <ValistContext.Provider value={valist}>
      <Component {...pageProps} />
    </ValistContext.Provider>
  ) : magic ? <LoginForm magic={magic} setLoggedIn={setLoggedIn} {...pageProps} /> : <div>Loading...</div>
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
