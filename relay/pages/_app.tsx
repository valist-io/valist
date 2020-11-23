import { AppProps } from 'next/app';
import React, { useEffect, useState } from 'react';
import Valist from 'valist';
import '../styles/main.css';
import ValistContext from '../components/ValistContext/ValistContext';
import LoadingDialog from '../components/LoadingDialog/LoadingDialog';

function App({ Component, pageProps }: AppProps) {

  const [valist, setValist] = useState<Valist>();

  const [ethereumEnabled, setEthereumEnabled] = useState(false);

  // initialize web3 and valist object on document load (this effect is only triggered once)
  useEffect(() => {
    (async function () {
        try {
          // @ts-ignore
          if (window.ethereum) {
            // @ts-ignore
            window.ethereum.enable();
            // @ts-ignore
            let valist = new Valist(window.ethereum, true);
            await valist.connect();
            setValist(valist);
            // @ts-ignore
            window.valist = valist;
            setEthereumEnabled(true);
          }

        } catch (error) {
            console.log(error);
        }
    })();
  }, []);

  return (
    // @ts-ignore
    <ValistContext.Provider value={valist}>
      <Component {...pageProps} />
      { !ethereumEnabled &&
        <LoadingDialog>
          <p>MetaMask is currently required to use this app. Please install at <a href="https://metamask.io" className="text-blue-700">metamask.io</a>.</p>
          <p>We are working hard to remove this requirement.</p><br/>
          <p>Sign up for our beta at <a href="https://valist.io" className="text-blue-700">valist.io</a> to be notified!</p>
        </LoadingDialog>
      }
    </ValistContext.Provider>
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

export default App
