import { AppProps } from 'next/app';
import React, { useEffect, useState } from 'react';
import Valist from 'valist';
import '../styles/main.css';

function App({ Component, pageProps }: AppProps) {

  const [valist, setValist] = useState<Valist>();

  // initialize web3 and valist object on document load (this effect is only triggered once)
  useEffect(() => {
    (async function () {
        try {
          // @ts-ignore
          if (window.ethereum){
            // @ts-ignore
            window.ethereum.enable();
            // @ts-ignore
            let valist = new Valist(window.ethereum, true);
            await valist.connect();
            setValist(valist);
            // @ts-ignore
            window.valist = valist;
          }

        } catch (error) {
            console.log(error);
        }
    })();
  }, []);

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
