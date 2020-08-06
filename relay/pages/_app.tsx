import { AppProps } from 'next/app';
import React, { useState, useEffect } from 'react';
import Valist, {Web3Providers} from 'valist';
import '../styles/main.css';

function App({ Component, pageProps }: AppProps) {

  const [valist, setValist] = useState<Valist>();

  const providerOptions = {}

  // initialize web3 and valist object on document load
  useEffect(() => {
    (async function () {
        try {
          // @ts-ignore
          
          if (window.ethereum){
            window.ethereum.enable();
            setValist(new Valist(window.ethereum, false));
          }

        } catch (error) {
            alert(
                `Failed to load web3, accounts, or contract. Check console for details.`
            )
            console.log(error)
        }
    })();
  }, []);

  // activate valist object
  useEffect(() => {
    (async function() {
      if (valist) {
        try {
          await valist.connect();
        } catch (e) {
          console.log(e);
          alert(`Failed to connect to the Valist contracts!`);
        }

        //if (process.env.NODE_ENV == 'development') {
          // @ts-ignore
          window.valist = valist;
        //}
      }
    })();
  }, [valist]);

  return <Component {...pageProps} valist={valist} />
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
